package pkg

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenShortUrl(lenght int) string {
	return GenString(lenght, charset)
}

func GenString(lenght int, charset string) string {
	res := make([]byte, lenght)
	for i := range res {
		res[i] = charset[seededRand.Intn(lenght)]
	}
	return string(res)
}
