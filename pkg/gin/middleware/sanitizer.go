package middleware

import (
	"bytes"
	"encoding/json"
	"html"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

// InputSanitizerMiddleware sanitizes incoming request bodies
func InputSanitizerMiddleware() gin.HandlerFunc {
	policy := CustomUGCPolicy()

	return func(c *gin.Context) {
		// Read the body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Parse JSON and sanitize
		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err == nil {
			data = sanitizeMap(data, policy)

			// Replace the body with sanitized JSON
			sanitizedBody, _ := json.Marshal(data)
			c.Request.Body = io.NopCloser(bytes.NewReader(sanitizedBody))
		}

		c.Next() // Continue to the next handler
	}
}

// CustomUGCPolicy creates a bluemonday policy allowing certain attributes
func CustomUGCPolicy() *bluemonday.Policy {
	policy := bluemonday.UGCPolicy()

	// Allow `class` attribute on <img>
	policy.AllowAttrs("src", "alt", "class").OnElements("img")

	// Allow `class` attribute on <p>
	policy.AllowAttrs("class").OnElements("p")

	return policy
}

func sanitizeMap(data map[string]interface{}, policy *bluemonday.Policy) map[string]interface{} {
	for key, value := range data {
		switch v := value.(type) {
		case string:
			// Sanitize and unescape the string
			data[key] = unescapeEntities(policy.Sanitize(v))
		case map[string]interface{}:
			data[key] = sanitizeMap(v, policy)
		case []interface{}:
			data[key] = sanitizeSlice(v, policy)
		}
	}
	return data
}

func sanitizeSlice(data []interface{}, policy *bluemonday.Policy) []interface{} {
	for i, value := range data {
		switch v := value.(type) {
		case string:
			// Sanitize and unescape the string
			data[i] = unescapeEntities(policy.Sanitize(v))
		case map[string]interface{}:
			data[i] = sanitizeMap(v, policy)
		case []interface{}:
			data[i] = sanitizeSlice(v, policy)
		}
	}
	return data
}

// unescapeEntities unescapes HTML entities such as &#39; back to original characters
func unescapeEntities(input string) string {
	return html.UnescapeString(input)
}
