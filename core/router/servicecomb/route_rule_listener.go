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

package servicecomb

import (
	"github.com/go-chassis/go-archaius/event"
	"github.com/go-chassis/go-chassis/v2/core/config"
)

type routeRuleEventListener struct{}

// update route rule of a service
func (r *routeRuleEventListener) Event(e *event.Event) { _ = "STUB: not implemented"; return }

// SaveRouteRule save event rule to local cache
func SaveRouteRule(service string, raw string, isV2 bool) { _ = "STUB: not implemented"; return }

func validateAndUpdate(routeRules []*config.RouteRule, service string) {
	_ = "STUB: not implemented"
	return
}
