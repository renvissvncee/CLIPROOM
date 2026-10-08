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

func (r *ClipRepository) FindClipByTitle(title string) (models.Clip, error) {
	clipMap := r.storage.Clips
	for _, value := range clipMap {
		if value.Title == title {
			return value, nil
		}
	}
	return models.Clip{}, ErrNotFound
}

func (r *ClipRepository) FindClipByArtist(artist string) (models.Clip, error) {
	clipMap := r.storage.Clips
	for _, value := range clipMap {
		if value.Artist == artist {
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
	clipMap := r.storage.Clips
	if _, ok := clipMap[id]; !ok {
		return ErrNotFound
	}
	delete(clipMap, id)
	return nil
}
