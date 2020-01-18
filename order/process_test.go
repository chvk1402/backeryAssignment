package order

import (
	"github.com/backery/structs"
	"reflect"
	"testing"
)

func TestProcessOrder(t *testing.T) {
	type args struct {
		code     string
		quantity int
	}
	tests := []struct {
		name    string
		args    args
		want    *structs.OrderResp
		wantErr bool
	}{
		{
			name: "Test VS5",
			args: args{
				code:     "VS5",
				quantity: 10,
			},
			want: &structs.OrderResp{
				Code:       "VS5",
				TotalPrice: 17.98,
				Packs: []structs.Price{
					{
						Pack:   5,
						QtySet: 2,
						Price:  8.99,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test MB11",
			args: args{
				code:     "MB11",
				quantity: 14,
			},
			want: &structs.OrderResp{
				Code:       "MB11",
				TotalPrice: 54.8,
				Packs: []structs.Price{
					{
						Pack:   8,
						QtySet: 1,
						Price:  24.95,
					},
					{
						Pack:   2,
						QtySet: 3,
						Price:  9.95,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test CF",
			args: args{
				code:     "CF",
				quantity: 15,
			},
			want: &structs.OrderResp{
				Code:       "CF",
				TotalPrice: 28.89,
				Packs: []structs.Price{
					{
						Pack:   9,
						QtySet: 1,
						Price:  16.99,
					},
					{
						Pack:   3,
						QtySet: 2,
						Price:  5.95,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "non existent code",
			args: args{
				code:     "VJ",
				quantity: 15,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Packs cannot be divided equally",
			args: args{
				code:     "CF",
				quantity: 13,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessOrder(tt.args.code, tt.args.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
