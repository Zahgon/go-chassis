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

package chassis

import (
	"os"
	"sync"

	"github.com/go-chassis/go-chassis/v2/core/server"
)

type chassis struct {
	schemas     []*Schema
	mu          sync.Mutex
	Initialized bool

	DefaultConsumerChainNames map[string]string
	DefaultProviderChainNames map[string]string

	sigs                   []os.Signal
	preShutDownFuncs       map[string]func(os.Signal)
	postShutDownFuncs      map[string]func(os.Signal)
	hijackGracefulShutdown func(os.Signal)
}

// Schema struct for to represent schema info
type Schema struct {
	serverName string
	schema     interface{}
	opts       []server.RegisterOption
}

func (c *chassis) initChains(chainType string) error { _ = "STUB: not implemented"; return nil }

func (c *chassis) initHandler() error { _ = "STUB: not implemented"; return nil }

// Init
func (c *chassis) initialize() error { _ = "STUB: not implemented"; return nil }

// router needs get configs from config-server when init
// so it must init after bootstrap

func initTooling() error { _ = "STUB: not implemented"; return nil }

func initBackendPlugins() error { _ = "STUB: not implemented"; return nil }

func (c *chassis) registerSchema(serverName string, structPtr interface{}, opts ...server.RegisterOption) {
	_ = "STUB: not implemented"
	return
}

func (c *chassis) start(options ...server.RunOption) error { _ = "STUB: not implemented"; return nil }
