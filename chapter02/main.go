package main

import (
	"log"
	"os"
)

// init在main之前调用
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
