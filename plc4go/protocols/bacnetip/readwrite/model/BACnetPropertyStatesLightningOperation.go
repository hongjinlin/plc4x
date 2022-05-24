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

// BACnetPropertyStatesLightningOperation is the data-structure of this message
type BACnetPropertyStatesLightningOperation struct {
	*BACnetPropertyStates
	LightningOperation *BACnetLightingOperationTagged
}

// IBACnetPropertyStatesLightningOperation is the corresponding interface of BACnetPropertyStatesLightningOperation
type IBACnetPropertyStatesLightningOperation interface {
	IBACnetPropertyStates
	// GetLightningOperation returns LightningOperation (property field)
	GetLightningOperation() *BACnetLightingOperationTagged
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

func (m *BACnetPropertyStatesLightningOperation) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesLightningOperation) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesLightningOperation) GetLightningOperation() *BACnetLightingOperationTagged {
	return m.LightningOperation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLightningOperation factory function for BACnetPropertyStatesLightningOperation
func NewBACnetPropertyStatesLightningOperation(lightningOperation *BACnetLightingOperationTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesLightningOperation {
	_result := &BACnetPropertyStatesLightningOperation{
		LightningOperation:   lightningOperation,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesLightningOperation(structType interface{}) *BACnetPropertyStatesLightningOperation {
	if casted, ok := structType.(BACnetPropertyStatesLightningOperation); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLightningOperation); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLightningOperation(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLightningOperation(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesLightningOperation) GetTypeName() string {
	return "BACnetPropertyStatesLightningOperation"
}

func (m *BACnetPropertyStatesLightningOperation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesLightningOperation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lightningOperation)
	lengthInBits += m.LightningOperation.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesLightningOperation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLightningOperationParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesLightningOperation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLightningOperation"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightningOperation)
	if pullErr := readBuffer.PullContext("lightningOperation"); pullErr != nil {
		return nil, pullErr
	}
	_lightningOperation, _lightningOperationErr := BACnetLightingOperationTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _lightningOperationErr != nil {
		return nil, errors.Wrap(_lightningOperationErr, "Error parsing 'lightningOperation' field")
	}
	lightningOperation := CastBACnetLightingOperationTagged(_lightningOperation)
	if closeErr := readBuffer.CloseContext("lightningOperation"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLightningOperation"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesLightningOperation{
		LightningOperation:   CastBACnetLightingOperationTagged(lightningOperation),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesLightningOperation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLightningOperation"); pushErr != nil {
			return pushErr
		}

		// Simple Field (lightningOperation)
		if pushErr := writeBuffer.PushContext("lightningOperation"); pushErr != nil {
			return pushErr
		}
		_lightningOperationErr := m.LightningOperation.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lightningOperation"); popErr != nil {
			return popErr
		}
		if _lightningOperationErr != nil {
			return errors.Wrap(_lightningOperationErr, "Error serializing 'lightningOperation' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLightningOperation"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesLightningOperation) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
