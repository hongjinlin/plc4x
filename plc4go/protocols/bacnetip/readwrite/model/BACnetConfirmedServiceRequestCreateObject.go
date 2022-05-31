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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestCreateObject is the data-structure of this message
type BACnetConfirmedServiceRequestCreateObject struct {
	*BACnetConfirmedServiceRequest
	ObjectSpecifier *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
	ListOfValues    *BACnetPropertyValues

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestCreateObject is the corresponding interface of BACnetConfirmedServiceRequestCreateObject
type IBACnetConfirmedServiceRequestCreateObject interface {
	IBACnetConfirmedServiceRequest
	// GetObjectSpecifier returns ObjectSpecifier (property field)
	GetObjectSpecifier() *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() *BACnetPropertyValues
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

func (m *BACnetConfirmedServiceRequestCreateObject) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CREATE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestCreateObject) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestCreateObject) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestCreateObject) GetObjectSpecifier() *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier {
	return m.ObjectSpecifier
}

func (m *BACnetConfirmedServiceRequestCreateObject) GetListOfValues() *BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestCreateObject factory function for BACnetConfirmedServiceRequestCreateObject
func NewBACnetConfirmedServiceRequestCreateObject(objectSpecifier *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, listOfValues *BACnetPropertyValues, serviceRequestLength uint16) *BACnetConfirmedServiceRequestCreateObject {
	_result := &BACnetConfirmedServiceRequestCreateObject{
		ObjectSpecifier:               objectSpecifier,
		ListOfValues:                  listOfValues,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestCreateObject(structType interface{}) *BACnetConfirmedServiceRequestCreateObject {
	if casted, ok := structType.(BACnetConfirmedServiceRequestCreateObject); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestCreateObject); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestCreateObject(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestCreateObject(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestCreateObject) GetTypeName() string {
	return "BACnetConfirmedServiceRequestCreateObject"
}

func (m *BACnetConfirmedServiceRequestCreateObject) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestCreateObject) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectSpecifier)
	lengthInBits += m.ObjectSpecifier.GetLengthInBits()

	// Optional Field (listOfValues)
	if m.ListOfValues != nil {
		lengthInBits += (*m.ListOfValues).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestCreateObject) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestCreateObjectParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestCreateObject, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestCreateObject"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectSpecifier)
	if pullErr := readBuffer.PullContext("objectSpecifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectSpecifier, _objectSpecifierErr := BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParse(readBuffer, uint8(uint8(0)))
	if _objectSpecifierErr != nil {
		return nil, errors.Wrap(_objectSpecifierErr, "Error parsing 'objectSpecifier' field")
	}
	objectSpecifier := CastBACnetConfirmedServiceRequestCreateObjectObjectSpecifier(_objectSpecifier)
	if closeErr := readBuffer.CloseContext("objectSpecifier"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (listOfValues) (Can be skipped, if a given expression evaluates to false)
	var listOfValues *BACnetPropertyValues = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetPropertyValuesParse(readBuffer, uint8(1), CastBACnetObjectType(utils.InlineIf(objectSpecifier.GetIsObjectType(), func() interface{} { return CastBACnetObjectType(objectSpecifier.GetObjectType()) }, func() interface{} { return CastBACnetObjectType(objectSpecifier.GetObjectIdentifier().GetObjectType()) })))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'listOfValues' field")
		default:
			listOfValues = CastBACnetPropertyValues(_val)
			if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestCreateObject"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestCreateObject{
		ObjectSpecifier:               CastBACnetConfirmedServiceRequestCreateObjectObjectSpecifier(objectSpecifier),
		ListOfValues:                  CastBACnetPropertyValues(listOfValues),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestCreateObject) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestCreateObject"); pushErr != nil {
			return pushErr
		}

		// Simple Field (objectSpecifier)
		if pushErr := writeBuffer.PushContext("objectSpecifier"); pushErr != nil {
			return pushErr
		}
		_objectSpecifierErr := m.ObjectSpecifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("objectSpecifier"); popErr != nil {
			return popErr
		}
		if _objectSpecifierErr != nil {
			return errors.Wrap(_objectSpecifierErr, "Error serializing 'objectSpecifier' field")
		}

		// Optional Field (listOfValues) (Can be skipped, if the value is null)
		var listOfValues *BACnetPropertyValues = nil
		if m.ListOfValues != nil {
			if pushErr := writeBuffer.PushContext("listOfValues"); pushErr != nil {
				return pushErr
			}
			listOfValues = m.ListOfValues
			_listOfValuesErr := listOfValues.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("listOfValues"); popErr != nil {
				return popErr
			}
			if _listOfValuesErr != nil {
				return errors.Wrap(_listOfValuesErr, "Error serializing 'listOfValues' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestCreateObject"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestCreateObject) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}