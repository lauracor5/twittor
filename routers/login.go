package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lauracor5/twittor.git/bd"
	"github.com/lauracor5/twittor.git/jwt"
	"github.com/lauracor5/twittor.git/models"
)

func Login(ctx context.Context) models.ResponseApi {
	var requestUser models.Usuario
	var responseUser models.ResponseApi
	responseUser.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &requestUser)

	if err != nil {
		responseUser.Message = "Usuario y o contraseña invalido" + err.Error()
		return responseUser
	}

	if len(requestUser.Email) == 0 {
		responseUser.Message = "El email es requerido"
		return responseUser
	}

	userData, isExist := bd.IntentoLogin(requestUser.Email, requestUser.Password)
	if !isExist {
		responseUser.Message = "Usuario y o contraseña invalido"
		return responseUser
	}

	jwtKey, err := jwt.GeneroJsonWebToken(ctx, userData)
	if err != nil {
		responseUser.Message = "Error al intentar generar el token " + err.Error()
		return responseUser
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		responseUser.Message = "Ocurrio un error al intentar formatear el token a json " + err2.Error()
		return responseUser
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}

	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	responseUser.Status = 200
	responseUser.Message = string(token)
	responseUser.CustomRepsonse = res

	return responseUser
}
