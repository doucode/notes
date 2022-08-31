package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(999)
	fmt.Println(i)
}
