package model

type Characteristic struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

func NewCharacteristic(strength int, dexterity int, constitution int,
	intelligence int, wisdom int, charisma int) *Characteristic {
	return &Characteristic{Strength: strength, Dexterity: dexterity,
		Constitution: constitution, Intelligence: intelligence, Wisdom: wisdom, Charisma: charisma}
}
