package stats

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Average(T *testing.T) {
	cases := []struct {
		Input    Floats
		Expected float64
	}{
		{
			Input:    Floats{1.1, 5.666666666, 234, -6.3, -44, -1, -0.3, 0.5, 34.34},
			Expected: 24.889629629555554,
		},

		{
			Input:    Floats{},
			Expected: 0,
		},
	}

	for cI, c := range cases {
		T.Run(fmt.Sprintf("Case #%d", cI), func(T *testing.T) {
			assert.Equal(T, c.Expected, Average(&c.Input))
		})
	}
}

func Test_Dispersion(T *testing.T) {
	cases := []struct {
		Input    Floats
		Expected float64
	}{
		{
			Input:    Floats{1.1, 5.666666666, 234, -6.3, -44, -1, -0.3, 0.5, 34.34},
			Expected: 6546.267967904439,
		},

		{
			Input:    Floats{},
			Expected: 0,
		},
	}

	for cI, c := range cases {
		T.Run(fmt.Sprintf("Case #%d", cI), func(T *testing.T) {
			assert.Equal(T, c.Expected, Dispersion(&c.Input))
		})
	}
}

func Test_StdDev(T *testing.T) {
	cases := []struct {
		Input    Floats
		Expected float64
	}{
		{
			Input:    Floats{97, 98, 99, 100, 101, 102, 103},
			Expected: 2.160246899469287,
		},

		{
			Input:    Floats{70, 80, 90, 100, 110, 120, 130},
			Expected: 21.602468994692867,
		},

		{
			Input:    Floats{},
			Expected: 0,
		},
	}

	for cI, c := range cases {
		T.Run(fmt.Sprintf("Case #%d", cI), func(T *testing.T) {
			assert.Equal(T, c.Expected, StdDev(&c.Input))
		})
	}
}
