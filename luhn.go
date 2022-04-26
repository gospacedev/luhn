package luhn

//Checks if the string is a luhn number
func Check(cardNo string) bool {
	nDigits := len(cardNo)

	nSum := 0
	isSecond := false

	for i := nDigits - 1; i >= 0; i-- {

		d := int(cardNo[i] - '0')

		if isSecond == true {
			d = d * 2
		}

		nSum += d / 10
		nSum += d % 10

		isSecond = !isSecond
	}
	return nSum%10 == 0
}