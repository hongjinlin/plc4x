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
    "github.com/apache/plc4x/plc4go/internal/plc4go/s7/readwrite/model"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi"
)

type S7XmlParserHelper struct {
}

func (m S7XmlParserHelper) Parse(typeName string, xmlString string) (spi.Message, error) {
    switch typeName {
    case "SzlId":
        var obj *model.SzlId
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "S7Message":
        var obj *model.S7Message
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "S7VarPayloadStatusItem":
        var obj *model.S7VarPayloadStatusItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "S7Parameter":
        var obj *model.S7Parameter
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "SzlDataTreeItem":
        var obj *model.SzlDataTreeItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "COTPPacket":
        var obj *model.COTPPacket
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "S7PayloadUserDataItem":
        var obj *model.S7PayloadUserDataItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "COTPParameter":
        var obj *model.COTPParameter
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "TPKTPacket":
        var obj *model.TPKTPacket
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "S7Payload":
        var obj *model.S7Payload
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "S7VarRequestParameterItem":
        var obj *model.S7VarRequestParameterItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "S7VarPayloadDataItem":
        var obj *model.S7VarPayloadDataItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "S7Address":
        var obj *model.S7Address
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "S7ParameterUserDataItem":
        var obj *model.S7ParameterUserDataItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    }
    return nil, errors.New("Unsupported type " + typeName)
}
