package user

import (
	"context"
	"database/sql"

	"github.com/aghaghiamh/ava/entity"
	"github.com/aghaghiamh/ava/pkg/errmsg"
	"github.com/aghaghiamh/ava/pkg/richerr"
)

type User struct {
	ID          uint
	Name        string
	PhoneNumber sql.NullString
	CreatedAt   []uint8
}

func userScanner(row *sql.Row, fetchedUser *User) error {
	return row.Scan(&fetchedUser.ID, &fetchedUser.Name, &fetchedUser.PhoneNumber, &fetchedUser.CreatedAt)
}

func (mysql Storage) IsAlreadyExist(phoneNumber string) (bool, error) {
	const op = "mysql.IsAlreadyExist"
	var fetchedUser User

	query := `SELECT * FROM users WHERE phone_number = ?`
	row := mysql.db.QueryRow(query, phoneNumber)

	sErr := userScanner(row, &fetchedUser)
	if sErr != nil {
		if sErr == sql.ErrNoRows {
			return false, nil
		}

		return false, richerr.New(op).
			WithError(sErr).
			WithCode(richerr.ErrServer).
			WithMessage(errmsg.ErrMsgCantScanQueryResult)
	}

	return true, nil
}

func (s Storage) Register(user entity.User) (entity.User, error) {
	const op = "RegisterUser"

	var query string
	var res sql.Result
	var err error

	// TODO: Implement using GORM to handle the nullable values smoothly.
	if len(user.PhoneNumber) > 0 {
		query = `INSERT INTO users(name, phone_number) VALUES (?, ?)`
		res, err = s.db.Exec(query, user.Name, user.PhoneNumber)
	} else {
		query = `INSERT INTO users(name) VALUES (?)`
		res, err = s.db.Exec(query, user.Name)
	}

	if err != nil {

		return entity.User{}, richerr.New(op).
			WithError(err).
			WithCode(richerr.ErrServer).
			WithMessage(errmsg.ErrMsgCantScanQueryResult)
	}

	lastID, _ := res.LastInsertId()
	user.ID = uint(lastID)

	return user, nil
}

func (s Storage) GetUserByID(ctx context.Context, user_id uint) (entity.User, error) {
	const op = "GetUserByID"
	var fetchedUser User

	query := `SELECT * FROM users WHERE id = ?`
	row := s.db.QueryRowContext(ctx, query, user_id)

	sErr := userScanner(row, &fetchedUser)
	if sErr != nil {
		richErr := richerr.New(op).WithError(sErr)

		if sErr == sql.ErrNoRows {
			return entity.User{}, richErr.
				WithCode(richerr.ErrEntityNotFound).
				WithMessage(errmsg.ErrMsgNotFound)
		}

		return entity.User{}, richErr.
			WithCode(richerr.ErrServer).
			WithMessage(errmsg.ErrMsgCantScanQueryResult)
	}

	user := entity.User{
		ID:          fetchedUser.ID,
		Name:        fetchedUser.Name,
		PhoneNumber: fetchedUser.PhoneNumber.String,
	}

	return user, nil
}
