package models

// TemplateData holds data sent from handlers to templagtes
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string /* tool tip message */
	Warning   string
	Error     string
}
