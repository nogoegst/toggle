// toggle.go - readiness toggle.
//
// To the extent possible under law, Ivan Markin has waived all copyright
// and related or neighboring rights to toggle, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package toggle

import (
	"fmt"
	"sync/atomic"
)

// Toggle is acctually a uint32 variable.
type Toggle = uint32

var StatusReady = uint32(1<<32 - 1)
var StatusNotReady = uint32(0)

// Ready sets status of toggle t to StatusReady.
func Ready(t *Toggle) {
	atomic.StoreUint32(t, StatusReady)
}

// Ready sets status of toggle t to StatusNotReady.
func NotReady(t *Toggle) {
	atomic.StoreUint32(t, StatusNotReady)
}

// Status returns current status of toggle t.
func Status(t *Toggle) uint32 {
	return atomic.LoadUint32(t)
}

// Check returns function that returns error when status of toggle t
// is not StatusReady.
func Check(t *Toggle) func() error {
	return func() error {
		if status := Status(t); status != StatusReady {
			return fmt.Errorf("not ready: status %u", status)
		}
		return nil
	}
}
