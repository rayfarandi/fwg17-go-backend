package service

import (
	"database/sql"
	"time"
)

//GLOBAL RESPONSE

type PageInfo struct {
	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
	NextPage    int `json:"nextPage"`
	PrevPage    int `json:"prevPage"`
	Limit       int `json:"limit"`
	TotalData   int `json:"totalData"`
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

// VARIANT RESPONSE
type Variants struct {
	Id              int          `db:"id" json:"id"`
	Name            string       `db:"name" json:"name" form:"name" binding:"required,min=3"`
	AdditionalPrice int          `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice" binding:"required,numeric"`
	CreatedAt       time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoV struct {
	Data  []Variants
	Count int
}

// USER RESPONSE
type User struct {
	Id          int            `db:"id" json:"id"`
	FullName    string         `db:"fullName" json:"fullName" form:"fullName"`
	Email       string         `db:"email" json:"email" form:"email"`
	Password    string         `db:"password" json:"-" form:"password"`
	Address     sql.NullString `db:"address" json:"address" form:"address"`
	Picture     string         `db:"picture" json:"picture"`
	PhoneNumber sql.NullString `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        string         `db:"role" json:"role" form:"role"`
	CreatedAt   time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

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

type Info struct {
	Data  []User
	Count int
}

// FORM RESET
type FormReset struct {
	Email           string `form:"email"`
	Otp             string `form:"otp"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword" binding:"eqfield=Password"`
}

// RESET PASSWORD RESPONSE
type ForgotPassword struct {
	Id        int           `db:"id" json:"id"`
	Otp       string        `db:"otp" json:"otp"`
	UserId    sql.NullInt64 `db:"userId" json:"userId"`
	Email     string        `db:"email" json:"email"`
	CreatedAt time.Time     `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime  `db:"updatedAt" json:"updatedAt"`
}
type FPForm struct {
	Id        int          `db:"id" json:"id"`
	Otp       *string      `db:"otp" json:"otp" form:"otp" binding:"required"`
	UserId    *int         `db:"userId" json:"userId" form:"userId" binding:"required"`
	Email     *string      `db:"email" json:"email" form:"email" binding:"required"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoFP struct {
	Data  []ForgotPassword
	Count int
}

// SIZE RESPONSE
type Sizes struct {
	Id              int          `db:"id" json:"id"`
	Size            string       `db:"size" json:"size" form:"size"`
	AdditionalPrice int          `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type SizesForm struct {
	Id   int    `db:"id" json:"id"`
	Size string `db:"size" json:"size" form:"size" binding:"required,min=3"`
	// AdditionalPrice *int         `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice" binding:"required,numeric"`
	AdditionalPrice int          `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice" binding:"required,numeric"`
	CreatedAt       time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoS struct {
	Data  []Sizes
	Count int
}

// TAGS RESPONSE
type Tags struct {
	Id        int          `db:"id" json:"id"`
	Name      string       `db:"name" json:"name" form:"name" binding:"required,min=3"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoT struct {
	Data  []Tags
	Count int
}

// PROMO RESPONSE
type Promo struct {
	Id            int            `db:"id" json:"id"`
	Name          string         `db:"name" json:"name" form:"name"`
	Code          string         `db:"code" json:"code" form:"code"`
	Description   sql.NullString `db:"description" json:"description" form:"description"`
	Percentage    float64        `db:"percentage" json:"percentage" form:"percentage"`
	IsExpired     sql.NullBool   `db:"isExpired" json:"isExpired" form:"isExpired"`
	MaximumPromo  int            `db:"maximumPromo" json:"maximumPromo" form:"maximumPromo"`
	MinimumAmount int            `db:"minimumAmount" json:"minimumAmount" form:"minimumAmount"`
	CreatedAt     time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type PromoForm struct {
	Id            int          `db:"id" json:"id"`
	Name          *string      `db:"name" json:"name" form:"name" binding:"required,min=3"`
	Code          *string      `db:"code" json:"code" form:"code" binding:"required,min=3"`
	Description   *string      `db:"description" json:"description" form:"description"`
	Percentage    *float64     `db:"percentage" json:"percentage" form:"percentage" binding:"required"`
	IsExpired     *bool        `db:"isExpired" json:"isExpired" form:"isExpired"`
	MaximumPromo  *int         `db:"maximumPromo" json:"maximumPromo" form:"maximumPromo" binding:"required"`
	MinimumAmount *int         `db:"minimumAmount" json:"minimumAmount" form:"minimumAmount" binding:"required"`
	CreatedAt     time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoPo struct {
	Data  []Promo
	Count int
}

// PRODUCTVARIANT RESPONSE
type ProductVariants struct {
	Id        int          `db:"id" json:"id"`
	ProductId int          `db:"productId" json:"productId" form:"productId" binding:"required"`
	VariantId int          `db:"variantId" json:"variantId" form:"variantId" binding:"required"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoPV struct {
	Data  []ProductVariants
	Count int
}

// PRODUCTRATING RESPONSE
type ProductRatings struct {
	Id            int            `db:"id" json:"id"`
	ProductId     int            `db:"productId" json:"productId"`
	Rate          int            `db:"rate" json:"rate"`
	ReviewMessage sql.NullString `db:"reviewMessage" json:"reviewMessage"`
	UserId        int            `db:"userId" json:"userId"`
	CreatedAt     time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type PRForm struct {
	Id            int          `db:"id" json:"id"`
	ProductId     *int         `db:"productId" json:"productId" form:"productId" binding:"required,numeric"`
	Rate          *int         `db:"rate" json:"rate" form:"rate" binding:"required,eq=5|eq=4|eq=3|eq=2|eq=1"`
	ReviewMessage *string      `db:"reviewMessage" json:"reviewMessage" form:"reviewMessage"`
	UserId        *int         `db:"userId" json:"userId" form:"userId" binding:"required,numeric"`
	CreatedAt     time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoPR struct {
	Data  []ProductRatings
	Count int
}

// CATEGORIES RESPONSE
type Categories struct {
	Id        int          `db:"id" json:"id"`
	Name      *string      `db:"name" json:"name" form:"name"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoC struct {
	Data  []Categories
	Count int
}

// MESSAGE RESPONSE
type Message struct {
	Id          int          `db:"id" json:"id"`
	RecipientId int          `db:"recipientId" json:"recipientId" form:"recipientId" binding:"required,numeric"`
	SenderId    int          `db:"senderId" json:"senderId" form:"senderId" binding:"required,numeric"`
	Text        string       `db:"text" json:"text" form:"text" binding:"required"`
	CreatedAt   time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoM struct {
	Data  []Message
	Count int
}

// PRODUCT CATEGORIE RESPONSE
type ProductCategories struct {
	Id         int          `db:"id" json:"id"`
	ProductId  int          `db:"productId" json:"productId" form:"productId" binding:"required,numeric"`
	CategoryId int          `db:"categoryId" json:"categoryId" form:"categoryId" binding:"required,numeric"`
	CreatedAt  time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt  sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoPC struct {
	Data  []ProductCategories
	Count int
}
