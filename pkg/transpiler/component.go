package transpiler

import (
	"bytes"
	"github.com/pubgo/errors"
	"io"
	"io/ioutil"
	"math/rand"
	"strings"
	"text/template"
	"time"

	"github.com/tdewolff/parse"
	"github.com/tdewolff/parse/css"
)

// ErrComponentNotParsed is returned when an attempt is made
// to Transform() a Component before calling Parse()

// A Component represents a web component in an HTML file
type Component struct {
	Name     string
	Template string
	Style    string
	Package  string
	Imports  []string
	parsed   bool
	Struct   bool
	UniqueID string
}

func (c *Component) WriteImports() string {
	return strings.Join(c.Imports, "\n\t")
}
func (c *Component) QuotedStyle() string {
	return "`" + c.Style + "`"
}

func (c *Component) QuotedTemplate() string {
	return "`" + c.Template + "`"
}

func (c *Component) TransformStyle() {
	p := css.NewParser(bytes.NewBufferString(c.Style), false)

	output := ""
	for {
		grammar, _, data := p.Next()
		data = parse.Copy(data)
		if grammar == css.ErrorGrammar {
			if err := p.Err(); err != io.EOF {
				for _, val := range p.Values() {
					data = append(data, val.Data...)
				}
				if perr, ok := err.(*parse.Error); ok && perr.Message == "unexpected token in declaration" {
					data = append(data, ";"...)
				}
			} else {
				break
			}
		} else if grammar == css.AtRuleGrammar || grammar == css.BeginAtRuleGrammar || grammar == css.QualifiedRuleGrammar || grammar == css.BeginRulesetGrammar || grammar == css.DeclarationGrammar || grammar == css.CustomPropertyGrammar {
			if grammar == css.DeclarationGrammar || grammar == css.CustomPropertyGrammar {
				data = append(data, ":"...)
			}
			for _, val := range p.Values() {
				data = append(data, val.Data...)
			}
			if grammar == css.BeginAtRuleGrammar || grammar == css.BeginRulesetGrammar {

				data = append(data, "."...)
				data = append(data, c.UniqueID...)
				data = append(data, "{"...)
			} else if grammar == css.AtRuleGrammar || grammar == css.DeclarationGrammar || grammar == css.CustomPropertyGrammar {
				data = append(data, ";"...)
			} else if grammar == css.QualifiedRuleGrammar {
				data = append(data, ","...)
			}
		}
		output += string(data)
	}

	c.Style = output

}

// Parse reads a component file like Nav.html into
// a Component structure
func Parse(r io.Reader, name string) *Component {
	var template, style string

	bb, err := ioutil.ReadAll(r)
	errors.Wrap(err, "reading component template")

	s := string(bb)
	styleStart := strings.Index(s, "<style>")
	if styleStart == -1 {
		styleStart = len(s)
	}
	template = strings.TrimSpace(s[:styleStart])
	if styleStart != len(s) {
		style = strings.TrimSpace(s[styleStart:])
		style = strings.Replace(style, "<style>", "", -1)
		style = strings.Replace(style, "</style>", "", -1)
	}
	return &Component{
		Name:     name,
		Template: template,
		Style:    style,
		UniqueID: randSeq(10),
		parsed:   true,
	}
}

func (c *Component) Transform(w io.Writer) {
	errors.T(!c.parsed, "transform must be called after parse")

	tpl := template.Must(template.New("component").Parse(comptpl))
	errors.Panic(tpl.Execute(w, c))
}

func removeStyleTags(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "<style>", "", -1)
	s = strings.Replace(s, "</style>", "", -1)
	return s
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
