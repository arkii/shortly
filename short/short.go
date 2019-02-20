package short

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

var seeds = strings.Split("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

func New(url string) []string {
	salt := "arkii"
	hex := fmt.Sprintf("%x", md5.Sum([]byte(salt+url)))
	keys := make([]string, 4)
	for i := 0; i < 4; i++ {
		v, _ := strconv.ParseInt(hex[i*8:i*8+8], 16, 0)
		hexLong := v & 0x3fffffff
		key := ""
		for j := 0; j < 6; j++ {
			key += seeds[0x0000003D&hexLong]
			hexLong >>= 5
		}
		keys[i] = key
	}
	// fmt.Println(keys)
	return keys
}
