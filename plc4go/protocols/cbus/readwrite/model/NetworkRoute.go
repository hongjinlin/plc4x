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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// NetworkRoute is the corresponding interface of NetworkRoute
type NetworkRoute interface {
	utils.LengthAware
	utils.Serializable
	// GetNetworkPCI returns NetworkPCI (property field)
	GetNetworkPCI() NetworkProtocolControlInformation
	// GetAdditionalBridgeAddresses returns AdditionalBridgeAddresses (property field)
	GetAdditionalBridgeAddresses() []BridgeAddress
}

// NetworkRouteExactly can be used when we want exactly this type and not a type which fulfills NetworkRoute.
// This is useful for switch cases.
type NetworkRouteExactly interface {
	NetworkRoute
	isNetworkRoute() bool
}

// _NetworkRoute is the data-structure of this message
type _NetworkRoute struct {
	NetworkPCI                NetworkProtocolControlInformation
	AdditionalBridgeAddresses []BridgeAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NetworkRoute) GetNetworkPCI() NetworkProtocolControlInformation {
	return m.NetworkPCI
}

func (m *_NetworkRoute) GetAdditionalBridgeAddresses() []BridgeAddress {
	return m.AdditionalBridgeAddresses
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNetworkRoute factory function for _NetworkRoute
func NewNetworkRoute(networkPCI NetworkProtocolControlInformation, additionalBridgeAddresses []BridgeAddress) *_NetworkRoute {
	return &_NetworkRoute{NetworkPCI: networkPCI, AdditionalBridgeAddresses: additionalBridgeAddresses}
}

// Deprecated: use the interface for direct cast
func CastNetworkRoute(structType interface{}) NetworkRoute {
	if casted, ok := structType.(NetworkRoute); ok {
		return casted
	}
	if casted, ok := structType.(*NetworkRoute); ok {
		return *casted
	}
	return nil
}

func (m *_NetworkRoute) GetTypeName() string {
	return "NetworkRoute"
}

func (m *_NetworkRoute) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NetworkRoute) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (networkPCI)
	lengthInBits += m.NetworkPCI.GetLengthInBits()

	// Array field
	if len(m.AdditionalBridgeAddresses) > 0 {
		for i, element := range m.AdditionalBridgeAddresses {
			last := i == len(m.AdditionalBridgeAddresses)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_NetworkRoute) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NetworkRouteParse(readBuffer utils.ReadBuffer) (NetworkRoute, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NetworkRoute"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkRoute")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkPCI)
	if pullErr := readBuffer.PullContext("networkPCI"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkPCI")
	}
	_networkPCI, _networkPCIErr := NetworkProtocolControlInformationParse(readBuffer)
	if _networkPCIErr != nil {
		return nil, errors.Wrap(_networkPCIErr, "Error parsing 'networkPCI' field of NetworkRoute")
	}
	networkPCI := _networkPCI.(NetworkProtocolControlInformation)
	if closeErr := readBuffer.CloseContext("networkPCI"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkPCI")
	}

	// Array field (additionalBridgeAddresses)
	if pullErr := readBuffer.PullContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for additionalBridgeAddresses")
	}
	// Count array
	additionalBridgeAddresses := make([]BridgeAddress, uint16(networkPCI.GetStackDepth())-uint16(uint16(1)))
	// This happens when the size is set conditional to 0
	if len(additionalBridgeAddresses) == 0 {
		additionalBridgeAddresses = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(uint16(networkPCI.GetStackDepth())-uint16(uint16(1))); curItem++ {
			_item, _err := BridgeAddressParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'additionalBridgeAddresses' field of NetworkRoute")
			}
			additionalBridgeAddresses[curItem] = _item.(BridgeAddress)
		}
	}
	if closeErr := readBuffer.CloseContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for additionalBridgeAddresses")
	}

	if closeErr := readBuffer.CloseContext("NetworkRoute"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkRoute")
	}

	// Create the instance
	return &_NetworkRoute{
		NetworkPCI:                networkPCI,
		AdditionalBridgeAddresses: additionalBridgeAddresses,
	}, nil
}

func (m *_NetworkRoute) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("NetworkRoute"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NetworkRoute")
	}

	// Simple Field (networkPCI)
	if pushErr := writeBuffer.PushContext("networkPCI"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for networkPCI")
	}
	_networkPCIErr := writeBuffer.WriteSerializable(m.GetNetworkPCI())
	if popErr := writeBuffer.PopContext("networkPCI"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for networkPCI")
	}
	if _networkPCIErr != nil {
		return errors.Wrap(_networkPCIErr, "Error serializing 'networkPCI' field")
	}

	// Array Field (additionalBridgeAddresses)
	if pushErr := writeBuffer.PushContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for additionalBridgeAddresses")
	}
	for _, _element := range m.GetAdditionalBridgeAddresses() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'additionalBridgeAddresses' field")
		}
	}
	if popErr := writeBuffer.PopContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for additionalBridgeAddresses")
	}

	if popErr := writeBuffer.PopContext("NetworkRoute"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NetworkRoute")
	}
	return nil
}

func (m *_NetworkRoute) isNetworkRoute() bool {
	return true
}

func (m *_NetworkRoute) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
