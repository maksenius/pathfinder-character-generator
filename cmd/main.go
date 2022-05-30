package main

import (
	"fmt"

	"pathfinderCharacterGenerator/internal/service/characteristic"
	"pathfinderCharacterGenerator/internal/service/generator"
)

func main() {
	characteristicDiceGenerator := generator.NewCharacteristicDiceGenerator("heroic")
	characteristicService := characteristic.NewService(characteristicDiceGenerator)
	fmt.Println(characteristicService)
}
