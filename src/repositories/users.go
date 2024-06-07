package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"heartvoice/src/models"
)

type users struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, statementError := repository.db.Prepare("INSERT INTO users(name, nickname, email, password) VALUES (?, ?, ?, ?)")
	if statementError != nil {
		return 0, statementError
	}

	defer statement.Close()

	result, statementError := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if statementError != nil {
		return 0, statementError
	}

	userInsertedId, statementError := result.LastInsertId()
	if statementError != nil {
		return 0, statementError
	}

	return uint64(userInsertedId), nil
}

func (repository users) FindBy(column string, value string) (models.User, error) {
	allowedColumns := map[string]bool{
		"id":       true,
		"email":    true,
		"nickname": true,
	}

	if !allowedColumns[column] {
		return models.User{}, errors.New("column not allowed on find by")
	}

	query := fmt.Sprintf("SELECT id, name, password, email, nickname, createdAt FROM users WHERE %s = ?", column)
	lines, linesError := repository.db.Query(query, value)
	if linesError != nil {
		return models.User{}, linesError
	}

	defer lines.Close()

	var user models.User
	if lines.Next() {
		if linesError = lines.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Nickname, &user.CreatedAt); linesError != nil {
			return models.User{}, linesError
		}
	}

	return user, nil
}
