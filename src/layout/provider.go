package layout


type LayoutProvider interface {
	TitleH2(title string) string
	TitleH3(title string) string
	ColourText(text, color string) string
	Table(rows [][]string) string
}