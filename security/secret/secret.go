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

package secret

import (
	"crypto/rsa"
)

// GenRSAPrivateKey generate a rsa private key
func GenRSAPrivateKey(bits int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GenRSAKeyPair create rsa key pair
func GenRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RSAPrivate2Bytes expose bytes of private key
func RSAPrivate2Bytes(privateKey *rsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RSAPublicKey2Bytes expose bytes of public key
func RSAPublicKey2Bytes(publicKey *rsa.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseRSAPrivateKey convert string to private key
func ParseRSAPrivateKey(key string) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseRSAPPublicKey convert string to pub key
func ParseRSAPPublicKey(key string) (*rsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
