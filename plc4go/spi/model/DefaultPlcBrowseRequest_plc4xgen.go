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

// Code generated by "plc4xgenerator -type=DefaultPlcBrowseRequest"; DO NOT EDIT.

package model

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *DefaultPlcBrowseRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("PlcBrowseRequest"); err != nil {
		return err
	}
	if err := d.DefaultRequest.Serialize(writeBuffer); err != nil {
		return err
	}

	if d.browser != nil {
		if serializableField, ok := d.browser.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("browser"); err != nil {
				return err
			}
			if err := serializableField.Serialize(writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("browser"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.browser)
			if err := writeBuffer.WriteString("browser", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("PlcBrowseRequest"); err != nil {
		return err
	}
	return nil
}

func (d *DefaultPlcBrowseRequest) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
