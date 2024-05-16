package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Global struct {
	Database Database
	Jwt      Jwt
	Bcrypt   Bcrypt
	S3       S3
}

type Bcrypt struct {
	Salt int `env:"BCRYPT_SALT,required"`
}

type Jwt struct {
	Secret string `env:"JWT_SECRET,required"`
}

type Database struct {
	Username string `env:"DB_USERNAME,required"`
	Password string `env:"DB_PASSWORD,required"`
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	DbName   string `env:"DB_NAME,required"`
	Params   string `env:"DB_PARAMS,required"`
}

type S3 struct {
	ID         string `env:"S3_ID,required"`
	SecretKey  string `env:"S3_SECRET_KEY,required"`
	BucketName string `env:"S3_BUCKET_NAME,required"`
	Region     string `env:"S3_REGION,required"`
}

func NewConfig() *Global {
	var c Global
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
