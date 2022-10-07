package users

import (
	"math"
	"sync"
)

type User struct {
	ID       int64
	Name     string
	Tryings  int64
	RightAns int64
}

func (u *User) GetStat() int64 {
	return int64(math.Round(float64(u.RightAns) / float64(u.RightAns)))
}

type SyncMap struct {
	Mutex sync.RWMutex
	Data  map[int64]User
}

func New() *SyncMap {
	return &SyncMap{}
}
