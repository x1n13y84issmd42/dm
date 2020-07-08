package ut_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph/contract"
	"github.com/x1n13y84issmd42/gog/graph/ut"
)

func Test_Node(T *testing.T) {
	n := ut.Node("A")
	assert.Equal(T, contract.NodeID("A"), n.ID())
}
