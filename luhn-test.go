package luhn

import "testing"

func testLuhn(t *testing.T)  {
	validNo := []string{"79927398713", "4929972884676289", "379537021417898", "373494930335082", "379203612454689", "6387065788050980", "6388464094939979"}

	for _, v := range validNo {
		if !Check(v) {
			t.Errorf("%v: Should be valid", v)
		}
	}
}