package base

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

const (
	ENEMY_PATH = "etc/config/enemies/"
)

// Enemy represents an enemy character
type Enemy struct {
	Name  string
	Stats Stats
}

// Stats holds the statistics of an enemy

// EnemyConfig holds the configuration for an enemy type
type EnemyConfig struct {
	BaseStats       Stats           `json:"BaseStats"`
	LevelMultiplier LevelMultiplier `json:"LevelMultiplier"`
}

// Config is a map of enemy names to their configurations
type Config map[string]EnemyConfig

// loadConfig loads enemy configurations from a JSON file
func loadConfig(level int) (Config, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s%d%s", ENEMY_PATH, level, ".json"))
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// newEnemy creates a new enemy with calculated stats based on level and configuration
func newEnemy(enemyType int, level int, config Config) (*Enemy, error) {
	var enemyName string
	switch enemyType {
	case Ghoul:
		enemyName = "Ghoul"
	case Vampire:
		enemyName = "Vampire"
	// Add cases for other enemy types
	default:
		return nil, fmt.Errorf("invalid enemy type: %d", enemyType)
	}

	enemyConfig, ok := config[enemyName]
	if !ok {
		return nil, fmt.Errorf("enemy not found in config: %s", enemyName)
	}

	stats := enemyConfig.BaseStats

	// Apply level multipliers
	stats.Hp = int(float64(stats.Hp) * math.Pow(float64(enemyConfig.LevelMultiplier.Hp), float64(level-1)))
	stats.AttackPower = int(float64(stats.AttackPower) * math.Pow(float64(enemyConfig.LevelMultiplier.AttackPower), float64(level-1)))
	stats.DefensePower = int(float64(stats.DefensePower) * math.Pow(float64(enemyConfig.LevelMultiplier.DefensePower), float64(level-1)))
	stats.Accuracy = int(float64(stats.Accuracy) * math.Pow(float64(enemyConfig.LevelMultiplier.Accuracy), float64(level-1)))

	stats.Level = level

	return &Enemy{
		Name:  enemyName,
		Stats: stats,
	}, nil
}

// EnemyAI manages a group of enemies
type EnemyAI struct {
	Enemies []Enemy
	Level   int
}

// newEnemyAI creates a new EnemyAI with enemies based on the player's level
func NewEnemyAI(p *Player) (*EnemyAI, error) {
	config, err := loadConfig(p.Level)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return nil, err // Or handle thse error appropriately
	}

	var enemies []Enemy
	switch p.Level {
	case 1:
		ghoul, err := newEnemy(Ghoul, 1, config)
		if err != nil {
			fmt.Println("Error creating Ghoul:", err)
			return nil, err // Or handle the error appropriately
		}
		enemies = append(enemies, *ghoul)
	case 2:
		ghoul, err := newEnemy(Ghoul, 2, config)
		if err != nil {
			fmt.Println("Error creating Ghoul:", err)
			return nil, err // Or handle the error appropriately
		}
		enemies = append(enemies, *ghoul)

		vampire, err := newEnemy(Vampire, 2, config)
		if err != nil {
			fmt.Println("Error creating Vampire:", err)
			return nil, err // Or handle the error appropriately
		}
		enemies = append(enemies, *vampire)
		// Add cases for other levels
	}

	return &EnemyAI{
		Enemies: enemies,
		Level:   p.Level,
	}, nil
}

// Enemy type constants
const (
	Ghoul   = iota // 0
	Vampire        // 1
	Zombie         // 2
	maxNumEnemies
)

// LoadEnemies is no longer needed, so it's removed
// func LoadEnemies(p *Player) *Enemies { ... }

// newGhoul and newVampire are no longer needed, so they're removed
// func newGhoul(cfg EnemyConfig) *Enemy { ... }
// func newVampire(cfg EnemyConfig) *Enemy { ... }
