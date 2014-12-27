package libphonenumber

import "testing"

func TestParse(t *testing.T) {
	var tests = []struct {
		input       string
		err         error
		expectedNum uint64
		region      string
	}{
		{
			input:       "4437990238",
			err:         nil,
			expectedNum: 4437990238,
			region:      "US",
		}, {
			input:       "(443) 799-0238",
			err:         nil,
			expectedNum: 4437990238,
			region:      "US",
		}, {
			input:       "((443) 799-023asdfghjk8",
			err:         ErrNumTooLong,
			expectedNum: 0,
			region:      "US",
		}, {
			input:       "+441932567890",
			err:         nil,
			expectedNum: 1932567890,
			region:      "GB",
		}, {
			input:       "45",
			err:         nil,
			expectedNum: 45,
			region:      "US",
		}, {
			input:       "1800AWWCUTE",
			err:         nil,
			expectedNum: 8002992883,
			region:      "US",
		},
	}

	for i, test := range tests {
		num, err := Parse(test.input, test.region)
		if err != test.err {
			t.Errorf("[test %d:err] failed: %v != %v\n", i, err, test.err)
		}
		if num.GetNationalNumber() != test.expectedNum {
			t.Errorf("[test %d:num] failed: %v != %v\n", i, err, test.err)
		}
	}
}

func TestConvertAlphaCharactersInNumber(t *testing.T) {
	var tests = []struct {
		input, output string
	}{
		{
			input:  "1800AWWPOOP",
			output: "18002997667",
		}, {
			input:  "(800) DAW-ORLD",
			output: "(800) 329-6753",
		},
	}

	for i, test := range tests {
		out := ConvertAlphaCharactersInNumber(test.input)
		if out != test.output {
			t.Errorf("[test %d] failed, %s != %s\n", i, out, test.output)
		}
	}
}

func Test_normalizeDigits(t *testing.T) {
	var tests = []struct {
		input         string
		expected      []byte
		keepNonDigits bool
	}{
		{
			input:         "4445556666",
			expected:      []byte("4445556666"),
			keepNonDigits: false,
		}, {
			input:         "(444)5556666",
			expected:      []byte("4445556666"),
			keepNonDigits: false,
		}, {
			input:         "(444)555a6666",
			expected:      []byte("4445556666"),
			keepNonDigits: false,
		}, {
			input:         "(444)555a6666",
			expected:      []byte("(444)555a6666"),
			keepNonDigits: true,
		},
	}

	for i, test := range tests {
		out := normalizeDigits(test.input, test.keepNonDigits)
		if string(out) != string(test.expected) {
			t.Errorf("[test %d] failed: %s != %s\n",
				i, string(out), string(test.expected))
		}
	}
}

func Test_extractPossibleNumber(t *testing.T) {
	var (
		input    = "(530) 583-6985 x302/x2303"
		expected = "530) 583-6985 x302" // yes, the leading '(' is missing
	)

	output := extractPossibleNumber(input)
	if output != expected {
		t.Error(output, "!=", expected)
	}
}

func Test_isViablePhoneNumer(t *testing.T) {
	var tests = []struct {
		input    string
		isViable bool
	}{
		{
			input:    "4445556666",
			isViable: true,
		}, {
			input:    "+441932123456",
			isViable: true,
		}, {
			input:    "4930123456",
			isViable: true,
		}, {
			input:    "2",
			isViable: false,
		}, {
			input:    "helloworld",
			isViable: false,
		},
	}

	for i, test := range tests {
		result := isViablePhoneNumber(test.input)
		if result != test.isViable {
			t.Errorf("[test %d] %v != %v\n", i, result, test.isViable)
		}
	}
}

func Test_normalize(t *testing.T) {
	var tests = []struct {
		in  string
		exp string
	}{
		{
			in:  "4431234567",
			exp: "4431234567",
		}, {
			in:  "443 1234567",
			exp: "4431234567",
		}, {
			in:  "(443)123-4567",
			exp: "4431234567",
		}, {
			in:  "800yoloFOO",
			exp: "8009656366",
		}, {
			in:  "444111a2222",
			exp: "4441112222",
		},
	}

	// TODO(ttacon): the above commented out test are because we hacked the crap
	// out of normalizeDigits, fix it

	for i, test := range tests {
		res := normalize(test.in)
		if res != test.exp {
			t.Errorf("[test %d] %s != %s\n", i, res, test.exp)
		}
	}
}

func Test_isValidNumber(t *testing.T) {
	var tests = []struct {
		input   string
		err     error
		isValid bool
		region  string
	}{
		{
			input:   "4437990238",
			err:     nil,
			isValid: true,
			region:  "US",
		}, {
			input:   "(443) 799-0238",
			err:     nil,
			isValid: true,
			region:  "US",
		}, {
			input:   "((443) 799-023asdfghjk8",
			err:     ErrNumTooLong,
			isValid: false,
			region:  "US",
		}, {
			input:   "+441932567890",
			err:     nil,
			isValid: true,
			region:  "GB",
		}, {
			input:   "45",
			err:     nil,
			isValid: false,
			region:  "US",
		}, {
			input:   "1800AWWCUTE",
			err:     nil,
			isValid: true,
			region:  "US",
		},
	}

	for i, test := range tests {
		num, err := Parse(test.input, test.region)
		if err != test.err {
			t.Errorf("[test %d:err] failed: %v != %v\n", i, err, test.err)
		}
		if test.err != nil {
			continue
		}
		if isValidNumber(num) != test.isValid {
			t.Errorf("[test %d:validity] failed: %v != %v\n",
				i, isValidNumber(num), test.isValid)
		}
	}
}

func Test_isValidNumberForRegion(t *testing.T) {
	var tests = []struct {
		input            string
		err              error
		isValid          bool
		validationRegion string
		region           string
	}{
		{
			input:            "4437990238",
			err:              nil,
			isValid:          true,
			validationRegion: "US",
			region:           "US",
		}, {
			input:            "(443) 799-0238",
			err:              nil,
			isValid:          true,
			region:           "US",
			validationRegion: "US",
		}, {
			input:            "((443) 799-023asdfghjk8",
			err:              ErrNumTooLong,
			isValid:          false,
			region:           "US",
			validationRegion: "US",
		}, {
			input:            "+441932567890",
			err:              nil,
			isValid:          true,
			region:           "GB",
			validationRegion: "GB",
		}, {
			input:            "45",
			err:              nil,
			isValid:          false,
			region:           "US",
			validationRegion: "US",
		}, {
			input:            "1800AWWCUTE",
			err:              nil,
			isValid:          true,
			region:           "US",
			validationRegion: "US",
		}, {
			input:            "+441932567890",
			err:              nil,
			isValid:          false,
			region:           "GB",
			validationRegion: "US",
		}, {
			input:            "1800AWWCUTE",
			err:              nil,
			isValid:          false,
			region:           "US",
			validationRegion: "GB",
		},
	}

	for i, test := range tests {
		num, err := Parse(test.input, test.region)
		if err != test.err {
			t.Errorf("[test %d:err] failed: %v != %v\n", i, err, test.err)
		}
		if test.err != nil {
			continue
		}
		if isValidNumberForRegion(num, test.validationRegion) != test.isValid {
			t.Errorf("[test %d:validity] failed: %v != %v\n",
				i, isValidNumberForRegion(num, test.validationRegion), test.isValid)
		}
	}
}
