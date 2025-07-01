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

func (s Storage) GetUserByID(ctx context.Context, userID uint) (entity.User, error) {
	const op = "GetUserByID"
	var fetchedUser User

	query := `SELECT * FROM users WHERE id = ?`
	row := s.db.QueryRowContext(ctx, query, userID)

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

func (s Storage) DelByID(userID uint) error {
	const op = "DelUserByID"

	query := `DELETE FROM users WHERE id = ?`
	res, qErr := s.db.Exec(query, userID)

	if qErr != nil {
		return richerr.New(op).WithError(qErr).
			WithCode(richerr.ErrServer).
			WithMessage(errmsg.ErrMsgExecutingQuery)
	}

	// Try not to act aggressively and catch any error
	rowsAffected, _ := res.RowsAffected()

	// Error since no user with that ID was found.
	if rowsAffected == 0 {
		return richerr.New(op).
			WithCode(richerr.ErrEntityNotFound).
			WithMessage(errmsg.ErrMsgNotFound)
	}

	// If we reach here, the deletion was successful.
	return nil
}

func (s Storage) ListWithPagination(ctx context.Context, page, pageSize int) ([]entity.User, error) {
	const op = "ListUsers"

	offset := (page - 1) * pageSize
	// TODO: better to order by another column && limit the columns to fetch.
	query := `SELECT * FROM users ORDER BY id ASC LIMIT ? OFFSET ?`

	rows, qErr := s.db.QueryContext(ctx, query, pageSize, offset)
	if qErr != nil {
		return nil, richerr.New(op).WithError(qErr).
			WithCode(richerr.ErrServer).
			WithMessage(errmsg.ErrMsgExecutingQuery)
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var fetchedUser User
		if err := rows.Scan(&fetchedUser.ID, &fetchedUser.Name, &fetchedUser.PhoneNumber, &fetchedUser.CreatedAt); err != nil {
			return nil, richerr.New(op).WithError(err).
				WithCode(richerr.ErrServer).
				WithMessage(errmsg.ErrMsgCantScanQueryResult)
		}

		user := entity.User{
			ID:          fetchedUser.ID,
			Name:        fetchedUser.Name,
			PhoneNumber: fetchedUser.PhoneNumber.String,
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, richerr.New(op).WithError(qErr).
			WithCode(richerr.ErrServer).
			WithMessage(errmsg.ErrMsgCantScanQueryResult)
	}

	return users, nil
}
