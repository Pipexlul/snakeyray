package config

type GameWindowConfig struct {
	Title     string
	Width     int32
	Height    int32
	MaxWidth  int
	MaxHeight int
	MinWidth  int
	MinHeight int
	TargetFPS int
}

func getDefaultWindowConfig() *GameWindowConfig {
	return &GameWindowConfig{
		Title:     "SnakeyRay",
		Width:     800,
		Height:    800,
		MinWidth:  400,
		MinHeight: 400,
		MaxWidth:  1080,
		MaxHeight: 1080,
		TargetFPS: 60,
	}
}

type WindowOption func(*GameWindowConfig)

func WithTitle(title string) WindowOption {
	return func(config *GameWindowConfig) {
		config.Title = title
	}
}

func WithWindowSize(width, height int32) WindowOption {
	return func(config *GameWindowConfig) {
		config.Width = width
		config.Height = height
	}
}

func WithMaxWindowSize(width, height int) WindowOption {
	return func(config *GameWindowConfig) {
		config.MaxWidth = width
		config.MaxHeight = height
	}
}

func WithMinWindowSize(width, height int) WindowOption {
	return func(config *GameWindowConfig) {
		config.MinWidth = width
		config.MinHeight = height
	}
}

func WithTargetFPS(fps int) WindowOption {
	return func(config *GameWindowConfig) {
		config.TargetFPS = fps
	}
}

func NewWindowConfig(opts ...WindowOption) *GameWindowConfig {
	config := getDefaultWindowConfig()

	for _, opt := range opts {
		opt(config)
	}

	return config
}
