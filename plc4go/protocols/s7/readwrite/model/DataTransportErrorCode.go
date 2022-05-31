/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// DataTransportErrorCode is an enum
type DataTransportErrorCode uint8

type IDataTransportErrorCode interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	DataTransportErrorCode_RESERVED                DataTransportErrorCode = 0x00
	DataTransportErrorCode_OK                      DataTransportErrorCode = 0xFF
	DataTransportErrorCode_ACCESS_DENIED           DataTransportErrorCode = 0x03
	DataTransportErrorCode_INVALID_ADDRESS         DataTransportErrorCode = 0x05
	DataTransportErrorCode_DATA_TYPE_NOT_SUPPORTED DataTransportErrorCode = 0x06
	DataTransportErrorCode_NOT_FOUND               DataTransportErrorCode = 0x0A
)

var DataTransportErrorCodeValues []DataTransportErrorCode

func init() {
	_ = errors.New
	DataTransportErrorCodeValues = []DataTransportErrorCode{
		DataTransportErrorCode_RESERVED,
		DataTransportErrorCode_OK,
		DataTransportErrorCode_ACCESS_DENIED,
		DataTransportErrorCode_INVALID_ADDRESS,
		DataTransportErrorCode_DATA_TYPE_NOT_SUPPORTED,
		DataTransportErrorCode_NOT_FOUND,
	}
}

func DataTransportErrorCodeByValue(value uint8) DataTransportErrorCode {
	switch value {
	case 0x00:
		return DataTransportErrorCode_RESERVED
	case 0x03:
		return DataTransportErrorCode_ACCESS_DENIED
	case 0x05:
		return DataTransportErrorCode_INVALID_ADDRESS
	case 0x06:
		return DataTransportErrorCode_DATA_TYPE_NOT_SUPPORTED
	case 0x0A:
		return DataTransportErrorCode_NOT_FOUND
	case 0xFF:
		return DataTransportErrorCode_OK
	}
	return 0
}

func DataTransportErrorCodeByName(value string) DataTransportErrorCode {
	switch value {
	case "RESERVED":
		return DataTransportErrorCode_RESERVED
	case "ACCESS_DENIED":
		return DataTransportErrorCode_ACCESS_DENIED
	case "INVALID_ADDRESS":
		return DataTransportErrorCode_INVALID_ADDRESS
	case "DATA_TYPE_NOT_SUPPORTED":
		return DataTransportErrorCode_DATA_TYPE_NOT_SUPPORTED
	case "NOT_FOUND":
		return DataTransportErrorCode_NOT_FOUND
	case "OK":
		return DataTransportErrorCode_OK
	}
	return 0
}

func DataTransportErrorCodeKnows(value uint8) bool {
	for _, typeValue := range DataTransportErrorCodeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDataTransportErrorCode(structType interface{}) DataTransportErrorCode {
	castFunc := func(typ interface{}) DataTransportErrorCode {
		if sDataTransportErrorCode, ok := typ.(DataTransportErrorCode); ok {
			return sDataTransportErrorCode
		}
		return 0
	}
	return castFunc(structType)
}

func (m DataTransportErrorCode) GetLengthInBits() uint16 {
	return 8
}

func (m DataTransportErrorCode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DataTransportErrorCodeParse(readBuffer utils.ReadBuffer) (DataTransportErrorCode, error) {
	val, err := readBuffer.ReadUint8("DataTransportErrorCode", 8)
	if err != nil {
		return 0, nil
	}
	return DataTransportErrorCodeByValue(val), nil
}

func (e DataTransportErrorCode) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("DataTransportErrorCode", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e DataTransportErrorCode) name() string {
	switch e {
	case DataTransportErrorCode_RESERVED:
		return "RESERVED"
	case DataTransportErrorCode_ACCESS_DENIED:
		return "ACCESS_DENIED"
	case DataTransportErrorCode_INVALID_ADDRESS:
		return "INVALID_ADDRESS"
	case DataTransportErrorCode_DATA_TYPE_NOT_SUPPORTED:
		return "DATA_TYPE_NOT_SUPPORTED"
	case DataTransportErrorCode_NOT_FOUND:
		return "NOT_FOUND"
	case DataTransportErrorCode_OK:
		return "OK"
	}
	return ""
}

func (e DataTransportErrorCode) String() string {
	return e.name()
}