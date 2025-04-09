package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lauracor5/twittor.git/jwt"
	"github.com/lauracor5/twittor.git/models"
	"github.com/lauracor5/twittor.git/routers"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.ResponseApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + "> " +
		ctx.Value(models.Key("method")).(string))

	var res models.ResponseApi
	res.Status = 400

	isOk, statusCode, message, _ := validateAuthorization(ctx, request)

	if !isOk {
		res.Status = statusCode
		res.Message = message
		return res
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)
		case "login":
			return routers.Login(ctx)
		}

	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "verperfil":
			return routers.VerPerfil(request)
		}

	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		}

	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		}
	}

	res.Message = "Method invalid"
	return res

}

func validateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 400, "Token requerido", models.Claim{}
	}

	claim, todoOK, msg, err := jwt.ProcesarToken(token, ctx.Value(models.Key("jwtSign")).(string))

	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim

}
