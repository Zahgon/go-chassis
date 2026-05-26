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

package marker

import (
	"sync"

	"github.com/go-chassis/go-chassis/v2/core/config"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

const (
	Once       = "once"
	PerService = "perService"
)

var matches sync.Map

// Operate decide value match expression or not
type Operate func(value, expression string) bool

var operatorPlugin = map[string]Operate{
	"exact":     exact,
	"contains":  contains,
	"regex":     regex,
	"noEqu":     noEqu,
	"less":      less,
	"noLess":    noLess,
	"greater":   greater,
	"noGreater": noGreater,
}

// Install a strategy
func Install(name string, m Operate) { _ = "STUB: not implemented"; return }

// Mark mark an invocation with matchName by match policy
func Mark(inv *invocation.Invocation) { _ = "STUB: not implemented"; return }

// the invocation math policy

func isMatch(inv *invocation.Invocation, matchPolicy config.MatchPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func methodMatch(reqMethod string, methods []string) bool { _ = "STUB: not implemented"; return false }

func apiMatch(apiPath string, apiPolicy map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

func queryMatch(query string, queryPolicy map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

func headsMatch(headers map[string]string, headPolicy map[string]map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

// match compare value and expression
func Match(operator, value, expression string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SaveMatchPolicy saves match policy
func SaveMatchPolicy(name, value string, k string) error { _ = "STUB: not implemented"; return nil }

// Policy return policy
func Policy(name string) *config.MatchPolicies { _ = "STUB: not implemented"; return nil }
