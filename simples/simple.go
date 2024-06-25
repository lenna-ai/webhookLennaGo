package simples

type SimpleRepository struct {
	errors bool
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		errors: isError,
	}
}

type SimpleService struct {
	Simple *SimpleRepository
}

func NewSimpleService(simple *SimpleRepository) *SimpleService {
	return &SimpleService{
		Simple: simple,
	}
}
