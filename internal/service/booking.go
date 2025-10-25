package service

import (
	"time"

	"github.com/Nonameipal/P2P/internal/errs"
	"github.com/Nonameipal/P2P/internal/models/domain"
)

func (s *Service) CreateBooking(booking domain.Booking) error {
	// Проверяем доступность вещи на выбранные даты
	available, err := s.repository.CheckItemAvailability(booking.ItemID, booking.StartDate, booking.EndDate)
	if err != nil {
		return err
	}
	if !available {
		return errs.ErrItemNotAvailable
	}

	// Получаем информацию о вещи для расчета стоимости
	item, err := s.repository.GetItemByID(booking.ItemID)
	if err != nil {
		return err
	}

	// Рассчитываем количество дней аренды
	days := booking.EndDate.Sub(booking.StartDate).Hours() / 24
	if days < 1 {
		days = 1
	}

	// Рассчитываем полную стоимость
	booking.TotalPrice = item.PricePerDay * float64(days)

	// Устанавливаем статус "в ожидании"
	booking.Status = "pending"

	// Устанавливаем временные метки
	now := time.Now()
	booking.CreatedAt = now
	booking.UpdatedAt = now

	return s.repository.CreateBooking(booking)
}

func (s *Service) GetBookingByID(id string) (domain.Booking, error) {
	return s.repository.GetBookingByID(id)
}

func (s *Service) GetBookingsByItemID(itemID int) ([]domain.Booking, error) {
	return s.repository.GetBookingsByItemID(itemID)
}

func (s *Service) GetBookingsByUserID(userID string) ([]domain.Booking, error) {
	return s.repository.GetBookingsByUserID(userID)
}

func (s *Service) UpdateBookingStatus(id string, status string) error {
	// Проверяем, что статус валидный
	validStatuses := map[string]bool{
		"pending":   true,
		"approved":  true,
		"rejected":  true,
		"cancelled": true,
		"completed": true,
	}

	if !validStatuses[status] {
		return errs.ErrInvalidStatus
	}

	return s.repository.UpdateBookingStatus(id, status)
}

func (s *Service) CheckItemAvailability(itemID int, startDate, endDate time.Time) (domain.BookingAvailability, error) {
	// Проверяем, что даты валидны
	if startDate.After(endDate) {
		return domain.BookingAvailability{}, errs.ErrInvalidDateRange
	}

	// Проверяем, что дата начала не в прошлом
	if startDate.Before(time.Now()) {
		return domain.BookingAvailability{}, errs.ErrPastDate
	}

	// Проверяем доступность в репозитории
	available, err := s.repository.CheckItemAvailability(itemID, startDate, endDate)
	if err != nil {
		return domain.BookingAvailability{}, err
	}

	// Если вещь доступна, получаем существующие бронирования для информации
	var bookings []domain.Booking
	if available {
		bookings, err = s.repository.GetBookingsByItemID(itemID)
		if err != nil {
			return domain.BookingAvailability{}, err
		}
	}

	return domain.BookingAvailability{
		Available: available,
		Bookings:  bookings,
	}, nil
}
