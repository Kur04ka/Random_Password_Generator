package variables

var (
	Chars_EN, Chars_en            string = "ABCDEFGHIJKLMNOPQRRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"
	Chars_specials, Chars_digits  string = "~=+%^*/()[]{}/!@#$?|", "0123456789"
	Chars_RU, Chars_ru            string = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЭЮЯ", "абвгдеёжзийклмнопрстуфхцчшщэюя"
	YesOrNo, UserInput, All_chars string
	PasswordLength                int
	Regexp_map                    = map[string]string{
		`[A-Z]`: Chars_EN, `[a-z]`: Chars_en, `[А-Я]`: Chars_RU, `[а-я]`: Chars_ru,
		`\d`: Chars_digits, `[~=\+\%\^\*/\(\)\[\]\{\}/!@#\$?\|]`: Chars_specials,
	}
	Dialog_map = map[string]string{
		"A-Z": Chars_EN, "a-z": Chars_en, "А-Я": Chars_RU, "а-я": Chars_ru,
		"0-9": Chars_digits, "~=+%^*/()[]{}/!@#$?|": Chars_specials,
	}
)
