package errortypes

//import "github.com/pkg/errors"

type notFoundErr struct { //--> private, this won-t acoplate to the rest of the code
	error
}

func WrapNotFoundErr(err error, format string, args ...interface{}) error {
	return &notFoundErr{errors.Wrapf(err, format, args...)}
}

func NewNotFoundErr(format string, args ...interface{}) error {
	return &notFoundErr{errors.Errorf(format, args...)}
}

// This func creates an error that is a notFoundEroor

func IsNotFoundErr(err error) bool { //--> controlamos si el error es del tipo notFoundErr
	err = errors.Cause(err)
	_, ok := err.(*notFoundErr)
	return ok
}
