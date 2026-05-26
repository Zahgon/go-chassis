/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package quota is a alpha feature. it manage service quota
// quota management can not assure you strong consistency
package codec

import (
	"github.com/go-chassis/cari/codec"
)

type newCodec func(opts Options) (codec.Codec, error)

var plugins = map[string]newCodec{
	"encoding/json": newDefault,
}

var defaultCodec codec.Codec = &StdJson{}

// Install install codec plugin
func Install(name string, f newCodec) { _ = "STUB: not implemented"; return }

// Init init codec
func Init(opts Options) error { _ = "STUB: not implemented"; return nil }

func Encode(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func Decode(data []byte, v any) error { _ = "STUB: not implemented"; return nil }
