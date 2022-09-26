/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

var (
	typeNames = flag.String("type", "", "comma-separated list of type names; must be set")
	output    = flag.String("output", "", "output file name; default srcdir/<type>_plc4xgen.go")
	buildTags = flag.String("tags", "", "comma-separated list of build tags to apply")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of plc4xgenerator:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tplc4xgenerator [flags] -type T [directory]\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tplc4xgenerator [flags] -type T files... # Must be a single package\n")
	_, _ = fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("plc4xgenerator: ")
	flag.Usage = Usage
	flag.Parse()
	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	typeList := strings.Split(*typeNames, ",")
	var tags []string
	if len(*buildTags) > 0 {
		tags = strings.Split(*buildTags, ",")
	}

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	// Parse the package once.
	var dir string
	generator := Generator{}
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		if len(tags) != 0 {
			log.Fatal("-tags option applies only to directories, not when files are specified")
		}
		dir = filepath.Dir(args[0])
	}

	generator.parsePackage(args, tags)

	// Print the header and package clause.
	generator.Printf(asfHeader)
	generator.Printf("// Code generated by \"plc4xgenerator %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " "))
	generator.Printf("\n")
	generator.Printf("package %s", generator.pkg.name)
	generator.Printf("\n")
	generator.Printf("import (\n")
	generator.Printf("\t\"github.com/apache/plc4x/plc4go/spi/utils\"\n")
	generator.Printf("\t\"fmt\"\n")
	generator.Printf(")\n")
	generator.Printf("\n")
	generator.Printf("var _ = fmt.Printf\n")

	// Run generate for each type.
	for _, typeName := range typeList {
		generator.generate(typeName)
	}

	// Format the output.
	src := generator.format()

	// Write to file.
	outputName := *output
	if outputName == "" {
		baseName := fmt.Sprintf("%s_plc4xgen.go", typeList[0])
		outputName = filepath.Join(dir, baseName)
	}
	err := ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string  // Name of the type.
	fields   []Field // Accumulator for fields of that type.

	trimPrefix  string
	lineComment bool
}

type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}

// parsePackage analyzes the single package constructed from the patterns and tags.
// parsePackage exits if there is an error.
func (g *Generator) parsePackage(patterns []string, tags []string) {
	cfg := &packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	g.addPackage(pkgs[0])
}

// addPackage adds a type checked Package and its syntax files to the generator.
func (g *Generator) addPackage(pkg *packages.Package) {
	g.pkg = &Package{
		name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}

	for i, file := range pkg.Syntax {
		g.pkg.files[i] = &File{
			file: file,
			pkg:  g.pkg,
		}
	}
}

// generate produces the String method for the named type.
func (g *Generator) generate(typeName string) {
	fields := make([]Field, 0, 100)
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		file.typeName = typeName
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			fields = append(fields, file.fields...)
		}
	}
	if len(fields) == 0 {
		log.Fatalf("no fields defined for type %s", typeName)
	}
	// TODO: for now we remove Default from the start (maybe move that to an option)
	logicalTypeName := strings.TrimPrefix(typeName, "Default")

	// Generate code that will fail if the constants change value.
	g.Printf("func (d *%s) Serialize(writeBuffer utils.WriteBuffer) error {\n", typeName)
	g.Printf("\tif err := writeBuffer.PushContext(\"%s\"); err != nil {\n", logicalTypeName)
	g.Printf("\t\treturn err\n")
	g.Printf("\t}\n")
	for _, field := range fields {
		if field.isDelegate {
			g.Printf("\t\t\tif err := d.%s.Serialize(writeBuffer); err != nil {\n", field.fieldType.(*ast.Ident).Name)
			g.Printf("\t\t\t\treturn err\n")
			g.Printf("\t\t\t}")
			continue
		}
		fieldName := field.name
		fieldNameUntitled := unTitle(fieldName)
		switch fieldType := field.fieldType.(type) {
		case *ast.SelectorExpr:
			g.Printf(serializableFieldTemplate, "d."+field.name, fieldNameUntitled)
		case *ast.Ident:
			switch fieldType.Name {
			case "bool":
				g.Printf(boolFieldSerialize, "d."+field.name, fieldNameUntitled)
			case "string":
				g.Printf(stringFieldSerialize, "d."+field.name, fieldNameUntitled)
			case "error":
				g.Printf(errorFieldSerialize, "d."+field.name, fieldNameUntitled)
			default:
				fmt.Printf("no support implemented %v\n", fieldType)
			}
		case *ast.ArrayType:
			g.Printf("if err := writeBuffer.PushContext(\"%s\", utils.WithRenderAsList(true)); err != nil {\n\t\treturn err\n\t}\n", fieldNameUntitled)
			g.Printf("for _, elem := range d.%s {", field.name)
			switch eltType := fieldType.Elt.(type) {
			case *ast.SelectorExpr:
				g.Printf(serializableFieldTemplate, "elem", "value")
			case *ast.Ident:
				switch eltType.Name {
				case "bool":
					g.Printf(boolFieldSerialize, "elem", "")
				case "string":
					g.Printf(stringFieldSerialize, "elem", "")
				case "error":
					g.Printf(errorFieldSerialize, "elem", "")
				default:
					fmt.Printf("no support implemented %v\n", fieldType)
				}
			}
			g.Printf("}\n")
			g.Printf("if err := writeBuffer.PopContext(\"%s\", utils.WithRenderAsList(true)); err != nil {\n\t\treturn err\n\t}\n", fieldNameUntitled)
		case *ast.MapType:
			if ident, ok := fieldType.Key.(*ast.Ident); !ok || ident.Name != "string" {
				fmt.Printf("Only string types are supported for maps right now")
			}
			g.Printf("if err := writeBuffer.PushContext(\"%s\", utils.WithRenderAsList(true)); err != nil {\n\t\treturn err\n\t}\n", fieldNameUntitled)
			// TODO: we use serializable or strings as we don't want to over-complex this
			g.Printf("for name, elemValue := range d.%s {\n", fieldName)
			g.Printf("\t\tif serializable, ok := elemValue.(utils.Serializable); ok {\n")
			g.Printf("\t\t\tif err := writeBuffer.PushContext(name); err != nil {\n")
			g.Printf("\t\t\t\treturn err\n")
			g.Printf("\t\t\t}\n")
			g.Printf("\t\t\tif err := serializable.Serialize(writeBuffer); err != nil {\n")
			g.Printf("\t\t\t\treturn err\n")
			g.Printf("\t\t\t}\n")
			g.Printf("\t\t\tif err := writeBuffer.PopContext(name); err != nil {\n")
			g.Printf("\t\t\t\treturn err\n")
			g.Printf("\t\t\t}\n")
			g.Printf("\t\t} else {\n")
			g.Printf("\t\t\telemValueAsString := fmt.Sprintf(\"%%v\", elemValue)\n")
			g.Printf("\t\t\tif err := writeBuffer.WriteString(name, uint32(len(elemValueAsString)*8), \"UTF-8\", elemValueAsString); err != nil {\n")
			g.Printf("\t\t\t\treturn err\n")
			g.Printf("\t\t\t}\n")
			g.Printf("\t\t}\n\t}\n")
			g.Printf("if err := writeBuffer.PopContext(\"%s\", utils.WithRenderAsList(true)); err != nil {\n\t\treturn err\n\t}\n", fieldNameUntitled)
		default:
			fmt.Printf("no support implemented %#v\n", fieldType)
		}
	}
	g.Printf("\tif err := writeBuffer.PopContext(\"%s\"); err != nil {\n", logicalTypeName)
	g.Printf("\t\treturn err\n")
	g.Printf("\t}\n")
	g.Printf("\treturn nil\n")
	g.Printf("}\n")
	g.Printf("\n")
	g.Printf(stringerTemplate, typeName)
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// Field represents a declared field.
type Field struct {
	name       string
	fieldType  ast.Expr
	isDelegate bool
}

func (f *Field) String() string {
	return f.name
}

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.TYPE {
		// We only care about type declarations.
		return true
	}
	for _, spec := range decl.Specs {
		typeSpec := spec.(*ast.TypeSpec)
		structDecl, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		if typeSpec.Name.Name != f.typeName {
			continue
		}
		fmt.Printf("Handling %s\n", typeSpec.Name.Name)
		for _, field := range structDecl.Fields.List {
			if len(field.Names) == 0 {
				fmt.Printf("\t adding delegate\n")
				if _, ok := field.Type.(*ast.Ident); ok {
					f.fields = append(f.fields, Field{
						fieldType:  field.Type,
						isDelegate: true,
					})
					continue
				} else {
					panic(fmt.Sprintf("Only struct delegates supported now. Type %T", field.Type))
				}
			}
			fmt.Printf("\t adding field %s %v\n", field.Names[0].Name, field.Type)
			f.fields = append(f.fields, Field{
				name:      field.Names[0].Name,
				fieldType: field.Type,
			})
		}
	}
	return false
}

var asfHeader = `/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

`

var stringerTemplate = `
func (d *%s) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
`

var serializableFieldTemplate = `
	if %[1]s != nil {
		if serializableField, ok := %[1]s.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("%[2]s"); err != nil {
				return err
			}
			if err := serializableField.Serialize(writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("%[2]s"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%%v", %[1]s)
			if err := writeBuffer.WriteString("%[2]s", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}
`

var boolFieldSerialize = `
	if err := writeBuffer.WriteBit("%[2]s", %[1]s); err != nil {
		return err
	}
`

var stringFieldSerialize = `
	if err := writeBuffer.WriteString("%[2]s", uint32(len(%[1]s)*8), "UTF-8", %[1]s); err != nil {
		return err
	}
`

var errorFieldSerialize = `
	if %[1]s != nil {
		if err := writeBuffer.WriteString("%[2]s", uint32(len(%[1]s.Error())*8), "UTF-8", %[1]s.Error()); err != nil {
			return err
		}
	}
`

func unTitle(value string) string {
	return strings.ToLower(string(value[0])) + value[1:]
}