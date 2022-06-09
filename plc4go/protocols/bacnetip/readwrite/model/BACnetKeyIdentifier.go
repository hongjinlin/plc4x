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

// BACnetKeyIdentifier is the data-structure of this message
type BACnetKeyIdentifier struct {
	Algorithm *BACnetContextTagUnsignedInteger
	KeyId     *BACnetContextTagUnsignedInteger
}

// IBACnetKeyIdentifier is the corresponding interface of BACnetKeyIdentifier
type IBACnetKeyIdentifier interface {
	// GetAlgorithm returns Algorithm (property field)
	GetAlgorithm() *BACnetContextTagUnsignedInteger
	// GetKeyId returns KeyId (property field)
	GetKeyId() *BACnetContextTagUnsignedInteger
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

func (m *BACnetKeyIdentifier) GetAlgorithm() *BACnetContextTagUnsignedInteger {
	return m.Algorithm
}

func (m *BACnetKeyIdentifier) GetKeyId() *BACnetContextTagUnsignedInteger {
	return m.KeyId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetKeyIdentifier factory function for BACnetKeyIdentifier
func NewBACnetKeyIdentifier(algorithm *BACnetContextTagUnsignedInteger, keyId *BACnetContextTagUnsignedInteger) *BACnetKeyIdentifier {
	return &BACnetKeyIdentifier{Algorithm: algorithm, KeyId: keyId}
}

func CastBACnetKeyIdentifier(structType interface{}) *BACnetKeyIdentifier {
	if casted, ok := structType.(BACnetKeyIdentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetKeyIdentifier); ok {
		return casted
	}
	return nil
}

func (m *BACnetKeyIdentifier) GetTypeName() string {
	return "BACnetKeyIdentifier"
}

func (m *BACnetKeyIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetKeyIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (algorithm)
	lengthInBits += m.Algorithm.GetLengthInBits()

	// Simple field (keyId)
	lengthInBits += m.KeyId.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetKeyIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetKeyIdentifierParse(readBuffer utils.ReadBuffer) (*BACnetKeyIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetKeyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (algorithm)
	if pullErr := readBuffer.PullContext("algorithm"); pullErr != nil {
		return nil, pullErr
	}
	_algorithm, _algorithmErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _algorithmErr != nil {
		return nil, errors.Wrap(_algorithmErr, "Error parsing 'algorithm' field")
	}
	algorithm := CastBACnetContextTagUnsignedInteger(_algorithm)
	if closeErr := readBuffer.CloseContext("algorithm"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (keyId)
	if pullErr := readBuffer.PullContext("keyId"); pullErr != nil {
		return nil, pullErr
	}
	_keyId, _keyIdErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _keyIdErr != nil {
		return nil, errors.Wrap(_keyIdErr, "Error parsing 'keyId' field")
	}
	keyId := CastBACnetContextTagUnsignedInteger(_keyId)
	if closeErr := readBuffer.CloseContext("keyId"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetKeyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetKeyIdentifier(algorithm, keyId), nil
}

func (m *BACnetKeyIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetKeyIdentifier"); pushErr != nil {
		return pushErr
	}

	// Simple Field (algorithm)
	if pushErr := writeBuffer.PushContext("algorithm"); pushErr != nil {
		return pushErr
	}
	_algorithmErr := m.Algorithm.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("algorithm"); popErr != nil {
		return popErr
	}
	if _algorithmErr != nil {
		return errors.Wrap(_algorithmErr, "Error serializing 'algorithm' field")
	}

	// Simple Field (keyId)
	if pushErr := writeBuffer.PushContext("keyId"); pushErr != nil {
		return pushErr
	}
	_keyIdErr := m.KeyId.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("keyId"); popErr != nil {
		return popErr
	}
	if _keyIdErr != nil {
		return errors.Wrap(_keyIdErr, "Error serializing 'keyId' field")
	}

	if popErr := writeBuffer.PopContext("BACnetKeyIdentifier"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetKeyIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}