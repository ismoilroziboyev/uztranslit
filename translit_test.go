package uztranslit

import "testing"

func TestTransliterate(t *testing.T) {
	cases := []struct {
		Input  string
		Result string
	}{
		{"Салом дунё", "Salom dunyo"},
		{"Ўзбекистон Республикаси", "Oʻzbekiston Respublikasi"},
		{"Машина тез юради", "Mashina tez yuradi"},
		{"Китоб ўқиш жуда фойдали", "Kitob oʻqish juda foydali"},
		{"Ёмғир ёғяпти", "Yomgʻir yogʻyapti"},
		{"Менинг отам шифокор", "Mening otam shifokor"},
		{"Бу ерда ҳеч ким йўқ", "Bu yerda hech kim yoʻq"},
		{"Ўқувчилар дарс қилишмоқда", "Oʻquvchilar dars qilishmoqda"},
		{"Менинг туғилган куним январда", "Mening tugʻilgan kunim yanvarda"},
	}

	for _, c := range cases {
		result := Transliterate(c.Input)

		if result != c.Result {
			t.Errorf("For input '%s', expected '%s', but got '%s'", c.Input, c.Result, result)
		}
	}
}
