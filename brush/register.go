package brush

var (
	actions     = map[string]func() Action{}
	actionNames []string
	shapes      = map[string]func(r int) Shape{}
	shapeNames  []string
)

// RegisterAction registers a function that returns an act.Action with the name of the action passed.
func RegisterAction(name string, v func() Action) {
	actions[name] = v
	actionNames = append(actionNames, name)
}

// RegisterShape registers a function that returns a geo.Shape with the name of the shape passed.
func RegisterShape(name string, v func(r int) Shape) {
	shapes[name] = v
	shapeNames = append(shapeNames, name)
}

// shapeByName returns a geo.Shape by the name passed, giving it a radius of the r passed.
func shapeByName(name string, r int) Shape {
	return shapes[name](r)
}

// actionByName returns an act.Action by the name passed.
func actionByName(name string) Action {
	return actions[name]()
}
