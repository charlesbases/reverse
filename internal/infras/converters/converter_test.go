package converters

import (
	"testing"
)

func TestConvertColumnName(t *testing.T) {
	tests := []struct {
		name     string
		acronyms []string
		args     string
		want     string
	}{
		{name: "1", acronyms: []string{"id"}, args: "id", want: "ID"},
		{name: "2", acronyms: []string{"id"}, args: "user_id", want: "UserID"},
		{name: "3", acronyms: []string{"ip"}, args: "ip", want: "IP"},
		{name: "4", acronyms: []string{"ip"}, args: "create_time", want: "CreateTime"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConverterRepository(tt.acronyms)
			if got := c.ConvertColumnName(tt.args); got != tt.want {
				t.Errorf("ConvertColumnName() = %v, want %v", got, tt.want)
			}
		})
	}
}
