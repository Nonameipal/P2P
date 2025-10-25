package repository

import (
	"os"

	dbModels "github.com/Nonameipal/P2P/internal/models/db"
	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/rs/zerolog"
)

func (r *Repository) CreateCategory(category domain.Category) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateCategory").Logger()

	_, err := r.db.Exec(`INSERT INTO categories (name, description)
					VALUES ($1, $2)`,
		category.Name,
		category.Description,
	)
	if err != nil {
		logger.Err(err).Msg("error inserting category")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) GetAllCategories() ([]domain.Category, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetAllCategories").Logger()

	var dbCategories []dbModels.Category
	if err := r.db.Select(&dbCategories, `
		SELECT id, name, description, created_at, updated_at
		FROM categories
		ORDER BY name`); err != nil {
		logger.Err(err).Msg("error selecting categories")
		return nil, r.translateError(err)
	}

	var categories []domain.Category
	for _, dbCategory := range dbCategories {
		categories = append(categories, dbCategory.ToDomain())
	}

	return categories, nil
}

func (r *Repository) GetCategoryByID(id int) (domain.Category, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetCategoryByID").Logger()

	var dbCategory dbModels.Category
	if err := r.db.Get(&dbCategory, `
		SELECT id, name, description, created_at, updated_at
		FROM categories
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("error selecting category")
		return domain.Category{}, r.translateError(err)
	}

	return dbCategory.ToDomain(), nil
}

func (r *Repository) UpdateCategory(category domain.Category) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.UpdateCategory").Logger()

	query := `UPDATE categories SET
					name = $1,
					description = $2,
					updated_at = CURRENT_TIMESTAMP
				WHERE id = $3`

	_, err := r.db.Exec(query,
		category.Name,
		category.Description,
		category.ID)
	if err != nil {
		logger.Error().Err(err).Msg("Error during category update")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) DeleteCategory(id int) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.DeleteCategory").Logger()

	_, err := r.db.Exec(`DELETE FROM categories WHERE id = $1`, id)
	if err != nil {
		logger.Err(err).Msg("error deleting category")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) GetCategoryByName(name string) (domain.Category, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetCategoryByName").Logger()

	var dbCategory dbModels.Category
	if err := r.db.Get(&dbCategory, `
		SELECT id, name, description, created_at, updated_at
		FROM categories
		WHERE name = $1`, name); err != nil {
		logger.Err(err).Msg("error selecting category by name")
		return domain.Category{}, r.translateError(err)
	}

	return dbCategory.ToDomain(), nil
}
