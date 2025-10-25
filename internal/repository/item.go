package repository

import (
	"os"

	dbModels "github.com/Nonameipal/P2P/internal/models/db"
	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/rs/zerolog"
)

func (r *Repository) CreateItem(item domain.Item) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateItem").Logger()

	dbItem := dbModels.Item{}
	dbItem.FromDomain(item)

	_, err := r.db.Exec(`INSERT INTO items (owner_id, category_id, title, description, price_per_day, status, available_from, available_to)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		dbItem.OwnerID,
		dbItem.CategoryID,
		dbItem.Title,
		dbItem.Description,
		dbItem.PricePerDay,
		dbItem.Status,
		dbItem.AvailableFrom,
		dbItem.AvailableTo,
	)
	if err != nil {
		logger.Err(err).Msg("error inserting item")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) UpdateItemByID(item domain.Item) error {
	dbItem := dbModels.Item{}
	dbItem.FromDomain(item)

	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.UpdateItemByID").Logger()

	query := `UPDATE items SET
					title = $1,
					description = $2,
					price_per_day = $3,
					category_id = $4,
					status = $5,
					updated_at = CURRENT_TIMESTAMP
				WHERE id = $6`

	_, err := r.db.Exec(query,
		dbItem.Title,
		dbItem.Description,
		dbItem.PricePerDay,
		dbItem.CategoryID,
		dbItem.Status,
		dbItem.ID)
	if err != nil {
		logger.Error().Err(err).Msg("Error during update")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) GetAllItems() (items []domain.Item, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetAllItems").Logger()

	var dbItems []dbModels.Item
	if err = r.db.Select(&dbItems, `
		SELECT id, owner_id, category_id, title, description, price_per_day, status, created_at, updated_at
		FROM items
		ORDER BY created_at DESC`); err != nil {
		logger.Err(err).Msg("error selecting items")
		return nil, r.translateError(err)
	}

	for _, dbItem := range dbItems {
		items = append(items, dbItem.ToDomain())
	}

	return items, nil
}

func (r *Repository) GetItemByID(id int) (item domain.Item, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetItemByID").Logger()

	var dbItem dbModels.Item
	if err = r.db.Get(&dbItem, `
		SELECT id, owner_id, category_id, title, description, price_per_day, status, created_at, updated_at
		FROM items
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("error selecting item")
		return domain.Item{}, r.translateError(err)
	}

	return dbItem.ToDomain(), nil
}

func (r *Repository) DeleteItemByID(id int, ownerID string) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.DeleteItemByID").Logger()

	_, err = r.db.Exec(`DELETE FROM items WHERE id = $1 AND owner_id = $2`, id, ownerID)
	if err != nil {
		logger.Err(err).Msg("error deleting item")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) GetItemsByCategory(category string) (items []domain.Item, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetItemsByCategory").Logger()

	var dbItems []dbModels.Item
	if err = r.db.Select(&dbItems, `
		SELECT id, owner_id, category_id, title, description, price_per_day, status, created_at, updated_at
		FROM items
		WHERE category_id = $1
		ORDER BY created_at DESC`, category); err != nil {
		logger.Err(err).Msg("error selecting items by category")
		return nil, r.translateError(err)
	}

	for _, dbItem := range dbItems {
		items = append(items, dbItem.ToDomain())
	}

	return items, nil
}

func (r *Repository) GetMyItems(userID string) (items []domain.Item, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetMyItems").Logger()

	var dbItems []dbModels.Item
	if err = r.db.Select(&dbItems, `
		SELECT id, owner_id, category_id, title, description, price_per_day, status, created_at, updated_at
		FROM items
		WHERE owner_id = $1
		ORDER BY created_at DESC`, userID); err != nil {
		logger.Err(err).Msg("error selecting user items")
		return nil, r.translateError(err)
	}

	for _, dbItem := range dbItems {
		items = append(items, dbItem.ToDomain())
	}

	return items, nil
}

func (r *Repository) GetItemsByStatus(status string) (items []domain.Item, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetItemsByStatus").Logger()

	var dbItems []dbModels.Item
	if err = r.db.Select(&dbItems, `
		SELECT id, owner_id, category_id, title, description, price_per_day, status, created_at, updated_at
		FROM items
		WHERE status = $1
		ORDER BY created_at DESC`, status); err != nil {
		logger.Err(err).Msg("error selecting items by status")
		return nil, r.translateError(err)
	}

	for _, dbItem := range dbItems {
		items = append(items, dbItem.ToDomain())
	}

	return items, nil
}

func (r *Repository) GetItemCount() (int, error) {
	var count int
	err := r.db.Get(&count, `SELECT COUNT(*) FROM items`)
	return count, err
}
