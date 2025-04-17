// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package ipnamelookup represents the imported interface "wasi:sockets/ip-name-lookup@0.2.0".
package ipnamelookup

import (
	"github.com/golemcloud/golem-go/binding/wasi/io/poll"
	"github.com/golemcloud/golem-go/binding/wasi/sockets/network"
	"go.bytecodealliance.org/cm"
)

// Pollable represents the imported type alias "wasi:sockets/ip-name-lookup@0.2.0#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// Network represents the imported type alias "wasi:sockets/ip-name-lookup@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// ErrorCode represents the type alias "wasi:sockets/ip-name-lookup@0.2.0#error-code".
//
// See [network.ErrorCode] for more information.
type ErrorCode = network.ErrorCode

// IPAddress represents the type alias "wasi:sockets/ip-name-lookup@0.2.0#ip-address".
//
// See [network.IPAddress] for more information.
type IPAddress = network.IPAddress

// ResolveAddressStream represents the imported resource "wasi:sockets/ip-name-lookup@0.2.0#resolve-address-stream".
//
//	resource resolve-address-stream
type ResolveAddressStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "resolve-address-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self ResolveAddressStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ResolveAddressStreamResourceDrop((uint32)(self0))
	return
}

// ResolveNextAddress represents the imported method "resolve-next-address".
//
// Returns the next address from the resolver.
//
// This function should be called multiple times. On each call, it will
// return the next address in connection order preference. If all
// addresses have been exhausted, this function returns `none`.
//
// This function never returns IPv4-mapped IPv6 addresses.
//
// # Typical errors
// - `name-unresolvable`:          Name does not exist or has no suitable associated
// IP addresses. (EAI_NONAME, EAI_NODATA, EAI_ADDRFAMILY)
// - `temporary-resolver-failure`: A temporary failure in name resolution occurred.
// (EAI_AGAIN)
// - `permanent-resolver-failure`: A permanent failure in name resolution occurred.
// (EAI_FAIL)
// - `would-block`:                A result is not available yet. (EWOULDBLOCK, EAGAIN)
//
//	resolve-next-address: func() -> result<option<ip-address>, error-code>
//
//go:nosplit
func (self ResolveAddressStream) ResolveNextAddress() (result cm.Result[OptionIPAddressShape, cm.Option[IPAddress], ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ResolveAddressStreamResolveNextAddress((uint32)(self0), &result)
	return
}

// Subscribe represents the imported method "subscribe".
//
// Create a `pollable` which will resolve once the stream is ready for I/O.
//
// Note: this function is here for WASI Preview2 only.
// It's planned to be removed when `future` is natively supported in Preview3.
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self ResolveAddressStream) Subscribe() (result Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_ResolveAddressStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[Pollable]((uint32)(result0))
	return
}

// ResolveAddresses represents the imported function "resolve-addresses".
//
// Resolve an internet host name to a list of IP addresses.
//
// Unicode domain names are automatically converted to ASCII using IDNA encoding.
// If the input is an IP address string, the address is parsed and returned
// as-is without making any external requests.
//
// See the wasi-socket proposal README.md for a comparison with getaddrinfo.
//
// This function never blocks. It either immediately fails or immediately
// returns successfully with a `resolve-address-stream` that can be used
// to (asynchronously) fetch the results.
//
// # Typical errors
// - `invalid-argument`: `name` is a syntactically invalid domain name or IP address.
//
// # References:
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/getaddrinfo.html>
// - <https://man7.org/linux/man-pages/man3/getaddrinfo.3.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/ws2tcpip/nf-ws2tcpip-getaddrinfo>
// - <https://man.freebsd.org/cgi/man.cgi?query=getaddrinfo&sektion=3>
//
//	resolve-addresses: func(network: borrow<network>, name: string) -> result<resolve-address-stream,
//	error-code>
//
//go:nosplit
func ResolveAddresses(network_ Network, name string) (result cm.Result[ResolveAddressStream, ResolveAddressStream, ErrorCode]) {
	network0 := cm.Reinterpret[uint32](network_)
	name0, name1 := cm.LowerString(name)
	wasmimport_ResolveAddresses((uint32)(network0), (*uint8)(name0), (uint32)(name1), &result)
	return
}
