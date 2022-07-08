package user

import (
	"fmt"
	"orderuserapi/utils"

	"orderuserapi/my_struct"
)

func CheckUser(userid int) *utils.RestErr {
	var user my_struct.User
	result := Db.Table("user").First(&user, my_struct.User{UserId: userid})
	if result.RowsAffected == 0 {
		fmt.Println("Already not exist", result.RowsAffected, result)
		return utils.AlreadyExistError("Already Not Exist")

	} else {
		fmt.Println("Already exist", result.RowsAffected, result)

		return nil
	}
}
func Save(OrderBucket *my_struct.Order) *utils.RestErr {
	result := Db.Table("order").Create(&OrderBucket)
	if result.Error != nil {
		return utils.NewInternalServerError("Error in create Query")
	}
	return nil

}

func Get(userid int) ([]my_struct.Order, *utils.RestErr) {
	var order []my_struct.Order
	result := Db.Table("order").Where("user_id=?", userid).Find(&order)
	fmt.Println(order)
	if result.RowsAffected == 0 {
		return nil, utils.NewInternalServerError("Error in Get Query")
	} else {

		return order, nil
	}

}

var user my_struct.User
var order my_struct.Order

func Delete(userid int) *utils.RestErr {
	fmt.Println(userid)
	resulto := Db.Table("order").Where("user_id=?", userid).Delete(&order)
	resultu := Db.Table("user").Where("user_id=?", userid).Delete(&user)
	if resulto.RowsAffected == 0 && resultu.RowsAffected == 0 {
		return utils.NewInternalServerError("No record Found")
	} else {
		return nil
	}
	// if resulto.RowsAffected != 0 {
	// 	fmt.Println("deleting user", resulto)
	// 	resultu := Db.Table("user").Where("user_id=?", userid).Delete(&user)
	// 	fmt.Println(resultu)
	// 	if resultu.RowsAffected == 0 {
	// 		return utils.NewInternalServerError("No Record Found")
	// 	} else {
	// 		return nil
	// 	}
	// } else {
	// 	return nil
	// }
}
