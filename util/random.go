package util

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"

	"github.com/mmcloughlin/spherand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomFloat generates float64
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomSQLNullFloat64 generates sql.NullFloat64 value
func RandomSQLNullFloat64() sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: RandomFloat(1.00, 1.00),
		Valid:   true,
	}
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomBool generates boolean
func RandomBool() bool {
	data := []bool{true, false}
	n := len(data)
	return data[rand.Intn(n)]
}

// RandomName generates a random Name
func RandomName() string {
	return RandomString(6)
}

// RandomShortName generates a random ShortName
func RandomShortName() string {
	return RandomString(3)
}

func RandomLongLat() (float64, float64) {
	lat, lng := spherand.Geographical()
	return lat, lng
}
