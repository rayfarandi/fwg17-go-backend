package services

import (
	"database/sql"
	"time"
)

type InfoUser struct {
	Data  []User
	Count int
}

// User Models
type User struct {
	Id          int            `db:"id" json:"id"`
	FullName    string         `db:"fullName" json:"fullName" form:"fullName"`
	Email       string         `db:"email" json:"email" form:"email" form:"email"`
	Password    string         `db:"password" json:"-" form:"password" form:"password"`
	Address     sql.NullString `db:"address" json:"address" form:"address"`
	Picture     string         `db:"picture" json:"picture" form:"picture"`
	PhoneNumber sql.NullString `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        string         `db:"role" json:"role" form:"role"`
	CreatedAt   time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

// User Form
type UserForm struct {
	Id          int          `db:"id" json:"id"`
	FullName    *string      `db:"fullName" json:"fullName" form:"fullName"`
	Email       *string      `db:"email" json:"email" form:"email"`
	Password    string       `db:"password" json:"-" form:"password"`
	Address     *string      `db:"address" json:"address" form:"address"`
	Picture     string       `db:"picture" json:"picture"`
	PhoneNumber *string      `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        *string      `db:"role" json:"role" form:"role"`
	CreatedAt   time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoProduct struct {
	Data []Product
	// Data  interface{}
	Count int
}

// Product models
type Product struct {
	Id            int            `db:"id" json:"id"`
	Name          string         `db:"name" json:"name" form:"name"`
	Description   sql.NullString `db:"description" json:"description"`
	Image         string         `db:"image" json:"image" form:"image"`
	Discount      sql.NullInt64  `db:"discount" json:"discount"`
	IsRecommended bool           `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	BasePrice     int            `db:"basePrice" json:"basePrice" form:"basePrice"`
	// Category      sql.NullString `db:"category" json:"category"`
	// Tag           sql.NullString `db:"tag" json:"tag"`
	// Rating        sql.NullInt64  `db:"rating" json:"rating"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

// Product Form
type ProductForm struct {
	Id            int          `db:"id" json:"id"`
	Name          *string      `db:"name" json:"name" form:"name"`
	Description   *string      `db:"description" json:"description"`
	BasePrice     *int         `db:"basePrice" json:"basePrice" form:"basePrice"`
	Image         string       `db:"image" json:"image"`
	Discount      *int         `db:"discount" json:"discount" form:"discount"`
	IsRecommended *bool        `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	CreatedAt     time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoProductSize struct {
	Data  interface{}
	Count int
}
type InfoProductVariant struct {
	Data  interface{}
	Count int
}

type InfoTags struct {
	Data  interface{}
	Count int
}
type InfoProductTags struct {
	Data  interface{}
	Count int
}

type InfoProductRatings struct {
	Data  interface{}
	Count int
}
type InfoCategories struct {
	Data  interface{}
	Count int
}
type InfoProductCategories struct {
	Data  interface{}
	Count int
}
type InfoPromo struct {
	Data  interface{}
	Count int
}

// General response //
type PageInfo struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	NextPage  int `json:"nextPage"`
	PrevPage  int `json:"prevPage"`
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
