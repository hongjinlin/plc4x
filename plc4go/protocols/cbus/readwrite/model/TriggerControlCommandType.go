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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// TriggerControlCommandType is an enum
type TriggerControlCommandType uint8

type ITriggerControlCommandType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	TriggerControlCommandType_TRIGGER_EVENT  TriggerControlCommandType = 0x00
	TriggerControlCommandType_TRIGGER_MIN    TriggerControlCommandType = 0x01
	TriggerControlCommandType_TRIGGER_MAX    TriggerControlCommandType = 0x02
	TriggerControlCommandType_INDICATOR_KILL TriggerControlCommandType = 0x03
	TriggerControlCommandType_LABEL          TriggerControlCommandType = 0x04
)

var TriggerControlCommandTypeValues []TriggerControlCommandType

func init() {
	_ = errors.New
	TriggerControlCommandTypeValues = []TriggerControlCommandType{
		TriggerControlCommandType_TRIGGER_EVENT,
		TriggerControlCommandType_TRIGGER_MIN,
		TriggerControlCommandType_TRIGGER_MAX,
		TriggerControlCommandType_INDICATOR_KILL,
		TriggerControlCommandType_LABEL,
	}
}

func TriggerControlCommandTypeByValue(value uint8) (enum TriggerControlCommandType, ok bool) {
	switch value {
	case 0x00:
		return TriggerControlCommandType_TRIGGER_EVENT, true
	case 0x01:
		return TriggerControlCommandType_TRIGGER_MIN, true
	case 0x02:
		return TriggerControlCommandType_TRIGGER_MAX, true
	case 0x03:
		return TriggerControlCommandType_INDICATOR_KILL, true
	case 0x04:
		return TriggerControlCommandType_LABEL, true
	}
	return 0, false
}

func TriggerControlCommandTypeByName(value string) (enum TriggerControlCommandType, ok bool) {
	switch value {
	case "TRIGGER_EVENT":
		return TriggerControlCommandType_TRIGGER_EVENT, true
	case "TRIGGER_MIN":
		return TriggerControlCommandType_TRIGGER_MIN, true
	case "TRIGGER_MAX":
		return TriggerControlCommandType_TRIGGER_MAX, true
	case "INDICATOR_KILL":
		return TriggerControlCommandType_INDICATOR_KILL, true
	case "LABEL":
		return TriggerControlCommandType_LABEL, true
	}
	return 0, false
}

func TriggerControlCommandTypeKnows(value uint8) bool {
	for _, typeValue := range TriggerControlCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTriggerControlCommandType(structType interface{}) TriggerControlCommandType {
	castFunc := func(typ interface{}) TriggerControlCommandType {
		if sTriggerControlCommandType, ok := typ.(TriggerControlCommandType); ok {
			return sTriggerControlCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m TriggerControlCommandType) GetLengthInBits() uint16 {
	return 4
}

func (m TriggerControlCommandType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TriggerControlCommandTypeParse(readBuffer utils.ReadBuffer) (TriggerControlCommandType, error) {
	val, err := readBuffer.ReadUint8("TriggerControlCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TriggerControlCommandType")
	}
	if enum, ok := TriggerControlCommandTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return TriggerControlCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e TriggerControlCommandType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("TriggerControlCommandType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TriggerControlCommandType) PLC4XEnumName() string {
	switch e {
	case TriggerControlCommandType_TRIGGER_EVENT:
		return "TRIGGER_EVENT"
	case TriggerControlCommandType_TRIGGER_MIN:
		return "TRIGGER_MIN"
	case TriggerControlCommandType_TRIGGER_MAX:
		return "TRIGGER_MAX"
	case TriggerControlCommandType_INDICATOR_KILL:
		return "INDICATOR_KILL"
	case TriggerControlCommandType_LABEL:
		return "LABEL"
	}
	return ""
}

func (e TriggerControlCommandType) String() string {
	return e.PLC4XEnumName()
}