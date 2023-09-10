package main

import "fmt"

func main() {

	player := NewPlayer()
	fmt.Println(player.Position)
	player.Move(1.1, 10.5)
	fmt.Println(player.Position)
	player.Teleport(5.0, 7.3)
	fmt.Println(player.Position)
	fmt.Println("-----------------")

	enemy := NewEnemy()
	fmt.Println(enemy.Position)
	enemy.Move(1.1, 10.5)
	fmt.Println(enemy.Position)
	enemy.Teleport(5.0, 7.3)
	enemy.MoveSpecial(1.1, 5.5)
	fmt.Println(enemy.Position)
}

type SpecialPosition struct {
	Position
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.x += x * x
	sp.y += y * y
}

type Position struct {
	x float64
	y float64
}

type Player struct {
	*Position
}

type Enemy struct {
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}
