package contactsbook

import "fmt"

func main() {
	fmt.Println("Satarting Contact Books API")
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
