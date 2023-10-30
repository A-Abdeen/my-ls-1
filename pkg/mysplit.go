package Myls

func MySplit(word string, toSplit string) (string, string) {
	wordAsBytes := []byte(word)
	dataBefore := ""
	dataAfter := ""
	before := false
	count := 0
	for i := len(wordAsBytes) - 1; i >= 0; i-- {
		if string(wordAsBytes[i]) == toSplit && count == 0{
			count++
			before = true
			continue
		}
		if before {
			dataBefore = string(wordAsBytes[i]) + dataBefore
		} else {
			dataAfter = string(wordAsBytes[i]) + dataAfter
		}
	}
	return dataBefore, dataAfter
}
