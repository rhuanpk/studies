package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const bytesCount = 16

func md5Salt(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}

func randSalt() string {
	bytes := make([]byte, bytesCount)

	_, err := rand.Read(bytes)
	if err != nil {
		return strconv.Itoa(int(time.Now().UnixNano()))
	}

	return hex.EncodeToString(bytes)
}

// go playground: https://go.dev/play/p/tivu-SNwpUg
func main() {
	println("md5 salt:", md5Salt(`input`))
	println("rand salt:", randSalt())
}
