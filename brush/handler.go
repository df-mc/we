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
	p    *player.Player
	undo []func()
}

func NewHandler(p *player.Player) *Handler {
	h := &Handler{p: p}
	handlers.Store(p, h)
	return h
}

// UndoLatest undoes the latest brush action. If no action was left to undo, false is returned.
func (h *Handler) UndoLatest() bool {
	if len(h.undo) == 0 {
		return false
	}
	offset := len(h.undo) - 1
	h.undo[offset]()
	h.undo = h.undo[:offset]
	return true
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
