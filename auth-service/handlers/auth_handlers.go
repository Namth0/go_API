package handlers

import (
	"auth-service/models"
	"auth-service/repository"
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
	}
}

// Health vérifie l'état du service d'authentification
func (h *AuthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "auth-service",
	})
}

// Login authentifie un utilisateur
func (h *AuthHandler) Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userRepo.FindByUsernameOrEmail(loginData.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Identifiants invalides"})
		return
	}

	if user.Password != loginData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Identifiants invalides"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Connexion réussie",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}

// GetAllUsers récupère tous les utilisateurs
func (h *AuthHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des utilisateurs"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID récupère un utilisateur par son ID
func (h *AuthHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser crée un nouvel utilisateur
func (h *AuthHandler) CreateUser(c *gin.Context) {
	// Afficher les headers de la requête pour déboguer CORS ou problèmes de content-type
	log.Println("Headers de la requête:", c.Request.Header)

	// Lire le body brut pour debugging
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	// Restaurer le body pour que c.ShouldBindJSON fonctionne
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	// Log du body pour voir ce qui est reçu exactement
	log.Println("Body de la requête:", string(bodyBytes))

	// Créer une struct distincte pour les créations d'utilisateurs
	type UserCreate struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userCreate UserCreate
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		log.Printf("Erreur bind JSON: %v", err)
		log.Printf("Body reçu: %s", string(bodyBytes))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide: " + err.Error()})
		return
	}

	// Transférer les données à l'objet User
	user := models.User{
		Username: userCreate.Username,
		Email:    userCreate.Email,
		Password: userCreate.Password,
		Role:     "user", // Rôle par défaut
	}

	// Debug log pour voir ce qui est reçu
	log.Printf("Demande d'inscription reçue: nom_utilisateur=%s, email=%s, longueur_mot_de_passe=%d",
		user.Username, user.Email, len(user.Password))

	// Valider les champs requis
	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userRepo.Create(&user)
	if err != nil {
		// Gérer les types d'erreurs spécifiques
		if strings.Contains(err.Error(), "uni_users_username") ||
			strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "Ce nom d'utilisateur existe déjà"})
			return
		} else if strings.Contains(err.Error(), "email") {
			c.JSON(http.StatusConflict, gin.H{"error": "Cette adresse email est déjà utilisée"})
			return
		}

		log.Printf("Erreur de base de données: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de base de données: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// ValidateToken valide un jeton d'authentification
func (h *AuthHandler) ValidateToken(c *gin.Context) {
	var tokenData struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&tokenData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dans une vraie application, vous valideriez un jeton JWT
	// Pour cet exemple, nous allons simplement retourner une réponse simulée

	c.JSON(http.StatusOK, gin.H{
		"valid": true,
		"user": gin.H{
			"id":       uuid.New(), // Utiliser un UUID approprié
			"username": "admin",
			"role":     "admin",
		},
	})
}
