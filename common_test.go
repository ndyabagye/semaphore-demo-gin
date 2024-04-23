package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

// this function is used for setup before executing the test function
func TestMain(m *testing.M) {
	// set gin to test mode
	gin.SetMode(gin.TestMode)

	// run the other tests
	os.Exit(m.Run())
}

// helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// create a response recorder
	w := httptest.NewRecorder()

	// create the service and process the above request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// this function is used to store the main lists into the temporary one for testing
func saveLists() {
	tmpArticleList = articleList
}

// this function is used to restore the main lists from the temporary one
func restoreLists() {
	articleList = tmpArticleList
}
