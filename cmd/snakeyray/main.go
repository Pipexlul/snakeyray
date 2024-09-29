package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/pipexlul/snakeyray/cmd/config"
	gamewindow "github.com/pipexlul/snakeyray/src/game_window"
)

func main() {
	rl.SetTraceLogLevel(rl.LogError)

	gameWindow := gamewindow.NewGameWindow(
		config.NewWindowConfig(
			config.WithTargetFPS(90),
		),
	)

	log.Print("Hello Snakey!")
	gameWindow.Create()

	defer func() {
		log.Print("Bye Snakey!")
		gameWindow.Dispose()
	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
