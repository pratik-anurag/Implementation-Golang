package array

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		numericArray []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Test for non numeric array 1",
			args: args{[]interface{}{"Hello", "World", "This", "Is", "Test 1"}},
			want: 0,
			wantErr: true,
		},
		{
			name: "Test for non numeric array 2",
			args: args{[]interface{}{1.2, 2, "This", "Is", "Test 1"}},
			want: 0,
			wantErr: true,
		},
		{
			name: "Test for float and int array",
			args: args{[]interface{}{1.0, 3.2, 6, 9.0, 12.0}},
			want: 31.2,
			wantErr: false,
		},
		{
			name: "Test for float array",
			args: args{[]interface{}{1.0, 3.2, 12.9}},
			want: 17.1,
			wantErr: false,
		},
		{
			name: "Test for int array",
			args: args{[]interface{}{10, 32, 129, 500}},
			want: 671,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sum(tt.args.numericArray)
			if (err != nil) != tt.wantErr {
				t.Errorf("sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sum() got = %v, want %v", got, tt.want)
			}
		})
	}
}