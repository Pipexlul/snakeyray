package gamewindow

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/pipexlul/snakeyray/cmd/config"
)

type GameWindow struct {
	Config *config.GameWindowConfig
}

func NewGameWindow(config *config.GameWindowConfig) *GameWindow {
	return &GameWindow{
		Config: config,
	}
}

func (w *GameWindow) Create() {
	rl.InitWindow(w.Config.Width, w.Config.Height, w.Config.Title)

	rl.SetWindowState(rl.FlagWindowResizable)

	rl.SetWindowMinSize(w.Config.MinWidth, w.Config.MinHeight)
	rl.SetWindowMaxSize(w.Config.MaxWidth, w.Config.MaxHeight)
}

func (w *GameWindow) Dispose() {
	rl.CloseWindow()
}
