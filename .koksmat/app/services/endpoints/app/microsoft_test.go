package app

import (
	"testing"
)

func TestOutlook(t *testing.T) {

	_, err := Process([]string{})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

}
