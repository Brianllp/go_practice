package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func GetEntries(c echo.Context) error {
	SPACE_ID := os.Getenv("CONTENTFUL_SPACE_ID")
	API_KEY := os.Getenv("CONTENTFUL_DELIVERY_API_KEY")
	ENVIRONMENT_ID := os.Getenv("CONTENTFUL_ENVIRONMENT_ID")

	url := "https://cdn.contentful.com/spaces/" + SPACE_ID + "/environments/" + ENVIRONMENT_ID + "/entries?access_token=" + API_KEY

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("aaaaaaaaa")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(url)
	fmt.Println(string(body))

	return c.String(http.StatusOK, string(body))
}
