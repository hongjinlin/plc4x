//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type GroupObjectDescriptorRealisationType6 struct {
}

// The corresponding interface
type IGroupObjectDescriptorRealisationType6 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewGroupObjectDescriptorRealisationType6() *GroupObjectDescriptorRealisationType6 {
	return &GroupObjectDescriptorRealisationType6{}
}

func CastGroupObjectDescriptorRealisationType6(structType interface{}) *GroupObjectDescriptorRealisationType6 {
	castFunc := func(typ interface{}) *GroupObjectDescriptorRealisationType6 {
		if casted, ok := typ.(GroupObjectDescriptorRealisationType6); ok {
			return &casted
		}
		if casted, ok := typ.(*GroupObjectDescriptorRealisationType6); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *GroupObjectDescriptorRealisationType6) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType6"
}

func (m *GroupObjectDescriptorRealisationType6) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *GroupObjectDescriptorRealisationType6) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *GroupObjectDescriptorRealisationType6) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func GroupObjectDescriptorRealisationType6Parse(io *utils.ReadBuffer) (*GroupObjectDescriptorRealisationType6, error) {

	// Create the instance
	return NewGroupObjectDescriptorRealisationType6(), nil
}

func (m *GroupObjectDescriptorRealisationType6) Serialize(io utils.WriteBuffer) error {

	return nil
}

func (m *GroupObjectDescriptorRealisationType6) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			}
		}
	}
}

func (m *GroupObjectDescriptorRealisationType6) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.knxnetip.readwrite.GroupObjectDescriptorRealisationType6"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m GroupObjectDescriptorRealisationType6) String() string {
	return string(m.Box("GroupObjectDescriptorRealisationType6", utils.DefaultWidth*2))
}

func (m GroupObjectDescriptorRealisationType6) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "GroupObjectDescriptorRealisationType6"
	}
	boxes := make([]utils.AsciiBox, 0)
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
