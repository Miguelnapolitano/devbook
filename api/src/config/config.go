package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String to connect with MySQL
	StringDBConnection = ""
	//Port where API will be running
	Port = 0
	//Key to sign token
	SecretKey []byte

)

//Load will init environments
func Load(){
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv(("API_PORT")))
	if err != nil {
		Port = 9000
	}

	StringDBConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_KEY"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}