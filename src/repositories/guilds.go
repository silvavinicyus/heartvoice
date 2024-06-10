package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"heartvoice/src/models"
)

type guilds struct {
	db *sql.DB
}

type PaginationParams struct {
	Page  uint64
	Limit uint64
}

func GuildRepository(db *sql.DB) *guilds {
	return &guilds{db}
}

func (repository guilds) Create(guild models.Guild) (uint64, error) {
	statement, statementError := repository.db.Prepare("INSERT INTO guilds(name, imageKey, description) VALUES (?, ?, ?)")

	if statementError != nil {
		return 0, statementError
	}
	defer statement.Close()

	result, statementError := statement.Exec(guild.Name, guild.ImageKey, guild.Description)

	if statementError != nil {
		return 0, statementError
	}

	guildInsertedId, statementError := result.LastInsertId()

	if statementError != nil {
		return 0, statementError
	}

	return uint64(guildInsertedId), nil
}

func (repository guilds) FindBy(column string, value interface{}) (models.Guild, error) {
	allowedColumns := map[string]bool{
		"id":   true,
		"name": true,
	}

	if !allowedColumns[column] {
		return models.Guild{}, errors.New("column not allowed on find by")
	}

	query := fmt.Sprintf("SELECT id, name, imageKey, description, createdAt, updatedAt FROM guilds WHERE %s = ?", column)
	lines, linesError := repository.db.Query(query, value)
	if linesError != nil {
		return models.Guild{}, linesError
	}

	defer lines.Close()

	var guild models.Guild

	if lines.Next() {
		if linesError = lines.Scan(&guild.ID, &guild.Name, &guild.ImageKey, &guild.Description, &guild.CreatedAt, &guild.UpdatedAt); linesError != nil {
			return models.Guild{}, linesError
		}
	}

	return guild, nil
}

func (repository guilds) FindAll(pagination PaginationParams, name string) ([]models.Guild, error) {
	name = fmt.Sprintf("%%%s%%", name)

	lines, queryError := repository.db.Query(
		"SELECT id, name, imageKey, description, createdAt, updatedAt FROM guilds WHERE name LIKE ? limit ? offset ?", name, pagination.Limit, (pagination.Limit * pagination.Page),
	)

	if queryError != nil {
		return nil, queryError
	}

	defer lines.Close()

	var guilds []models.Guild

	for lines.Next() {
		var guild models.Guild

		if scanError := lines.Scan(&guild.ID, &guild.Name, &guild.ImageKey, &guild.Description, &guild.CreatedAt, &guild.UpdatedAt); scanError != nil {
			return nil, scanError
		}

		guilds = append(guilds, guild)
	}

	return guilds, nil
}
