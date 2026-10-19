package service

import (
	"cliproom/models"
	"cliproom/repository"
	"errors"
	"fmt"
)

// Структура, отвечающая за бизнес-лоигку, связанную с пользователем
type ClipService struct {
	clipRepository *repository.ClipRepository
}

func NewClipService(p_repositroy *repository.ClipRepository) *ClipService {
	return &ClipService{clipRepository: p_repositroy}
}

// ХЕЛПЕР-функции

func (s *ClipService) clipIsExists(artist string, title string) (bool, error) {
	_, err := s.clipRepository.FindClipByArtistAndTitle(artist, title)
	if errors.Is(err, repository.ErrNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// func (s *ClipService) ensureLoginFree(login string, excludeId uint) error {
// 	taken, err := s.clipIsTaken(func() (models.Clip, error) {
// 		return s.clipRepository.FindClipByLogin(login)
// 	}, excludeId)

// 	if err != nil {
// 		return err
// 	} else if taken {
// 		return ErrLoginTaken
// 	}
// 	return nil
// }

// func (s *ClipService) ensureEmailFree(email string, excludeId uint) error {
// 	taken, err := s.clipIsTaken(func() (models.Clip, error) {
// 		return s.clipRepository.FindClipByEmail(email)
// 	}, excludeId)

// 	if err != nil {
// 		return err
// 	} else if taken {
// 		return ErrEmailTaken
// 	}
// 	return nil
// }

func fetchClips(find func() ([]models.Clip, error)) ([]models.Clip, error) {
	clips, err := find()
	if errors.Is(err, repository.ErrNotFound) {
		return []models.Clip{}, ErrClipNotFound
	}
	if err != nil {
		return []models.Clip{}, err
	}
	return clips, nil
}

// VALIDATE-функции

func validateTitle(title string) error {
	if title == "" {
		return fmt.Errorf("%w: is empty", ErrInvalidTitle)
	}
	return nil
}

func validateArtist(artist string) error {
	if artist == "" {
		return fmt.Errorf("%w: is empty", ErrInvalidArtist)
	}
	return nil
}

func validateDescription(description string) error {
	if len(description) > 1000 {
		return fmt.Errorf("%w: too long, max 1000 chars", ErrInvalidDescription)
	}
	return nil
}

func validateLink(link string) error {
	return nil
}

func parseTags(raw []string) ([]models.Tag, error) {
	tags := make([]models.Tag, 0, len(raw))
	for _, str := range raw {
		t := models.Tag(str)
		if !t.IsValidTag() {
			return []models.Tag{}, fmt.Errorf("%w: a non-existent tag (%s)", ErrInvalidTags, str)
		}
		tags = append(tags, t)
	}
	return tags, nil
}

func parseMoods(raw []string) ([]models.Mood, error) {
	moods := make([]models.Mood, 0, len(raw))
	for _, str := range raw {
		m := models.Mood(str)
		if !m.IsValidMood() {
			return []models.Mood{}, fmt.Errorf("%w: a non-existent mood (%s)", ErrInvalidMoods, str)
		}
		moods = append(moods, m)
	}
	return moods, nil
}

// CREATE-функции

func (s *ClipService) CreateClip(title, artist, description, link string, tagsRaw, moodsRaw []string) (models.Clip, error) {
	// Валидация логина
	if err := validateTitle(title); err != nil {
		return models.Clip{}, err
	}
	// Валидация пароля
	if err := validateArtist(artist); err != nil {
		return models.Clip{}, err
	}
	// Валидация почты
	if err := validateDescription(description); err != nil {
		return models.Clip{}, err
	}
	// Валидация имени
	if err := validateLink(link); err != nil {
		return models.Clip{}, err
	}
	// Валидация тэгов
	tags, err := parseTags(tagsRaw)
	if err != nil {
		return models.Clip{}, err
	}
	// Валидация мудов
	moods, err := parseMoods(moodsRaw)
	if err != nil {
		return models.Clip{}, err
	}

	// Проверка уникальности:
	exists, err := s.clipIsExists(artist, title)
	if err != nil {
		return models.Clip{}, err
	}
	if exists {
		return models.Clip{}, ErrClipExists
	}

	id := s.clipRepository.NextClipId()
	newClip := models.Clip{
		ID:          id,
		Artist:      artist,
		Title:       title,
		Description: description,
		Likes:       0,
		Tags:        tags,
		Moods:       moods,
		Link:        link,
	}

	err = s.clipRepository.SaveClip(&newClip)
	if err != nil {
		return models.Clip{}, err
	}
	return newClip, nil
}

// READ-функции

func (s *ClipService) GetClipsByTitle(title string) ([]models.Clip, error) {
	// Валидация полученного логина
	if err := validateTitle(title); err != nil {
		return []models.Clip{}, err
	}

	return fetchClips(func() ([]models.Clip, error) {
		return s.clipRepository.FindClipsByTitle(title)
	})
}

func (s *ClipService) GetClipById(id uint) (models.Clip, error) {
	// Валидация полученного Id
	if id == 0 {
		return models.Clip{}, errors.New("invalid id (0)")
	}

	clip, err := s.clipRepository.FindClipById(id)
	if errors.Is(err, repository.ErrNotFound) {
		return models.Clip{}, ErrClipNotFound
	}
	if err != nil {
		return models.Clip{}, err
	}
	return clip, nil

}

func (s *ClipService) GetClipsByArtist(artist string) ([]models.Clip, error) {
	// Валидация полученной эл. почты
	if err := validateArtist(artist); err != nil {
		return []models.Clip{}, err
	}

	return fetchClips(func() ([]models.Clip, error) {
		return s.clipRepository.FindClipsByArtist(artist)
	})
}

func (s *ClipService) GetAllClips() ([]models.Clip, error) {
	return s.clipRepository.FindAllClips()
}

// UPDATE-функции

func (s *ClipService) UpdateClipInfo(id uint,
	newTitle, newArtist, newDescription, newLink *string,
	newTagsRaw []string, newMoodsRaw []string,
) (models.Clip, error) {
	// Получим копию юзера, которую будем менять
	clip, err := s.clipRepository.FindClipById(id)
	if errors.Is(err, repository.ErrNotFound) {
		return models.Clip{}, ErrClipNotFound
	}
	if err != nil {
		return models.Clip{}, err
	}

	// Валидация полученных данных
	if newTitle == nil {
		newTitle = &(clip.Title)
	}
	if newArtist == nil {
		newArtist = &(clip.Artist)
	}
	if err := validateTitle(*newTitle); err != nil {
		return models.Clip{}, err
	}
	if err := validateArtist(*newArtist); err != nil {
		return models.Clip{}, err
	}
	if *newTitle != clip.Title || *newArtist != clip.Artist {
		exists, err := s.clipIsExists(*newArtist, *newTitle)
		if err != nil {
			return models.Clip{}, err
		}
		if exists {
			return models.Clip{}, ErrClipExists
		}
	}
	clip.Artist = *newArtist
	clip.Title = *newTitle

	if newDescription != nil {
		if err := validateDescription(*newDescription); err != nil {
			return models.Clip{}, err
		}
		clip.Description = *newDescription
	}
	if newLink != nil {
		if err := validateLink(*newLink); err != nil {
			return models.Clip{}, err
		}
		clip.Link = *newLink
	}
	if newTagsRaw != nil {
		newTags, err := parseTags(newTagsRaw)
		if err != nil {
			return models.Clip{}, err
		}
		clip.Tags = newTags
	}
	if newMoodsRaw != nil {
		newMoods, err := parseMoods(newMoodsRaw)
		if err != nil {
			return models.Clip{}, err
		}
		clip.Moods = newMoods
	}

	if err := s.clipRepository.SaveClip(&clip); err != nil {
		return models.Clip{}, err
	}

	return clip, nil
}

// // DELETE-функции

// func (s *ClipService) DeleteClip(id uint) error {
// 	if id == 0 {
// 		return errors.New("ID не может быть 0")
// 	}
// 	err := s.clipRepository.DeleteClip(id)
// 	if errors.Is(err, repository.ErrNotFound) {
// 		return ErrClipNotFound
// 	}
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
