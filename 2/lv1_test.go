package main

import "testing"

func TestProduct_Sell(t *testing.T) {
	type fields struct {
		Name  string
		Price float64
		Stock int
	}
	type args struct {
		amount int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantSuccess bool
		wantMessage string
	}{
		// TODO: Add test cases.
		{name: "message", fields: fields{Name: "message", Stock: 200}, args: args{amount: 100}, wantSuccess: true, wantMessage: "售买成功"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Message := &Product{
				Name:  tt.fields.Name,
				Price: tt.fields.Price,
				Stock: tt.fields.Stock,
			}
			gotSuccess, gotMessage := Message.Sell(tt.args.amount)
			if gotSuccess != tt.wantSuccess {
				t.Errorf("Sell() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
			if gotMessage != tt.wantMessage {
				t.Errorf("Sell() gotMessage = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
