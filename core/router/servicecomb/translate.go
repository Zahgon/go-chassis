package servicecomb

import (
	"github.com/go-chassis/go-chassis/v2/core/config"
)

// ConvertJSON2RouteRule parse raw json from cse server to route rule config
func ConvertJSON2RouteRule(raw string) ([]*config.RouteRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DarkLaunchRule2RouteRule translates dark launch rule to route rule
func DarkLaunchRule2RouteRule(rule *config.DarkLaunchRule) []*config.RouteRule {
	_ = "STUB: not implemented"
	return nil
}

// generateRouteTags generate route tags
func generateRouteTags(weights int, versions []string) []*config.RouteTag {
	_ = "STUB: not implemented"
	return nil
}

func caseInsensitiveToString(isCaseInsensitive bool) string { _ = "STUB: not implemented"; return "" }

func setHeadersAndHTTPHeaders(match *config.Match, isCaseInsensitive bool, cKey, con, sp string) {
	_ = "STUB: not implemented"
	return
}

func toCamelCase(s string) string { _ = "STUB: not implemented"; return "" }
