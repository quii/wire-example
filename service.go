package main

import "fmt"

type DB interface {
	Get() int
}

func NewService(db DB) Service {
	return Service{db}
}

type Service struct{
	db DB
}

func (s *Service) Print(){
	fmt.Println("PRINTING!", s.db.Get())
}
