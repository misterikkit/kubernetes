package predicates

// FitError is an error that describes why a pod did not fit on a node.
type FitError interface {
	error
	GetReason() string
}

// fitError implements the FitError interface.
type fitError string

func (f fitError) Error() string     { return string(f) }
func (f fitError) GetReason() string { return string(f) }

// TODO: we don't want to add a list of error constants here. Error details should be checked through interfaces.
