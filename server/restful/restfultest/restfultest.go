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

package restfultest

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/go-chassis/go-chassis/v2/core/handler"
)

// Container is unit test solution for rest api method
type Container struct {
	container *restful.Container
	ws        *restful.WebService
}

// New create a isolated test container,
// you can register a struct, and it will be registered to a isolated container
func New(schema interface{}, chain *handler.Chain) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ServeHTTP accept native httptest, after process, response writer will write response
func (c *Container) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}
