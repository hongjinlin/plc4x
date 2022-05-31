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

// BACnetRecipientEnclosed is the data-structure of this message
type BACnetRecipientEnclosed struct {
	OpeningTag *BACnetOpeningTag
	Recipient  *BACnetRecipient
	ClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetRecipientEnclosed is the corresponding interface of BACnetRecipientEnclosed
type IBACnetRecipientEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetRecipient returns Recipient (property field)
	GetRecipient() *BACnetRecipient
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetRecipientEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetRecipientEnclosed) GetRecipient() *BACnetRecipient {
	return m.Recipient
}

func (m *BACnetRecipientEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRecipientEnclosed factory function for BACnetRecipientEnclosed
func NewBACnetRecipientEnclosed(openingTag *BACnetOpeningTag, recipient *BACnetRecipient, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetRecipientEnclosed {
	return &BACnetRecipientEnclosed{OpeningTag: openingTag, Recipient: recipient, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetRecipientEnclosed(structType interface{}) *BACnetRecipientEnclosed {
	if casted, ok := structType.(BACnetRecipientEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetRecipientEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetRecipientEnclosed) GetTypeName() string {
	return "BACnetRecipientEnclosed"
}

func (m *BACnetRecipientEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetRecipientEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetRecipientEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRecipientEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetRecipientEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientEnclosed"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (recipient)
	if pullErr := readBuffer.PullContext("recipient"); pullErr != nil {
		return nil, pullErr
	}
	_recipient, _recipientErr := BACnetRecipientParse(readBuffer)
	if _recipientErr != nil {
		return nil, errors.Wrap(_recipientErr, "Error parsing 'recipient' field")
	}
	recipient := CastBACnetRecipient(_recipient)
	if closeErr := readBuffer.CloseContext("recipient"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipientEnclosed"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetRecipientEnclosed(openingTag, recipient, closingTag, tagNumber), nil
}

func (m *BACnetRecipientEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetRecipientEnclosed"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (recipient)
	if pushErr := writeBuffer.PushContext("recipient"); pushErr != nil {
		return pushErr
	}
	_recipientErr := m.Recipient.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("recipient"); popErr != nil {
		return popErr
	}
	if _recipientErr != nil {
		return errors.Wrap(_recipientErr, "Error serializing 'recipient' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRecipientEnclosed"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetRecipientEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}