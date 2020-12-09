package gohttp

import (
	"math/rand"
	"time"

	"github.com/blackcrw/wprecon/pkg/wordlist"
)

// RandomUserAgent : this function "generates" user-agents randomly.
func RandomUserAgent() string {
	timeUnix := time.Now().Unix()

	rand.Seed(timeUnix)
	randomValue := rand.Intn(len(wordlist.UserAgents))

	return wordlist.UserAgents[randomValue]
}
