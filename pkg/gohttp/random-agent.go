package gohttp

import (
	"math/rand"
	"time"

	"github.com/blackbinn/wprecon/pkg/wordlist"
)

// randomuseragent : this function "generates" user-agents randomly.
func randomuseragent() string {
	timeUnix := time.Now().Unix()

	rand.Seed(timeUnix)
	randomValue := rand.Intn(len(wordlist.UserAgents))

	return wordlist.UserAgents[randomValue]
}
