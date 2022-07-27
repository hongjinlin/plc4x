/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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

// SALDataReserved is the corresponding interface of SALDataReserved
type SALDataReserved interface {
	utils.LengthAware
	utils.Serializable
	SALData
}

// SALDataReservedExactly can be used when we want exactly this type and not a type which fulfills SALDataReserved.
// This is useful for switch cases.
type SALDataReservedExactly interface {
	SALDataReserved
	isSALDataReserved() bool
}

// _SALDataReserved is the data-structure of this message
type _SALDataReserved struct {
	*_SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataReserved) GetApplicationId() ApplicationId {
	return ApplicationId_RESERVED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataReserved) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataReserved) GetParent() SALData {
	return m._SALData
}

// NewSALDataReserved factory function for _SALDataReserved
func NewSALDataReserved(salData SALData) *_SALDataReserved {
	_result := &_SALDataReserved{
		_SALData: NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataReserved(structType interface{}) SALDataReserved {
	if casted, ok := structType.(SALDataReserved); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataReserved); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataReserved) GetTypeName() string {
	return "SALDataReserved"
}

func (m *_SALDataReserved) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataReserved) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SALDataReserved) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataReservedParse(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataReserved, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataReserved"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataReserved")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{"RESERVED Not yet implemented"})
	}

	if closeErr := readBuffer.CloseContext("SALDataReserved"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataReserved")
	}

	// Create a partially initialized instance
	_child := &_SALDataReserved{
		_SALData: &_SALData{},
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataReserved) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataReserved"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataReserved")
		}

		if popErr := writeBuffer.PopContext("SALDataReserved"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataReserved")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataReserved) isSALDataReserved() bool {
	return true
}

func (m *_SALDataReserved) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}