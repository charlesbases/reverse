package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charlesbases/generator"
)

type Plugin struct {
	schema  string
	structs []*Struct

	imports   map[string]*generator.ExternalPackage
	externals []*generator.ExternalPackage
}

type Struct struct {
	Name   string
	Desc   string
	Table  *Table
	Fields []*StructField
}

type StructField struct {
	Name    string // 字段
	Tag     tag    // 标签
	Type    string // 类型
	Comment string // 注释
}

// parse table to struct
func (plugin *Plugin) parse() {
	for _, st := range plugin.structs {
		if st.Table != nil {
			st.parse(plugin)
		}
	}
}

// parse struct fields
func (st *Struct) parse(plugin *Plugin) {
	st.Fields = make([]*StructField, 0, len(st.Table.Fields))
	for _, tf := range st.Table.Fields {
		field := tf.parse()
		switch field.Type {
		case "time.Time":
			if _, find := plugin.imports[field.Type]; !find {
				plugin.imports[field.Type] = generator.NewExternalPackage("time")
				plugin.externals = append(plugin.externals, plugin.imports[field.Type])
			}
		}

		st.Fields = append(st.Fields, field)
	}
}

// parse table field to struct field
func (tf *TableField) parse() *StructField {
	return &StructField{
		Name:    camelcase(tf.FieldName),
		Tag:     tf.parseTag(),
		Type:    tf.parseType(),
		Comment: tf.FieldComment,
	}
}

// parseTag .
func (tf *TableField) parseTag() tag {
	var builder = strings.Builder{}

	{
		// json tag
		builder.WriteString(fmt.Sprintf(`json:"%s"`, tf.FieldName))
	}

	{
		// orm tag
		builder.WriteString(fmt.Sprintf(` gorm:"column:%s;type:%s`, tf.FieldName, tf.FieldType))
		if tf.IsNull == "NO" {
			builder.WriteString(";not null")
		}
		if tf.FieldKey == "PRI" {
			builder.WriteString(";primary_key")
		}
		if tf.Extra == "auto_increment" {
			builder.WriteString(";auto_increment")
		}
		builder.WriteString(`"`)
	}
	return tag(builder.String())
}

// parseType .
func (tf *TableField) parseType() string {
	var gotype = mysqltype[tf.DataType]
	if strings.HasSuffix(tf.FieldType, "unsigned") {
		return "u" + gotype
	} else {
		return gotype
	}
}

// camelcase aaa_bbb to AaaBbb
func camelcase(s string) string {
	if b, find := abbre[s]; find {
		return b
	} else {
		var bs strings.Builder
		bs.Grow(len(s))
		for _, item := range strings.Split(s, "_") {
			if len(item) > 0 {
				if b, find := abbre[item]; find {
					bs.WriteString(b)
				} else {
					if c := item[0]; isASCIILower(c) {
						bs.WriteByte(c - ('a' - 'A'))
					} else {
						bs.WriteByte(c)
					}
					if len(item) > 1 {
						bs.WriteString(item[1:])
					}
				}
			}
		}
		return bs.String()
	}
}

// encamelcase AaaBbb to aaa_bbb
func encamelcase(s string) string {
	builder := strings.Builder{}
	for index, letter := range []byte(s) {
		if isASCIIUpper(letter) {
			if index != 0 {
				builder.WriteString("_")
			}
			builder.WriteByte(letter + 'a' - 'A')
		} else {
			builder.WriteByte(letter)
		}
	}
	return builder.String()
}

func today() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}
