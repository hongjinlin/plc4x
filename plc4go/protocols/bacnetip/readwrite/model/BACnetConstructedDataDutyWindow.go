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

// BACnetConstructedDataDutyWindow is the data-structure of this message
type BACnetConstructedDataDutyWindow struct {
	*BACnetConstructedData
	DutyWindow *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataDutyWindow is the corresponding interface of BACnetConstructedDataDutyWindow
type IBACnetConstructedDataDutyWindow interface {
	IBACnetConstructedData
	// GetDutyWindow returns DutyWindow (property field)
	GetDutyWindow() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataDutyWindow) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDutyWindow) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DUTY_WINDOW
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDutyWindow) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDutyWindow) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDutyWindow) GetDutyWindow() *BACnetApplicationTagUnsignedInteger {
	return m.DutyWindow
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDutyWindow factory function for BACnetConstructedDataDutyWindow
func NewBACnetConstructedDataDutyWindow(dutyWindow *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataDutyWindow {
	_result := &BACnetConstructedDataDutyWindow{
		DutyWindow:            dutyWindow,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDutyWindow(structType interface{}) *BACnetConstructedDataDutyWindow {
	if casted, ok := structType.(BACnetConstructedDataDutyWindow); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDutyWindow); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDutyWindow(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDutyWindow(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDutyWindow) GetTypeName() string {
	return "BACnetConstructedDataDutyWindow"
}

func (m *BACnetConstructedDataDutyWindow) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDutyWindow) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dutyWindow)
	lengthInBits += m.DutyWindow.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDutyWindow) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDutyWindowParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataDutyWindow, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDutyWindow"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dutyWindow)
	if pullErr := readBuffer.PullContext("dutyWindow"); pullErr != nil {
		return nil, pullErr
	}
	_dutyWindow, _dutyWindowErr := BACnetApplicationTagParse(readBuffer)
	if _dutyWindowErr != nil {
		return nil, errors.Wrap(_dutyWindowErr, "Error parsing 'dutyWindow' field")
	}
	dutyWindow := CastBACnetApplicationTagUnsignedInteger(_dutyWindow)
	if closeErr := readBuffer.CloseContext("dutyWindow"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDutyWindow"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDutyWindow{
		DutyWindow:            CastBACnetApplicationTagUnsignedInteger(dutyWindow),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDutyWindow) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDutyWindow"); pushErr != nil {
			return pushErr
		}

		// Simple Field (dutyWindow)
		if pushErr := writeBuffer.PushContext("dutyWindow"); pushErr != nil {
			return pushErr
		}
		_dutyWindowErr := m.DutyWindow.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dutyWindow"); popErr != nil {
			return popErr
		}
		if _dutyWindowErr != nil {
			return errors.Wrap(_dutyWindowErr, "Error serializing 'dutyWindow' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDutyWindow"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDutyWindow) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}