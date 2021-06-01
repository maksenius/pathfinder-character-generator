package main

import (
	"pathfinderCharacterGenerator/internal/service/characteristic"
	"pathfinderCharacterGenerator/internal/service/generator"
)

func main() {
	characteristicDiceGenerator := generator.NewCharacteristicDiceGenerator("heroic")
	characteristicService := characteristic.NewService(characteristicDiceGenerator)
	characteristicService.SetCharacteristic()
}
