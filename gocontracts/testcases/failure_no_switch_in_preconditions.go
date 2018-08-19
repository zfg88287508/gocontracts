package testcases

// FailureNoSwitchInPrecondition tests that we correctly detect when comment implying pre-conditions is not followed by
// a switch statement in the function body.
var FailureNoSwitchInPrecondition = Failure{
	ID: "no_switch_in_preconditions",
	Text: `package somepkg

// SomeFunc does something.
//
// SomeFunc requires:
// * x > 0
//
// SomeFunc ensures:
// * result == "oi"
//
// Some text here.
func SomeFunc(x int, y int) (result string, err error) {
	// Pre-conditions
	panic("hello")

	return
}
`,
	Error: "expected a 'switch' statement after the comment \"Pre-conditions\" in function SomeFunc on line 14"}
