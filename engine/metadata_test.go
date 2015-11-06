package engine

import (
	"testing"
)

// metadata should return that a step is retriable
func TestIsRetriable(t *testing.T) {
	metadata := generateMetadata(10, 9)
	if !metadata.isRetriable() {
		t.Fatalf("Step should be retriable")
	}

	metadata = generateMetadata(10, 10)
	if metadata.isRetriable() {
		t.Fatalf("Step should not be retriable")
	}
}

func generateMetadata(max int, num int) *Metadata {

	metadata := Metadata{
		workflowId:       "workflow-id",
		workflowStepId:   "workflow-step-id",
		workflowName:     "workflow-name",
		workflowStepName: "workflowStepName",
		maxRetries:       max,
		numRetries:       num,
	}

	return &metadata
}
