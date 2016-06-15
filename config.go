package config

import "os"

// Global variables
var (
	Mode string
	Port string
)

// Setup config
func Setup(port string, mode string) {
	if mode != "" {
		// explicitly set the mode
		Mode = mode
		Port = port
	} else {
		dbURL := os.Getenv("DATABASE_URL")

		if dbURL == "" {
			// dev mode
			Mode = "dev"
			Port = port
		} else {
			Mode = "prod"
			Port = ":" + os.Getenv("PORT")
		}
	}
}
