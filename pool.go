package engine

// !! pool not effective

func Get() *Engine {
	return NewEngine()
}
func Put(e *Engine) {

}
