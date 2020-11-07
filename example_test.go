package snippets

import (
	"fmt"
	"log"
)

func ExampleCryptoRandInt() {
	x, err := RandInt(1, 4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(x)
}
