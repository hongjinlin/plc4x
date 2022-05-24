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

// BACnetPropertyStatesBinaryLightningValue is the data-structure of this message
type BACnetPropertyStatesBinaryLightningValue struct {
	*BACnetPropertyStates
	BinaryLightningValue *BACnetBinaryLightingPVTagged
}

// IBACnetPropertyStatesBinaryLightningValue is the corresponding interface of BACnetPropertyStatesBinaryLightningValue
type IBACnetPropertyStatesBinaryLightningValue interface {
	IBACnetPropertyStates
	// GetBinaryLightningValue returns BinaryLightningValue (property field)
	GetBinaryLightningValue() *BACnetBinaryLightingPVTagged
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPropertyStatesBinaryLightningValue) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesBinaryLightningValue) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesBinaryLightningValue) GetBinaryLightningValue() *BACnetBinaryLightingPVTagged {
	return m.BinaryLightningValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesBinaryLightningValue factory function for BACnetPropertyStatesBinaryLightningValue
func NewBACnetPropertyStatesBinaryLightningValue(binaryLightningValue *BACnetBinaryLightingPVTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesBinaryLightningValue {
	_result := &BACnetPropertyStatesBinaryLightningValue{
		BinaryLightningValue: binaryLightningValue,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesBinaryLightningValue(structType interface{}) *BACnetPropertyStatesBinaryLightningValue {
	if casted, ok := structType.(BACnetPropertyStatesBinaryLightningValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBinaryLightningValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBinaryLightningValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBinaryLightningValue(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesBinaryLightningValue) GetTypeName() string {
	return "BACnetPropertyStatesBinaryLightningValue"
}

func (m *BACnetPropertyStatesBinaryLightningValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesBinaryLightningValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (binaryLightningValue)
	lengthInBits += m.BinaryLightningValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesBinaryLightningValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesBinaryLightningValueParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesBinaryLightningValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBinaryLightningValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (binaryLightningValue)
	if pullErr := readBuffer.PullContext("binaryLightningValue"); pullErr != nil {
		return nil, pullErr
	}
	_binaryLightningValue, _binaryLightningValueErr := BACnetBinaryLightingPVTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _binaryLightningValueErr != nil {
		return nil, errors.Wrap(_binaryLightningValueErr, "Error parsing 'binaryLightningValue' field")
	}
	binaryLightningValue := CastBACnetBinaryLightingPVTagged(_binaryLightningValue)
	if closeErr := readBuffer.CloseContext("binaryLightningValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBinaryLightningValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesBinaryLightningValue{
		BinaryLightningValue: CastBACnetBinaryLightingPVTagged(binaryLightningValue),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesBinaryLightningValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBinaryLightningValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (binaryLightningValue)
		if pushErr := writeBuffer.PushContext("binaryLightningValue"); pushErr != nil {
			return pushErr
		}
		_binaryLightningValueErr := m.BinaryLightningValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("binaryLightningValue"); popErr != nil {
			return popErr
		}
		if _binaryLightningValueErr != nil {
			return errors.Wrap(_binaryLightningValueErr, "Error serializing 'binaryLightningValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBinaryLightningValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesBinaryLightningValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
