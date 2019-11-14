package middleware

import (
	"project/src/utils/codewrap"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func parseMessage(message string) string {
	switch message {
	case mongo.ErrNoDocuments.Error():
		return "数据库未查到此数据"
	default:
		return message
	}
}

//Response 后置处理中间件
func Response(c *gin.Context) {
	c.Next()
	statusCode := c.Writer.Status()

	if statusCode == 200 || statusCode == 201 {
		code := 0
		message := ""

		if codeStr, exists := c.Get("code"); exists {
			code = codeStr.(int)
		}

		if messageStr, exists := c.Get("message"); exists {
			if strMess, ok := messageStr.(string); ok {
				message = strMess
			} else if errMess, ok := messageStr.(error); ok {
				message = errMess.Error()
			}
			eCode, eMessage := codewrap.Parse(message)
			if eCode != 0 {
				code = eCode
			}
			message = eMessage
		}

		data, _ := c.Get("data")
		c.JSON(http.StatusOK, gin.H{"code": code, "data": data, "message": parseMessage(message)})
	}
}
