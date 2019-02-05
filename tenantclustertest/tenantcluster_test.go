package tenantclustertest

import (
	"testing"
)

func Test_New(t *testing.T) {
	// Test that New doesn't panic and tenantcluster.Interface is implemented.
	New(Config{})
}
