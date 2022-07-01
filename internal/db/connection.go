package db

import (
	"better-auth/configs"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	uriPass = configs.LoadConfig("pass")
	uriPem  = configs.LoadConfig("pem")
)

type DB struct {
	Client         *mongo.Client
	Ctx            context.Context
	DBOptions      *options.ClientOptions
	DatabaseName   string
	CollectionName string
}

func (d *DB) GetCollection(collectionName string) *mongo.Collection {
	return d.Client.Database(d.DatabaseName).Collection(collectionName)
}

func InitDB(ctx context.Context) DB {
	client := DB{
		Ctx:            ctx,
		DatabaseName:   configs.DBName,
		CollectionName: configs.CollectionName,
	}
	client.ConnectDB("pem")
	return client
}

/**
CONNECT DATABASE SWITCH B/W PASSWORD AND X.509 AUTH
*/
func (d *DB) ConnectDB(authFlag string) {
	var err error
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	switch authFlag {
	case "pem":
		d.DBOptions = options.Client().ApplyURI(uriPem).SetServerAPIOptions(serverAPIOptions)
	case "pass":
		d.DBOptions = options.Client().ApplyURI(uriPass).SetServerAPIOptions(serverAPIOptions)
	}
	d.Client, err = mongo.Connect(d.Ctx, d.DBOptions)
	if err != nil {
		log.Fatalf("Database connection failed:\nreason:%v", err.Error())
	}

	log.Printf("Database connected Successfully!")
}

/**
PING  DATABASE
*/
func (d *DB) PingDB(readP string) {
	log.Printf("Pinging Mongo Client with context %v ", d.Ctx)
	switch readP {
	case "primary":
		err := d.Client.Ping(d.Ctx, readpref.Primary())
		if err != nil {
			log.Fatalf("Error while pinging Primary %v", err)
		}
		log.Printf("Ping Primary Success with context %v ", d.Ctx)
	case "nearest":
		err := d.Client.Ping(d.Ctx, readpref.Nearest())
		if err != nil {
			log.Fatalf("Error while pinging Nearest %v", err)
		}
		log.Printf("Ping Nearest Success with context %v ", d.Ctx)
	case "secondary":
		err := d.Client.Ping(d.Ctx, readpref.Secondary())
		if err != nil {
			log.Fatalf("Error while pinging Secondary %v", err)
		}
		log.Printf("Ping Secondary Success with context %v ", d.Ctx)
	}

}
func (d *DB) PingAll() {
	d.PingDB("primary")
	d.PingDB("secondary")
	d.PingDB("nearest")
}

/**
DISCONNECT DATABASE
*/
func (d *DB) DisconnectDB() {
	err := d.Client.Disconnect(d.Ctx)
	if err != nil {
		log.Fatalf("Database disconnect failed:\nreason:%v", err.Error())
	}
	log.Printf("Database disconnected Successfully!")
}

/**
QUIT DATABASE
*/
func (d *DB) QuitDB() {
	defer func(cdb *DB) {
		err := cdb.Client.Disconnect(d.Ctx)
		if err != nil {
			log.Fatalf("Database disconnect failed:\nreason:%v", err.Error())
		}
		log.Printf("Database disconnected Successfully!")
	}(d)
}
