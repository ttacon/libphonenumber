package libphonenumber

type PhoneNumberMatcher struct {
}

func NewPhoneNumberMatcher(seq string) *PhoneNumberMatcher {
	// TODO(ttacon): to be implemented
	return nil
}

func ContainsOnlyValidXChars(
	number *PhoneNumber,
	candidate string,
	util *PhoneNumberUtil) bool {
	// TODO(ttacon): to be implemented
	return false
}

func IsNationalPrefixPresentIfRequired(
	number *PhoneNumber,
	util *PhoneNumberUtil) bool {
	// TODO(ttacon): to be implemented
	return false
}

func ContainsMoreThanOneSlashInNationalNumber(
	number *PhoneNumber,
	candidate string) bool {
	// TODO(ttacon): to be implemented
	return false
}

func CheckNumberGroupingIsValid(
	number *PhoneNumber,
	candidate string,
	util *PhoneNumberUtil,
	fn func(*PhoneNumberUtil, *PhoneNumber, string, []string) bool) bool {
	// TODO(ttacon): to be implemented
	return false
}

func AllNumberGroupsRemainGrouped(
	util *PhoneNumberUtil,
	number *PhoneNumber,
	normalizedCandidate string,
	expectedNumberGroups []string) bool {
	// TODO(ttacon): to be implemented
	return false
}

func AllNumberGroupsAreExactlyPresent(
	util *PhoneNumberUtil,
	number *PhoneNumber,
	normalizedCandidate string,
	expectedNumberGroups []string) bool {
	// TODO(ttacon): to be implemented
	return false
}
