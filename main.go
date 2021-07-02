package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 01
	max := 1000
	randNum := rand.Intn(max-min+100) + min
	fmt.Println(randNum)
}
