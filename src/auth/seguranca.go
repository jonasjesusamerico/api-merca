package auth

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(senha string) (cript []byte, err error) {
	cript, err = bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return
}

// VerificarSenha compara uma senha e um hash e retorna se elas são iguais
func VerificarSenha(senhaComHash, senhaString string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
	return
}
