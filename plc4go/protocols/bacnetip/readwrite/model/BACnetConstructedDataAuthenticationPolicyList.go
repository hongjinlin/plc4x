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

// BACnetConstructedDataAuthenticationPolicyList is the data-structure of this message
type BACnetConstructedDataAuthenticationPolicyList struct {
	*BACnetConstructedData
	AuthenticationPolicyList []*BACnetAuthenticationPolicy

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataAuthenticationPolicyList is the corresponding interface of BACnetConstructedDataAuthenticationPolicyList
type IBACnetConstructedDataAuthenticationPolicyList interface {
	IBACnetConstructedData
	// GetAuthenticationPolicyList returns AuthenticationPolicyList (property field)
	GetAuthenticationPolicyList() []*BACnetAuthenticationPolicy
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

func (m *BACnetConstructedDataAuthenticationPolicyList) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAuthenticationPolicyList) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHENTICATION_POLICY_LIST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAuthenticationPolicyList) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAuthenticationPolicyList) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAuthenticationPolicyList) GetAuthenticationPolicyList() []*BACnetAuthenticationPolicy {
	return m.AuthenticationPolicyList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAuthenticationPolicyList factory function for BACnetConstructedDataAuthenticationPolicyList
func NewBACnetConstructedDataAuthenticationPolicyList(authenticationPolicyList []*BACnetAuthenticationPolicy, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataAuthenticationPolicyList {
	_result := &BACnetConstructedDataAuthenticationPolicyList{
		AuthenticationPolicyList: authenticationPolicyList,
		BACnetConstructedData:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAuthenticationPolicyList(structType interface{}) *BACnetConstructedDataAuthenticationPolicyList {
	if casted, ok := structType.(BACnetConstructedDataAuthenticationPolicyList); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthenticationPolicyList); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAuthenticationPolicyList(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAuthenticationPolicyList(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAuthenticationPolicyList) GetTypeName() string {
	return "BACnetConstructedDataAuthenticationPolicyList"
}

func (m *BACnetConstructedDataAuthenticationPolicyList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAuthenticationPolicyList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.AuthenticationPolicyList) > 0 {
		for _, element := range m.AuthenticationPolicyList {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataAuthenticationPolicyList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAuthenticationPolicyListParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataAuthenticationPolicyList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthenticationPolicyList"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (authenticationPolicyList)
	if pullErr := readBuffer.PullContext("authenticationPolicyList", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	authenticationPolicyList := make([]*BACnetAuthenticationPolicy, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetAuthenticationPolicyParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'authenticationPolicyList' field")
			}
			authenticationPolicyList = append(authenticationPolicyList, CastBACnetAuthenticationPolicy(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("authenticationPolicyList", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthenticationPolicyList"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAuthenticationPolicyList{
		AuthenticationPolicyList: authenticationPolicyList,
		BACnetConstructedData:    &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAuthenticationPolicyList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthenticationPolicyList"); pushErr != nil {
			return pushErr
		}

		// Array Field (authenticationPolicyList)
		if m.AuthenticationPolicyList != nil {
			if pushErr := writeBuffer.PushContext("authenticationPolicyList", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.AuthenticationPolicyList {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'authenticationPolicyList' field")
				}
			}
			if popErr := writeBuffer.PopContext("authenticationPolicyList", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthenticationPolicyList"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAuthenticationPolicyList) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}