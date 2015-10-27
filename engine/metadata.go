package engine

import ()

type Metadata struct {
   // uuids 
    workflowId string
    workflowStepId string

    // names
    workflowName string
    workflowStepName string

    // retries
    maxRetries int
    numRetries int

}

// Returns whether the workflow step has exhausted all retries
func (m *Metadata) isRetriable() bool {
    return m.numRetries < m.maxRetries;
}


