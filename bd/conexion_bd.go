package bd

import (
	"context"
	"fmt"

	"github.com/lauracor5/twittor.git/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabseName string

func ConectarBD(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error al conectar a la BD" + err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error al hacer ping a la BD" + err.Error())
	}

	fmt.Println("Conexión exitosa a la BD")
	MongoCN = client
	DatabseName = ctx.Value(models.Key("databse")).(string)

	return nil
}

func BaseConectada() bool {
	error := MongoCN.Ping(context.Background(), nil)
	return error == nil
}
