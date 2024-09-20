package main

import (
	"strings"
)

var (
	RomansMap = map[string]int{ // карта для перевода из римских в арабские
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	someInts = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} // срезы для перевода из арабских в римские
	// (стандартный алгоритм)
	someRomans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

func romanToInt(rom string) (result int) { // римское число/цифра в арабское для расчета выражения

	romanStringSlice := strings.Split(rom, "") // со срезом в этом случае проще работать

	for index := 0; index < len(romanStringSlice)-1; index++ {

		currentInt, currentOk := RomansMap[romanStringSlice[index]]
		nextInt, nextOk := RomansMap[romanStringSlice[index+1]]

		if currentOk && nextOk { // если текущий и следующий элементы присутствуют в мапе
			if currentInt >= nextInt { // например VI
				result += currentInt
			} else { // например IV
				result -= currentInt
			}
		} else {
			panic("Roman value doesn't exist")
		}
	}

	result += RomansMap[romanStringSlice[len(romanStringSlice)-1]] // + последний символ, так как в цикле его не трогаем

	return
}

func intToRoman(number int) (roman string) { // для перевода результата в римское число/цифру

	for i := 0; i < len(someInts); i++ { // от большего к меньшему по срезам
		for number >= someInts[i] {
			roman += someRomans[i]
			number -= someInts[i]
		}
	}
	return roman
}

func EvaluateRomansExpression(romans []string, operation string) string { // просто получение результата с использованием
	// ранее реализованных функций

	firstDig, SecondDig := romanToInt(romans[0]), romanToInt(romans[1])

	numericalResult := EvaluateDigitsExpression([]int{firstDig, SecondDig}, operation)

	if numericalResult < 0 {
		panic("The result of the roman expression is negative")
	}

	return intToRoman(numericalResult)

}
