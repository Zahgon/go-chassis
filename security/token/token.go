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

package token

import (
	"errors"
)

// DefaultManager can be replaced
var DefaultManager Manager = &jwtTokenManager{}

// token pkg common errors
var (
	ErrInvalidExp = errors.New("expire time is illegal")
)

// jwt claims RFC 7519
// https://tools.ietf.org/html/rfc7519#section-4.1.2
const (
	JWTClaimsExp = "exp"
	JWTClaimsSub = "sub"
)

// SecretFunc is a callback function to supply
// the key for verification.  The function receives the parsed,
// but unverified claims in Token.  This allows you to use properties in the
// claims of the token (such as `username`) to identify which key to use.
type SecretFunc func(claims interface{}, method SigningMethod) (interface{}, error)

// Sign gen token
func Sign(claims map[string]interface{}, secret interface{}, opts ...Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Verify return claims
func Verify(tokenString string, f SecretFunc, opts ...Option) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Manager manages token
type Manager interface {
	Sign(claims map[string]interface{}, secret interface{}, option ...Option) (string, error)
	Verify(tokenString string, f SecretFunc, opts ...Option) (map[string]interface{}, error)
}
type jwtTokenManager struct {
}

// Sign signature a token
func (j *jwtTokenManager) Sign(claims map[string]interface{}, secret interface{}, opts ...Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Verify return claims
func (j *jwtTokenManager) Verify(tokenString string, f SecretFunc, opts ...Option) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Token is either expired or not active yet
