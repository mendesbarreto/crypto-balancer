package environment

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadVariables() {
	if err := godotenv.Load(".env/prod"); err != nil {
		fmt.Println(err.Error())
	}
}
