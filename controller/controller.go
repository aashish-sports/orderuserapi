package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"orderuserapi/my_struct"
	"orderuserapi/services"
	"orderuserapi/utils"

	"strconv"

	//"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	// var orderObject my_struct.Order
	// if err := c.ShouldBindJSON(&orderObject); err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// }
	// c.JSON(http.StatusOK, &orderObject)
	var usersBucket my_struct.UserArray
	byteArray, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.Unmarshal(byteArray, &usersBucket); err != nil {
		fmt.Println(err)
		c.JSON(400, "error is in shouldBindJson")
		fmt.Println(err)
		return
	}
	for _, i := range usersBucket {
		var userIdBucket my_struct.User
		userIdBucket.UserId = i.UserId
		fmt.Println(i.UserId)
		if err := services.CheckUser(i.UserId); err == nil {
			for _, j := range i.Orders {
				var OrderBucket my_struct.Order
				OrderBucket.OrderId = j.OrderId
				OrderBucket.Price = j.Price
				OrderBucket.ProductName = j.ProductName
				OrderBucket.Qty = j.Qty
				OrderBucket.UserId = i.UserId
				fmt.Println("=>", OrderBucket, "<=>", userIdBucket)
				err := services.Save(&OrderBucket)
				if err != nil {
					c.JSON(http.StatusBadRequest, err)
					continue
				}
				c.JSON(http.StatusOK, "Sucessfully saved")

			}
		} else {
			continue
		}
	}

}

func GetUser(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErr := utils.NewBadRequest("Json bind error")
		c.JSON(400, newErr)
	}
	result, saveErr := services.Get(userid)
	if saveErr != nil {
		fmt.Println("Error Occur")
		c.JSON(saveErr.Status, saveErr)
	} else {
		fmt.Println("Data Get Successfully")
		c.JSON(http.StatusOK, result)
	}

}
func DeleteUser(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		fmt.Println(err)
	}
	er := services.Delete(userid)
	fmt.Println(er)
	if er != nil {
		fmt.Println(er)
		c.JSON(er.Status, er)
		return
	} else {
		c.JSON(http.StatusOK, "Sucessfully Deleted")
	}

}
