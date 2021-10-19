package worldedit

import "github.com/df-mc/dragonfly/server/player"

// Session is a type that contains []*Editor to keep track of all player world *Positions
type Session struct {
	Editors []*Editor
}

// NewSession returns a new *Session
func NewSession() *Session { return &Session{make([]*Editor, 0)} }

// EditorByPlayer returns the *Editor found by searching for the *player.Player passed
// if it is not yet registered, it will be registered and returned
func (s *Session) EditorByPlayer(p *player.Player) *Editor {
	for _, p2 := range s.Editors {
		if p2.Player == p {
			return p2
		}
	}
	editor := NewEditor(p)
	s.Editors = append(s.Editors, editor)
	return editor
}
