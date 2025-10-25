package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Nonameipal/P2P/internal/errs"
	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createBookingRequest struct {
	ItemID    int       `json:"item_id" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

type checkAvailabilityRequest struct {
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

func (c *Controller) CreateBooking(ctx *gin.Context) {
	var req createBookingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	userID := ctx.GetString("user_id")
	if userID == "" {
		c.handleError(ctx, errs.ErrInvalidToken)
		return
	}

	booking := domain.Booking{
		ID:        uuid.New().String(),
		ItemID:    req.ItemID,
		UserID:    userID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	if err := c.service.CreateBooking(booking); err != nil {
		c.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "booking created successfully"})
}

func (c *Controller) GetMyBookings(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		c.handleError(ctx, errs.ErrInvalidToken)
		return
	}

	bookings, err := c.service.GetBookingsByUserID(userID)
	if err != nil {
		c.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, bookings)
}

func (c *Controller) GetItemBookings(ctx *gin.Context) {
	itemID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	bookings, err := c.service.GetBookingsByItemID(itemID)
	if err != nil {
		c.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, bookings)
}

func (c *Controller) CheckItemAvailability(ctx *gin.Context) {
	itemID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	var req checkAvailabilityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	availability, err := c.service.CheckItemAvailability(itemID, req.StartDate, req.EndDate)
	if err != nil {
		c.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, availability)
}

func (c *Controller) UpdateBookingStatus(ctx *gin.Context) {
	bookingID := ctx.Param("id")
	if bookingID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking id"})
		return
	}

	status := ctx.Query("status")
	if status == "" {
		c.handleError(ctx, errs.ErrInvalidStatus)
		return
	}

	if err := c.service.UpdateBookingStatus(bookingID, status); err != nil {
		c.handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "booking status updated successfully"})
}
