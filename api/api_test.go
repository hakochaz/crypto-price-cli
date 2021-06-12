package api

import "testing"

func TestGetCurrentPriceData(t *testing.T) {
	var tests = []struct {
		id      string
		vc      string
		wantErr bool
		name    string
	}{
		{"bitcoin", "gbp", false, "compatible paramaters"},
		{"bitcoin", "gbs", true, "incompatible versus currency"},
		{"bitco", "gbs", true, "incompatible id"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetCurrentPriceData(tt.id, tt.vc)
			if (err == nil) && tt.wantErr {
				t.Error("returned no error but one was expected")
			}
			if (err != nil) && !tt.wantErr {
				t.Errorf("returned an error but one was not expected: %s", err)
			}
		})
	}
}

func TestGetHistoricalPriceData(t *testing.T) {
	var tests = []struct {
		id      string
		vc      string
		date    string
		wantErr bool
		name    string
	}{
		{"bitcoin", "gbp", "05-06-2018", false, "compatible paramaters"},
		{"bitcoin", "gbh", "05-06-2018", true, "incompatible versus currency"},
		{"bitcoin", "gbp", "05-19-2018", true, "incompatible date"},
		{"bitco", "gbp", "05-06-2018", true, "incompatible id"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetHistoricalPriceData(tt.id, tt.vc, tt.date)
			if (err == nil) && tt.wantErr {
				t.Error("returned no error but one was expected")
			}
			if (err != nil) && !tt.wantErr {
				t.Errorf("returned an error but one was not expected: %s", err)
			}
		})
	}
}

func TestSearchSupportedCoin(t *testing.T) {
	var tests = []struct {
		searchName string
		wantErr    bool
		name       string
	}{
		{"bitcoin", false, "compatible id"},
		{"Bitcoin", false, "compatible search name"},
		{"wada", true, "incompatible search name"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SearchSupportedCoin(tt.searchName)
			if (err == nil) && tt.wantErr {
				t.Errorf("returned no error but one was expected")
			}
			if (err != nil) && !tt.wantErr {
				t.Errorf("returned an error but one was not expected: %s", err)
			}
		})
	}
}

func TestGetVersusCurrencysList(t *testing.T) {
	_, err := GetVersusCurrencysList()

	if err != nil {
		t.Error(err)
	}
}

func TestCalculateAmount(t *testing.T) {
	pd := PriceData{"Test", "Test", 0, "Test", 9, 0}
	pd.CalculateAmount()

	if pd.Amount == 0 {
		t.Error("amount should be greater than 0")
	}
}
