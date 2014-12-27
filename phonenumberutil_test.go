package libphonenumber

import "testing"

func Test_parse(t *testing.T) {
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
		num, err := parse(test.input, test.region)
		if err != test.err {
			t.Errorf("[test %d:err] failed: %v != %v\n", i, err, test.err)
		}
		if num.GetNationalNumber() != test.expectedNum {
			t.Errorf("[test %d:num] failed: %v != %v\n", i, err, test.err)
		}
	}
}

func Test_convertAlphaCharactersInNumber(t *testing.T) {
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
		out := convertAlphaCharactersInNumber(test.input)
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
