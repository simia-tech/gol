package gol

type mapper func(string) string

func buildTerminalColorMapper(color color) mapper {
	return func (input string) string {
		return surroundColor(input, color)
	}
}
