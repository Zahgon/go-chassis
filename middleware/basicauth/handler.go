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

package basicauth

import (
	"errors"

	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// errors
var (
	ErrInvalidBase64 = errors.New("invalid base64")
	ErrNoHeader      = errors.New("not authorized")
	ErrInvalidAuth   = errors.New("invalid authentication")
)

// HeaderAuth is common auth header
const HeaderAuth = "Authorization"

// Handler is is a basic auth pre process raw data in handler
type Handler struct {
}

// Handle pre process raw data in handler
func (ph *Handler) Handle(chain *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

func newBasicAuth() handler.Handler {
	_ = "STUB: not implemented"

	// Name returns the router string
	return *new(handler.Handler)
}

func (ph *Handler) Name() string { _ = "STUB: not implemented"; return "" }

func decode(subject string) (user string, pwd string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
