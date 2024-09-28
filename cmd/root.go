package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/Nathene/h3_backend/cmd/common"
	"github.com/Nathene/h3_backend/pkg/base"
)

const (
	SAVE_TO_FILE = "etc/saves/"
)

// unpack recursively unpacks the object and marshals it to JSON
func unpack(obj interface{}) (string, error) {
	v := reflect.ValueOf(obj)
	var data interface{} // To hold the unpacked data

	switch v.Kind() {
	case reflect.Ptr:
		if !v.IsNil() {
			// Recursively unpack the underlying value
			return unpack(v.Elem().Interface())
		}
		return "null", nil // Handle nil pointers
	case reflect.Struct:
		data = make(map[string]interface{})
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldName := v.Type().Field(i).Name
			fieldJSON, err := unpack(field.Interface()) // Recursively unpack the field
			if err != nil {
				return "", err
			}
			data.(map[string]interface{})[fieldName] = json.RawMessage(fieldJSON) // Use RawMessage
		}
	case reflect.Array, reflect.Slice:
		data = make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			elemJSON, err := unpack(v.Index(i).Interface()) // Recursively unpack the element
			if err != nil {
				return "", err
			}
			data.([]interface{})[i] = json.RawMessage(elemJSON) // Use RawMessage
		}
	case reflect.Map:
		data = make(map[string]interface{})
		for _, key := range v.MapKeys() {
			valueJSON, err := unpack(v.MapIndex(key).Interface()) // Recursively unpack the value
			if err != nil {
				return "", err
			}
			data.(map[string]interface{})[key.String()] = json.RawMessage(valueJSON) // Use RawMessage
		}
	default:
		data = v.Interface()
	}

	jsonData, err := json.MarshalIndent(data, "", "  ") // Use MarshalIndent for pretty printing
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func Run() {
	mainCharacter := base.NewPlayer("Nathan")
	fmt.Println(mainCharacter.Level)

	enemyAI, err := base.NewEnemyAI(mainCharacter) // Check for error from NewEnemyAI
	if err != nil {
		fmt.Println("Error creating EnemyAI:", err)
		return
	}

	game := common.NewGame(mainCharacter, enemyAI)

	gameOutput, _ := unpack(game)
	fmt.Println(gameOutput)

	if err := game.LevelUp(); err != nil {
		log.Fatal(err)
	}

	gameOutput, _ = unpack(game)
	fmt.Println(gameOutput)

	// fmt.Println(jsonOutput)

	// ghoul := base.NewEnemy(base.Ghoul)
	// ghoulOutput, err := unpack(ghoul)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(ghoulOutput)

	// if err := save(jsonOutput, "test1"); err != nil {
	// 	fmt.Printf("err: %v", err)
	// }
}

func save(obj string, filename string) error {
	err := os.WriteFile(SAVE_TO_FILE+filename+".json", []byte(obj), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	println(obj)
	return nil
}
