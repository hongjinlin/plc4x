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
    "errors"
    "io"
    "github.com/apache/plc4x/plc4go/internal/plc4go/utils"
)

// The data-structure of this message
type ModbusPDUDiagnosticRequest struct {
    SubFunction uint16
    Data uint16
    Parent *ModbusPDU
    IModbusPDUDiagnosticRequest
}

// The corresponding interface
type IModbusPDUDiagnosticRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUDiagnosticRequest) ErrorFlag() bool {
    return false
}

func (m *ModbusPDUDiagnosticRequest) FunctionFlag() uint8 {
    return 0x08
}

func (m *ModbusPDUDiagnosticRequest) Response() bool {
    return false
}


func (m *ModbusPDUDiagnosticRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUDiagnosticRequest(subFunction uint16, data uint16, ) *ModbusPDU {
    child := &ModbusPDUDiagnosticRequest{
        SubFunction: subFunction,
        Data: data,
        Parent: NewModbusPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastModbusPDUDiagnosticRequest(structType interface{}) *ModbusPDUDiagnosticRequest {
    castFunc := func(typ interface{}) *ModbusPDUDiagnosticRequest {
        if casted, ok := typ.(ModbusPDUDiagnosticRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*ModbusPDUDiagnosticRequest); ok {
            return casted
        }
        if casted, ok := typ.(ModbusPDU); ok {
            return CastModbusPDUDiagnosticRequest(casted.Child)
        }
        if casted, ok := typ.(*ModbusPDU); ok {
            return CastModbusPDUDiagnosticRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ModbusPDUDiagnosticRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (subFunction)
    lengthInBits += 16

    // Simple field (data)
    lengthInBits += 16

    return lengthInBits
}

func (m *ModbusPDUDiagnosticRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUDiagnosticRequestParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

    // Simple Field (subFunction)
    subFunction, _subFunctionErr := io.ReadUint16(16)
    if _subFunctionErr != nil {
        return nil, errors.New("Error parsing 'subFunction' field " + _subFunctionErr.Error())
    }

    // Simple Field (data)
    data, _dataErr := io.ReadUint16(16)
    if _dataErr != nil {
        return nil, errors.New("Error parsing 'data' field " + _dataErr.Error())
    }

    // Create a partially initialized instance
    _child := &ModbusPDUDiagnosticRequest{
        SubFunction: subFunction,
        Data: data,
        Parent: &ModbusPDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ModbusPDUDiagnosticRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (subFunction)
    subFunction := uint16(m.SubFunction)
    _subFunctionErr := io.WriteUint16(16, (subFunction))
    if _subFunctionErr != nil {
        return errors.New("Error serializing 'subFunction' field " + _subFunctionErr.Error())
    }

    // Simple Field (data)
    data := uint16(m.Data)
    _dataErr := io.WriteUint16(16, (data))
    if _dataErr != nil {
        return errors.New("Error serializing 'data' field " + _dataErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUDiagnosticRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "subFunction":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.SubFunction = data
            case "data":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Data = data
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *ModbusPDUDiagnosticRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.SubFunction, xml.StartElement{Name: xml.Name{Local: "subFunction"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    return nil
}

