package core

import (
    "math"
    "time"
)

func ParseDateToCoords(date time.Time) HypercubeCoords {
    day := date.Day()
    month := int(date.Month())
    year := date.Year()

    x := toThreeDigitPreserve(day)
    y := toThreeDigitPreserve(month)
    z := toThreeDigitPreserve(year)

    w := sumDigits(x) + sumDigits(y) + sumDigits(z)

    return HypercubeCoords{
        X: int32(x),
        Y: int32(y),
        Z: int32(z),
        W: int32(w),
    }
}

func toThreeDigitPreserve(n int) int {
    if n < 10 {
        return n * 100
    }
    if n < 100 {
        return n * 10
    }
    return n % 1000
}

func sumDigits(n int) int {
    sum := 0
    n = int(math.Abs(float64(n)))
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}

func ExtractDigits(n int32) []int32 {
    n = int32(math.Abs(float64(n)))
    return []int32{
        (n / 100) % 10,
        (n / 10) % 10,
        n % 10,
    }
}

func CalculateVectors(coords HypercubeCoords) PersonVectors {
    digitsX := ExtractDigits(coords.X)
    digitsY := ExtractDigits(coords.Y)
    digitsZ := ExtractDigits(coords.Z)

    return PersonVectors{
        X: []int32{
            digitsX[0] - digitsX[1],
            digitsX[1] - digitsX[2],
            digitsX[2] - digitsX[0],
        },
        Y: []int32{
            digitsY[0] - digitsY[1],
            digitsY[1] - digitsY[2],
            digitsY[2] - digitsY[0],
        },
        Z: []int32{
            digitsZ[0] - digitsZ[1],
            digitsZ[1] - digitsZ[2],
            digitsZ[2] - digitsZ[0],
        },
        Full: [3][]int32{
            {
                digitsX[0] - digitsX[1],
                digitsX[1] - digitsX[2],
                digitsX[2] - digitsX[0],
            },
            {
                digitsY[0] - digitsY[1],
                digitsY[1] - digitsY[2],
                digitsY[2] - digitsY[0],
            },
            {
                digitsZ[0] - digitsZ[1],
                digitsZ[1] - digitsZ[2],
                digitsZ[2] - digitsZ[0],
            },
        },
    }
}

func VectorAmplitude(vectors PersonVectors) float64 {
    sum := 0.0
    count := 0
    for _, vec := range vectors.Full {
        for _, v := range vec {
            sum += math.Abs(float64(v))
            count++
        }
    }
    if count == 0 {
        return 0
    }
    return sum / float64(count)
}

func ParseDate(dateStr string) (time.Time, error) {
    formats := []string{
        "02.01.2006",
        "2.1.2006",
        "02.1.2006",
        "2.01.2006",
        "2006-01-02",
        "2006-1-2",
    }
    var t time.Time
    var err error
    for _, f := range formats {
        t, err = time.Parse(f, dateStr)
        if err == nil {
            return t, nil
        }
    }
    return t, err
}
