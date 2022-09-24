package http

import (
	"math/rand"
	"time"

	"github.com/blackcrw/wprecon/internal/wordlist"
)

// randomUserAgent :: this function "generates" user-agents randomly.
func randomUserAgent() string {
	rand.Seed(time.Now().Unix())

	return wordlist.UserAgents[rand.Intn(len(wordlist.UserAgents))]
}