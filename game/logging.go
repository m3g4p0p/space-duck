package game

import (
	"bytes"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi/ecs"
)

var (
	buf    bytes.Buffer
	logger = slog.New(slog.NewJSONHandler(&buf, nil))
)

func init() {
	slog.SetDefault(logger)
}

func FlushLogs(_ *ecs.ECS, image *ebiten.Image) {
	ebitenutil.DebugPrint(image, buf.String())
	buf.Reset()
}
