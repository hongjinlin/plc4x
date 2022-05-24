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

// BACnetPropertyStatesPolarity is the data-structure of this message
type BACnetPropertyStatesPolarity struct {
	*BACnetPropertyStates
	Polarity *BACnetPolarityTagged
}

// IBACnetPropertyStatesPolarity is the corresponding interface of BACnetPropertyStatesPolarity
type IBACnetPropertyStatesPolarity interface {
	IBACnetPropertyStates
	// GetPolarity returns Polarity (property field)
	GetPolarity() *BACnetPolarityTagged
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

func (m *BACnetPropertyStatesPolarity) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesPolarity) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesPolarity) GetPolarity() *BACnetPolarityTagged {
	return m.Polarity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesPolarity factory function for BACnetPropertyStatesPolarity
func NewBACnetPropertyStatesPolarity(polarity *BACnetPolarityTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesPolarity {
	_result := &BACnetPropertyStatesPolarity{
		Polarity:             polarity,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesPolarity(structType interface{}) *BACnetPropertyStatesPolarity {
	if casted, ok := structType.(BACnetPropertyStatesPolarity); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesPolarity); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesPolarity(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesPolarity(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesPolarity) GetTypeName() string {
	return "BACnetPropertyStatesPolarity"
}

func (m *BACnetPropertyStatesPolarity) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesPolarity) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (polarity)
	lengthInBits += m.Polarity.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesPolarity) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesPolarityParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesPolarity, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesPolarity"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (polarity)
	if pullErr := readBuffer.PullContext("polarity"); pullErr != nil {
		return nil, pullErr
	}
	_polarity, _polarityErr := BACnetPolarityTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _polarityErr != nil {
		return nil, errors.Wrap(_polarityErr, "Error parsing 'polarity' field")
	}
	polarity := CastBACnetPolarityTagged(_polarity)
	if closeErr := readBuffer.CloseContext("polarity"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesPolarity"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesPolarity{
		Polarity:             CastBACnetPolarityTagged(polarity),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesPolarity) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesPolarity"); pushErr != nil {
			return pushErr
		}

		// Simple Field (polarity)
		if pushErr := writeBuffer.PushContext("polarity"); pushErr != nil {
			return pushErr
		}
		_polarityErr := m.Polarity.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("polarity"); popErr != nil {
			return popErr
		}
		if _polarityErr != nil {
			return errors.Wrap(_polarityErr, "Error serializing 'polarity' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesPolarity"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesPolarity) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
