package repositories

import (
	"database/sql"
	"heartvoice/src/models"
)

type users struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, statementError := repository.db.Prepare("INSERT INTO users(name, nick, email, password) VALUES (?, ?, ?, ?)")
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
