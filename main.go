package main

import (
    "fmt"
    "math/rand"
    "time"
)

var randomWords = []string{
    "error", "critical", "warning", "info", "information", "debug", "trace", "error", "critical", "panic",
}

func main() {
    rand.Seed(time.Now().UnixNano())
    now := time.Now()

    for {
        fmt.Println(now,generateRandomText())
	time.Sleep(1 * time.Second)
    }
}

func generateRandomText() string {
    randomIndex := rand.Intn(len(randomWords))
    return randomWords[randomIndex]
}
