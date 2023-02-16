package gotest

import "math"

// Pi é uma porpocao numérica definida
const Pi = 3.14

// Circ é um método para calcular raio
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// React calcular a partir da base e altura
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
