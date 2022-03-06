package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpha string = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

func RandomString(n int) string {
	var s strings.Builder
	k := 26

	for i := 0; i < n; i++ {
		c := alpha[rand.Intn(k)]
		s.WriteByte(c)
	}

	return s.String()
}

func RandomOwner() string {
	return RandomString(10)
}

func RandomMoney() int64 {
	return RandomInt(1, 100)
}

func RandomCurrency() string {
	cur := []string{"EUR", "USD", "CAD"}
	n := len(cur)
	return cur[rand.Intn(n)]
}
