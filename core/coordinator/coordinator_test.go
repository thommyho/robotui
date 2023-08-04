package coordinator

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/core/loadpoint"
	"github.com/thommyho/robotui/mock"
	"github.com/thommyho/robotui/util"
)

func TestVehicleDetectByStatus(t *testing.T) {
	ctrl := gomock.NewController(t)

	type vehicle struct {
		*mock.MockVehicle
		*mock.MockChargeState
	}

	v1 := &vehicle{mock.NewMockVehicle(ctrl), mock.NewMockChargeState(ctrl)}
	v2 := &vehicle{mock.NewMockVehicle(ctrl), mock.NewMockChargeState(ctrl)}

	type testcase struct {
		string
		v1, v2 api.ChargeStatus
		res    api.Vehicle
	}
	tc := []testcase{
		{"A/A->0", api.StatusA, api.StatusA, nil},
		{"B/A->1", api.StatusB, api.StatusA, v1},
		{"B/A->1", api.StatusB, api.StatusA, v1},
		{"A/B->2", api.StatusA, api.StatusB, v2},
		{"A/B->2", api.StatusA, api.StatusB, v2},
		{"B/B->1", api.StatusB, api.StatusB, nil},
	}

	log := util.NewLogger("foo")
	vehicles := []api.Vehicle{v1, v2}

	v1.MockVehicle.EXPECT().Title().Return("v1").AnyTimes()
	v2.MockVehicle.EXPECT().Title().Return("v2").AnyTimes()
	v1.MockVehicle.EXPECT().Identifiers().Return(nil).AnyTimes()
	v2.MockVehicle.EXPECT().Identifiers().Return([]string{"it's me"}).AnyTimes()

	var lp loadpoint.API
	c := New(log, vehicles)

	for _, tc := range tc {
		t.Logf("%+v", tc)

		v1.MockChargeState.EXPECT().Status().Return(tc.v1, nil)
		v2.MockChargeState.EXPECT().Status().Return(tc.v2, nil)

		available := c.availableDetectibleVehicles(lp) // include id-able vehicles
		res := c.identifyVehicleByStatus(available)
		if tc.res != res {
			t.Errorf("expected %v, got %v", tc.res, res)
		}

		if res != nil {
			c.acquire(lp, res)
		} else {
			c.release(res)
		}
	}
}
