package main

import (
	"sync"

	"gorm.io/gorm"
)

type Bridge struct {
	DB        *gorm.DB
	WaitGroup *sync.WaitGroup
}
