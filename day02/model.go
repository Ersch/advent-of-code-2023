package day02

type Round struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	id     int
	rounds []Round
}
