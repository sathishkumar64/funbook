package durable

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// ConnectionInfo is the info of the postgres
type ConnectionInfo struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// Database is wrapped struct of *sql.DB
type Database struct {
	db *mongo.Client
}

// OpenDatabaseClient generate a database client
func OpenDatabaseClient(ctx context.Context, c *ConnectionInfo) *mongo.Client {
	//connStr := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Name)
	connStr := fmt.Sprintf("mongodb://%s:%s%s", c.Host, c.Port,c.Name)
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if err := db.Ping(ctx, nil); err != nil {
		log.Fatal(fmt.Errorf("\nFail to connect the database.\nPlease make sure the connection info is valid %#v", c))
		return nil
	}
	return db
}

// WrapDatabase create a *Database
func WrapDatabase(db *mongo.Client) *Database {
	return &Database{db: db}
}

// Close the *sql.DB
func (d *Database) Close(ctx context.Context) error {
	return d.db.Disconnect(ctx)
}