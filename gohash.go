package gohash

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"strings"
)

func Hash80(s string) string {
	full := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", full[:10])
}

// hash63 returns a non-negative 63-bit integer hash
// same rationale of using int64 to store non-negative 63-bit as in math/rand.Int63
func Hash63(s string) int64 {
	full := md5.Sum([]byte(s))
	eight := full[:8]
	eight[0] = eight[0] & 127
	u := binary.BigEndian.Uint64(eight)
	if u > 9223372036854775807 {
		panic(fmt.Sprintf("gohash.Hash63 u out of range. u=%d, s=%s", u, s))
	}
	return int64(u)
}

// strip 'http://' or 'https://' or '//' from url
// also remove the trailing slash
func StripHttp(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasSuffix(s, "/") {
		s = s[:len(s)-1]
	}
	if strings.HasPrefix(s, "http://") {
		s = s[7:]
	}
	if strings.HasPrefix(s, "https://") {
		s = s[8:]
	}
	if strings.HasPrefix(s, "//") {
		s = s[2:]
	}
	return s
}

func IsHttpStripped(s string) bool {
	return !(strings.HasSuffix(s, "/") ||
		strings.HasPrefix(s, "http://") ||
		strings.HasPrefix(s, "https://") ||
		strings.HasPrefix(s, "//"))
}
