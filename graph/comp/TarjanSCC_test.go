package comp_test

import (
	"testing"

	"github.com/x1n13y84issmd42/gog/graph"
	"github.com/x1n13y84issmd42/gog/graph/comp"
)

func Test_TarjanSCC(T *testing.T) {
	g := graph.NewDGraph()
	comp.TarjanSCC(g)
}
