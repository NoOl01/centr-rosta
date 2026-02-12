package config

import (
	"log"
	"os"
	"strconv"
)

type envStorage struct {
	ServerPort string
	DbUser     string
	DbPass     string
	DbName     string
	DbHost     string
	DbPort     string
	JwtSecret  string
	Mail       string
	MailPass   string
	SmtpHost   string
	SmtpPort   string
	RedisHost  string
	RedisPort  string
	RedisPass  string
	Debug      bool
}

var Env *envStorage

func LoadEnv() {
	Env = &envStorage{}
	Env.ServerPort = os.Getenv("SERVER_PORT")
	Env.DbUser = os.Getenv("DB_USER")
	Env.DbPass = os.Getenv("DB_PASS")
	Env.DbName = os.Getenv("DB_NAME")
	Env.DbHost = os.Getenv("DB_HOST")
	Env.DbPort = os.Getenv("DB_PORT")
	Env.JwtSecret = os.Getenv("JWT_SECRET")
	Env.Mail = os.Getenv("MAIL")
	Env.MailPass = os.Getenv("MAIL_PASS")
	Env.SmtpHost = os.Getenv("SMTP_HOST")
	Env.SmtpPort = os.Getenv("SMTP_PORT")
	Env.RedisHost = os.Getenv("REDIS_HOST")
	Env.RedisPort = os.Getenv("REDIS_PORT")
	Env.RedisPass = os.Getenv("REDIS_PASS")

	debugStr := os.Getenv("DEBUG")
	debug, err := strconv.ParseBool(debugStr)
	if err != nil {
		log.Fatalln(err)
	}
	Env.Debug = debug
}
