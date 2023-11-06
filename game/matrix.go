package game

import "math"

type Matrix [3][3]float64

func (m Matrix) Mul(n Matrix) Matrix {
	var l Matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				l[i][j] += m[i][k] * n[k][j]
			}
		}
	}
	return l
}

func (m Matrix) MulVec(v Vector) Vector {
	var w Vector
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			w[i] += m[i][k] * v[k]
		}
	}
	return w
}

func LongitudeRotate(theta float64) Matrix {
	var m Matrix
	m[0][0] = math.Cos(theta)
	m[0][1] = -math.Sin(theta)
	m[1][0] = math.Sin(theta)
	m[1][1] = math.Cos(theta)
	m[2][2] = 1
	return m
}

func LatitudeRotate(theta float64) Matrix {
	var m Matrix
	m[0][0] = 1
	m[1][1] = math.Cos(theta)
	m[1][2] = -math.Sin(theta)
	m[2][1] = math.Sin(theta)
	m[2][2] = math.Cos(theta)
	return m
}
