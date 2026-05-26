package session

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/patrickmn/go-cache"
)

// ErrResponseNil used for to represent the error response, when it is nil
var ErrResponseNil = errors.New("can not set session, resp is nil")

// Cache session cache variable
var Cache *cache.Cache

// SessionStickinessCache key: go-chassisLB , value is cookie
var SessionStickinessCache *cache.Cache

func init() {
	Cache = initCache()
	SessionStickinessCache = initCache()
	cookieMap = make(map[string]string)
}
func initCache() *cache.Cache { _ = "STUB: not implemented"; return nil }

var cookieMap map[string]string

// getLBCookie gets cookie from local map
func getLBCookie(key string) string { _ = "STUB: not implemented"; return "" }

// setLBCookie sets cookie to local map
func setLBCookie(key, value string) { _ = "STUB: not implemented"; return }

// GetContextMetadata gets data from context
func GetContextMetadata(ctx context.Context, key string) string {
	_ = "STUB: not implemented"
	return ""
}

// SetContextMetadata sets data to context
func SetContextMetadata(ctx context.Context, key string, value string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GetSessionFromResp return session uuid in resp if there is
func GetSessionFromResp(cookieKey string, resp *http.Response) string {
	_ = "STUB: not implemented"
	return ""
}

// SaveSessionIDFromContext check session id in response ctx and save it to session storage
func SaveSessionIDFromContext(ctx context.Context, ep string, autoTimeout int) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Temporary responsewriter for SetCookie
type cookieResponseWriter http.Header

// Header implements ResponseWriter Header interface
func (c cookieResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *

	// Write is a dummy function
	new(http.Header)
}

func (c cookieResponseWriter) Write([]byte) (int, error) {
	_ = "STUB: not implemented"

	// WriteHeader is a dummy function
	return 0, nil
}

func (c cookieResponseWriter) WriteHeader(int) {
	_ = "STUB: not implemented"

	// setCookie appends cookie with already present cookie with ';' in between
	return
}

func setCookie(resp *http.Response, value string) { _ = "STUB: not implemented"; return }

//If cookie is already set, append it with ';'

// SaveSessionIDFromHTTP check session id
func SaveSessionIDFromHTTP(ep string, autoTimeout int, resp *http.Response, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

//if session is in resp, then just save it

// GenerateSessionID generate a session id
func GenerateSessionID() (string, error) { _ = "STUB: not implemented"; return "", nil }

// DeletingKeySuccessiveFailure deleting key successes and failures
func DeletingKeySuccessiveFailure(resp *http.Response) { _ = "STUB: not implemented"; return }

// GetSessionCookie getting session cookie
func GetSessionCookie(ctx context.Context, resp *http.Response) string {
	_ = "STUB: not implemented"
	return ""
}

// AddSessionStickinessToCache add new cookie or refresh old cookie
func AddSessionStickinessToCache(cookie, namespace string) { _ = "STUB: not implemented"; return }

// GetSessionID get sessionID from cache
func GetSessionID(namespace string) string { _ = "STUB: not implemented"; return "" }

func getSessionStickinessCacheKey(namespace string) string { _ = "STUB: not implemented"; return "" }

// GetSessionIDFromInv when use  SessionStickiness , get session id from inv
func GetSessionIDFromInv(inv invocation.Invocation, key string) string {
	_ = "STUB: not implemented"
	return ""
}
