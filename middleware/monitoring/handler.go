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

package monitoring

import (
	"github.com/go-chassis/openlog"

	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// errors
const (
	MetricsLatency = "scb_request_process_duration"
	MetricsRequest = "scb_request_count"
	MetricsErrors  = "scb_error_response_count"
	Name           = "monitoring"
)

var labels = []string{"service", "instance", "version", "app", "env", "API", "method"}
var labels4Resp = []string{"service", "instance", "version", "app", "env", "code", "API", "method"}

// Handler monitor server side metrics, the key metrics is latency, QPS, Errors, do not use it in consumer chain
type Handler struct {
}

// Handle record metrics
func (ph *Handler) Handle(chain *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

//skip monitoring

func newHandler() handler.Handler { _ = "STUB: not implemented"; return *new(handler.Handler) }

func GetUrlPath(i *invocation.Invocation) string { _ = "STUB: not implemented"; return "" }

// Name returns the router string
func (ph *Handler) Name() string { _ = "STUB: not implemented"; return "" }

func init() {
	err := handler.RegisterHandler(Name, newHandler)
	if err != nil {
		openlog.Error(err.Error())
	}
}
