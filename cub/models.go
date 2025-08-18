package main

import "time"

type User struct {
	Id         int
	Email      string
	FirstName  string
	LastName   string
	Password   string
	UserActive int
	IsAdmin    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PlanId     int
	Plan       Plan
}

type Plan struct {
	Id         int
	PlanName   string
	PlanAmount int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
