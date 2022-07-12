package configs

import "os"

var (
	Host           = GetEnv("HOST")
	Port           = "8080"
	DBName         = GetEnv("DB_NAME")
	CollectionName = GetEnv("COLLECTION_NAME")
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func LoadConfig(flag string) string {
	var uri string

	if flag == "pass" {
		/* statement(s) will execute if the boolean expression is true */
		uri := GetEnv("MONGO_URI")
		return uri
	}
	if flag == "pem" {
		pemKey := GetEnv("MONGO_PEM_FILE")
		uri = "mongodb+srv://cluster0.rh7ld.mongodb.net/?" +
			"authSource=%24external&authMechanism=MONGODB-X509" +
			"&retryWrites=true&w=majority" +
			"&tlsCertificateKeyFile="
		uri = uri + pemKey
		return uri
	}
	if flag == "url" {
		if Host == "" {
			uri = ":" + Port
		} else {
			uri = Host + ":" + Port
		}
	}
	return uri
}
