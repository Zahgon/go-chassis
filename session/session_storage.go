package session

import (
	"time"
)

// Save for setting the session uuid, endpoint, timeout
func Save(sid string, ep string, timeOut time.Duration) { _ = "STUB: not implemented"; return }

// Get return endpoint based on session uuid
func Get(sid string) (ep interface{}, ok bool) { _ = "STUB: not implemented"; return nil, false }

// ClearExpired delete all expired session
func ClearExpired() { _ = "STUB: not implemented"; return }

// Delete delete the session uuid
func Delete(sid string) { _ = "STUB: not implemented"; return }
