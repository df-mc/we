package palette

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/we/internal/msg"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"sync"
	"time"
)

var handlers sync.Map

// LookupHandler finds the Handler of a specific player.Player, assuming it is currently online.
func LookupHandler(p *player.Player) (*Handler, bool) {
	v, _ := handlers.Load(p)
	h, ok := v.(*Handler)
	return h, ok
}

type Handler struct {
	p        *player.Player
	close    chan struct{}
	palettes sync.Map

	mu        sync.Mutex
	m         Selection
	selecting int
	first     cube.Pos
}

func NewHandler(p *player.Player) *Handler {
	h := &Handler{p: p, close: make(chan struct{})}
	go h.visualisePalette()
	handlers.Store(p, h)
	return h
}

func (h *Handler) Palette(name string) (Palette, bool) {
	if name == "m" || name == "M" {
		h.mu.Lock()
		defer h.mu.Unlock()
		return h.m, true
	}
	p, _ := h.palettes.Load(name)
	b, ok := p.(Blocks)
	return b, ok
}

func (h *Handler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, _ cube.Face, _ mgl64.Vec3) {
	h.handleSelection(ctx, pos)
}

func (h *Handler) HandleBlockBreak(ctx *event.Context, pos cube.Pos) {
	h.handleSelection(ctx, pos)
}

func (h *Handler) HandleQuit() {
	close(h.close)
	handlers.Delete(h.p)
}

func (h *Handler) handleSelection(ctx *event.Context, pos cube.Pos) {
	if h.selecting == 0 {
		// Not currently selecting, return immediately.
		return
	}
	ctx.Cancel()

	h.selecting--
	if h.selecting == 1 {
		// Selecting the first point: Store it in the handler and return.
		h.first = pos
		h.p.Message(fmt.Sprintf(msg.FirstPointSelected, pos))
		return
	}
	// First point was selected, we now have a second point so we can create a palette.
	h.p.Message(fmt.Sprintf(msg.SecondPointSelected, pos))
	h.m = NewSelection(h.first, pos, h.p.World())
	h.p.Message(text.Colourf("<green>"+msg.PaletteCreated+"</green>", h.m.Min, h.m.Max))
}

func (h *Handler) visualisePalette() {
	t := time.NewTicker(time.Second / 20)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			p, _ := h.Palette("m")
			m := p.(Selection)
			_ = m
		case <-h.close:
			return
		}
	}
}
