package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lauracor5/twittor.git/bd"
	"github.com/lauracor5/twittor.git/models"
)

func Registro(ctx context.Context) models.ResponseApi {
	var modelUser models.Usuario
	var responseApi models.ResponseApi

	responseApi.Status = 400

	fmt.Println("Enrte a registro")

	body := ctx.Value((models.Key("body"))).(string)
	err := json.Unmarshal([]byte(body), &modelUser)
	if err != nil {
		responseApi.Message = "Error al leer los datos " + err.Error()
		fmt.Println(responseApi.Message)
		return responseApi
	}

	if len(modelUser.Email) == 0 {
		responseApi.Message = "Debe especificar el email"
		fmt.Println(responseApi.Message)
		return responseApi
	}

	if len(modelUser.Password) <= 6 {
		responseApi.Message = "La contraseña debe tener al menos 6 caracteres"
		fmt.Println(responseApi.Message)
		return responseApi
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(modelUser.Email)

	if encontrado {
		responseApi.Message = "Ya existe un usuario registrado con ese email"
		fmt.Println(responseApi.Message)
		return responseApi
	}
	_, status, err := bd.InsertoRegistro(modelUser)
	if err != nil {
		responseApi.Message = "Ocurrió un error al intentar realizar el registro de usuario " + err.Error()
		fmt.Println(responseApi.Message)
		return responseApi
	}

	if status == false {
		responseApi.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println(responseApi.Message)
		return responseApi
	}

	responseApi.Status = 200
	responseApi.Message = "Usuario registrado correctamente"
	fmt.Println(responseApi.Message)
	return responseApi
}
