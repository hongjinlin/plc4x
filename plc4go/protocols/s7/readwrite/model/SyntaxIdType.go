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

// SyntaxIdType is an enum
type SyntaxIdType uint8

type ISyntaxIdType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	SyntaxIdType_S7ANY             SyntaxIdType = 0x01
	SyntaxIdType_PBC_ID            SyntaxIdType = 0x13
	SyntaxIdType_ALARM_LOCKFREESET SyntaxIdType = 0x15
	SyntaxIdType_ALARM_INDSET      SyntaxIdType = 0x16
	SyntaxIdType_ALARM_ACKSET      SyntaxIdType = 0x19
	SyntaxIdType_ALARM_QUERYREQSET SyntaxIdType = 0x1A
	SyntaxIdType_NOTIFY_INDSET     SyntaxIdType = 0x1C
	SyntaxIdType_NCK               SyntaxIdType = 0x82
	SyntaxIdType_NCK_METRIC        SyntaxIdType = 0x83
	SyntaxIdType_NCK_INCH          SyntaxIdType = 0x84
	SyntaxIdType_DRIVEESANY        SyntaxIdType = 0xA2
	SyntaxIdType_SYM1200           SyntaxIdType = 0xB2
	SyntaxIdType_DBREAD            SyntaxIdType = 0xB0
)

var SyntaxIdTypeValues []SyntaxIdType

func init() {
	_ = errors.New
	SyntaxIdTypeValues = []SyntaxIdType{
		SyntaxIdType_S7ANY,
		SyntaxIdType_PBC_ID,
		SyntaxIdType_ALARM_LOCKFREESET,
		SyntaxIdType_ALARM_INDSET,
		SyntaxIdType_ALARM_ACKSET,
		SyntaxIdType_ALARM_QUERYREQSET,
		SyntaxIdType_NOTIFY_INDSET,
		SyntaxIdType_NCK,
		SyntaxIdType_NCK_METRIC,
		SyntaxIdType_NCK_INCH,
		SyntaxIdType_DRIVEESANY,
		SyntaxIdType_SYM1200,
		SyntaxIdType_DBREAD,
	}
}

func SyntaxIdTypeByValue(value uint8) SyntaxIdType {
	switch value {
	case 0x01:
		return SyntaxIdType_S7ANY
	case 0x13:
		return SyntaxIdType_PBC_ID
	case 0x15:
		return SyntaxIdType_ALARM_LOCKFREESET
	case 0x16:
		return SyntaxIdType_ALARM_INDSET
	case 0x19:
		return SyntaxIdType_ALARM_ACKSET
	case 0x1A:
		return SyntaxIdType_ALARM_QUERYREQSET
	case 0x1C:
		return SyntaxIdType_NOTIFY_INDSET
	case 0x82:
		return SyntaxIdType_NCK
	case 0x83:
		return SyntaxIdType_NCK_METRIC
	case 0x84:
		return SyntaxIdType_NCK_INCH
	case 0xA2:
		return SyntaxIdType_DRIVEESANY
	case 0xB0:
		return SyntaxIdType_DBREAD
	case 0xB2:
		return SyntaxIdType_SYM1200
	}
	return 0
}

func SyntaxIdTypeByName(value string) SyntaxIdType {
	switch value {
	case "S7ANY":
		return SyntaxIdType_S7ANY
	case "PBC_ID":
		return SyntaxIdType_PBC_ID
	case "ALARM_LOCKFREESET":
		return SyntaxIdType_ALARM_LOCKFREESET
	case "ALARM_INDSET":
		return SyntaxIdType_ALARM_INDSET
	case "ALARM_ACKSET":
		return SyntaxIdType_ALARM_ACKSET
	case "ALARM_QUERYREQSET":
		return SyntaxIdType_ALARM_QUERYREQSET
	case "NOTIFY_INDSET":
		return SyntaxIdType_NOTIFY_INDSET
	case "NCK":
		return SyntaxIdType_NCK
	case "NCK_METRIC":
		return SyntaxIdType_NCK_METRIC
	case "NCK_INCH":
		return SyntaxIdType_NCK_INCH
	case "DRIVEESANY":
		return SyntaxIdType_DRIVEESANY
	case "DBREAD":
		return SyntaxIdType_DBREAD
	case "SYM1200":
		return SyntaxIdType_SYM1200
	}
	return 0
}

func SyntaxIdTypeKnows(value uint8) bool {
	for _, typeValue := range SyntaxIdTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastSyntaxIdType(structType interface{}) SyntaxIdType {
	castFunc := func(typ interface{}) SyntaxIdType {
		if sSyntaxIdType, ok := typ.(SyntaxIdType); ok {
			return sSyntaxIdType
		}
		return 0
	}
	return castFunc(structType)
}

func (m SyntaxIdType) GetLengthInBits() uint16 {
	return 8
}

func (m SyntaxIdType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SyntaxIdTypeParse(readBuffer utils.ReadBuffer) (SyntaxIdType, error) {
	val, err := readBuffer.ReadUint8("SyntaxIdType", 8)
	if err != nil {
		return 0, nil
	}
	return SyntaxIdTypeByValue(val), nil
}

func (e SyntaxIdType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("SyntaxIdType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e SyntaxIdType) name() string {
	switch e {
	case SyntaxIdType_S7ANY:
		return "S7ANY"
	case SyntaxIdType_PBC_ID:
		return "PBC_ID"
	case SyntaxIdType_ALARM_LOCKFREESET:
		return "ALARM_LOCKFREESET"
	case SyntaxIdType_ALARM_INDSET:
		return "ALARM_INDSET"
	case SyntaxIdType_ALARM_ACKSET:
		return "ALARM_ACKSET"
	case SyntaxIdType_ALARM_QUERYREQSET:
		return "ALARM_QUERYREQSET"
	case SyntaxIdType_NOTIFY_INDSET:
		return "NOTIFY_INDSET"
	case SyntaxIdType_NCK:
		return "NCK"
	case SyntaxIdType_NCK_METRIC:
		return "NCK_METRIC"
	case SyntaxIdType_NCK_INCH:
		return "NCK_INCH"
	case SyntaxIdType_DRIVEESANY:
		return "DRIVEESANY"
	case SyntaxIdType_DBREAD:
		return "DBREAD"
	case SyntaxIdType_SYM1200:
		return "SYM1200"
	}
	return ""
}

func (e SyntaxIdType) String() string {
	return e.name()
}