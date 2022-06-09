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

// BACnetEventLogRecordLogDatumLogStatus is the data-structure of this message
type BACnetEventLogRecordLogDatumLogStatus struct {
	*BACnetEventLogRecordLogDatum
	LogStatus *BACnetLogStatusTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetEventLogRecordLogDatumLogStatus is the corresponding interface of BACnetEventLogRecordLogDatumLogStatus
type IBACnetEventLogRecordLogDatumLogStatus interface {
	IBACnetEventLogRecordLogDatum
	// GetLogStatus returns LogStatus (property field)
	GetLogStatus() *BACnetLogStatusTagged
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

func (m *BACnetEventLogRecordLogDatumLogStatus) InitializeParent(parent *BACnetEventLogRecordLogDatum, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetEventLogRecordLogDatum.OpeningTag = openingTag
	m.BACnetEventLogRecordLogDatum.PeekedTagHeader = peekedTagHeader
	m.BACnetEventLogRecordLogDatum.ClosingTag = closingTag
}

func (m *BACnetEventLogRecordLogDatumLogStatus) GetParent() *BACnetEventLogRecordLogDatum {
	return m.BACnetEventLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventLogRecordLogDatumLogStatus) GetLogStatus() *BACnetLogStatusTagged {
	return m.LogStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventLogRecordLogDatumLogStatus factory function for BACnetEventLogRecordLogDatumLogStatus
func NewBACnetEventLogRecordLogDatumLogStatus(logStatus *BACnetLogStatusTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetEventLogRecordLogDatumLogStatus {
	_result := &BACnetEventLogRecordLogDatumLogStatus{
		LogStatus:                    logStatus,
		BACnetEventLogRecordLogDatum: NewBACnetEventLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetEventLogRecordLogDatumLogStatus(structType interface{}) *BACnetEventLogRecordLogDatumLogStatus {
	if casted, ok := structType.(BACnetEventLogRecordLogDatumLogStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatumLogStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetEventLogRecordLogDatum); ok {
		return CastBACnetEventLogRecordLogDatumLogStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatum); ok {
		return CastBACnetEventLogRecordLogDatumLogStatus(casted.Child)
	}
	return nil
}

func (m *BACnetEventLogRecordLogDatumLogStatus) GetTypeName() string {
	return "BACnetEventLogRecordLogDatumLogStatus"
}

func (m *BACnetEventLogRecordLogDatumLogStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventLogRecordLogDatumLogStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (logStatus)
	lengthInBits += m.LogStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventLogRecordLogDatumLogStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventLogRecordLogDatumLogStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetEventLogRecordLogDatumLogStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatumLogStatus"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (logStatus)
	if pullErr := readBuffer.PullContext("logStatus"); pullErr != nil {
		return nil, pullErr
	}
	_logStatus, _logStatusErr := BACnetLogStatusTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _logStatusErr != nil {
		return nil, errors.Wrap(_logStatusErr, "Error parsing 'logStatus' field")
	}
	logStatus := CastBACnetLogStatusTagged(_logStatus)
	if closeErr := readBuffer.CloseContext("logStatus"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatumLogStatus"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetEventLogRecordLogDatumLogStatus{
		LogStatus:                    CastBACnetLogStatusTagged(logStatus),
		BACnetEventLogRecordLogDatum: &BACnetEventLogRecordLogDatum{},
	}
	_child.BACnetEventLogRecordLogDatum.Child = _child
	return _child, nil
}

func (m *BACnetEventLogRecordLogDatumLogStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatumLogStatus"); pushErr != nil {
			return pushErr
		}

		// Simple Field (logStatus)
		if pushErr := writeBuffer.PushContext("logStatus"); pushErr != nil {
			return pushErr
		}
		_logStatusErr := m.LogStatus.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("logStatus"); popErr != nil {
			return popErr
		}
		if _logStatusErr != nil {
			return errors.Wrap(_logStatusErr, "Error serializing 'logStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatumLogStatus"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetEventLogRecordLogDatumLogStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}