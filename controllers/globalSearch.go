package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func GlobalSearch(c *gin.Context) {
	elasticSearchUrl := os.Getenv("ELASTICSEARCH_URL") + "/_search"

	request, err := http.NewRequest(http.MethodGet, elasticSearchUrl, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to create request object for /GET endpoint:" + err.Error(),
		})

		return
	}

	request.Header.Add("Content-type", "application/json; charset=utf-8")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error making request to elastic search " + err.Error(),
		})

		return
	}

	body, reqError := io.ReadAll(response.Body)

	if reqError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read response body: " + err.Error(),
		})

		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to close response body: " + err.Error(),
			})

			return
		}
	}(response.Body)

	// now we need to read this to our native Golang type
	// as the map data structure is close to JSON, we could use it
	// in fact we could this for most of the wire formats.
	data := make(map[string]interface{})

	// this step de-serializes JSON data to our native Golang data
	jsonErr := json.Unmarshal(body, &data)

	if jsonErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to unmarshal: " + jsonErr.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
