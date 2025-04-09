package routers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lauracor5/twittor.git/bd"
	"github.com/lauracor5/twittor.git/models"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.ResponseApi {

	var response models.ResponseApi
	response.Status = 400

	fmt.Println("Entré en verperfil")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		response.Message = "El ID del usuario es requerido"
		return response
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		response.Message = "Error al intentar buscar el perfil " + err.Error()
		return response
	}

	responseJson, err := json.Marshal(perfil)
	if err != nil {
		response.Status = 500
		response.Message = "Error al formatearlos datos del usuario como json " + err.Error()
		return response
	}

	response.Status = 200
	response.Message = string(responseJson)
	return response
}
