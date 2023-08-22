package controllers

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CreateIndex(c *gin.Context) {
	elasticSearchUrl := os.Getenv("ELASTICSEARCH_URL") + "users"

	fmt.Println("url: ", elasticSearchUrl)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	body := bytes.NewBuffer([]byte{})

	// Send a GET request to the API using the custom client
	response, err := client.Post(elasticSearchUrl, "application/json", body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch data from API: " + err.Error(),
		})
		return
	}
	defer response.Body.Close()

	responseData := make([]byte, 0)
	_, err = response.Body.Read(responseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read response data",
		})
		return
	}

	// Respond with the fetched data
	c.JSON(http.StatusOK, gin.H{
		"data": string(responseData),
	})
}
