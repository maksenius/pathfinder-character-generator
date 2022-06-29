package classes

type Service struct {
	ClassName string
}

func NewService(className string) *Service {
	return &Service{ClassName: className}
}
