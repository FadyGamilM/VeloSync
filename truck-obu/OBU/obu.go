package obu

import (
	"math"
	"math/rand"

	"github.com/FadyGamilM/truckobu/types"
)

func generateLatitude() float64 {
	return generateRandCoordinate()
}

func generateLongitude() float64 {
	return generateRandCoordinate()
}

func generateRandCoordinate() float64 {
	return rand.Float64() + float64(rand.Intn(100)+1) // +1 to start from 1 not 0
}

func GenerateObuData(n int) []types.ObuData {
	obuReads := []types.ObuData{}
	for i := 0; i < n; i++ {
		id := rand.Intn(math.MaxInt)
		obuReads = append(obuReads, types.ObuData{
			ID:        int64(id),
			Latitude:  generateLatitude(),
			Longitude: generateLongitude(),
		})
	}
	return obuReads
}
