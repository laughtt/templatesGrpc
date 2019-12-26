package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/laughtt/templatesGrpc/service+sql/pkg/server"
	v1 "github.com/laughtt/templatesGrpc/service+sql/pkg/service"
)

const (
	database   = "test"
	name       = "db"
	host       = "database:3306"
	apiVersion = "v1"
	password   = "password"
	port       = "8080"
	username   = "db"
)

//Config Database connection
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string
	// DB Datastore parameters selsction
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

//AUTHservice Crea el servicio para la conexion de la base de datos

func RunServer() error {
	ctx := context.Background()
	cfg := Config{
		GRPCPort:            port,
		DatastoreDBHost:     host,
		DatastoreDBPassword: password,
		DatastoreDBSchema:   database,
		DatastoreDBUser:     username,
	}

	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)

	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	err = db.Ping()
	for db.Ping() != nil{
		log.Println("mysqld is not alive")
		time.Sleep(5 * time.Second)	
	}
	db, err = sql.Open("mysql", dsn)
	v1API := v1.NewAuthServiceServer(db)

	if err != nil {
		return fmt.Errorf("failed to coonect: %v", err)
	}

	defer db.Close()

	return grpc.RunServer(ctx, v1API, port)
}

func main() {
	fmt.Println(RunServer())
}
