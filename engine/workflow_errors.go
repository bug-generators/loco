package engine;

// determine next action in workflow
type workflow_error interface {
    Error() error
}

// will ask engine to retry current step
type Retriable_error struct {
    info error
}

func (re *Retriable_error) Error() error {
    return re.info;
}

// will force engine to abort workflow
type Fatal_error struct {
    info error
}

func (fe *Fatal_error) Error() error {
    return fe.info;
}
