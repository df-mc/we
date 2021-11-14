package brush

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"sync"
)

var handlers sync.Map

// LookupHandler finds the Handler of a specific player.Player, assuming it is currently online.
func LookupHandler(p *player.Player) (*Handler, bool) {
	v, _ := handlers.Load(p)
	h, ok := v.(*Handler)
	return h, ok
}

type Handler struct {
	p *player.Player
}

func NewHandler(p *player.Player) *Handler {
	h := &Handler{p: p}
	handlers.Store(p, h)
	return h
}

func (h *Handler) HandleItemUse(ctx *event.Context) {
	held, _ := h.p.HeldItems()
	if b, ok := find(held); ok {
		ctx.Cancel()
		b.Use(h.p)
	}
}

func (h *Handler) HandleQuit() {
	handlers.Delete(h.p)
}
