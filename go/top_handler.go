package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var TagMap map[int64]string
var InvTagMap map[string]int64
var Tags []*Tag

func saveTags(ctx context.Context) error {
	clear(TagMap)
	clear(InvTagMap)
	clear(Tags)

	var tagModels []TagModel
	if err := dbConn.SelectContext(ctx, &tagModels, "SELECT * FROM tags"); err != nil {
		return err
	}
	Tags = make([]*Tag, 0, 200)
	TagMap = make(map[int64]string)
	InvTagMap = make(map[string]int64)

	for _, tag := range tagModels {
		TagMap[tag.ID] = tag.Name
		InvTagMap[tag.Name] = tag.ID
		Tags = append(Tags, &Tag{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}
	return nil
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TagModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type TagsResponse struct {
	Tags []*Tag `json:"tags"`
}

func getTagHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &TagsResponse{
		Tags: Tags,
	})
}

// 配信者のテーマ取得API
// GET /api/user/:username/theme
func getStreamerThemeHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if err := verifyUserSession(c); err != nil {
		// echo.NewHTTPErrorが返っているのでそのまま出力
		c.Logger().Printf("verifyUserSession: %+v\n", err)
		return err
	}

	username := c.Param("username")

	userModel := UserModel{}
	err := dbConn.GetContext(ctx, &userModel, "SELECT id, dark_mode FROM users WHERE name = ?", username)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "not found user that has the given username")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user: "+err.Error())
	}

	theme := Theme{
		ID:       userModel.ID,
		DarkMode: userModel.DarkMode,
	}

	return c.JSON(http.StatusOK, theme)
}
