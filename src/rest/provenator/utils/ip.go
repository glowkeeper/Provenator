package utils

import (
	"net"
	"net/http"
)

const (
	// XForwardedFor - forwarded IP
	XForwardedFor = "X-Forwarded-For"
	//XRealIP - Real IP
	XRealIP       = "X-Real-IP"
)

//IPAddress - get the IP
func IPAddress(r *http.Request) string {
	remoteAddr := r.RemoteAddr

	if ip := r.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = r.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
