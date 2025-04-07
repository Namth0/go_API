package models

import "errors"

// Erreurs de validation
var (
	ErrUsernameRequired = errors.New("le nom d'utilisateur est requis")
	ErrEmailRequired    = errors.New("l'email est requis")
	ErrPasswordRequired = errors.New("le mot de passe est requis")
)

// Erreurs d'authentification
var (
	ErrInvalidCredentials = errors.New("identifiants invalides")
)

// Erreurs de base de données
var (
	ErrUserAlreadyExists  = errors.New("ce nom d'utilisateur existe déjà")
	ErrEmailAlreadyExists = errors.New("cette adresse email est déjà utilisée")
	ErrUserNotFound       = errors.New("utilisateur non trouvé")
	ErrDatabaseConnection = errors.New("erreur de connexion à la base de données")
	ErrDatabaseOperation  = errors.New("erreur lors de l'opération sur la base de données")
)
