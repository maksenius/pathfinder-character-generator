package characteristic

import "fmt"

type DiceCharacteristicGenerator interface {
	GetNumbers() []int
}

type Service struct {
	diceCharacteristicGenerator DiceCharacteristicGenerator
}

func NewService(diceCharacteristicGenerator DiceCharacteristicGenerator) *Service {
	return &Service{diceCharacteristicGenerator: diceCharacteristicGenerator}
}

func (s *Service) SetCharacteristic() {
	fmt.Println(s.diceCharacteristicGenerator.GetNumbers())
}
