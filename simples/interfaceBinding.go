package simples

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
}

type SayHelloImpl struct {
}

func (shi *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{}
}
