package service

type SendEmail struct{
	Email string `json:"email" form:"email"`
	Passwd string `json:"passwd" form:"passwd"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}
