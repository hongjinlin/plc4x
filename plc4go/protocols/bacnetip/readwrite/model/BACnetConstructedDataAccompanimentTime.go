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

// BACnetConstructedDataAccompanimentTime is the data-structure of this message
type BACnetConstructedDataAccompanimentTime struct {
	*BACnetConstructedData
	AccompanimentTime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataAccompanimentTime is the corresponding interface of BACnetConstructedDataAccompanimentTime
type IBACnetConstructedDataAccompanimentTime interface {
	IBACnetConstructedData
	// GetAccompanimentTime returns AccompanimentTime (property field)
	GetAccompanimentTime() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataAccompanimentTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAccompanimentTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCOMPANIMENT_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAccompanimentTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAccompanimentTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAccompanimentTime) GetAccompanimentTime() *BACnetApplicationTagUnsignedInteger {
	return m.AccompanimentTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccompanimentTime factory function for BACnetConstructedDataAccompanimentTime
func NewBACnetConstructedDataAccompanimentTime(accompanimentTime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataAccompanimentTime {
	_result := &BACnetConstructedDataAccompanimentTime{
		AccompanimentTime:     accompanimentTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAccompanimentTime(structType interface{}) *BACnetConstructedDataAccompanimentTime {
	if casted, ok := structType.(BACnetConstructedDataAccompanimentTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccompanimentTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccompanimentTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccompanimentTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAccompanimentTime) GetTypeName() string {
	return "BACnetConstructedDataAccompanimentTime"
}

func (m *BACnetConstructedDataAccompanimentTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAccompanimentTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accompanimentTime)
	lengthInBits += m.AccompanimentTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAccompanimentTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccompanimentTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataAccompanimentTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccompanimentTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accompanimentTime)
	if pullErr := readBuffer.PullContext("accompanimentTime"); pullErr != nil {
		return nil, pullErr
	}
	_accompanimentTime, _accompanimentTimeErr := BACnetApplicationTagParse(readBuffer)
	if _accompanimentTimeErr != nil {
		return nil, errors.Wrap(_accompanimentTimeErr, "Error parsing 'accompanimentTime' field")
	}
	accompanimentTime := CastBACnetApplicationTagUnsignedInteger(_accompanimentTime)
	if closeErr := readBuffer.CloseContext("accompanimentTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccompanimentTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAccompanimentTime{
		AccompanimentTime:     CastBACnetApplicationTagUnsignedInteger(accompanimentTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAccompanimentTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccompanimentTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (accompanimentTime)
		if pushErr := writeBuffer.PushContext("accompanimentTime"); pushErr != nil {
			return pushErr
		}
		_accompanimentTimeErr := m.AccompanimentTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("accompanimentTime"); popErr != nil {
			return popErr
		}
		if _accompanimentTimeErr != nil {
			return errors.Wrap(_accompanimentTimeErr, "Error serializing 'accompanimentTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccompanimentTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAccompanimentTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}