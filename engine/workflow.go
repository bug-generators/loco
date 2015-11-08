package engine

// A workflow is an ordered set of steps
type Workflow struct {
	id    string // The unique id of each workflow.
	name  string
	steps []Step
}

// Every step in a workflow must define this interface.
type Step struct {
	id     string // The unique id of each workflow step
	name   string
	action Action
}

// Every step requires an action to be performed
type Action interface {
	exec(input *[]byte, metadata *Metadata) (output *[]byte, err *workflow_error)
}
