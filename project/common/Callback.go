package common

import (
	// "context"
	"net/http"
	"project/controllers"

	"github.com/gin-gonic/gin"
)

// func Login(c *gin.Context) {
// 	googleConfig := SetupConfig()
// 	url := googleConfig.AuthCodeURL(RandomState)
// 	// fmt.Println(url)
// 	c.Redirect(http.StatusSeeOther, url)
// }



func CallBack(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if err:=controllers.UpdateToken(accessToken);err!=nil{
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to login")
		return
	}
	
	c.IndentedJSON(http.StatusOK, accessToken)
}
