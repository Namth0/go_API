package service

import (
	"errors"
	"mon-api/internal/models"
	"mon-api/internal/repository"
	"strings"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() []models.User {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id string) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("utilisateur non trouvé")
	}
	return user, nil
}

func (s *UserService) GetByUsername(username string) (*models.User, error) {
	return s.repo.GetByUsername(username)
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	// Vérifier si le nom d'utilisateur existe déjà
	existingUser, _ := s.repo.GetByUsername(user.Username)
	if existingUser != nil {
		return errors.New("ce nom d'utilisateur est déjà pris")
	}

	// Vérifier si l'email existe déjà
	existingUser, _ = s.repo.GetByEmail(user.Email)
	if existingUser != nil {
		return errors.New("cet email est déjà utilisé")
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	return s.repo.Create(user)
}

func (s *UserService) Update(id string, user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	// Vérifier si l'utilisateur existe
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("utilisateur non trouvé")
	}

	// Vérifier si le nom d'utilisateur est pris par un autre utilisateur
	if user.Username != existingUser.Username {
		otherUser, err := s.repo.GetByUsername(user.Username)
		if err == nil && otherUser.ID != existingUser.ID {
			return errors.New("ce nom d'utilisateur est déjà pris")
		}
	}

	// Vérifier si l'email est pris par un autre utilisateur
	if user.Email != existingUser.Email {
		otherUser, err := s.repo.GetByEmail(user.Email)
		if err == nil && otherUser.ID != existingUser.ID {
			return errors.New("cet email est déjà utilisé")
		}
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	return s.repo.Update(id, user)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *UserService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("nom d'utilisateur ou mot de passe incorrect")
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, errors.New("nom d'utilisateur ou mot de passe incorrect")
	}

	return user, nil
}
