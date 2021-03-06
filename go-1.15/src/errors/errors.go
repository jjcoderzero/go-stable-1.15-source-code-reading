// Package errors implements functions to manipulate errors.
//
// The New function creates errors whose only content is a text message.
//
// The Unwrap, Is and As functions work on errors that may wrap other errors.
// An error wraps another error if its type has the method
//
//	Unwrap() error
//
// If e.Unwrap() returns a non-nil error w, then we say that e wraps w.
//
// Unwrap unpacks wrapped errors. If its argument's type has an
// Unwrap method, it calls the method once. Otherwise, it returns nil.
//
// A simple way to create wrapped errors is to call fmt.Errorf and apply the %w verb
// to the error argument:
//
//	errors.Unwrap(fmt.Errorf("... %w ...", ..., err, ...))
//
// returns err.
//
// Is unwraps its first argument sequentially looking for an error that matches the
// second. It reports whether it finds a match. It should be used in preference to
// simple equality checks:
//
//	if errors.Is(err, os.ErrExist)
//
// is preferable to
//
//	if err == os.ErrExist
//
// because the former will succeed if err wraps os.ErrExist.
//
// As unwraps its first argument sequentially looking for an error that can be
// assigned to its second argument, which must be a pointer. If it succeeds, it
// performs the assignment and returns true. Otherwise, it returns false. The form
//
//	var perr *os.PathError
//	if errors.As(err, &perr) {
//		fmt.Println(perr.Path)
//	}
//
// is preferable to
//
//	if perr, ok := err.(*os.PathError); ok {
//		fmt.Println(perr.Path)
//	}
//
// because the former will succeed if err wraps an *os.PathError.
package errors

// New返回一个错误，该错误将格式化为给定的文本。每次调用New都会返回一个不同的错误值，即使文本是相同的。
func New(text string) error {
	return &errorString{text}
}

// errorString是error的一个简单实现。
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
