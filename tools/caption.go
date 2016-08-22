package present

import "strings"

func init() {
	Register("caption", parseCaption)
}

type Caption struct {
	Text string
}

func (c Caption) TemplateName() string { return "caption" }

func parseCaption(_ *Context, _ string, _ int, text string) (Elem, error) {
	text = strings.TrimSpace(strings.TrimPrefix(text, ".caption"))
	return Caption{text}, nil
}
