package parser

import (
	"github.com/charlesbases/generator"
	"github.com/charlesbases/reverse/dialect"
	"github.com/charlesbases/reverse/types"
)

type (
	Struct struct {
		Name   string
		Desc   string
		Table  *dialect.Table
		Fields []*StructField
	}

	StructField struct {
		Name    string    // 字段
		Tag     types.Tag // 标签
		Type    string    // 类型
		Comment string    // 注释
	}
)

// parse .
func (p *Plugin) parse(tables []*dialect.Table) {
	p.structs = make([]*Struct, 0, len(tables))

	for _, table := range tables {
		var st = &Struct{Table: table}
		p.structs = append(p.structs, st.parse(p))
	}
}

// parse .
func (st *Struct) parse(p *Plugin) *Struct {
	st.Name = camelcase(st.Table.TableName)
	st.Desc = st.Table.TableComment
	st.Fields = make([]*StructField, 0, len(st.Table.Fields))

	for _, column := range st.Table.Fields {
		st.Fields = append(st.Fields, st.field(p, column))
	}

	return st
}

// fields .
func (st *Struct) field(p *Plugin, tf *dialect.TableColumn) *StructField {
	var sf = &StructField{
		Name:    camelcase(tf.ColumnName),
		Tag:     p.dialect.ParseColumnTag(tf),
		Type:    p.dialect.ParseColumnType(tf),
		Comment: tf.ColumnComment,
	}

	switch sf.Type {
	case "time.Time":
		if _, find := p.imports["time.Time"]; !find {
			p.imports["time.Time"] = generator.NewExternalPackage("time")
			p.externals = append(p.externals, p.imports["time.Time"])
		}
	}
	return sf
}
