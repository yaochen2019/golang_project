package dto

// 接收前端发送来的用户数据参数
type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,email,first_is_a" message:"userName is wrong" required_err:"userName is null"`
	Password string `json:"password" binding:"required" message:"password need value"`
}
