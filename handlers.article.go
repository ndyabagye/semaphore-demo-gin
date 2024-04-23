package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// call the HTML method of the context to render a template
	// Call the render function with the name of the template to render

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func getArticle(c *gin.Context) {
	// check if  the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// check if the article exisits
		if article, err := getArticleByID(articleID); err == nil {
			// call the html method of the context to render a template
			c.HTML(
				// set the http status to 200 (OK)
				http.StatusOK,
				// use the article.html template
				"article.html",
				// pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			// if the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// if an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// resonsd with xml
		c.XML(http.StatusOK, data["payload"])
	default:
		// respond with html
		c.HTML(http.StatusOK, templateName, data)
	}
}
