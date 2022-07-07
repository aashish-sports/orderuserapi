package services

import (
	"orderuserapi/my_struct"
	"orderuserapi/user"
	"orderuserapi/utils"
)

func Save(OrderBucket *my_struct.Order) *utils.RestErr {
	if err := user.Save(OrderBucket); err != nil {
		return nil
	}
	return utils.NewInternalServerError("Not saved ")
}

func CheckUser(userid int) *utils.RestErr {
	//fmt.Println("services", userid)
	if err := user.CheckUser(userid); err == nil {
		return nil
	}
	return utils.AlreadyExistError("Already Not Exist")
}

func Get(userid int64) {

}
