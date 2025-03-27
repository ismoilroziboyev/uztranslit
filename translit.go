package uztranslit

import "strings"

var cyrillicToLatin = map[rune]string{
	// capital letters
	'А': "A", 'Б': "B", 'В': "V", 'Г': "G", 'Д': "D", 'Е': "E", 'Ё': "Yo", 'Ғ': "Gʻ",
	'Ж': "J", 'З': "Z", 'И': "I", 'Й': "Y", 'К': "K", 'Қ': "Q", 'Л': "L", 'М': "M",
	'Н': "N", 'О': "O", 'П': "P", 'Р': "R", 'С': "S", 'Т': "T", 'У': "U", 'Ў': "Oʻ",
	'Ф': "F", 'Х': "Kh", 'Ц': "Ts", 'Ч': "Ch", 'Ш': "Sh", 'Щ': "Sh",
	'Ъ': "", 'Ы': "Y", 'Ь': "", 'Э': "E", 'Ю': "Yu", 'Я': "Ya", 'Ҳ': "H",

	//
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "yo", 'ғ': "gʻ",
	'ж': "j", 'з': "z", 'и': "i", 'й': "y", 'к': "k", 'қ': "q", 'л': "l", 'м': "m",
	'н': "n", 'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u", 'ў': "oʻ",
	'ф': "f", 'х': "kh", 'ц': "ts", 'ч': "ch", 'ш': "sh", 'щ': "sh",
	'ъ': "", 'ы': "y", 'ь': "", 'э': "e", 'ю': "yu", 'я': "ya", 'ҳ': "h",
}

func transliterate(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			var transliteratedWord string
			// If first letter is 'Е' or 'е', replace it
			if runes[0] == 'Е' {
				transliteratedWord = "Ye"
			} else if runes[0] == 'е' {
				transliteratedWord = "ye"
			} else if val, exists := cyrillicToLatin[runes[0]]; exists {
				transliteratedWord = val
			} else {
				transliteratedWord = string(runes[0])
			}

			// Process remaining letters
			for _, char := range runes[1:] {
				if val, exists := cyrillicToLatin[char]; exists {
					transliteratedWord += val
				} else {
					transliteratedWord += string(char)
				}
			}
			words[i] = transliteratedWord
		}
	}
	return strings.Join(words, " ")
}
