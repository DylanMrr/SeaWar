package ai

import (
	"math/rand"
	"time"
)

var current int
var start [100][2]interface{}

func BuildMoves() {
	rand.Seed(time.Now().UnixNano())
	current = 0
	k := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			start[k] = [2]interface{}{i, j}
			k++
		}
	}
	rand.Shuffle(len(start), func(i, j int) { start[i], start[j] = start[j], start[i] })
}

func GetMove() (int, int) {
	i := start[current][0].(int)
	j := start[current][1].(int)
	current++
	return i, j
}
