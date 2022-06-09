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

// BACnetConstructedDataDefaultSubordinateRelationship is the data-structure of this message
type BACnetConstructedDataDefaultSubordinateRelationship struct {
	*BACnetConstructedData
	DefaultSubordinateRelationship *BACnetRelationshipTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataDefaultSubordinateRelationship is the corresponding interface of BACnetConstructedDataDefaultSubordinateRelationship
type IBACnetConstructedDataDefaultSubordinateRelationship interface {
	IBACnetConstructedData
	// GetDefaultSubordinateRelationship returns DefaultSubordinateRelationship (property field)
	GetDefaultSubordinateRelationship() *BACnetRelationshipTagged
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

func (m *BACnetConstructedDataDefaultSubordinateRelationship) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDefaultSubordinateRelationship) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEFAULT_SUBORDINATE_RELATIONSHIP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDefaultSubordinateRelationship) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDefaultSubordinateRelationship) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDefaultSubordinateRelationship) GetDefaultSubordinateRelationship() *BACnetRelationshipTagged {
	return m.DefaultSubordinateRelationship
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDefaultSubordinateRelationship factory function for BACnetConstructedDataDefaultSubordinateRelationship
func NewBACnetConstructedDataDefaultSubordinateRelationship(defaultSubordinateRelationship *BACnetRelationshipTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataDefaultSubordinateRelationship {
	_result := &BACnetConstructedDataDefaultSubordinateRelationship{
		DefaultSubordinateRelationship: defaultSubordinateRelationship,
		BACnetConstructedData:          NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDefaultSubordinateRelationship(structType interface{}) *BACnetConstructedDataDefaultSubordinateRelationship {
	if casted, ok := structType.(BACnetConstructedDataDefaultSubordinateRelationship); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDefaultSubordinateRelationship); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDefaultSubordinateRelationship(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDefaultSubordinateRelationship(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDefaultSubordinateRelationship) GetTypeName() string {
	return "BACnetConstructedDataDefaultSubordinateRelationship"
}

func (m *BACnetConstructedDataDefaultSubordinateRelationship) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDefaultSubordinateRelationship) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (defaultSubordinateRelationship)
	lengthInBits += m.DefaultSubordinateRelationship.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDefaultSubordinateRelationship) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDefaultSubordinateRelationshipParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataDefaultSubordinateRelationship, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDefaultSubordinateRelationship"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (defaultSubordinateRelationship)
	if pullErr := readBuffer.PullContext("defaultSubordinateRelationship"); pullErr != nil {
		return nil, pullErr
	}
	_defaultSubordinateRelationship, _defaultSubordinateRelationshipErr := BACnetRelationshipTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _defaultSubordinateRelationshipErr != nil {
		return nil, errors.Wrap(_defaultSubordinateRelationshipErr, "Error parsing 'defaultSubordinateRelationship' field")
	}
	defaultSubordinateRelationship := CastBACnetRelationshipTagged(_defaultSubordinateRelationship)
	if closeErr := readBuffer.CloseContext("defaultSubordinateRelationship"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDefaultSubordinateRelationship"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDefaultSubordinateRelationship{
		DefaultSubordinateRelationship: CastBACnetRelationshipTagged(defaultSubordinateRelationship),
		BACnetConstructedData:          &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDefaultSubordinateRelationship) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDefaultSubordinateRelationship"); pushErr != nil {
			return pushErr
		}

		// Simple Field (defaultSubordinateRelationship)
		if pushErr := writeBuffer.PushContext("defaultSubordinateRelationship"); pushErr != nil {
			return pushErr
		}
		_defaultSubordinateRelationshipErr := m.DefaultSubordinateRelationship.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("defaultSubordinateRelationship"); popErr != nil {
			return popErr
		}
		if _defaultSubordinateRelationshipErr != nil {
			return errors.Wrap(_defaultSubordinateRelationshipErr, "Error serializing 'defaultSubordinateRelationship' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDefaultSubordinateRelationship"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDefaultSubordinateRelationship) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}