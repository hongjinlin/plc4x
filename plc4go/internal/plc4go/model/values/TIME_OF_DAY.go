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
package values

import (
	"encoding/xml"
	"time"
)

type PlcTIME_OF_DAY struct {
	value time.Time
	PlcSimpleValueAdapter
}

func NewPlcTIME_OF_DAY(value time.Time) PlcTIME_OF_DAY {
	safeValue := time.Date(0, 0, 0, value.Hour(), value.Minute(), value.Second(), value.Nanosecond(), value.Location())
	return PlcTIME_OF_DAY{
		value: safeValue,
	}
}

func (m PlcTIME_OF_DAY) IsTime() bool {
	return true
}

func (m PlcTIME_OF_DAY) GetTime() time.Time {
	return m.value
}

func (m PlcTIME_OF_DAY) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.value, xml.StartElement{Name: xml.Name{Local: "PlcTIME_OF_DAY"}}); err != nil {
		return err
	}
	return nil
}
