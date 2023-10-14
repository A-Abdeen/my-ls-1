package Myls

func Alphanumeric(word string) string {
	wordByte := []byte(word)
	var finalword []byte
	for i := 0; i < len(wordByte); i++ {
		if wordByte[i] > 47 && wordByte[i] < 58 {
			finalword = append(finalword, wordByte[i])
		}
		if wordByte[i] > 96 && wordByte[i] < 123 {
			finalword = append(finalword, wordByte[i])
		}
		if wordByte[i] > 64 && wordByte[i] < 91 {
			finalword = append(finalword, wordByte[i])
		}
	}
	return string(finalword)
}
