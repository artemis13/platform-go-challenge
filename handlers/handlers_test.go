package handlers_test

import (
	"encoding/json"
	"gwi-platformGoC/handlers"
	"gwi-platformGoC/models"
	"gwi-platformGoC/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupTestData() {
	storage.ClearUsers()
	storage.AddUser(&models.User{
		ID: 1,
		Favorites: []models.Asset{
			{ID: "1", Description: "Favorite 1"},
		},
	})
}

func TestGetUserFavorites(t *testing.T) {
	e := echo.New()
	setupTestData()

	req := httptest.NewRequest(http.MethodGet, "/users/1/favorites", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id/favorites")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, handlers.GetUserFavorites(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `[{"ID":"1","Type":0,"Description":"Favorite 1","Chart":null,"Insight":null,"Audience":null}]`, rec.Body.String())
	}
}

func TestAddUserFavorite(t *testing.T) {
	e := echo.New()
	setupTestData()

	assetJSON := `{"Type":0,"Description":"New Favorite"}`
	req := httptest.NewRequest(http.MethodPost, "/users/1/favorites", strings.NewReader(assetJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id/favorites")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, handlers.AddUserFavorite(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var asset models.Asset
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &asset))
		assert.NotEmpty(t, asset.ID)
		assert.Equal(t, "New Favorite", asset.Description)
	}
}

func TestRemoveUserFavorite(t *testing.T) {
	e := echo.New()
	setupTestData()

	req := httptest.NewRequest(http.MethodDelete, "/users/1/favorites/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id/favorites/:asset_id")
	c.SetParamNames("id", "asset_id")
	c.SetParamValues("1", "1")

	if assert.NoError(t, handlers.RemoveUserFavorite(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestEditUserFavorite(t *testing.T) {
	e := echo.New()
	setupTestData()

	newDescription := `{"description":"Updated Favorite"}`
	req := httptest.NewRequest(http.MethodPut, "/users/1/favorites/1", strings.NewReader(newDescription))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id/favorites/:asset_id")
	c.SetParamNames("id", "asset_id")
	c.SetParamValues("1", "1")

	if assert.NoError(t, handlers.EditUserFavorite(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var asset models.Asset
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &asset))
		assert.Equal(t, "Updated Favorite", asset.Description)
	}
}
