package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789" +
	"-_"

var randSeed *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randSeed.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return RandStringWithCharset(length, charset)
}

func main() {
	for {
		url := "https://hackmd.io/" + RandString(22)
		response, err := http.Get(url)
		if err != nil {
			log.Fatalf("error on get url %s: %v", url, err)
		}
		mark := ""
		if response.StatusCode == 200 {
			mark = "<<<"
		}
		log.Println(url, response.StatusCode, mark)
	}
}
