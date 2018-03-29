package gohash

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
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
