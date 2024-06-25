//go:build wireinject
// +build wireinject

package simples

import "github.com/google/wire"

func InitializeService(isError bool) *SimpleService {
	wire.Build(NewSimpleService, NewSimpleRepository)
	return nil
}

func InitializeDatabaseMultipleBinding() *Database {
	wire.Build(NewDatabaseRepo, NewDatabaseMongodbDsn, NewDatabasePostgresDsn)
	return nil
}

var fooService = wire.NewSet(NewFooService, NewFooRepository)
var barService = wire.NewSet(NewBarService, NewBarRepository)

func InitializeGroupingSetFooBarControlllerMultiple() *FooBarController {
	wire.Build(NewFooBarController, fooService, barService)
	return nil
}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

func InitializedFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}
