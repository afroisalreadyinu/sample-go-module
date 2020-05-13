package pathfinder_test

import (
	"github.com/stretchr/testify/assert"
	"myprinter/pathfinder"
	"testing"
)

func TestFind(t *testing.T) {
	assert.NotEqual(t, "", pathfinder.Find())
}
