package security

import "net/http"

// FakeSecurity always returns true for the given request.
type FakeSecurity struct {
}

func (s *FakeSecurity) Verify(_ *http.Request) bool {
	return true
}

func ProvideFakeSecurity() *FakeSecurity {
	return &FakeSecurity{}
}
