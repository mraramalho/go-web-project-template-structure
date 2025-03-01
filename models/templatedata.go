package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
