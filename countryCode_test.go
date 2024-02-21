package countryCode

import "testing"

func TestGetCountryCode(t *testing.T) {
	testCases := []struct {
		inputCountry string
		expectedCode string
	}{
		{"United States", "US"},
		{"United Kingdom", "GB"},
		{"United Arab Emirates", "AE"},
		{"China", "CN"},
		{"India", "IN"},
		{"Japan", "JP"},
		{"South Korea", "KR"},
		{"Germany", "DE"},
		{"France", "FR"},
		{"Italy", "IT"},
		{"Spain", "ES"},
		{"Brazil", "BR"},
		{"Australia", "AU"},
		{"Canada", "CA"},
		{"Mexico", "MX"},
		{"Russia", "RU"},
		{"South Africa", "ZA"},
		{"Argentina", "AR"},
		{"Egypt", "EG"},
		{"Pakistan", "PK"},
		{"Bangladesh", "BD"},
		{"Turkey", "TR"},
		{"Nigeria", "NG"},
		{"Vietnam", "VN"},
		{"Thailand", "TH"},
		{"Indonesia", "ID"},
		{"Philippines", "PH"},
		{"Malaysia", "MY"},
		{"Singapore", "SG"},
		{"Saudi Arabia", "SA"},
		{"Poland", "PL"},
		{"Netherlands", "NL"},
		{"Belgium", "BE"},
		{"Sweden", "SE"},
		{"Norway", "NO"},
		{"Denmark", "DK"},
		{"Finland", "FI"},
		{"Ireland", "IE"},
		{"Portugal", "PT"},
		{"Austria", "AT"},
		{"Greece", "GR"},
		{"Switzerland", "CH"},
		{"Romania", "RO"},
		{"Hungary", "HU"},
		{"Ukraine", "UA"},
		{"Croatia", "HR"},
		{"Bulgaria", "BG"},
		{"Serbia", "RS"},
		{"Slovenia", "SI"},
		{"Slovakia", "SK"},
		{"Lithuania", "LT"},
		{"Latvia", "LV"},
		{"Estonia", "EE"},
		{"Cyprus", "CY"},
		{"Malta", "MT"},
		{"Iceland", "IS"},
		{"Luxembourg", "LU"},
		{"Monaco", "MC"},
		{"Liechtenstein", "LI"},
		{"Andorra", "AD"},
		{"San Marino", "SM"},
		{"Vatican City", "VA"},
		{"Montenegro", "ME"},
		{"Albania", "AL"},
		{"Macedonia", "MK"},
		{"Bosnia and Herzegovina", "BA"},
		{"Uzbekistan", "UZ"},
		{"Kazakhstan", "KZ"},
		{"Turkmenistan", "TM"},
		{"Kyrgyzstan", "KG"},
		{"Tajikistan", "TJ"},
		{"Afghanistan", "AF"},
		{"Iraq", "IQ"},
		{"Iran", "IR"},
		{"Syria", "SY"},
		{"Lebanon", "LB"},
		{"Jordan", "JO"},
	}

	for _, tc := range testCases {
		countryCode, err := GetCountryCode(tc.inputCountry)
		if err != nil {
			t.Logf(err.Error())
			t.Errorf("Error for %s", tc.inputCountry)
		}
		if countryCode != tc.expectedCode {
			t.Logf(countryCode)
			t.Errorf("Error for %s", tc.inputCountry)
		}
	}
}
