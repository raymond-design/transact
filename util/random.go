package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// the following are utility functions for generating random test data
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generate random integer
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generate random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generate random name
func RandomOwner() string {
	return RandomString(6)
}

// generate random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// generate random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// generate random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
