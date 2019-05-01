package main

import (
	"math"
	"time"

	"github.com/hashicorp/serf/coordinate"
)

func main() {

	

}

func dist(a *coordinate.Coordinate, b *coordinate.Coordinate) time.Duration  {

	if len(a.Vec) != len(b.Vec) {
		panic("dimensions aren't compatible")
	}

	//Calculate the Euclidean distance plus the heights
	sumsq := 0.0
	for i := 0; i < len(a.Vec); i++ {
		diff := a.Vec[i] - b.Vec[i]
		sumsq += diff * diff
	}

	rtt := math.Sqrt(sumsq) + a.Height + b.Height

	adjusted := rtt + a.Adjustment + b.Adjustment
	if adjusted > 0.0 {
		rtt = adjusted
	}

	const secondsToNanoseconds = 1.0e9
	return time.Duration(rtt * secondsToNanoseconds)
}
