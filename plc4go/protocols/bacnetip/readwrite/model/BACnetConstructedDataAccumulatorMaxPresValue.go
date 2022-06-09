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

// BACnetConstructedDataAccumulatorMaxPresValue is the data-structure of this message
type BACnetConstructedDataAccumulatorMaxPresValue struct {
	*BACnetConstructedData
	MaxPresValue *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataAccumulatorMaxPresValue is the corresponding interface of BACnetConstructedDataAccumulatorMaxPresValue
type IBACnetConstructedDataAccumulatorMaxPresValue interface {
	IBACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataAccumulatorMaxPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *BACnetConstructedDataAccumulatorMaxPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAccumulatorMaxPresValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAccumulatorMaxPresValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAccumulatorMaxPresValue) GetMaxPresValue() *BACnetApplicationTagUnsignedInteger {
	return m.MaxPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccumulatorMaxPresValue factory function for BACnetConstructedDataAccumulatorMaxPresValue
func NewBACnetConstructedDataAccumulatorMaxPresValue(maxPresValue *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataAccumulatorMaxPresValue {
	_result := &BACnetConstructedDataAccumulatorMaxPresValue{
		MaxPresValue:          maxPresValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAccumulatorMaxPresValue(structType interface{}) *BACnetConstructedDataAccumulatorMaxPresValue {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorMaxPresValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccumulatorMaxPresValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccumulatorMaxPresValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAccumulatorMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorMaxPresValue"
}

func (m *BACnetConstructedDataAccumulatorMaxPresValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAccumulatorMaxPresValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAccumulatorMaxPresValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccumulatorMaxPresValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataAccumulatorMaxPresValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorMaxPresValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxPresValue)
	if pullErr := readBuffer.PullContext("maxPresValue"); pullErr != nil {
		return nil, pullErr
	}
	_maxPresValue, _maxPresValueErr := BACnetApplicationTagParse(readBuffer)
	if _maxPresValueErr != nil {
		return nil, errors.Wrap(_maxPresValueErr, "Error parsing 'maxPresValue' field")
	}
	maxPresValue := CastBACnetApplicationTagUnsignedInteger(_maxPresValue)
	if closeErr := readBuffer.CloseContext("maxPresValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorMaxPresValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAccumulatorMaxPresValue{
		MaxPresValue:          CastBACnetApplicationTagUnsignedInteger(maxPresValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAccumulatorMaxPresValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorMaxPresValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (maxPresValue)
		if pushErr := writeBuffer.PushContext("maxPresValue"); pushErr != nil {
			return pushErr
		}
		_maxPresValueErr := m.MaxPresValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("maxPresValue"); popErr != nil {
			return popErr
		}
		if _maxPresValueErr != nil {
			return errors.Wrap(_maxPresValueErr, "Error serializing 'maxPresValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorMaxPresValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAccumulatorMaxPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}