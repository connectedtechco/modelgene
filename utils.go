package modelgene

import (
	"fmt"
	"github.com/connectedtechco/modelgene/pkg/types"
)

// --- Pointer Helpers ---

// Helper functions for pointers to primitive types
func PtrString(s string) *string    { return &s }
func PtrInt(i int) *int             { return &i }
func PtrBool(b bool) *bool          { return &b }
func PtrFloat64(f float64) *float64 { return &f }

// --- Token Helpers ---

// GetMaxTokens returns the max tokens value, defaulting to 1024 if nil
func GetMaxTokens(max *int) int64 {
	if max != nil {
		return int64(*max)
	}
	return 1024
}

// --- Error Handling ---

// ModelGeneError represents an error originating from the modelgene library.
type ModelGeneError struct {
	Provider types.Provider
	Message  string
	Err      error // Underlying error, if any
}

func (e *ModelGeneError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("modelgene [%s]: %s: %v", e.Provider, e.Message, e.Err)
	}
	return fmt.Sprintf("modelgene [%s]: %s", e.Provider, e.Message)
}

func NewError(provider types.Provider, message string, err error) error {
	return &ModelGeneError{
		Provider: provider,
		Message:  message,
		Err:      err,
	}
}
