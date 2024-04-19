package repo

import (
	"context"
	"database-example/model"
	"log"

	// NoSQL: module containing Neo4J api client
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type FollowerRepository struct {
	driver neo4j.DriverWithContext
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment and creates a keyspace
func New(logger *log.Logger) (*FollowerRepository, error) {
	uri := "bolt://localhost:7687"
	user := "neo4j"
	pass := "neo4jteam25"
	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}

	logger.Println("Connected to Neo4j database")
	return &FollowerRepository{
		driver: driver,
		logger: logger,
	}, nil
}

func (fr *FollowerRepository) CheckConnection() {
	ctx := context.Background()
	err := fr.driver.VerifyConnectivity(ctx)
	if err != nil {
		fr.logger.Panic(err)
		return
	}

	fr.logger.Printf(`Neo4J server address: %s`, fr.driver.Target().Host)
}

func (mr *FollowerRepository) CreateUser(user *model.User) error {
	ctx := context.Background()
	session := mr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// ExecuteWrite for write transactions (Create/Update/Delete)
	savedPerson, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"CREATE (u:User {id: $id, username: $username}) RETURN u.username + ', from node ' + id(u)",
				map[string]interface{}{"id": user.ID, "username": user.Username})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		mr.logger.Println("Error inserting Person:", err)
		return err
	}
	mr.logger.Println(savedPerson.(string))
	return nil
}

func (mr *FollowerRepository) CloseDriverConnection(ctx context.Context) {
	mr.driver.Close(ctx)
}
