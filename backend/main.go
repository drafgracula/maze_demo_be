package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Maze struct {
	Size    int               `json:"size"`
	Content map[string]string `json:"content"`
}

func genMapJson(size int) (string, error) {
	if size <= 0 {
		return "", fmt.Errorf("Size must be a positive integer")
	}

	content := make(map[string]string)
	rand.Seed(time.Now().UnixNano())

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			key := fmt.Sprintf("%d.%d", x, y)
			// Подстраивайте вероятность по вашему усмотрению
			if rand.Float64() < 0.3 {
				content[key] = "1"
			} else {
				content[key] = "0"
			}
		}
	}

	maze := Maze{
		Size:    size,
		Content: content,
	}

	jsonString, err := json.Marshal(maze)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

func main() {
	size := 128
	mazeJSON, err := genMapJson(size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Maze JSON:")
	fmt.Println(mazeJSON)
}
