package generator

import (
	"fmt"
	"sort"

	"pathfinderCharacterGenerator/internal/helper"
)

type CharacteristicDiceGenerator struct {
	generatorMethod string
}

func NewCharacteristicDiceGenerator(generatorMethod string) *CharacteristicDiceGenerator {
	return &CharacteristicDiceGenerator{generatorMethod: generatorMethod}
}

func (c *CharacteristicDiceGenerator) GetNumbers() []int {
	fmt.Println("Generator method is", c.generatorMethod)

	switch c.generatorMethod {
	case StandardGeneratorMethod:
		return c.standardMethod()
	case ClassicGeneratorMethod:
		return c.classicMethod()
	case HeroicGeneratorMethod:
		return c.heroicMethod()
	default:
		return c.standardMethod()
	}
}

func (c *CharacteristicDiceGenerator) standardMethod() []int {
	results := make([]int, CountOfCharacteristics, CountOfCharacteristics)
	tmpList := make([]int, StandardMethodTmpSliceSize, StandardMethodTmpSliceSize)
	for j := 0; j < CountOfCharacteristics; j++ {
		for i := 0; i < StandardMethodTmpSliceSize; i++ {
			tmpList[i] = helper.D6()
		}
		sort.Slice(tmpList, func(i, j int) bool {
			if tmpList[i] > tmpList[j] {
				return true
			}
			return false
		})
		res := 0
		for i := 0; i < StandardMethodTmpSliceSize-1; i++ {
			res = res + tmpList[i]
		}
		results[j] = res
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i] > results[j] {
			return true
		}
		return false
	})
	return results
}

func (c *CharacteristicDiceGenerator) classicMethod() []int {
	results := make([]int, CountOfCharacteristics, CountOfCharacteristics)
	tmpList := make([]int, ClassicMethodTmpSliceSize, ClassicMethodTmpSliceSize)
	for j := 0; j < CountOfCharacteristics; j++ {
		for i := 0; i < ClassicMethodTmpSliceSize; i++ {
			tmpList[i] = helper.D6()
		}
		res := 0
		for i := 0; i < ClassicMethodTmpSliceSize; i++ {
			res = res + tmpList[i]
		}
		results[j] = res
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i] > results[j] {
			return true
		}
		return false
	})
	return results
}

func (c *CharacteristicDiceGenerator) heroicMethod() []int {
	results := make([]int, CountOfCharacteristics, CountOfCharacteristics)
	tmpList := make([]int, HeroicMethodTmpSliceSize, HeroicMethodTmpSliceSize)
	for j := 0; j < CountOfCharacteristics; j++ {
		for i := 0; i < HeroicMethodTmpSliceSize; i++ {
			tmpList[i] = helper.D6()
		}
		res := 0
		for i := 0; i < HeroicMethodTmpSliceSize; i++ {
			res = res + tmpList[i]
		}
		results[j] = res + HeroicMethodBuff
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i] > results[j] {
			return true
		}
		return false
	})
	return results
}

//TODO: add buy method
