package services

type PageInfo struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	LastPage  int `json:"lastPage"`
	TotalData int `json:"totalData"`
}

type ResponseList struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo PageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type User struct {
	Id          int    `json:"id" form:"id"`
	FullName    string `json:"fullName" form:"fullName"`
	Email       string `json:"email" form:"email" binding:"email"`
	Password    string `json:"password" form:"password"`
	Address     string `json:"address" form:"address"`
	Picture     string `json:"picture" form:"picture"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Role        string `json:"role" form:"role"`
}
type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
