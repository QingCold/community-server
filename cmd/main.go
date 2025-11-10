package main

import (
	"fmt"
	"log"

	"github.com/QingCold/community-server/cmd/internal/api"
)

func main() {
	r := api.NewRouter()
	log.Println("Server started on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("vim-go")
}
