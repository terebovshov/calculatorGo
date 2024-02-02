// roman.go
package main

// Карта для преобразования римских цифр в арабские.
var romanToArabic = map[string]int{
	"I":      1,
	"II":     2,
	"III":    3,
	"IV":     4,
	"V":      5,
	"VI":     6,
	"VII":    7,
	"VIII":   8,
	"IX":     9,
	"X":      10,
	"XI":     11,
	"XII":    12,
	"XIII":   13,
	"XIV":    14,
	"XV":     15,
	"XVI":    16,
	"XVII":   17,
	"XVIII":  18,
	"XIX":    19,
	"XX":     20,
	"XXI":    21,
	"XXIV":   24,
	"XXV":    25,
	"XXVIII": 28,
	"XXX":    30,
	"XXXII":  32,
	"XXXV":   35,
	"XXXVI":  36,
	"XL":     40,
	"XLII":   42,
	"XLV":    45,
	"XLVIII": 48,
	"L":      50,
	"LIV":    54,
	"LVI":    56,
	"LX":     60,
	"LXIII":  63,
	"LXIV":   64,
	"LXX":    70,
	"LXXII":  72,
	"LXXX":   80,
	"LXXXI":  81,
	"XC":     90,
	"C":      100,
}

// Карта для преобразования арабских цифр в римские.
var arabicToRoman = map[int]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	11:  "XI",
	12:  "XII",
	13:  "XIII",
	14:  "XIV",
	15:  "XV",
	16:  "XVI",
	17:  "XVII",
	18:  "XVIII",
	19:  "XIX",
	20:  "XX",
	21:  "XXI",
	24:  "XXIV",
	25:  "XXV",
	28:  "XXVIII",
	30:  "XXX",
	32:  "XXXII",
	35:  "XXXV",
	36:  "XXXVI",
	40:  "XL",
	42:  "XLII",
	45:  "XLV",
	48:  "XLVIII",
	50:  "L",
	54:  "LIV",
	56:  "LVI",
	60:  "LX",
	63:  "LXIII",
	64:  "LXIV",
	70:  "LXX",
	72:  "LXXII",
	80:  "LXXX",
	81:  "LXXXI",
	90:  "XC",
	100: "C",
}

// Функция для проверки, является ли введенная строка римским числом.
func isRoman(input string) bool {
	_, exists := romanToArabic[input]
	return exists
}

// Функция для преобразования римского числа в арабское.
func romanToArabicNumber(roman string) int {
	return romanToArabic[roman]
}

// Функция для преобразования арабского числа в римское.
func arabicToRomanNumber(arabic int) string {
	return arabicToRoman[arabic]
}
