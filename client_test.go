package gofb

import (
	"testing"
)

func TestClient_Query(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		want    *Response
		wantErr bool
	}{
		{
			name:    "DropTable",
			query:   "drop table if exists test",
			want:    &Response{},
			wantErr: false,
		},
		{
			name:    "CreateTable",
			query:   "create table test (_id id, name string)",
			want:    &Response{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(
				&Options{
					Host: "localhost",
					Port: "10101",
				})
			got, err := c.Query(tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !IsEqualResponse(got, tt.want) {
				t.Errorf("Client.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
