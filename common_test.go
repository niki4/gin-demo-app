package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpArticleList []article

// Setup before executing the test functions
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run()) // run the tests, then exit
}

// Helper to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	router := gin.Default()
	if withTemplates {
		router.LoadHTMLGlob("templates/*")
	}
	return router
}

// Helper to process a request and test its response
func testHTTPResponse(
	t *testing.T,
	r *gin.Engine,
	req *http.Request,
	f func(w *httptest.ResponseRecorder) bool,
) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the passed request
	r.ServeHTTP(w, req)

	if !f(w) { // check if test was successful
		t.Fail()
	}
}

// Helper to store the main lists into the temporary one for testing
func saveLists() {
	tmpArticleList = articleList
}

// Helper to restore the main lists from the temporary one
func restoreLists() {
	articleList = tmpArticleList
}
