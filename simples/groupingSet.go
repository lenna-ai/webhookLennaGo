package simples

type FooRepository struct {
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{
		FooRepository: fooRepository,
	}
}

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{
		BarRepository: barRepository,
	}
}

type FooBarController struct {
	*FooService
	*BarService
}

func NewFooBarController(fooService *FooService, barService *BarService) *FooBarController {
	return &FooBarController{
		FooService: fooService,
		BarService: barService,
	}
}
