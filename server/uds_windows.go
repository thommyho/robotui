//go:build windows

package server

import "github.com/thommyho/robotui/core/site"

// HealthListener attaches listener to unix domain socket
func HealthListener(_ site.API) {
	// nop
}
