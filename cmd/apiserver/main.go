package main

import (
	"fmt"
	"log"

	"github.com/RomanKazakevich/group-materials/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start main")
}
