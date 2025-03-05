// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium
package singleton

import (
	"github.com/cilium/cilium/pkg/fqdn/proxy"
	"github.com/cilium/cilium/pkg/lock"
)

// NewDefaultDNSProxy instantiates the DefaultDNSProxy singleton.
func NewDefaultDNSProxy() *DefaultDNSProxy {
	return &DefaultDNSProxy{
		mu: &lock.RWMutex{},
	}
}

// DefaultDNSProxy is the default dns proxy for
// the process.
type DefaultDNSProxy struct {
	mu    *lock.RWMutex
	proxy proxy.DNSProxier
}

// Set sets the DefaultDNSProxy
func (ddp *DefaultDNSProxy) Set(dnsProxy proxy.DNSProxier) {
	ddp.mu.Lock()
	defer ddp.mu.Unlock()
	ddp.proxy = dnsProxy
}

// Get gets the DefaultDNSProxy
func (ddp *DefaultDNSProxy) Get() proxy.DNSProxier {
	ddp.mu.RLock()
	defer ddp.mu.RUnlock()
	return ddp.proxy
}
