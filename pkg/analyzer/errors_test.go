package analyzer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMatching(t *testing.T) {
	assert.True(t, errors.Is(InvalidNumberOfFilesInAnalysis{1, 2}, InvalidNumberOfFilesInAnalysis{}))
}
