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

package circuit

const (
	regex       = "(Provider|Consumer)\\.(.*)"
	regexSource = "\\.(.+)\\.(Provider|Consumer)\\.(.*)"
)

// GetMetricsName get only metrics name from cmd name
func GetMetricsName(cmd string) (source string) { _ = "STUB: not implemented"; return "" }

// ParseCircuitCMD return metrics related infos
// example Consumer.ErrServer.rest./sayhimessage.rejects
// the first and last string consist of metrics name
// second is ErrServer
// 3th and 4th is schema and operation
func ParseCircuitCMD(cmd string) (source string, target string, schema string, op string) {
	_ = "STUB: not implemented"
	return "", "", "", ""
}

// ExtractServiceSchemaOperationMetrics parse service,schema and operation
// key example Microservice.SchemaID.OperationId.metrics
func ExtractServiceSchemaOperationMetrics(raw string) (target, schemaID, operation, metrics string) {
	_ = "STUB: not implemented"
	return "", "", "", ""
}

// GetEventType get metrics suffix
func GetEventType(cmdName string) string { _ = "STUB: not implemented"; return "" }
