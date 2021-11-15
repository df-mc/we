package msg

const (
	BlockPalette   = "Block palette"
	InvalidPalette = "Palette must exist and be non-empty."
)

const (
	BrushSelection = "Brush selection"
	BrushShape     = "Brush shape"
	BrushRadius    = "Brush radius"
	BrushAction    = "Brush action"
	FillMenu       = "Fill menu"
)

const (
	BindNeedsItem  = "You must hold an item to bind a brush."
	AlreadyBound   = "You can only bind one brush to an item."
	NotBound       = "The item held is not currently bound to a brush."
	BrushUnbound   = "Unbound brush from held item."
	NoUndo         = "No actions left to undo."
	UndoSuccessful = "Undid the last brush action. (%v left)"
)

const (
	StartPaletteSelection = "Select 2 points to create a palette."
	FirstPointSelected    = "Selected point 1 %v."
	SecondPointSelected   = "Selected point 2 %v."
	PaletteCreated        = "Palette created %v-%v."
	NoPaletteSelected     = "You have not currently selected a palette."
	PaletteSaved          = "Saved palette %v-%v to disk as '%v'."
	PaletteExists         = "Palette with name '%v' already exists. Use /palette delete %v to remove it."
	PaletteDoesNotExist   = "Palette with name '%v' does not exist."
	PaletteDeleted        = "Deleted palette '%v'."
)
