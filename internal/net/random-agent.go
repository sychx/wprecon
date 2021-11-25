package net

import (
	"math/rand"
	"time"

	"github.com/blackcrw/wprecon/internal/wordlist"
)

// random_user_agent :: this function "generates" user-agents randomly.
func random_user_agent() string {
	var time_unix = time.Now().Unix()

	rand.Seed(time_unix)
	var random_value = rand.Intn(len(wordlist.UserAgents))

	return wordlist.UserAgents[random_value]
}