package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Brianllp/go_practice/database"
	"github.com/Brianllp/go_practice/models"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Sys struct {
		Type string `json:"type"`
	} `json:"sys"`
	Total int `json:"total"`
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
	Items []struct {
		Metadata struct {
			Tags []interface{} `json:"tags"`
		} `json:"metadata"`
		Sys struct {
			Space struct {
				Sys struct {
					Type     string `json:"type"`
					LinkType string `json:"linkType"`
					ID       string `json:"id"`
				} `json:"sys"`
			} `json:"space"`
			ID          string    `json:"id"`
			Type        string    `json:"type"`
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			Environment struct {
				Sys struct {
					ID       string `json:"id"`
					Type     string `json:"type"`
					LinkType string `json:"linkType"`
				} `json:"sys"`
			} `json:"environment"`
			Revision    int `json:"revision"`
			ContentType struct {
				Sys struct {
					Type     string `json:"type"`
					LinkType string `json:"linkType"`
					ID       string `json:"id"`
				} `json:"sys"`
			} `json:"contentType"`
			Locale string `json:"locale"`
		} `json:"sys"`
		Fields struct {
			Title string `json:"title"`
			Body  struct {
				NodeType string `json:"nodeType"`
				Data     struct {
				} `json:"data"`
				Content []struct {
					NodeType string `json:"nodeType"`
					Content  []struct {
						NodeType string        `json:"nodeType"`
						Value    string        `json:"value"`
						Marks    []interface{} `json:"marks"`
						Data     struct {
						} `json:"data"`
					} `json:"content"`
					Data struct {
					} `json:"data"`
				} `json:"content"`
			} `json:"body"`
		} `json:"fields"`
	} `json:"items"`
}

func GetEntries(c echo.Context) error {
	db, _ := database.ConnectDB(false)

	SPACE_ID := os.Getenv("CONTENTFUL_SPACE_ID")
	API_KEY := os.Getenv("CONTENTFUL_DELIVERY_API_KEY")
	ENVIRONMENT_ID := os.Getenv("CONTENTFUL_ENVIRONMENT_ID")

	url := "https://cdn.contentful.com/spaces/" + SPACE_ID + "/environments/" + ENVIRONMENT_ID + "/entries?access_token=" + API_KEY

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	jsonBytes := []byte(body)
	data := new(Response)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	for _, item := range data.Items {
		uuid := item.Sys.ID
		title := item.Fields.Title
		body := item.Fields.Body.Content[0].Content[0].Value

		entry := models.Entry{UUID: uuid, Title: title, Body: body}
		models.CreateOrUpdateEntry(db, entry)
	}

	return c.String(http.StatusOK, string(body))
}

func GetContentfulEntries() {
	db, _ := database.ConnectDB(false)

	SPACE_ID := os.Getenv("CONTENTFUL_SPACE_ID")
	API_KEY := os.Getenv("CONTENTFUL_DELIVERY_API_KEY")
	ENVIRONMENT_ID := os.Getenv("CONTENTFUL_ENVIRONMENT_ID")

	url := "https://cdn.contentful.com/spaces/" + SPACE_ID + "/environments/" + ENVIRONMENT_ID + "/entries?access_token=" + API_KEY

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	jsonBytes := []byte(body)
	data := new(Response)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	for _, item := range data.Items {
		uuid := item.Sys.ID
		title := item.Fields.Title
		body := item.Fields.Body.Content[0].Content[0].Value

		entry := models.Entry{UUID: uuid, Title: title, Body: body}
		models.CreateOrUpdateEntry(db, entry)
	}
}

func GetEntriesFromDB(c echo.Context) error {
	db, _ := database.ConnectDB(false)

	entries := models.IndexEntries(db)
	return c.JSON(http.StatusOK, entries)
}
