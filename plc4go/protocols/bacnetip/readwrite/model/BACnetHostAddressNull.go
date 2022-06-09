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

// BACnetHostAddressNull is the data-structure of this message
type BACnetHostAddressNull struct {
	*BACnetHostAddress
	None *BACnetContextTagNull
}

// IBACnetHostAddressNull is the corresponding interface of BACnetHostAddressNull
type IBACnetHostAddressNull interface {
	IBACnetHostAddress
	// GetNone returns None (property field)
	GetNone() *BACnetContextTagNull
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

func (m *BACnetHostAddressNull) InitializeParent(parent *BACnetHostAddress, peekedTagHeader *BACnetTagHeader) {
	m.BACnetHostAddress.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetHostAddressNull) GetParent() *BACnetHostAddress {
	return m.BACnetHostAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetHostAddressNull) GetNone() *BACnetContextTagNull {
	return m.None
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostAddressNull factory function for BACnetHostAddressNull
func NewBACnetHostAddressNull(none *BACnetContextTagNull, peekedTagHeader *BACnetTagHeader) *BACnetHostAddressNull {
	_result := &BACnetHostAddressNull{
		None:              none,
		BACnetHostAddress: NewBACnetHostAddress(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetHostAddressNull(structType interface{}) *BACnetHostAddressNull {
	if casted, ok := structType.(BACnetHostAddressNull); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetHostAddressNull); ok {
		return casted
	}
	if casted, ok := structType.(BACnetHostAddress); ok {
		return CastBACnetHostAddressNull(casted.Child)
	}
	if casted, ok := structType.(*BACnetHostAddress); ok {
		return CastBACnetHostAddressNull(casted.Child)
	}
	return nil
}

func (m *BACnetHostAddressNull) GetTypeName() string {
	return "BACnetHostAddressNull"
}

func (m *BACnetHostAddressNull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetHostAddressNull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (none)
	lengthInBits += m.None.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetHostAddressNull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetHostAddressNullParse(readBuffer utils.ReadBuffer) (*BACnetHostAddressNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressNull"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (none)
	if pullErr := readBuffer.PullContext("none"); pullErr != nil {
		return nil, pullErr
	}
	_none, _noneErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_NULL))
	if _noneErr != nil {
		return nil, errors.Wrap(_noneErr, "Error parsing 'none' field")
	}
	none := CastBACnetContextTagNull(_none)
	if closeErr := readBuffer.CloseContext("none"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetHostAddressNull"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetHostAddressNull{
		None:              CastBACnetContextTagNull(none),
		BACnetHostAddress: &BACnetHostAddress{},
	}
	_child.BACnetHostAddress.Child = _child
	return _child, nil
}

func (m *BACnetHostAddressNull) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetHostAddressNull"); pushErr != nil {
			return pushErr
		}

		// Simple Field (none)
		if pushErr := writeBuffer.PushContext("none"); pushErr != nil {
			return pushErr
		}
		_noneErr := m.None.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("none"); popErr != nil {
			return popErr
		}
		if _noneErr != nil {
			return errors.Wrap(_noneErr, "Error serializing 'none' field")
		}

		if popErr := writeBuffer.PopContext("BACnetHostAddressNull"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetHostAddressNull) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}