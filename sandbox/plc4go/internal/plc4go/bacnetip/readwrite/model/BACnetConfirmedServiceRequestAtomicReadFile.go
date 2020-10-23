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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetConfirmedServiceRequestAtomicReadFile struct {
    BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestAtomicReadFile interface {
    IBACnetConfirmedServiceRequest
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceRequestAtomicReadFile) ServiceChoice() uint8 {
    return 0x06
}

func (m BACnetConfirmedServiceRequestAtomicReadFile) initialize() spi.Message {
    return m
}

func NewBACnetConfirmedServiceRequestAtomicReadFile() BACnetConfirmedServiceRequestInitializer {
    return &BACnetConfirmedServiceRequestAtomicReadFile{}
}

func CastIBACnetConfirmedServiceRequestAtomicReadFile(structType interface{}) IBACnetConfirmedServiceRequestAtomicReadFile {
    castFunc := func(typ interface{}) IBACnetConfirmedServiceRequestAtomicReadFile {
        if iBACnetConfirmedServiceRequestAtomicReadFile, ok := typ.(IBACnetConfirmedServiceRequestAtomicReadFile); ok {
            return iBACnetConfirmedServiceRequestAtomicReadFile
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetConfirmedServiceRequestAtomicReadFile(structType interface{}) BACnetConfirmedServiceRequestAtomicReadFile {
    castFunc := func(typ interface{}) BACnetConfirmedServiceRequestAtomicReadFile {
        if sBACnetConfirmedServiceRequestAtomicReadFile, ok := typ.(BACnetConfirmedServiceRequestAtomicReadFile); ok {
            return sBACnetConfirmedServiceRequestAtomicReadFile
        }
        if sBACnetConfirmedServiceRequestAtomicReadFile, ok := typ.(*BACnetConfirmedServiceRequestAtomicReadFile); ok {
            return *sBACnetConfirmedServiceRequestAtomicReadFile
        }
        return BACnetConfirmedServiceRequestAtomicReadFile{}
    }
    return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestAtomicReadFile) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetConfirmedServiceRequest.LengthInBits()

    return lengthInBits
}

func (m BACnetConfirmedServiceRequestAtomicReadFile) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestAtomicReadFileParse(io *utils.ReadBuffer) (BACnetConfirmedServiceRequestInitializer, error) {

    // Create the instance
    return NewBACnetConfirmedServiceRequestAtomicReadFile(), nil
}

func (m BACnetConfirmedServiceRequestAtomicReadFile) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return BACnetConfirmedServiceRequestSerialize(io, m.BACnetConfirmedServiceRequest, CastIBACnetConfirmedServiceRequest(m), ser)
}