package main

import (
	"fmt"
	"time"

	"github.com/syrinsecurity/cryptotime"
)

func main() {

	fmt.Println("Unix:", cryptotime.Now().Unix())
	fmt.Println("Works exactly the same to the time package", cryptotime.Now().Add(time.Hour).Unix())
}
