package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

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
	var s strings.Builder

	for line := range bytes.SplitSeq(buf.Bytes(), []byte{'\n'}) {
		var data any

		err := json.Unmarshal(line, &data)
		if err != nil {
			continue
		}

		val, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			continue
		}

		fmt.Fprintln(&s, string(val))
	}

	ebitenutil.DebugPrint(image, s.String())
	buf.Reset()
}
