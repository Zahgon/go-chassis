package config

// constant for hystrix keys
const (
	FixedPrefix                       = "cse"
	NamespaceIsolation                = "isolation"
	NamespaceCircuitBreaker           = "circuitBreaker"
	NamespaceFallback                 = "fallback" //降级
	NamespaceFallbackpolicy           = "fallbackpolicy"
	PropertyTimeoutInMilliseconds     = "timeoutInMilliseconds"
	PropertyTimeoutEnabled            = "timeout.enabled"
	PropertyMaxConcurrentRequests     = "maxConcurrentRequests"
	PropertyErrorThresholdPercentage  = "errorThresholdPercentage"  //失败率
	PropertyRequestVolumeThreshold    = "requestVolumeThreshold"    //窗口请求数
	PropertySleepWindowInMilliseconds = "sleepWindowInMilliseconds" //熔断时间窗
	PropertyEnabled                   = "enabled"
	PropertyForce                     = "force"
	PropertyPolicy                    = "policy"
	PropertyForceClosed               = "forceClosed"
	PropertyForceOpen                 = "forceOpen"
	PropertyFault                     = "fault"
	PropertyGlobal                    = "_global"
	PropertyGovernance                = "governance"
	PropertyConsumer                  = "Consumer"
	PropertySchema                    = "schemas"
	PropertyOperations                = "operations"
	PropertyProtocol                  = "protocols"
	PropertyAbort                     = "abort"
	PropertyPercent                   = "percent"
	PropertyFixedDelay                = "fixedDelay"
	PropertyDelay                     = "delay"
	PropertyHTTPStatus                = "httpStatus"

	LoadBalance = "loadbalance"
)

/*
Hystrix Keys
*/

// GetHystrixSpecificKey get hystrix specific key
func GetHystrixSpecificKey(namespace, cmd, property string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetForceFallbackKey get force fallback key
func GetForceFallbackKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultForceFallbackKey get default force fallback key
func GetDefaultForceFallbackKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetTimeoutKey get timeout key
func GetTimeoutKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultTimeoutKey get default timeout key
func GetDefaultTimeoutKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetMaxConcurrentKey get maximum concurrent key
func GetMaxConcurrentKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultMaxConcurrentKey get default maximum concurrent key
func GetDefaultMaxConcurrentKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetErrorPercentThresholdKey get error percentage threshold key
func GetErrorPercentThresholdKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultErrorPercentThreshold get default error percentage threshold value
func GetDefaultErrorPercentThreshold(t string) string { _ = "STUB: not implemented"; return "" }

// GetRequestVolumeThresholdKey get request volume threshold key
func GetRequestVolumeThresholdKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultRequestVolumeThresholdKey get default request volume threshold key
func GetDefaultRequestVolumeThresholdKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetSleepWindowKey get sleep window key
func GetSleepWindowKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultSleepWindowKey get default sleep window key
func GetDefaultSleepWindowKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetForceCloseKey get force close key
func GetForceCloseKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultForceCloseKey get default force close key
func GetDefaultForceCloseKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetForceOpenKey get force open key
func GetForceOpenKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultForceOpenKey get default force open key
func GetDefaultForceOpenKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetCircuitBreakerEnabledKey get circuit breaker enabled key
func GetCircuitBreakerEnabledKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultCircuitBreakerEnabledKey get default circuit breaker enabled key
func GetDefaultCircuitBreakerEnabledKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetFallbackEnabledKey get fallback enabled key
func GetFallbackEnabledKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultGetFallbackEnabledKey get default fallback enabled key
func GetDefaultGetFallbackEnabledKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetFallbackPolicyKey get fallback policy key
func GetFallbackPolicyKey(command string) string { _ = "STUB: not implemented"; return "" }

// GetDefaultFallbackPolicyKey get default fallback policy key
func GetDefaultFallbackPolicyKey(t string) string { _ = "STUB: not implemented"; return "" }

// GetFilterNamesKey get filer name and key
func GetFilterNamesKey() string { _ = "STUB: not implemented"; return "" }

// GetFaultInjectionOperationKey get fault injection operation key
func GetFaultInjectionOperationKey(microServiceName, schema, operation string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetFaultInjectionSchemaKey get fault injection schema key
func GetFaultInjectionSchemaKey(microServiceName, schema string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetFaultInjectionServiceKey get fault injection service key
func GetFaultInjectionServiceKey(microServiceName string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetFaultInjectionGlobalKey get fault injection global key
func GetFaultInjectionGlobalKey() string { _ = "STUB: not implemented"; return "" }

// GetFaultAbortPercentKey get fault abort percentage key
func GetFaultAbortPercentKey(key, protocol string) string { _ = "STUB: not implemented"; return "" }

// GetFaultAbortHTTPStatusKey get fault abort http status key
func GetFaultAbortHTTPStatusKey(key, protocol string) string { _ = "STUB: not implemented"; return "" }

// GetFaultDelayPercentKey get fault daley percentage key
func GetFaultDelayPercentKey(key, protocol string) string { _ = "STUB: not implemented"; return "" }

// GetFaultFixedDelayKey get fault fixed delay key
func GetFaultFixedDelayKey(key, protocol string) string { _ = "STUB: not implemented"; return "" }
