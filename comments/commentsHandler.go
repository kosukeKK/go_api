package comments

import "github.com/gin-gonic/gin"

type Comment struct {
	Title string `json: title`
	Text  string `json: text`
}


// v1/comments
func Handler(c *gin.Context) {
	comment := []Comment{
		{Title: "おもしろい", Text: "良い本だと思います"},
		{Title: "つまらない", Text: "ひどい本でした"},
	}
	var jsonMap map[string]interface{} = make(map[string]interface{})
	jsonMap["comments"] = comment
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")

	c.JSON(200, jsonMap)
}
