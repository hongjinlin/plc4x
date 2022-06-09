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

// BACnetConstructedDataTrigger is the data-structure of this message
type BACnetConstructedDataTrigger struct {
	*BACnetConstructedData
	Trigger *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataTrigger is the corresponding interface of BACnetConstructedDataTrigger
type IBACnetConstructedDataTrigger interface {
	IBACnetConstructedData
	// GetTrigger returns Trigger (property field)
	GetTrigger() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataTrigger) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataTrigger) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRIGGER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTrigger) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTrigger) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTrigger) GetTrigger() *BACnetApplicationTagBoolean {
	return m.Trigger
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTrigger factory function for BACnetConstructedDataTrigger
func NewBACnetConstructedDataTrigger(trigger *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataTrigger {
	_result := &BACnetConstructedDataTrigger{
		Trigger:               trigger,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTrigger(structType interface{}) *BACnetConstructedDataTrigger {
	if casted, ok := structType.(BACnetConstructedDataTrigger); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrigger); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTrigger(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTrigger(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTrigger) GetTypeName() string {
	return "BACnetConstructedDataTrigger"
}

func (m *BACnetConstructedDataTrigger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTrigger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (trigger)
	lengthInBits += m.Trigger.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTrigger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTriggerParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataTrigger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrigger"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (trigger)
	if pullErr := readBuffer.PullContext("trigger"); pullErr != nil {
		return nil, pullErr
	}
	_trigger, _triggerErr := BACnetApplicationTagParse(readBuffer)
	if _triggerErr != nil {
		return nil, errors.Wrap(_triggerErr, "Error parsing 'trigger' field")
	}
	trigger := CastBACnetApplicationTagBoolean(_trigger)
	if closeErr := readBuffer.CloseContext("trigger"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrigger"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTrigger{
		Trigger:               CastBACnetApplicationTagBoolean(trigger),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTrigger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrigger"); pushErr != nil {
			return pushErr
		}

		// Simple Field (trigger)
		if pushErr := writeBuffer.PushContext("trigger"); pushErr != nil {
			return pushErr
		}
		_triggerErr := m.Trigger.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("trigger"); popErr != nil {
			return popErr
		}
		if _triggerErr != nil {
			return errors.Wrap(_triggerErr, "Error serializing 'trigger' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrigger"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTrigger) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}