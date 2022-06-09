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

// BACnetConstructedDataRestoreCompletionTime is the data-structure of this message
type BACnetConstructedDataRestoreCompletionTime struct {
	*BACnetConstructedData
	CompletionTime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataRestoreCompletionTime is the corresponding interface of BACnetConstructedDataRestoreCompletionTime
type IBACnetConstructedDataRestoreCompletionTime interface {
	IBACnetConstructedData
	// GetCompletionTime returns CompletionTime (property field)
	GetCompletionTime() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataRestoreCompletionTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataRestoreCompletionTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESTORE_COMPLETION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataRestoreCompletionTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataRestoreCompletionTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataRestoreCompletionTime) GetCompletionTime() *BACnetApplicationTagUnsignedInteger {
	return m.CompletionTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRestoreCompletionTime factory function for BACnetConstructedDataRestoreCompletionTime
func NewBACnetConstructedDataRestoreCompletionTime(completionTime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataRestoreCompletionTime {
	_result := &BACnetConstructedDataRestoreCompletionTime{
		CompletionTime:        completionTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataRestoreCompletionTime(structType interface{}) *BACnetConstructedDataRestoreCompletionTime {
	if casted, ok := structType.(BACnetConstructedDataRestoreCompletionTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRestoreCompletionTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataRestoreCompletionTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataRestoreCompletionTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataRestoreCompletionTime) GetTypeName() string {
	return "BACnetConstructedDataRestoreCompletionTime"
}

func (m *BACnetConstructedDataRestoreCompletionTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataRestoreCompletionTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (completionTime)
	lengthInBits += m.CompletionTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataRestoreCompletionTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRestoreCompletionTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataRestoreCompletionTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRestoreCompletionTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (completionTime)
	if pullErr := readBuffer.PullContext("completionTime"); pullErr != nil {
		return nil, pullErr
	}
	_completionTime, _completionTimeErr := BACnetApplicationTagParse(readBuffer)
	if _completionTimeErr != nil {
		return nil, errors.Wrap(_completionTimeErr, "Error parsing 'completionTime' field")
	}
	completionTime := CastBACnetApplicationTagUnsignedInteger(_completionTime)
	if closeErr := readBuffer.CloseContext("completionTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRestoreCompletionTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataRestoreCompletionTime{
		CompletionTime:        CastBACnetApplicationTagUnsignedInteger(completionTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataRestoreCompletionTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRestoreCompletionTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (completionTime)
		if pushErr := writeBuffer.PushContext("completionTime"); pushErr != nil {
			return pushErr
		}
		_completionTimeErr := m.CompletionTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("completionTime"); popErr != nil {
			return popErr
		}
		if _completionTimeErr != nil {
			return errors.Wrap(_completionTimeErr, "Error serializing 'completionTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRestoreCompletionTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataRestoreCompletionTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}