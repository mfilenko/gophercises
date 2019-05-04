package arrays

import "testing"

func TestEquilibrium(t *testing.T) {
	t.Run("True positive", func(t *testing.T) {
		balanced := [7]int{3, 2, 1, 0, 1, 2, 3}
		want := 3
		got := Equilibrium(balanced[:])

		if got != want {
			t.Errorf("Equilibrium(%v) = %d; want %d", balanced, got, want)
		}
	})

	t.Run("True negative", func(t *testing.T) {
		notBalanced := [6]int{3, 2, 1, 1, 2, 3}
		want := -1
		got := Equilibrium(notBalanced[:])

		if got != want {
			t.Errorf("Equilibrium(%v) = %d; want %d", notBalanced, got, want)
		}

	})
}
