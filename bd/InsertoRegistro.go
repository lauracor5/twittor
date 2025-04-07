package bd

import (
	"context"

	"github.com/lauracor5/twittor.git/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx := context.TODO()
	var responseApi models.ResponseApi

	db := MongoCN.Database(DatabseName)
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)
	responseApi.Message = u.Apellidos + " " + u.Nombre + " " + u.Email + " " + u.Password + " " + u.Avatar + " " + u.Banner + " " + u.Biografia + " " + u.Ubicacion + " " + u.SitioWeb

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil

}
