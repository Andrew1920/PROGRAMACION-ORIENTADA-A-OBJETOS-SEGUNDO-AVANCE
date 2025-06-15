package utils

import "golang.org/x/crypto/bcrypt"

// Con esta función, genero un hash bcrypt de la contraseña para guardarla de forma segura.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Aquí comparo de forma segura la contraseña que me envían con el hash que tengo guardado.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
