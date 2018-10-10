//+build wireinject

package main

import "github.com/google/go-cloud/wire"

func InitialiseService() Service {
	wire.Build(NewDB, NewService)
	return Service{}
}