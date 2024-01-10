package nils

import "testing"

func TestMapsGoBoom(t *testing.T) {
	t.Run("will go boom", func(t *testing.T) {
		MapsGoBoom()
	})

	t.Run("recover", func(t *testing.T) {
		defer func() {
			if recover() != nil {
				t.Error("map went boom")
			}
		}()

		MapsGoBoom()
	})

}

func TestSlicesAreSafe(t *testing.T) {
	SlicesAreSafe()
}
