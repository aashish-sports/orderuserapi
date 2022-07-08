package services

import (
	"fmt"
	"orderuserapi/my_struct"
	"orderuserapi/user"
	"orderuserapi/utils"
)

func Save(OrderBucket *my_struct.Order) *utils.RestErr {
	if err := user.Save(OrderBucket); err != nil {
		return err
	}
	return nil
}

func CheckUser(userid int) *utils.RestErr {
	//fmt.Println("services", userid)
	if err := user.CheckUser(userid); err == nil {
		return nil
	}
	return utils.AlreadyExistError("Already Not Exist")
}

func Get(userid int) (*[]my_struct.Order, *utils.RestErr) {
	getOrder, err := user.Get(userid)
	if err == nil {
		fmt.Println("Services: Data get SucessFully")
		return &getOrder, nil
	} else {
		fmt.Println("Services: Error Occur")
		return nil, err
	}

}
func Delete(userid int) *utils.RestErr {
	if err := user.Delete(userid); err != nil {
		return err
	} else {
		return nil
	}
}
