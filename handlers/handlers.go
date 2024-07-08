package handlers

import (
	"net/http"

	"github.com/artemis13/platform-go-challenge/models"
	"github.com/artemis13/platform-go-challenge/storage"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetUserFavorites(c echo.Context) error {
	userIDStr := c.Param("id")

	userID, err := storage.ConvertStringToUint(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	// Use the storage package's method to safely retrieve the user
	user, exists := storage.GetUser(userID)
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user.Favorites)
}

func AddUserFavorite(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := storage.ConvertStringToUint(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	var asset models.Asset
	if err := c.Bind(&asset); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	asset.ID = uuid.New().String()

	storage.UpdateUserFavorites(userID, asset)
	return c.JSON(http.StatusCreated, asset)
}

func RemoveUserFavorite(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := storage.ConvertStringToUint(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	assetID := c.Param("asset_id")

	user, exists := storage.GetUser(userID)
	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	for i, asset := range user.Favorites {
		if asset.ID == assetID {
			user.Favorites = append(user.Favorites[:i], user.Favorites[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "asset not found")
}

func EditUserFavorite(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := storage.ConvertStringToUint(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}
	assetID := c.Param("asset_id")
	var newDescription struct {
		Description string `json:"description"`
	}
	if err := c.Bind(&newDescription); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, exists := storage.GetUser(userID)
	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	for i, asset := range user.Favorites {
		if asset.ID == assetID {
			user.Favorites[i].Description = newDescription.Description
			return c.JSON(http.StatusOK, user.Favorites[i])
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "asset not found")
}
