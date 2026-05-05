package main

import (
	"fmt"

	"github.com/abg216te/tgbt/internal/config"
)

// Bot
func main() {
	сfg, err := config.Load()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	
	token := сfg.Token

	fmt.Println("Token:", token)
}
