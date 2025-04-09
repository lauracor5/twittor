package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/lauracor5/twittor.git/bd"
	"github.com/lauracor5/twittor.git/models"
)

var Email string
var IDUsuario string

func ProcesarToken(token string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tokenParserClaims, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		// Rutina que Cheqeuea contra la BD
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return &claims, encontrado, claims.ID.Hex(), nil

	}

	if !tokenParserClaims.Valid {
		return &claims, false, string(""), errors.New("token invalid")
	}

	return &claims, true, claims.ID.Hex(), nil

}
