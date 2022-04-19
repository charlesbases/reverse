package parser

import (
	"github.com/charlesbases/generator"
	"github.com/charlesbases/reverse/dialer"
	"github.com/charlesbases/reverse/types"
)

type (
	Struct struct {
		Name   string
		Desc   string
		Table  *dialer.Table
		Fields []*StructField
	}

	StructField struct {
		Name string    // 字段
		Tag  types.Tag // 标签
		Type string    // 类型
		Desc string    // 注释
	}
)

// parse .
func (p *Plugin) parse(tables []*dialer.Table) {
	p.structs = make([]*Struct, 0, len(tables))

	for _, table := range tables {
		var st = &Struct{Table: table}
		p.structs = append(p.structs, st.parse(p))
	}
}

// parse .
func (st *Struct) parse(p *Plugin) *Struct {
	st.Name = camelcase(st.Table.TableName)
	st.Desc = st.Table.TableDesc
	st.Fields = make([]*StructField, 0, len(st.Table.Columns))

	for _, column := range st.Table.Columns {
		st.Fields = append(st.Fields, st.field(p, column))
	}

	return st
}

// fields .
func (st *Struct) field(p *Plugin, tf *dialer.TableColumn) *StructField {
	var sf = &StructField{
		Name: camelcase(tf.ColumnName),
		Tag:  p.dialer.ParseColumnTag(tf),
		Type: p.dialer.ParseColumnType(tf),
		Desc: tf.ColumnDesc,
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
