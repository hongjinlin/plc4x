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

// BACnetConstructedDataIPv6DHCPServer is the data-structure of this message
type BACnetConstructedDataIPv6DHCPServer struct {
	*BACnetConstructedData
	DhcpServer *BACnetApplicationTagOctetString

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataIPv6DHCPServer is the corresponding interface of BACnetConstructedDataIPv6DHCPServer
type IBACnetConstructedDataIPv6DHCPServer interface {
	IBACnetConstructedData
	// GetDhcpServer returns DhcpServer (property field)
	GetDhcpServer() *BACnetApplicationTagOctetString
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

func (m *BACnetConstructedDataIPv6DHCPServer) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataIPv6DHCPServer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DHCP_SERVER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIPv6DHCPServer) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIPv6DHCPServer) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIPv6DHCPServer) GetDhcpServer() *BACnetApplicationTagOctetString {
	return m.DhcpServer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6DHCPServer factory function for BACnetConstructedDataIPv6DHCPServer
func NewBACnetConstructedDataIPv6DHCPServer(dhcpServer *BACnetApplicationTagOctetString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataIPv6DHCPServer {
	_result := &BACnetConstructedDataIPv6DHCPServer{
		DhcpServer:            dhcpServer,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIPv6DHCPServer(structType interface{}) *BACnetConstructedDataIPv6DHCPServer {
	if casted, ok := structType.(BACnetConstructedDataIPv6DHCPServer); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DHCPServer); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPv6DHCPServer(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPv6DHCPServer(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIPv6DHCPServer) GetTypeName() string {
	return "BACnetConstructedDataIPv6DHCPServer"
}

func (m *BACnetConstructedDataIPv6DHCPServer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIPv6DHCPServer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dhcpServer)
	lengthInBits += m.DhcpServer.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataIPv6DHCPServer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPv6DHCPServerParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataIPv6DHCPServer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DHCPServer"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dhcpServer)
	if pullErr := readBuffer.PullContext("dhcpServer"); pullErr != nil {
		return nil, pullErr
	}
	_dhcpServer, _dhcpServerErr := BACnetApplicationTagParse(readBuffer)
	if _dhcpServerErr != nil {
		return nil, errors.Wrap(_dhcpServerErr, "Error parsing 'dhcpServer' field")
	}
	dhcpServer := CastBACnetApplicationTagOctetString(_dhcpServer)
	if closeErr := readBuffer.CloseContext("dhcpServer"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DHCPServer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIPv6DHCPServer{
		DhcpServer:            CastBACnetApplicationTagOctetString(dhcpServer),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIPv6DHCPServer) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DHCPServer"); pushErr != nil {
			return pushErr
		}

		// Simple Field (dhcpServer)
		if pushErr := writeBuffer.PushContext("dhcpServer"); pushErr != nil {
			return pushErr
		}
		_dhcpServerErr := m.DhcpServer.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dhcpServer"); popErr != nil {
			return popErr
		}
		if _dhcpServerErr != nil {
			return errors.Wrap(_dhcpServerErr, "Error serializing 'dhcpServer' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DHCPServer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIPv6DHCPServer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}