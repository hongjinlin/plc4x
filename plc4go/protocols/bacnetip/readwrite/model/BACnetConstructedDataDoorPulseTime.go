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

// BACnetConstructedDataDoorPulseTime is the data-structure of this message
type BACnetConstructedDataDoorPulseTime struct {
	*BACnetConstructedData
	DoorPulseTime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataDoorPulseTime is the corresponding interface of BACnetConstructedDataDoorPulseTime
type IBACnetConstructedDataDoorPulseTime interface {
	IBACnetConstructedData
	// GetDoorPulseTime returns DoorPulseTime (property field)
	GetDoorPulseTime() *BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataDoorPulseTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDoorPulseTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_PULSE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDoorPulseTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDoorPulseTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDoorPulseTime) GetDoorPulseTime() *BACnetApplicationTagUnsignedInteger {
	return m.DoorPulseTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorPulseTime factory function for BACnetConstructedDataDoorPulseTime
func NewBACnetConstructedDataDoorPulseTime(doorPulseTime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataDoorPulseTime {
	_result := &BACnetConstructedDataDoorPulseTime{
		DoorPulseTime:         doorPulseTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDoorPulseTime(structType interface{}) *BACnetConstructedDataDoorPulseTime {
	if casted, ok := structType.(BACnetConstructedDataDoorPulseTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorPulseTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDoorPulseTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDoorPulseTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDoorPulseTime) GetTypeName() string {
	return "BACnetConstructedDataDoorPulseTime"
}

func (m *BACnetConstructedDataDoorPulseTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDoorPulseTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorPulseTime)
	lengthInBits += m.DoorPulseTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDoorPulseTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDoorPulseTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataDoorPulseTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorPulseTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorPulseTime)
	if pullErr := readBuffer.PullContext("doorPulseTime"); pullErr != nil {
		return nil, pullErr
	}
	_doorPulseTime, _doorPulseTimeErr := BACnetApplicationTagParse(readBuffer)
	if _doorPulseTimeErr != nil {
		return nil, errors.Wrap(_doorPulseTimeErr, "Error parsing 'doorPulseTime' field")
	}
	doorPulseTime := CastBACnetApplicationTagUnsignedInteger(_doorPulseTime)
	if closeErr := readBuffer.CloseContext("doorPulseTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorPulseTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDoorPulseTime{
		DoorPulseTime:         CastBACnetApplicationTagUnsignedInteger(doorPulseTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDoorPulseTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorPulseTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (doorPulseTime)
		if pushErr := writeBuffer.PushContext("doorPulseTime"); pushErr != nil {
			return pushErr
		}
		_doorPulseTimeErr := m.DoorPulseTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("doorPulseTime"); popErr != nil {
			return popErr
		}
		if _doorPulseTimeErr != nil {
			return errors.Wrap(_doorPulseTimeErr, "Error serializing 'doorPulseTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorPulseTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDoorPulseTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}