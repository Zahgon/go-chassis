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

package jwt

import (
	"errors"
	"net/http"

	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/openlog"
)

// errors
var (
	ErrNoHeader    = errors.New("no authorization in header")
	ErrInvalidAuth = errors.New("invalid authentication")
)

// Handler is is a jwt interceptor
type Handler struct {
}

// Handle intercept unauthorized request
func (h *Handler) Handle(chain *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"

	//jwt is not initialized, then skip authentication, do not report error
	return
}

func mustAuth(req *http.Request) bool { _ = "STUB: not implemented"; return false }

func newHandler() handler.Handler {
	_ = "STUB: not implemented"

	// Name returns the router string
	return *new(handler.Handler)
}

func (h *Handler) Name() string { _ = "STUB: not implemented"; return "" }

func init() {
	err := handler.RegisterHandler("jwt", newHandler)
	if err != nil {
		openlog.Error(err.Error())
	}
}
