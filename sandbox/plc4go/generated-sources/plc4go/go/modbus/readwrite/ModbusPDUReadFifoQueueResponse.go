//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package readwrite

import "plc4x.apache.org/plc4go-modbus-driver/0.8.0/src/plc4go/spi"

type ModbusPDUReadFifoQueueResponse struct {
	fifoValue []uint16
	ModbusPDU
}

func (m ModbusPDUReadFifoQueueResponse) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Implicit Field (byteCount)
	lengthInBits += 16

	// Implicit Field (fifoCount)
	lengthInBits += 16

	// Array field
	if len(m.fifoValue) > 0 {
		lengthInBits += 16 * uint16(len(m.fifoValue))
	}

	return lengthInBits
}

func (m ModbusPDUReadFifoQueueResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func (m ModbusPDUReadFifoQueueResponse) Parse(io spi.ReadBuffer) {
	// TODO: Implement ...
}

func (m ModbusPDUReadFifoQueueResponse) Serialize(io spi.WriteBuffer) {
	// TODO: Implement ...
}