package services

type InfoUser struct {
	// Data []User
	Data  interface{}
	Count int
}

type InfoProduct struct {
	Data  interface{}
	Count int
}

type InfoProductSize struct {
	Data  interface{}
	Count int
}
type InfoProductSizeVariant struct {
	Data  interface{}
	Count int
}
type InfoTags struct {
	Data  interface{}
	Count int
}

// General response //
type PageInfo struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"totalPage"`
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

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
