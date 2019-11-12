package clients

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/packethost/packngo"
)

func TestNewClient(t *testing.T) {
	testClient := packngo.NewClientWithAuth("crossplane", "api-token", nil)
	type want struct {
		res *packngo.Client
	}
	tests := []struct {
		name string
		args string
		want want
	}{
		{name: "Test", args: "api-token", want: want{res: testClient}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(context.Background(), []byte(tt.args))
			if diff := cmp.Diff(got.APIKey, tt.want.res.APIKey); diff != "" {
				t.Errorf("TestNewClient() = %v, want %v\n%s", got, tt.want.res, diff)
			}
		})
	}
}
