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

// BACnetConstructedDataShedDuration is the data-structure of this message
type BACnetConstructedDataShedDuration struct {
	*BACnetConstructedData
	ShedDuration *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataShedDuration is the corresponding interface of BACnetConstructedDataShedDuration
type IBACnetConstructedDataShedDuration interface {
	IBACnetConstructedData
	// GetShedDuration returns ShedDuration (property field)
	GetShedDuration() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataShedDuration) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataShedDuration) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SHED_DURATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataShedDuration) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataShedDuration) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataShedDuration) GetShedDuration() *BACnetApplicationTagUnsignedInteger {
	return m.ShedDuration
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataShedDuration factory function for BACnetConstructedDataShedDuration
func NewBACnetConstructedDataShedDuration(shedDuration *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataShedDuration {
	_result := &BACnetConstructedDataShedDuration{
		ShedDuration:          shedDuration,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataShedDuration(structType interface{}) *BACnetConstructedDataShedDuration {
	if casted, ok := structType.(BACnetConstructedDataShedDuration); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataShedDuration); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataShedDuration(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataShedDuration(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataShedDuration) GetTypeName() string {
	return "BACnetConstructedDataShedDuration"
}

func (m *BACnetConstructedDataShedDuration) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataShedDuration) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (shedDuration)
	lengthInBits += m.ShedDuration.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataShedDuration) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataShedDurationParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataShedDuration, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataShedDuration"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (shedDuration)
	if pullErr := readBuffer.PullContext("shedDuration"); pullErr != nil {
		return nil, pullErr
	}
	_shedDuration, _shedDurationErr := BACnetApplicationTagParse(readBuffer)
	if _shedDurationErr != nil {
		return nil, errors.Wrap(_shedDurationErr, "Error parsing 'shedDuration' field")
	}
	shedDuration := CastBACnetApplicationTagUnsignedInteger(_shedDuration)
	if closeErr := readBuffer.CloseContext("shedDuration"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataShedDuration"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataShedDuration{
		ShedDuration:          CastBACnetApplicationTagUnsignedInteger(shedDuration),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataShedDuration) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataShedDuration"); pushErr != nil {
			return pushErr
		}

		// Simple Field (shedDuration)
		if pushErr := writeBuffer.PushContext("shedDuration"); pushErr != nil {
			return pushErr
		}
		_shedDurationErr := m.ShedDuration.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("shedDuration"); popErr != nil {
			return popErr
		}
		if _shedDurationErr != nil {
			return errors.Wrap(_shedDurationErr, "Error serializing 'shedDuration' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataShedDuration"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataShedDuration) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}