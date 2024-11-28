package middleware

import (
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/sharvatic/BookMyHotel/firebase"
	"net/http"
	"strings"
)

func VerifyToken(requiredRole string) gin.HandlerFunc {
	// Get the token from the Authorization header
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
			c.Abort()
			return
		}

		// Strip the "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Verify the token
		token, err := firebase.AuthClient.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract the role from the claims
		role, ok := token.Claims["role"].(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found in token claims"})
			c.Abort()
			return
		}

		// Verify that the role matches the required role
		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		// Set user ID and role in the context
		c.Set("userID", token.UID)
		c.Set("role", role)

		c.Next() // Proceed to the next handler
	}
}

func SetClaims(c *gin.Context){
	var requestBody struct {
		Role string `json:"role"`
	}

	// Bind JSON payload from request
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Retrieve token from Authorization header
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	// Verify ID token
	client := firebase.AuthClient // Use the initialized AuthClient
	token, err := client.VerifyIDToken(context.Background(), tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Set custom claims
	claims := map[string]interface{}{
		"role": requestBody.Role,
	}
	if err := client.SetCustomUserClaims(context.Background(), token.UID, claims); err != nil {
		log.Printf("Failed to set claims: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set claims"})
		return
	}

	// Successfully set claims
	c.JSON(http.StatusCreated, gin.H{"message": "Claims set successfully"})

}

