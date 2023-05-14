package model

import "testing"

func TestWordNumber_String(t *testing.T) {
	type fields struct {
		seq   int
		point Point
		aLen  int
		dLen  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1 across", fields{1, Point{1, 2}, 3, 4}, `seq:1,point:{r:1,c:2},aLen:3,dLen:4`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wn := &WordNumber{
				seq:   tt.fields.seq,
				point: tt.fields.point,
				aLen:  tt.fields.aLen,
				dLen:  tt.fields.dLen,
			}
			if got := wn.String(); got != tt.want {
				t.Errorf("WordNumber.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
