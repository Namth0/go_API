package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Temps de début
		startTime := time.Now()

		// Traitement de la requête
		c.Next()

		// Temps de fin
		endTime := time.Now()

		// Informations de la requête
		latency := endTime.Sub(startTime)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		// Log au format: [TEMPS] | IP | MÉTHODE | STATUS | LATENCE | CHEMIN
		fmt.Printf("[%v] | %s | %s | %d | %v | %s\n",
			endTime.Format("2006-01-02 15:04:05"),
			clientIP,
			method,
			statusCode,
			latency,
			path,
		)
	}
}
