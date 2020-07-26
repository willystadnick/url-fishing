package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var charset = ""

var randSeed *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randSeed.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	charset = os.Getenv("CHARSET")
	length, err := strconv.Atoi(os.Getenv("LENGTH"))
	if err != nil {
		log.Fatal("Error loading .env length")
	}
	for {
		url := fmt.Sprintf(os.Getenv("URL"), RandStringWithCharset(length, charset))
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
