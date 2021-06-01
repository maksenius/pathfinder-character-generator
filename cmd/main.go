package main

import (
	"pathfinderCharacter/internal/service/characteristic"
	"pathfinderCharacter/internal/service/generator"
)

func main() {
	characteristicDiceGenerator := generator.NewCharacteristicDiceGenerator("heroic")
	characteristicService := characteristic.NewService(characteristicDiceGenerator)
	characteristicService.SetCharacteristic()
}
