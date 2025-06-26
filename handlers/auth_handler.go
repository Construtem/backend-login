// Aqui deberia ir toda la logica de autenticacion, 

package handlers

import (
	"backend-login/db"
	"backend-login/services"
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// VerifyToken maneja la verificación del token de autenticación enviada por el frontend
func VerifyToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token faltante"})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
		return
	}

	idToken := parts[1]
	token, err := services.FirebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Se obtiene el correo electrónico del token
	email := token.Claims["email"].(string)

	// Verificar que el correo electrónico tenga el dominio @utem.cl
	if !strings.HasSuffix(email, "@utem.cl") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Correo no autorizado"})
		return
	}

	// Consultar la base de datos para obtener el rol del usuario
	// OJO : Asegúrate de que la tabla 'usuarios' y la columna 'rol' existan en la base de datos
	var rol string
	err = db.DB.QueryRow("SELECT r.nombre FROM usuarios u JOIN rols r ON u.rol_id = r.id WHERE u.correo = $1", email).Scan(&rol)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusForbidden, gin.H{"error": "Usuario no registrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error consultando rol"})
		}
		return
	}

	// Entregar el rol al frontend
	c.JSON(http.StatusOK, gin.H{"rol": rol})
}