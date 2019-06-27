package main

import "testing"

func TestDisplacement(t *testing.T) {
	tables := []struct {
		params       [3]float64
		time         float64
		displacement float64
	}{
		{[3]float64{10, 2, 1}, 3, 52},
		{[3]float64{2, 2, 1}, 5, 36},
	}

	for _, table := range tables {
		fn := GenDisplaceFn(table.params[0], table.params[1], table.params[2])
		displacement := fn(table.time)
		if displacement != table.displacement {
			t.Errorf("Displacement failed. Got %v, expected %v.", displacement, table.displacement)
		}
	}
}
