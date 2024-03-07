package ts_test

import (
	"fmt"
	"testing"

	"github.com/nullbull/gotch"
	"github.com/nullbull/gotch/ts"
)

func TestTensor_Format(t *testing.T) {
	shape := []int64{8, 8, 8}
	numels := int64(8 * 8 * 8)

	x := ts.MustArange(ts.IntScalar(numels), gotch.Float, gotch.CPU).MustView(shape, true)

	fmt.Printf("%0.1f", x) // print truncated data
	fmt.Printf("%i", x)    // print truncated data

	// fmt.Printf("%#0.1f", x) // print full data
}
