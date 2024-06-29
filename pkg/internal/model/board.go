package model

type Board struct {
	Snakes map[int]int
	Ladder map[int]int
	Size   int
}
