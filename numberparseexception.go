package libphonenumber

import "fmt"

type ErrorType int

func (e ErrorType) String() string {
	switch e {
	case InvalidCountryCode:
		return "InvalidCountryCode"
	case NotANumber:
		return "NotANumber"
	case TooShortAfterIDD:
		return "TooShortAfterIDD"
	case TooShortNSN:
		return "TooShortNSN"
	case TooLong:
		return "TooLong"
	default:
		return ""
	}
}

const (
	InvalidCountryCode ErrorType = iota
	NotANumber
	TooShortAfterIDD
	TooShortNSN
	TooLong
)

type NumberParseError struct {
	errorType ErrorType
	message   string
}

func (n NumberParseError) ErrorType() ErrorType {
	return n.errorType
}

func (n NumberParseError) Error() string {
	return n.String()
}

func (n NumberParseError) String() string {
	return fmt.Sprintf("Error type: %s.%s", n.errorType.String(), n.message)
}
