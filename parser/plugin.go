package parser

import (
	"path"

	"github.com/charlesbases/generator"
	"github.com/charlesbases/reverse/dialer"
	"github.com/charlesbases/reverse/logger"
)

type Plugin struct {
	dialer dialer.Dialect

	schema  string
	structs []*Struct

	imports   map[string]*generator.ExternalPackage
	externals []*generator.ExternalPackage
}

// Run .
func Run(d dialer.Dialect) {
	var p = &Plugin{
		dialer:  d,
		schema:  d.Schema(),
		imports: make(map[string]*generator.ExternalPackage),
	}

	// tables to structs
	p.parse(d.Tables())

	// generate models
	p.generate()
}

// generate generate models
func (p *Plugin) generate() {
	generator.Run(func(gen *generator.Plugin) error {
		f := gen.NewGeneratedFile("models.go", p.dialer.Options().Dir, p.externals...)
		f.Writer("// Code generated by https://github.com/charlesbases/reverse. DO DOT EDIT")
		f.Writer("// schema: [", p.dialer.Options().Type, "] ", p.schema)
		f.Writer("// date:   ", today())
		f.Writer()
		f.Writer("package ", path.Base(p.dialer.Options().Dir))
		f.Writer()
		for _, item := range p.structs {
			f.Writer("// ", item.Name, " ", item.Desc)
			f.Writer("type ", item.Name, " struct {")
			for _, field := range item.Fields {
				f.Writer(field.Name, " ", field.Type, " ", field.Tag.String())
			}
			f.Writer("}")
			f.Writer()

			f.Writer("func (*", item.Name, ") TableName() string {")
			f.Writer(`return "` + item.Table.TableName + `"`)
			f.Writer("}")
			f.Writer()
		}
		return nil
	})

	logger.Inforf("complete [%d rows]", len(p.structs))
}
