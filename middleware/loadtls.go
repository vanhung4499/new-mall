package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			// If an error occurs, do not continue
			fmt.Println(err)
			return
		}
		// Continue to process
		c.Next()
	}
}
