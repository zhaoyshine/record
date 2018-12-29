package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "报废欠赵耀的钱已经还清，好借好还，再借不难！给报废一个赞！",
		})
	})
	r.Run(":4000")
}
