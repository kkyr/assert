package assert

import (
	"fmt"
	"text/template"
)

const (
	DefaultTmpl = `{{if .FieldVal}}{{ .FieldVal }}: {{end}}{{ .GotText }} {{ .GotVerb }}, {{ .WantText }} {{ .WantVerb }}`

	GotTextKey  = "GotText"
	GotVerbKey  = "GotVerb"
	WantTextKey = "WantText"
	WantVerbKey = "WantVerb"
	FieldValKey = "FieldVal"

	DefaultGotText  = "got"
	DefaultGotVerb  = "%v"
	DefaultWantText = "want"
	DefaultWantVerb = "%v"
)

func (a *Asserter) SetGotText(text string) {
	a.gotText = text
}

func (a *Asserter) SetGotVerb(verb string) {
	a.gotVerb = verb
}

func (a *Asserter) SetWantText(text string) {
	a.wantText = text
}

func (a *Asserter) SetWantVerb(verb string) {
	a.wantVerb = verb
}

func (a *Asserter) SetTemplate(text string) error {
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	a.tmpl = tmpl

	return nil
}

func SetGotText(text string) {
	std.SetGotText(text)
}

func SetGotVerb(verb string) {
	std.SetGotVerb(verb)
}

func SetWantText(text string) {
	std.SetWantText(text)
}

func SetWantVerb(verb string) {
	std.SetWantVerb(verb)
}

func SetTemplate(text string) error {
	return std.SetTemplate(text)
}
