package utils

import (
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{
				path: "example.txt",
			},
			want: []string{
				"61087   87490",
				"31697   16584",
				"57649   82503",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadInput(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadInput() = %v, want %v", got, tt.want)
			}
		})
	}
}