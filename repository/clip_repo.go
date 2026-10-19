package repository

import (
	"cliproom/models"
	"cliproom/storage"
	"errors"
)

// Структура, отвечающая за запросы к "БД", связанные с клипами
type ClipRepository struct {
	storage *storage.Storage
}

func NewClipRepository(p_storage *storage.Storage) *ClipRepository {
	return &ClipRepository{storage: p_storage}
}

func (r *ClipRepository) NextClipId() uint {
	return r.storage.NextClipId()
}

// CREATE-функции

func (r *ClipRepository) SaveClip(clip *models.Clip) error {
	if clip == nil {
		return errors.New("Ссылка на клип не найдена")
	}
	if clip.ID == 0 {
		return errors.New("Id клипа не может быть 0")
	}
	r.storage.Clips[clip.ID] = *clip
	return nil
}

// READ-функции

func (r *ClipRepository) FindClipById(id uint) (models.Clip, error) {
	if clip, ok := r.storage.Clips[id]; ok {
		return clip, nil
	}
	return models.Clip{}, ErrNotFound
}

func (r *ClipRepository) FindClipsByTitle(title string) ([]models.Clip, error) {
	clipSlice := []models.Clip{}
	for _, value := range r.storage.Clips {
		if value.Title == title {
			clipSlice = append(clipSlice, value)
		}
	}
	if len(clipSlice) == 0 {
		return []models.Clip{}, ErrNotFound
	}
	return clipSlice, nil
}

func (r *ClipRepository) FindClipsByArtist(artist string) ([]models.Clip, error) {
	clipSlice := []models.Clip{}
	for _, value := range r.storage.Clips {
		if value.Artist == artist {
			clipSlice = append(clipSlice, value)
		}
	}
	if len(clipSlice) == 0 {
		return []models.Clip{}, ErrNotFound
	}
	return clipSlice, nil
}

func (r *ClipRepository) FindClipByArtistAndTitle(artist string, title string) (models.Clip, error) {
	for _, value := range r.storage.Clips {
		if value.Artist == artist && value.Title == title {
			return value, nil
		}
	}
	return models.Clip{}, ErrNotFound
}

func (r *ClipRepository) FindAllClips() ([]models.Clip, error) {
	clipsSlice := make([]models.Clip, 0, len(r.storage.Clips))
	for _, clip := range r.storage.Clips {
		clipsSlice = append(clipsSlice, clip)
	}
	return clipsSlice, nil
}

// DELETE-функции

func (r *ClipRepository) DeleteClip(id uint) error {
	if _, ok := r.storage.Clips[id]; !ok {
		return ErrNotFound
	}
	delete(r.storage.Clips, id)
	return nil
}
