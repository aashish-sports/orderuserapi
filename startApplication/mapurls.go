package startapplication

import "orderuserapi/controller"

func mapurls() {
	router.POST("/user", controller.CreateOrder)
	router.GET("/user/:user_id", controller.GetUser)
	router.DELETE("/user/:user_id", controller.DeleteUser)

}
