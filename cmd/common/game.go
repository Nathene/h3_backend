package common

import (
	"fmt"

	"github.com/Nathene/h3_backend/pkg/base"
)

type Game struct {
	Player  base.Player
	Enemies base.EnemyAI
}

func NewGame(p *base.Player, e *base.EnemyAI) *Game {
	return &Game{
		Player:  *p,
		Enemies: *e,
	}
}

func (g *Game) LevelUp() error {
	g.Player.Level++
	err := g.UpdateEnemies()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) UpdateEnemies() error {
	enemyAI, err := base.NewEnemyAI(&g.Player) // Create new EnemyAI with updated player level
	if err != nil {
		return fmt.Errorf("error creating EnemyAI: %w", err)
	}
	g.Enemies = *enemyAI
	return nil
}
