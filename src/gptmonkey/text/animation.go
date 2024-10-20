package text

import (
	"fmt"
	"time"
)

type Animation struct {
	Frames       []string
	CurrentFrame int
	FrameCount   int
}

func NewAnimation(frames []string) Animation {
	return Animation{Frames: frames, CurrentFrame: 0, FrameCount: len(frames)}
}

func (*Animation) Init() {
	fmt.Print("\033[s")
}

func (a *Animation) Animate() {
	fmt.Print("\033[u\033[K")
	fmt.Print(a.Frames[a.CurrentFrame])
	a.CurrentFrame = (a.CurrentFrame + 1) % a.FrameCount
	time.Sleep(time.Second / 100)
}
