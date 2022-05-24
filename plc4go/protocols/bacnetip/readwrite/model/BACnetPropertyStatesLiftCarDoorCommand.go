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

// BACnetPropertyStatesLiftCarDoorCommand is the data-structure of this message
type BACnetPropertyStatesLiftCarDoorCommand struct {
	*BACnetPropertyStates
	LiftCarDoorCommand *BACnetLiftCarDoorCommandTagged
}

// IBACnetPropertyStatesLiftCarDoorCommand is the corresponding interface of BACnetPropertyStatesLiftCarDoorCommand
type IBACnetPropertyStatesLiftCarDoorCommand interface {
	IBACnetPropertyStates
	// GetLiftCarDoorCommand returns LiftCarDoorCommand (property field)
	GetLiftCarDoorCommand() *BACnetLiftCarDoorCommandTagged
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

func (m *BACnetPropertyStatesLiftCarDoorCommand) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesLiftCarDoorCommand) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesLiftCarDoorCommand) GetLiftCarDoorCommand() *BACnetLiftCarDoorCommandTagged {
	return m.LiftCarDoorCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLiftCarDoorCommand factory function for BACnetPropertyStatesLiftCarDoorCommand
func NewBACnetPropertyStatesLiftCarDoorCommand(liftCarDoorCommand *BACnetLiftCarDoorCommandTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesLiftCarDoorCommand {
	_result := &BACnetPropertyStatesLiftCarDoorCommand{
		LiftCarDoorCommand:   liftCarDoorCommand,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesLiftCarDoorCommand(structType interface{}) *BACnetPropertyStatesLiftCarDoorCommand {
	if casted, ok := structType.(BACnetPropertyStatesLiftCarDoorCommand); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftCarDoorCommand); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLiftCarDoorCommand(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLiftCarDoorCommand(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesLiftCarDoorCommand) GetTypeName() string {
	return "BACnetPropertyStatesLiftCarDoorCommand"
}

func (m *BACnetPropertyStatesLiftCarDoorCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesLiftCarDoorCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (liftCarDoorCommand)
	lengthInBits += m.LiftCarDoorCommand.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesLiftCarDoorCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLiftCarDoorCommandParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesLiftCarDoorCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftCarDoorCommand"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (liftCarDoorCommand)
	if pullErr := readBuffer.PullContext("liftCarDoorCommand"); pullErr != nil {
		return nil, pullErr
	}
	_liftCarDoorCommand, _liftCarDoorCommandErr := BACnetLiftCarDoorCommandTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _liftCarDoorCommandErr != nil {
		return nil, errors.Wrap(_liftCarDoorCommandErr, "Error parsing 'liftCarDoorCommand' field")
	}
	liftCarDoorCommand := CastBACnetLiftCarDoorCommandTagged(_liftCarDoorCommand)
	if closeErr := readBuffer.CloseContext("liftCarDoorCommand"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftCarDoorCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesLiftCarDoorCommand{
		LiftCarDoorCommand:   CastBACnetLiftCarDoorCommandTagged(liftCarDoorCommand),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesLiftCarDoorCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftCarDoorCommand"); pushErr != nil {
			return pushErr
		}

		// Simple Field (liftCarDoorCommand)
		if pushErr := writeBuffer.PushContext("liftCarDoorCommand"); pushErr != nil {
			return pushErr
		}
		_liftCarDoorCommandErr := m.LiftCarDoorCommand.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("liftCarDoorCommand"); popErr != nil {
			return popErr
		}
		if _liftCarDoorCommandErr != nil {
			return errors.Wrap(_liftCarDoorCommandErr, "Error serializing 'liftCarDoorCommand' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftCarDoorCommand"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesLiftCarDoorCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
