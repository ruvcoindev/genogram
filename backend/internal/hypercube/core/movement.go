package core

import (
    "math"
)

func MoveRoom(coords HypercubeCoords, vectors PersonVectors, step int) HypercubeCoords {
    newCoords := coords
    vectorIdx := step % 3

    newCoords.X += vectors.Full[0][vectorIdx]
    newCoords.Y += vectors.Full[1][vectorIdx]
    newCoords.Z += vectors.Full[2][vectorIdx]

    newCoords.W = int32(sumDigits(int(newCoords.X)) + sumDigits(int(newCoords.Y)) + sumDigits(int(newCoords.Z)))

    newCoords.X = modulo(newCoords.X, CubeSize)
    newCoords.Y = modulo(newCoords.Y, CubeSize)
    newCoords.Z = modulo(newCoords.Z, CubeSize)

    return newCoords
}

func modulo(a, m int32) int32 {
    return ((a % m) + m) % m
}

func IsBridgeRoom(coords HypercubeCoords) bool {
    return coords.W == 27 || coords.X == 999 || coords.Y == 999 || coords.Z == 999
}

func IsTrapRoom(coords HypercubeCoords) bool {
    digitsX := ExtractDigits(coords.X)
    digitsY := ExtractDigits(coords.Y)
    digitsZ := ExtractDigits(coords.Z)

    isTrap := func(d []int32) bool {
        return len(d) == 3 && d[0] == d[1] && d[1] == d[2] && d[0] > 0
    }

    return isTrap(digitsX) || isTrap(digitsY) || isTrap(digitsZ)
}

func EuclideanDistance(a, b HypercubeCoords) float64 {
    dx := float64(a.X - b.X)
    dy := float64(a.Y - b.Y)
    dz := float64(a.Z - b.Z)
    dw := float64(a.W - b.W)
    return math.Sqrt(dx*dx + dy*dy + dz*dz + dw*dw)
}
