package repository

import (
	"discord/app/biz/internal/model"
	"discord/pkg/snowflakeutil"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type SpaceRepository interface {
	Create(space model.Space) (model.Space, error)
	GetByID(id int64) (*model.Space, error)
	UserList(userID int64) ([]model.Space, error)
	Update(space model.Space) error
	Delete(id int64) error
}

type spaceRepository struct {
	db        *gorm.DB
	snowflake *snowflakeutil.Node
	cache     *redis.Client
}

func NewSpaceRepo(db *gorm.DB, snowflake *snowflakeutil.Node, cache *redis.Client) SpaceRepository {
	return &spaceRepository{
		db:        db,
		snowflake: snowflake,
		cache:     cache,
	}
}

func (s *spaceRepository) Create(space model.Space) (model.Space, error) {
	space.Id = s.snowflake.GenerateID()

	// 开启事务
	return space, s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&space).Error; err != nil {
			return err
		}

		spaceUser := model.SpaceUser{
			SpaceId:  space.Id,
			UserId:   space.Owner,
			Nickname: "admin",
		}

		if err := tx.Create(&spaceUser).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *spaceRepository) GetByID(id int64) (*model.Space, error) {
	var space model.Space
	result := s.db.Where("id = ?", id).Take(&space)
	if result.Error != nil {
		return nil, result.Error
	}
	return &space, nil
}

func (s *spaceRepository) UserList(userID int64) ([]model.Space, error) {
	var spaces []model.Space
	result := s.db.Model(&model.Space{}).
		Joins("JOIN space_user ON space.id = space_user.space_id").
		Where("space_user.user_id = ?", userID).
		Find(&spaces)
	if result.Error != nil {
		return nil, result.Error
	}

	return spaces, nil
}

func (s *spaceRepository) Update(space model.Space) error {
	return s.db.Model(&model.Space{}).Where("id = ?", space.Id).Updates(space).Error
}

func (s *spaceRepository) Delete(id int64) error {
	return s.db.Delete(&model.Space{}, id).Error
}
