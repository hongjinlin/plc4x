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

// SALDataEnableControl is the corresponding interface of SALDataEnableControl
type SALDataEnableControl interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetEnableControlData returns EnableControlData (property field)
	GetEnableControlData() EnableControlData
}

// SALDataEnableControlExactly can be used when we want exactly this type and not a type which fulfills SALDataEnableControl.
// This is useful for switch cases.
type SALDataEnableControlExactly interface {
	SALDataEnableControl
	isSALDataEnableControl() bool
}

// _SALDataEnableControl is the data-structure of this message
type _SALDataEnableControl struct {
	*_SALData
	EnableControlData EnableControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataEnableControl) GetApplicationId() ApplicationId {
	return ApplicationId_ENABLE_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataEnableControl) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataEnableControl) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataEnableControl) GetEnableControlData() EnableControlData {
	return m.EnableControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataEnableControl factory function for _SALDataEnableControl
func NewSALDataEnableControl(enableControlData EnableControlData, salData SALData) *_SALDataEnableControl {
	_result := &_SALDataEnableControl{
		EnableControlData: enableControlData,
		_SALData:          NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataEnableControl(structType interface{}) SALDataEnableControl {
	if casted, ok := structType.(SALDataEnableControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataEnableControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataEnableControl) GetTypeName() string {
	return "SALDataEnableControl"
}

func (m *_SALDataEnableControl) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataEnableControl) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (enableControlData)
	lengthInBits += m.EnableControlData.GetLengthInBits()

	return lengthInBits
}

func (m *_SALDataEnableControl) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataEnableControlParse(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataEnableControl, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataEnableControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataEnableControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (enableControlData)
	if pullErr := readBuffer.PullContext("enableControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enableControlData")
	}
	_enableControlData, _enableControlDataErr := EnableControlDataParse(readBuffer)
	if _enableControlDataErr != nil {
		return nil, errors.Wrap(_enableControlDataErr, "Error parsing 'enableControlData' field of SALDataEnableControl")
	}
	enableControlData := _enableControlData.(EnableControlData)
	if closeErr := readBuffer.CloseContext("enableControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enableControlData")
	}

	if closeErr := readBuffer.CloseContext("SALDataEnableControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataEnableControl")
	}

	// Create a partially initialized instance
	_child := &_SALDataEnableControl{
		EnableControlData: enableControlData,
		_SALData:          &_SALData{},
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataEnableControl) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataEnableControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataEnableControl")
		}

		// Simple Field (enableControlData)
		if pushErr := writeBuffer.PushContext("enableControlData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enableControlData")
		}
		_enableControlDataErr := writeBuffer.WriteSerializable(m.GetEnableControlData())
		if popErr := writeBuffer.PopContext("enableControlData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enableControlData")
		}
		if _enableControlDataErr != nil {
			return errors.Wrap(_enableControlDataErr, "Error serializing 'enableControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataEnableControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataEnableControl")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataEnableControl) isSALDataEnableControl() bool {
	return true
}

func (m *_SALDataEnableControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}