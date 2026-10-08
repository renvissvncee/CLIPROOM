package service

import (
	"cliproom/models"
	"cliproom/repository"
	"errors"
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

func (usserv *UserService) userIsTaken(find func() (models.User, error), excludeId uint) (bool, error) {
	user, err := find()
	if errors.Is(err, repository.ErrNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return user.ID != excludeId, nil
}

func (usserv *UserService) ensureLoginFree(login string, excludeId uint) error {
	taken, err := usserv.userIsTaken(func() (models.User, error) {
		return usserv.userRepository.FindUserByLogin(login)
	}, excludeId)

	if err != nil {
		return err
	} else if taken {
		return ErrLoginTaken
	}
	return nil
}

func (usserv *UserService) ensureEmailFree(email string, excludeId uint) error {
	taken, err := usserv.userIsTaken(func() (models.User, error) {
		return usserv.userRepository.FindUserByEmail(email)
	}, excludeId)

	if err != nil {
		return err
	} else if taken {
		return ErrLoginTaken
	}
	return nil
}

func (usserv *UserService) fetchUser(find func() (models.User, error)) (models.User, error) {
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

func (usserv *UserService) validateLogin(login string) error {
	if login == "" {
		return errors.New("Логин не может быть пустым")
	} else if len(login) < 3 || len(login) > 30 {
		return errors.New("Логин должен быть верной длины: 3 <= login <= 30")
	}
	return nil
}

func (usserv *UserService) validateEmail(email string) error {
	if email == "" {
		return errors.New("Эл. почта не может быть пустой")
	} else if !strings.Contains(email, "@") {
		return errors.New("Неверный формат эл. почты")
	}
	return nil
}

func (usserv *UserService) validatePassword(password string) error {
	if password == "" {
		return errors.New("Пароль не может быть пустым")
	} else if len(password) < 3 {
		return errors.New("Пароль не может быть меньше 3 символов")
	}
	return nil
}

func (usserv *UserService) validateName(name string) error {
	if strings.ContainsAny(name, "0123456789!@#$%^&*()_+~`./,{}[];:<>") {
		return errors.New("Имя не должно содержать чисел и(или) специальных символов")
	}
	return nil
}

// Пока пустой, потому что любой возраст подходит
func (usserv *UserService) validateAge(age uint) error {
	return nil
}

// CREATE-функции

func (usserv *UserService) CreateUser(login, password, email, name string, age uint) (models.User, error) {
	// Валидация логина
	if err := usserv.validateLogin(login); err != nil {
		return models.User{}, err
	}
	// Валидация пароля
	if err := usserv.validatePassword(password); err != nil {
		return models.User{}, err
	}
	// Валидация почты
	if err := usserv.validateEmail(email); err != nil {
		return models.User{}, err
	}
	// Валидация имени
	if err := usserv.validateName(name); err != nil {
		return models.User{}, err
	}

	// Проверка уникальности:
	if err := usserv.ensureLoginFree(login, noExcludeId); err != nil {
		return models.User{}, err
	}

	if err := usserv.ensureEmailFree(email, noExcludeId); err != nil {
		return models.User{}, err
	}

	id := usserv.userRepository.NextUserId()
	newUser := models.User{
		ID:       id,
		Login:    login,
		Password: password,
		Email:    email,
		Name:     name,
		Age:      age,
	}

	err := usserv.userRepository.SaveUser(&newUser)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

// READ-функции

func (usserv *UserService) GetUserByLogin(login string) (models.User, error) {
	// Валидация полученного логина
	if err := usserv.validateLogin(login); err != nil {
		return models.User{}, err
	}

	return usserv.fetchUser(func() (models.User, error) {
		return usserv.userRepository.FindUserByLogin(login)
	})
}

func (usserv *UserService) GetUserById(id uint) (models.User, error) {
	// Валидация полученного Id
	if id == 0 {
		return models.User{}, errors.New("Id не может быть равен 0")
	}

	return usserv.fetchUser(func() (models.User, error) {
		return usserv.userRepository.FindUserById(id)
	})
}

func (usserv *UserService) GetUserByEmail(email string) (models.User, error) {
	// Валидация полученной эл. почты
	if err := usserv.validateEmail(email); err != nil {
		return models.User{}, err
	}

	return usserv.fetchUser(func() (models.User, error) {
		return usserv.userRepository.FindUserByEmail(email)
	})
}

// UPDATE-функции

func (usserv *UserService) UpdateUserInfo(id uint,
	newLogin, newEmail, newName *string,
	newAge *uint,
) (models.User, error) {
	// Получим копию юзера, которую будем менять
	user, err := usserv.userRepository.FindUserById(id)
	if err != nil {
		return models.User{}, err
	}

	// Валидация поступивших данных
	if newLogin != nil {
		if err := usserv.validateLogin(*newLogin); err != nil {
			return models.User{}, err
		}

		// Проверка уникальности:
		if err := usserv.ensureLoginFree(*newLogin, noExcludeId); err != nil {
			return models.User{}, err
		}

		user.Login = *newLogin
	}
	if newEmail != nil {
		if err := usserv.validateEmail(*newEmail); err != nil {
			return models.User{}, err
		}

		// Проверка уникальности:
		if err := usserv.ensureEmailFree(*newEmail, noExcludeId); err != nil {
			return models.User{}, err
		}
		user.Email = *newEmail
	}
	if newName != nil {
		if err := usserv.validateName(*newName); err != nil {
			return models.User{}, err
		}
		user.Name = *newName
	}
	if newAge != nil {
		if err := usserv.validateAge(*newAge); err != nil {
			return models.User{}, err
		}
		user.Age = *newAge
	}

	if err := usserv.userRepository.SaveUser(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}
