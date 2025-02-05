// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package codec

type Record struct {
	RecordAttributes  byte
	RelativeTimestamp int64
	RelativeOffset    int
	Key               []byte
	Value             []byte
	Headers           []byte
}

func DecodeRecord(bytes []byte, version int16) *Record {
	record := &Record{}
	idx := 0
	record.RecordAttributes, idx = readRecordAttributes(bytes, idx)
	record.RelativeTimestamp, idx = readRelativeTimestamp(bytes, idx)
	record.RelativeOffset, idx = readRelativeOffset(bytes, idx)
	record.Key, idx = readVCompactBytes(bytes, idx)
	record.Value, idx = readVCompactBytes(bytes, idx)
	record.Headers, idx = readCompactNullableBytes(bytes, idx)
	return record
}

func (r *Record) BytesLength() int {
	result := 0
	result += LenRecordAttributes
	result += varint64Size(r.RelativeTimestamp)
	result += varintSize(r.RelativeOffset)
	result += CompactBytesLen(r.Key)
	result += CompactBytesLen(r.Value)
	result += CompactNullableBytesLen(r.Headers)
	return result
}

func (r *Record) Bytes() []byte {
	bytes := make([]byte, r.BytesLength())
	idx := 0
	idx = putRecordAttributes(bytes, idx, 0)
	idx = putRelativeTimestamp(bytes, idx, r.RelativeTimestamp)
	idx = putRelativeOffset(bytes, idx, r.RelativeOffset)
	idx = putVCompactBytes(bytes, idx, r.Key)
	idx = putVCompactBytes(bytes, idx, r.Value)
	idx = putCompactNullableBytes(bytes, idx, r.Headers)
	return bytes
}
