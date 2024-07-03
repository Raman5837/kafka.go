package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Utility Function To Get env Value
func GetEnv(key string) string {

	// Load .env File
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error While Loading .env File", err)
	}
	// Return The Value Of The Variable
	return os.Getenv(key)
}
