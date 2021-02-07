package environment

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	appName           = "crypto-balancer"
	envProduction     = "crypto-balancer-prod"
	envTest           = "crypto-balancer-test"
	envDevelopment    = "crypto-balancer-dev"
	envDevelopmentKey = "DEVELOPMENT"
)

func isDevelopment() bool {
	developmentVar := os.Getenv(envDevelopmentKey)

	if value, err := strconv.Atoi(developmentVar); err == nil && value == 1 {
		return true
	}

	return false
}

func LoadVariables() {
	if flag.Lookup("test.v") == nil {
		if isDevelopment() {
			load(envDevelopment)
			fmt.Println("[Crypto-balancer]: DEVELOPMENT environment variables loaded")
		} else {
			load(envProduction)
			fmt.Println("[Crypto-balancer]: PRODUCTION environment variables loaded")
		}
	} else {
		load(envTest)
		fmt.Println("[Crypto-balancer]: loading TESTS environment variables")
	}
}

func load(environmentName string) {
	findWorkSpaceRegex := regexp.MustCompile(`^(.*` + appName + `)`)
	currentExecutionDirectory, _ := os.Getwd()
	rootPath := findWorkSpaceRegex.Find([]byte(currentExecutionDirectory))
	filePath := string(rootPath) + `/.env/` + environmentName
	if err := godotenv.Load(filePath); err != nil {
		log.Fatal(fmt.Printf("Problem loading .env %s file on path: %s", environmentName, filePath))
	}
}
