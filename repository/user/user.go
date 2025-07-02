package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aghaghiamh/ava/entity"
	"github.com/aghaghiamh/ava/pkg/errmsg"
	"github.com/aghaghiamh/ava/pkg/richerr"
	"gorm.io/gorm"
)

type User struct {
	ID          uint `gorm:"primarykey"`
	Name        string	`gorm:"not null"`
	PhoneNumber *string
	CreatedAt   time.Time
}

func (s Storage) Register(user entity.User) (entity.User, error) {
	const op = "repository.user.RegisterUser"

	gormUser := User{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}

	result := s.db.Create(&gormUser)
	if result.Error != nil {
		return entity.User{}, richerr.New(op).WithError(result.Error).
			WithCode(richerr.ErrServer).WithMessage(errmsg.ErrMsgCantExecuteQuery)
	}

	user.ID = gormUser.ID
	return user, nil
}

func (s Storage) GetUserByID(ctx context.Context, userID uint) (entity.User, error) {
	const op = "repository.user.GetUserByID"
	var fetchedUser User

	result := s.db.WithContext(ctx).First(&fetchedUser, userID)
	if result.Error != nil {
		fmt.Println(result.Error)
		richErr := richerr.New(op).WithError(result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, richErr.WithCode(richerr.ErrEntityNotFound).WithMessage(errmsg.ErrMsgNotFound)
		}
		return entity.User{}, richErr.WithCode(richerr.ErrServer).WithMessage(errmsg.ErrMsgCantExecuteQuery)
	}

	// TODO: Define mapper
	user := entity.User{
		ID:   fetchedUser.ID,
		Name: fetchedUser.Name,
	}
	if fetchedUser.PhoneNumber != nil {
		user.PhoneNumber = fetchedUser.PhoneNumber
	}

	return user, nil
}

func (s Storage) DelByID(userID uint) error {
	const op = "repository.user.DelUserByID"

	result := s.db.Delete(&User{}, userID)
	if result.Error != nil {
		return richerr.New(op).WithError(result.Error).
			WithCode(richerr.ErrServer).WithMessage(errmsg.ErrMsgCantExecuteQuery)
	}

	if result.RowsAffected == 0 {
		return richerr.New(op).WithCode(richerr.ErrEntityNotFound).WithMessage(errmsg.ErrMsgNotFound)
	}
	
	// If we reach here, the deletion was successful.
	return nil
}

func (s Storage) ListWithPagination(ctx context.Context, page, pageSize int) ([]entity.User, error) {
	const op = "repository.user.ListUsers"
	var fetchedUsers []User

	offset := (page - 1) * pageSize

	result := s.db.WithContext(ctx).Order("id asc").Limit(pageSize).Offset(offset).Find(&fetchedUsers)
	if result.Error != nil {
		return nil, richerr.New(op).WithError(result.Error).
			WithCode(richerr.ErrServer).WithMessage(errmsg.ErrMsgCantExecuteQuery)
	}

	// TODO: Define Mapper
	users := make([]entity.User, len(fetchedUsers))
	for i, u := range fetchedUsers {
		users[i] = entity.User{
			ID:   u.ID,
			Name: u.Name,
		}
		if u.PhoneNumber != nil {
			users[i].PhoneNumber = u.PhoneNumber
		}
	}

	return users, nil
}
