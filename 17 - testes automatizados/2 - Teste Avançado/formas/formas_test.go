package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		ret := Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := ret.area()

		if expectedArea != receivedArea {
			t.Fatalf("o valor recebido não é compativel com o valor recebido, Recebido : %f, esperado: %f", receivedArea, expectedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		expectedArea := math.Pi * 100
		receivedArea := circle.area()

		if expectedArea != receivedArea {
			t.Fatalf("A area recebida: %f, é diferente da área esperada: %f", receivedArea, expectedArea)
		}
	})
}
