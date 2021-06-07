package util

import (
	"errors"
	"net/http"
)

var (
	// ErrIPAM indicates an unsupported IPAM driver was used
	ErrIPAM = errors.New("only the null IPAM driver is supported")
	// ErrBridgeRequired indicates a network bridge was not provided for network creation
	ErrBridgeRequired = errors.New("bridge required")
	// ErrBridgeNotFound indicates that a bridge could not be found
	ErrBridgeNotFound = errors.New("bridge not found")
	// ErrBridgeUsed indicates that a bridge is already in use
	ErrBridgeUsed = errors.New("bridge already in use by Docker")
	// ErrMACAddress indicates an invalid MAC address
	ErrMACAddress = errors.New("invalid MAC address")
	// ErrNoLease indicates a DHCP lease was not obtained from udhcpc
	ErrNoLease = errors.New("udhcpc did not output a lease")
)

func ErrToStatus(err error) int {
	switch {
	case errors.Is(err, ErrIPAM), errors.Is(err, ErrBridgeRequired), errors.Is(err, ErrBridgeNotFound),
		errors.Is(err, ErrBridgeUsed), errors.Is(err, ErrMACAddress):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
