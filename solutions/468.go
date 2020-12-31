package solutions

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if isIPv4(IP) {
		return "IPv4"
	}

	if isIPv6(IP) {
		return "IPv6"
	}

	return "Neither"
}

func isIPv4(IP string) bool {
	parts := strings.Split(IP, ".")

	if len(parts) != 4 {
		return false
	}

	for _, value := range parts {
		if number, ok := strconv.Atoi(value); ok != nil || number < 0 || number > 255 {
			return false
		} else if strconv.Itoa(number) != value {
			return false
		}
	}

	return true
}

func isIPv6(IP string) bool {
	parts := strings.Split(IP, ":")

	if len(parts) != 8 {
		return false
	}

	for _, value := range parts {
		if !isValidIPv6Part(value) {
			return false
		}
	}

	return true
}

func isValidIPv6Part(s string) bool {
	if len(s) == 0 || len(s) > 4 {
		return false
	}

	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}

	return true
}
