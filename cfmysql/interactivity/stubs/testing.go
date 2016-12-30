package stubs

// InteractiveStubDetector is for stubbing the Detector interface in tests
// with the one that is always returns true
type InteractiveStubDetector struct{}

// IsInteractive returns always true
func (InteractiveStubDetector) IsInteractive() bool {
	return true
}

// NonInteractiveStubDetector is for stubbing the Detector interface in tests
// with the one that is always returns false
type NonInteractiveStubDetector struct{}

// IsInteractive returns always false
func (NonInteractiveStubDetector) IsInteractive() bool {
	return false
}
