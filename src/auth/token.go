package auth

import (
	"api-merca/src/config"
	"api-merca/src/model/enum"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CriarToken(usuarioID uint64, isCustmizavel bool, bancoDados enum.BancoDados) (tokenNew string, err error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	permissoes["isCustmizavel"] = isCustmizavel
	permissoes["bancoDados"] = bancoDados
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	tokenNew, err = token.SignedString([]byte(config.SecretKey))
	return
}

func ValidarToken(r *gin.Context) (erro error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return
	}

	erro = errors.New("token inválido")
	return
}

func ExtrairUsuarioID(r *gin.Context) (usuarioID uint64, erro error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro = strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return
		}

		return
	}

	erro = errors.New("token inválido")
	return
}

func ExtrairBanco(r *gin.Context) (bancoDados string, erro error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		bancoDados = fmt.Sprintf("%s", permissoes["bancoDados"])
		if erro != nil {
			return
		}

		return
	}
	erro = errors.New("token inválido")
	return
}

func ExtrairIsCustomizavel(r *gin.Context) (isCustmizavel bool, erro error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		isCustmizavel = permissoes["isCustmizavel"].(bool)
		if erro != nil {
			return
		}
		return
	}
	erro = errors.New("token inválido")
	return
}

func extrairToken(r *gin.Context) (token string) {
	token = r.GetHeader("Authorization")

	if len(strings.Split(token, " ")) != 2 {
		return
	}
	token = strings.Split(token, " ")[1]
	return
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
