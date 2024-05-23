package utils

import (
	"encoding/json"
	"os"
	"os/exec"
)

func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func SaveData[T any](data T, filename string) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(filename, file, 0644)
}

func LoadData[T any](filename string) T {
	var data T
	file, _ := os.ReadFile(filename)
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
