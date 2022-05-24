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

// BACnetPropertyStatesAccessCredentialDisableReason is the data-structure of this message
type BACnetPropertyStatesAccessCredentialDisableReason struct {
	*BACnetPropertyStates
	AccessCredentialDisableReason *BACnetAccessCredentialDisableReasonTagged
}

// IBACnetPropertyStatesAccessCredentialDisableReason is the corresponding interface of BACnetPropertyStatesAccessCredentialDisableReason
type IBACnetPropertyStatesAccessCredentialDisableReason interface {
	IBACnetPropertyStates
	// GetAccessCredentialDisableReason returns AccessCredentialDisableReason (property field)
	GetAccessCredentialDisableReason() *BACnetAccessCredentialDisableReasonTagged
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

func (m *BACnetPropertyStatesAccessCredentialDisableReason) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesAccessCredentialDisableReason) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesAccessCredentialDisableReason) GetAccessCredentialDisableReason() *BACnetAccessCredentialDisableReasonTagged {
	return m.AccessCredentialDisableReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesAccessCredentialDisableReason factory function for BACnetPropertyStatesAccessCredentialDisableReason
func NewBACnetPropertyStatesAccessCredentialDisableReason(accessCredentialDisableReason *BACnetAccessCredentialDisableReasonTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesAccessCredentialDisableReason {
	_result := &BACnetPropertyStatesAccessCredentialDisableReason{
		AccessCredentialDisableReason: accessCredentialDisableReason,
		BACnetPropertyStates:          NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesAccessCredentialDisableReason(structType interface{}) *BACnetPropertyStatesAccessCredentialDisableReason {
	if casted, ok := structType.(BACnetPropertyStatesAccessCredentialDisableReason); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAccessCredentialDisableReason); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesAccessCredentialDisableReason(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesAccessCredentialDisableReason(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesAccessCredentialDisableReason) GetTypeName() string {
	return "BACnetPropertyStatesAccessCredentialDisableReason"
}

func (m *BACnetPropertyStatesAccessCredentialDisableReason) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesAccessCredentialDisableReason) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accessCredentialDisableReason)
	lengthInBits += m.AccessCredentialDisableReason.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesAccessCredentialDisableReason) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesAccessCredentialDisableReasonParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesAccessCredentialDisableReason, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAccessCredentialDisableReason"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessCredentialDisableReason)
	if pullErr := readBuffer.PullContext("accessCredentialDisableReason"); pullErr != nil {
		return nil, pullErr
	}
	_accessCredentialDisableReason, _accessCredentialDisableReasonErr := BACnetAccessCredentialDisableReasonTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _accessCredentialDisableReasonErr != nil {
		return nil, errors.Wrap(_accessCredentialDisableReasonErr, "Error parsing 'accessCredentialDisableReason' field")
	}
	accessCredentialDisableReason := CastBACnetAccessCredentialDisableReasonTagged(_accessCredentialDisableReason)
	if closeErr := readBuffer.CloseContext("accessCredentialDisableReason"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAccessCredentialDisableReason"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesAccessCredentialDisableReason{
		AccessCredentialDisableReason: CastBACnetAccessCredentialDisableReasonTagged(accessCredentialDisableReason),
		BACnetPropertyStates:          &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesAccessCredentialDisableReason) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAccessCredentialDisableReason"); pushErr != nil {
			return pushErr
		}

		// Simple Field (accessCredentialDisableReason)
		if pushErr := writeBuffer.PushContext("accessCredentialDisableReason"); pushErr != nil {
			return pushErr
		}
		_accessCredentialDisableReasonErr := m.AccessCredentialDisableReason.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("accessCredentialDisableReason"); popErr != nil {
			return popErr
		}
		if _accessCredentialDisableReasonErr != nil {
			return errors.Wrap(_accessCredentialDisableReasonErr, "Error serializing 'accessCredentialDisableReason' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAccessCredentialDisableReason"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesAccessCredentialDisableReason) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
