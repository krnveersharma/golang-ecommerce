package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID
	FirstName      *string
	LastName       *string
	Password       *string
	Email          *string
	Phone          *string
	Token          *string
	RefreshToken   *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         string
	UserCart       []ProductUser
	AddressDetails []Address
	OrderStatus    []Order
}

type Product struct {
	ProductID   primitive.ObjectID
	ProductName *string
	Price       *uint64
	Rating      *uint8
	Image       *string
}

type ProductUser struct {
	ProductID   primitive.ObjectID
	ProductName *string
	Price       int
	Rating      *uint
	Image       *string
}

type Address struct {
	AddressID primitive.ObjectID
	House     *string
	Street    *string
	City      *string
	Pincode   *string
}

type Order struct {
	OrderID       primitive.ObjectID
	OrderCart     []ProductUser
	OrderedAt     time.Time
	Price         int
	Discount      *int
	PaymentMethod Payment
}

type Payment struct {
	Digital bool
	COD     bool
}
