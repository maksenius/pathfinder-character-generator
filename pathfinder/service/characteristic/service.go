package characteristic

import (
	"pathfinderCharacterGenerator/pathfinder/model"
)

type Service struct {
	diceCharacteristicGenerator DiceCharacteristicGenerator
	classesCharacteristicSetup  ClassesCharacteristicSetup
}

type DiceCharacteristicGenerator interface {
	GetNumbers() []int
}

type ClassesCharacteristicSetup interface {
	SetPriorityCharacteristic([]int) *model.Characteristic
}

func NewService(diceCharacteristicGenerator DiceCharacteristicGenerator) *Service {
	return &Service{diceCharacteristicGenerator: diceCharacteristicGenerator}
}

func (s *Service) GetCharacteristic() *model.Characteristic {
	numbers := s.diceCharacteristicGenerator.GetNumbers()
	return s.classesCharacteristicSetup.SetPriorityCharacteristic(numbers)
}
