package main

import (
	"ShamanEstBanan-GophKeeper-server/internal/app"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, this is keeper!")
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(a.Run())
}
