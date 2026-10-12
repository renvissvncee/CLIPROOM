package service

import (
	"cliproom/models"
	"cliproom/repository"
	"errors"
	"fmt"
	"strings"
)

// Структура, отвечающая за бизнес-лоигку, связанную с пользователем
type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(p_repositroy *repository.UserRepository) *UserService {
	return &UserService{userRepository: p_repositroy}
}

// ХЕЛПЕР-функции

func userIsTaken(find func() (models.User, error), excludeId uint) (bool, error) {
	user, err := find()
	if errors.Is(err, repository.ErrNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return user.ID != excludeId, nil
}

func (s *UserService) ensureLoginFree(login string, excludeId uint) error {
	taken, err := userIsTaken(func() (models.User, error) {
		return s.userRepository.FindUserByLogin(login)
	}, excludeId)

	if err != nil {
		return err
	} else if taken {
		return ErrLoginTaken
	}
	return nil
}

func (s *UserService) ensureEmailFree(email string, excludeId uint) error {
	taken, err := userIsTaken(func() (models.User, error) {
		return s.userRepository.FindUserByEmail(email)
	}, excludeId)

	if err != nil {
		return err
	} else if taken {
		return ErrEmailTaken
	}
	return nil
}

func fetchUser(find func() (models.User, error)) (models.User, error) {
	user, err := find()
	if errors.Is(err, repository.ErrNotFound) {
		return models.User{}, ErrUserNotFound
	}
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// VALIDATE-функции

func validateLogin(login string) error {
	if login == "" {
		return fmt.Errorf("%w: is empty", ErrInvalidLogin)
	}
	if len(login) < 3 {
		return fmt.Errorf("%w: too short, min 3 chars", ErrInvalidLogin)
	}
	if len(login) > 30 {
		return fmt.Errorf("%w: too long, max 30 chars", ErrInvalidLogin)
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("%w: is empty", ErrInvalidEmail)
	}
	if !strings.Contains(email, "@") {
		return fmt.Errorf("%w: wrong format (must be ___@_._)", ErrInvalidEmail)
	}
	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("%w: is empty", ErrInvalidPassword)
	} else if len(password) < 3 {
		return fmt.Errorf("%w: too short, min 3 chars", ErrInvalidPassword)
	}
	return nil
}

func validateName(name string) error {
	if strings.ContainsAny(name, "0123456789!@#$%^&*()_+~`./,{}[];:<>") {
		return fmt.Errorf("%w: contains prohibited characters (0123456789!@#$%^&*()_+~`./,{}[];:<>)", ErrInvalidName)
	}
	return nil
}

// Пока пустой, потому что любой возраст подходит
func validateAge(age uint) error {
	return nil
}

// CREATE-функции

func (s *UserService) CreateUser(login, password, email, name string, age uint) (models.User, error) {
	// Валидация логина
	if err := validateLogin(login); err != nil {
		return models.User{}, err
	}
	// Валидация пароля
	if err := validatePassword(password); err != nil {
		return models.User{}, err
	}
	// Валидация почты
	if err := validateEmail(email); err != nil {
		return models.User{}, err
	}
	// Валидация имени
	if err := validateName(name); err != nil {
		return models.User{}, err
	}

	// Проверка уникальности:
	if err := s.ensureLoginFree(login, noExcludeId); err != nil {
		return models.User{}, err
	}

	if err := s.ensureEmailFree(email, noExcludeId); err != nil {
		return models.User{}, err
	}

	id := s.userRepository.NextUserId()
	newUser := models.User{
		ID:       id,
		Login:    login,
		Password: password,
		Email:    email,
		Name:     name,
		Age:      age,
	}

	err := s.userRepository.SaveUser(&newUser)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

// READ-функции

func (s *UserService) GetUserByLogin(login string) (models.User, error) {
	// Валидация полученного логина
	if err := validateLogin(login); err != nil {
		return models.User{}, err
	}

	return fetchUser(func() (models.User, error) {
		return s.userRepository.FindUserByLogin(login)
	})
}

func (s *UserService) GetUserById(id uint) (models.User, error) {
	// Валидация полученного Id
	if id == 0 {
		return models.User{}, errors.New("Id не может быть равен 0")
	}

	return fetchUser(func() (models.User, error) {
		return s.userRepository.FindUserById(id)
	})
}

func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	// Валидация полученной эл. почты
	if err := validateEmail(email); err != nil {
		return models.User{}, err
	}

	return fetchUser(func() (models.User, error) {
		return s.userRepository.FindUserByEmail(email)
	})
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.FindAllUsers()
}

// UPDATE-функции

func (s *UserService) UpdateUserInfo(id uint,
	newLogin, newEmail, newName *string,
	newAge *uint,
) (models.User, error) {
	// Получим копию юзера, которую будем менять
	user, err := s.userRepository.FindUserById(id)
	if err != nil {
		return models.User{}, err
	}

	// Валидация поступивших данных
	if newLogin != nil {
		if err := validateLogin(*newLogin); err != nil {
			return models.User{}, err
		}

		// Проверка уникальности:
		if err := s.ensureLoginFree(*newLogin, id); err != nil {
			return models.User{}, err
		}

		user.Login = *newLogin
	}
	if newEmail != nil {
		if err := validateEmail(*newEmail); err != nil {
			return models.User{}, err
		}

		// Проверка уникальности:
		if err := s.ensureEmailFree(*newEmail, id); err != nil {
			return models.User{}, err
		}
		user.Email = *newEmail
	}
	if newName != nil {
		if err := validateName(*newName); err != nil {
			return models.User{}, err
		}
		user.Name = *newName
	}
	if newAge != nil {
		if err := validateAge(*newAge); err != nil {
			return models.User{}, err
		}
		user.Age = *newAge
	}

	if err := s.userRepository.SaveUser(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// DELETE-функции

func (s *UserService) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("ID не может быть 0")
	}
	err := s.userRepository.DeleteUser(id)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrUserNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
