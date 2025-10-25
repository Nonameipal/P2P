package repository

import (
	"database/sql"
	"time"

	"github.com/Nonameipal/P2P/internal/models/db"
	"github.com/Nonameipal/P2P/internal/models/domain"
)

func (r *Repository) CreateBooking(booking domain.Booking) error {
	dbBooking := &db.Booking{}
	dbBooking.FromDomain(booking)

	query := `
		INSERT INTO bookings (id, item_id, user_id, start_date, end_date, total_price, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(query,
		dbBooking.ID,
		dbBooking.ItemID,
		dbBooking.UserID,
		dbBooking.StartDate,
		dbBooking.EndDate,
		dbBooking.TotalPrice,
		dbBooking.Status,
		dbBooking.CreatedAt,
		dbBooking.UpdatedAt,
	)

	return err
}

func (r *Repository) GetBookingByID(id string) (domain.Booking, error) {
	var dbBooking db.Booking

	query := `
		SELECT id, item_id, user_id, start_date, end_date, total_price, status, created_at, updated_at
		FROM bookings
		WHERE id = $1
	`

	err := r.db.Get(&dbBooking, query, id)
	if err == sql.ErrNoRows {
		return domain.Booking{}, err
	}

	return dbBooking.ToDomain(), nil
}

func (r *Repository) GetBookingsByItemID(itemID int) ([]domain.Booking, error) {
	var dbBookings []db.Booking

	query := `
		SELECT id, item_id, user_id, start_date, end_date, total_price, status, created_at, updated_at
		FROM bookings
		WHERE item_id = $1
		ORDER BY start_date DESC
	`

	err := r.db.Select(&dbBookings, query, itemID)
	if err != nil {
		return nil, err
	}

	bookings := make([]domain.Booking, len(dbBookings))
	for i, dbBooking := range dbBookings {
		bookings[i] = dbBooking.ToDomain()
	}

	return bookings, nil
}

func (r *Repository) GetBookingsByUserID(userID string) ([]domain.Booking, error) {
	var dbBookings []db.Booking

	query := `
		SELECT id, item_id, user_id, start_date, end_date, total_price, status, created_at, updated_at
		FROM bookings
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	err := r.db.Select(&dbBookings, query, userID)
	if err != nil {
		return nil, err
	}

	bookings := make([]domain.Booking, len(dbBookings))
	for i, dbBooking := range dbBookings {
		bookings[i] = dbBooking.ToDomain()
	}

	return bookings, nil
}

func (r *Repository) UpdateBookingStatus(id string, status string) error {
	query := `
		UPDATE bookings
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`

	_, err := r.db.Exec(query, status, id)
	return err
}

func (r *Repository) CheckItemAvailability(itemID int, startDate, endDate time.Time) (bool, error) {
	var count int

	query := `
		SELECT COUNT(*)
		FROM bookings
		WHERE item_id = $1
		AND status != 'cancelled'
		AND (
			(start_date <= $2 AND end_date >= $2)
			OR (start_date <= $3 AND end_date >= $3)
			OR (start_date >= $2 AND end_date <= $3)
		)
	`

	err := r.db.Get(&count, query, itemID, startDate, endDate)
	if err != nil {
		return false, err
	}

	return count == 0, nil
}
