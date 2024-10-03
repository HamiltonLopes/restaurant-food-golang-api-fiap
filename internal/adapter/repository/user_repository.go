package repository

import (
	"context"
	"github.com/tupizz/restaurant-food-golang-api-fiap/internal/domain"
	"github.com/tupizz/restaurant-food-golang-api-fiap/internal/domain/entity"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

func (r *userRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.QueryRow(ctx, "INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&user.ID)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
