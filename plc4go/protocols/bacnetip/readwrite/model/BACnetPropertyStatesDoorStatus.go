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

// BACnetPropertyStatesDoorStatus is the data-structure of this message
type BACnetPropertyStatesDoorStatus struct {
	*BACnetPropertyStates
	DoorStatus *BACnetDoorStatusTagged
}

// IBACnetPropertyStatesDoorStatus is the corresponding interface of BACnetPropertyStatesDoorStatus
type IBACnetPropertyStatesDoorStatus interface {
	IBACnetPropertyStates
	// GetDoorStatus returns DoorStatus (property field)
	GetDoorStatus() *BACnetDoorStatusTagged
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

func (m *BACnetPropertyStatesDoorStatus) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesDoorStatus) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesDoorStatus) GetDoorStatus() *BACnetDoorStatusTagged {
	return m.DoorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesDoorStatus factory function for BACnetPropertyStatesDoorStatus
func NewBACnetPropertyStatesDoorStatus(doorStatus *BACnetDoorStatusTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesDoorStatus {
	_result := &BACnetPropertyStatesDoorStatus{
		DoorStatus:           doorStatus,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesDoorStatus(structType interface{}) *BACnetPropertyStatesDoorStatus {
	if casted, ok := structType.(BACnetPropertyStatesDoorStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesDoorStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesDoorStatus(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesDoorStatus) GetTypeName() string {
	return "BACnetPropertyStatesDoorStatus"
}

func (m *BACnetPropertyStatesDoorStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesDoorStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorStatus)
	lengthInBits += m.DoorStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesDoorStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesDoorStatusParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesDoorStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorStatus"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorStatus)
	if pullErr := readBuffer.PullContext("doorStatus"); pullErr != nil {
		return nil, pullErr
	}
	_doorStatus, _doorStatusErr := BACnetDoorStatusTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _doorStatusErr != nil {
		return nil, errors.Wrap(_doorStatusErr, "Error parsing 'doorStatus' field")
	}
	doorStatus := CastBACnetDoorStatusTagged(_doorStatus)
	if closeErr := readBuffer.CloseContext("doorStatus"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorStatus"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesDoorStatus{
		DoorStatus:           CastBACnetDoorStatusTagged(doorStatus),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesDoorStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorStatus"); pushErr != nil {
			return pushErr
		}

		// Simple Field (doorStatus)
		if pushErr := writeBuffer.PushContext("doorStatus"); pushErr != nil {
			return pushErr
		}
		_doorStatusErr := m.DoorStatus.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("doorStatus"); popErr != nil {
			return popErr
		}
		if _doorStatusErr != nil {
			return errors.Wrap(_doorStatusErr, "Error serializing 'doorStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorStatus"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesDoorStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
