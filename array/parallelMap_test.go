package array

import (
	"reflect"
	"testing"
)

func TestParallelMap(t *testing.T) {
	type args struct {
		source    interface{}
		transform interface{}
		T         reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test ideal scenario",
			args: args{
				source:    []int{1, 2, 3},
				transform: func(num int) int { return num + 1 },
				T:         reflect.TypeOf(1),
			},
			want: []int{2, 3, 4},
			wantErr: false,
		},
		{
			name: "Test nil array",
			args: args{
				source:    nil,
				transform: func(num int) int { return num + 1 },
				T:         reflect.TypeOf(1),
			},
			wantErr: true,
		},
		{
			name: "Test nil function",
			args: args{
				source:    []int{1, 2, 3},
				transform: nil,
				T:         reflect.TypeOf(1),
			},
			wantErr: true,
		},
		{
			name: "Test nil Type",
			args: args{
				source:    []int{1, 2, 3},
				transform: func(num int) int { return num + 1 },
				T:         nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParallelMap(tt.args.source, tt.args.transform, tt.args.T)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParallelMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParallelMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}
