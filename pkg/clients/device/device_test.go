package device

import (
	"github.com/crossplaneio/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/packethost/packngo"
	"github.com/packethost/stack-packet/apis/server/v1alpha1"
	"testing"
)

const (
	deviceName = "test"
)

func TestCreateFromDevice(t *testing.T) {

	v1alphaReq := v1alpha1.Device{
		Spec: v1alpha1.DeviceSpec{ForProvider: v1alpha1.DeviceParameters{
			Hostname:     deviceName,
			Plan:         "baremetal_0",
			Facility:     "ams1",
			OS:           "ubuntu_16_04",
			ProjectID:    "fakeId",
			BillingCycle: "hourly",
		}},
	}
	got := CreateFromDevice(&v1alphaReq)

	want := packngo.DeviceCreateRequest{
		Hostname:     deviceName,
		Plan:         "baremetal_0",
		Facility:     []string{"ams1"},
		OS:           "ubuntu_16_04",
		ProjectID:    "fakeId",
		BillingCycle: "hourly",
	}

	if diff := cmp.Diff(&want, got, test.EquateErrors()); diff != "" {
		t.Errorf("client.Create(): -want, +got:\n%s", diff)
	}

}
