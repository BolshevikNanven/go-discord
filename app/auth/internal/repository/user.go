package repository

import (
	"discord/app/auth/internal/model"
	"discord/pkg/snowflakeutil"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUsername(username string) *model.User
	Create(username, password string) (*model.User, error)
}

type userRepository struct {
	db        *gorm.DB
	snowflake *snowflakeutil.Node
}

func NewUserRepo(db *gorm.DB, snowflake *snowflakeutil.Node) UserRepository {
	return &userRepository{
		db:        db,
		snowflake: snowflake,
	}
}

func (u *userRepository) GetByUsername(username string) *model.User {
	var user model.User
	result := u.db.Where("username = ?", username).Take(&user)
	if result.Error != nil {
		return nil
	}
	return &user
}

func (u *userRepository) Create(username, password string) (*model.User, error) {
	user := &model.User{
		Id:       u.snowflake.GenerateID(),
		Username: username,
		Password: password,
	}

	result := u.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
