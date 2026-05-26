package servicecomb

import (
	"github.com/go-chassis/go-chassis/v2/core/config"
)

// constant for route rule keys
const (
	DarkLaunchKey      = "^servicecomb\\.darklaunch\\.policy\\."
	DarkLaunchKeyV2    = "^servicecomb\\.routeRule\\."
	DarkLaunchPrefix   = "servicecomb.darklaunch.policy."
	DarkLaunchPrefixV2 = "servicecomb.routeRule."
	DarkLaunchTypeRule = "RULE"
	DarkLaunchTypeRate = "RATE"
)

// MergeLocalAndRemoteConfig get router config from archaius,
// including local file,memory and config server
func MergeLocalAndRemoteConfig() (map[string][]*config.RouteRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//then get config from archaius and simply overwrite rule from file

//filter out key:value pairs which are not route rules

func processV2Rule(ruleV2Map map[string]interface{}, destinations map[string][]*config.RouteRule) (map[string][]*config.RouteRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processV1Rule(ruleV1Map map[string]interface{}, destinations map[string][]*config.RouteRule) (map[string][]*config.RouteRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareRule(configMap map[string]interface{}, ruleV1Map map[string]interface{}, ruleV2Map map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}
