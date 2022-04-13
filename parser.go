package main

import (
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/charlesbases/generator"
	"github.com/charlesbases/reverse/logger"
)

// options .
type options struct {
	Name string
	Dsn  string
	Dir  string
}

type Table struct {
	TableName    string        // 表名
	TableComment string        // 表注释
	Fields       []*TableField // 表字段
}

type TableField struct {
	FieldName    string // 列名
	FieldKey     string // 键类别
	Extra        string // 自增
	IsNull       string // NOT NULL
	DataType     string // 类型
	FieldType    string // 类型+长度
	FieldComment string // 注释
}

// DefaultOptions .
func DefaultOptions() *options {
	return &options{
		Dir: "./models",
	}
}

// decode decode toml file
func decode(fpath string) *options {
	abspath, err := filepath.Abs(fpath)
	if err != nil {
		logger.Fatal(err)
	}

	var opts = DefaultOptions()
	if _, err := toml.DecodeFile(abspath, opts); err != nil {
		logger.Fatal(err)
	}

	if opts.Dsn == "" {
		logger.Fatal(ErrInvalidDsn)
	}

	return opts
}

// parse .
func (opts *options) parse() *Plugin {
	switch opts.Name {
	case "mysql":
		return MysqlDialector(opts.Dsn).Parse()
	case "postgres":
		return PostgresDialector(opts.Dsn).Parse()
	default:
		logger.Fatalf("unsupported %s", opts.Name)
	}
	return nil
}

func (opts *options) run() {
	var plugin = opts.parse()
	plugin.parse()

	generator.Run(func(gen *generator.Plugin) error {
		f := gen.NewGeneratedFile("models.go", opts.Dir, plugin.externals...)
		f.Writer("// Code generated by https://github.com/charlesbases/reverse. DO DOT EDIT")
		f.Writer("// schema: [", opts.Name, "] ", plugin.schema)
		f.Writer("// date:   ", today())
		f.Writer()
		f.Writer("package ", path.Base(opts.Dir))
		f.Writer()
		for _, item := range plugin.structs {
			f.Writer("// ", item.Name, " ", item.Desc)
			f.Writer("type ", item.Name, " struct {")
			for _, field := range item.Fields {
				f.Writer(field.Name, " ", field.Type, " ", field.Tag.string())
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

	logger.Inforf("complete [%d rows]", len(plugin.structs))
}
