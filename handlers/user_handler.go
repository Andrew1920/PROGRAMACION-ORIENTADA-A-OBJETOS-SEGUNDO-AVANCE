package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"tienda/models"
	"tienda/utils"

	"github.com/google/uuid"
)

// Defino un mapa global para guardar mis usuarios en memoria para este proyecto.
// La clave es el 'username' para búsquedas rápidas.
var users = make(map[string]models.User)

// RegisterHandler se encarga de crear nuevas cuentas de usuario.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Primero, decodifico el JSON que me llega en el cuerpo de la petición.
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}

	// Un paso de seguridad crucial: genero un hash de la contraseña
	// para no guardarla nunca en texto plano.
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error al procesar la contraseña", http.StatusInternalServerError)
		return
	}
	// Reemplazo la contraseña original con el hash seguro.
	user.Password = hashedPassword

	// Asigno un ID único universal a mi nuevo usuario.
	user.ID = uuid.NewString()
	users[user.Username] = user
	w.WriteHeader(http.StatusCreated)
}

// LoginHandler verifica las credenciales de un usuario.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}

	// Aquí verifico si el usuario que intenta iniciar sesión realmente existe en mi mapa.
	user, ok := users[credentials.Username]
	if !ok || !utils.CheckPasswordHash(credentials.Password, user.Password) {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}

	// Si todo es correcto, simplemente respondo con un mensaje de éxito.
	log.Printf("Inicio de sesión exitoso para el usuario: %s", user.Username)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Inicio de sesión exitoso"})
}
