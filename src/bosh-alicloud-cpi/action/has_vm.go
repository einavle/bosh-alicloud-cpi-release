package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"alibaba/bosh-alicloud-cpi/src/bosh-alicloud-cpi/alicloud"
)

type HasVMMethod struct {
	runner alicloud.Runner
}

func NewHasVMMethod(runner alicloud.Runner) HasVMMethod {
	return HasVMMethod{runner}
}

func (a HasVMMethod) HasVM(cid apiv1.VMCID) (bool, error) {
	//
	//
	instid := cid.AsString()
	inst, err := a.runner.GetInstance(instid)

	if err != nil {
		return false, bosherr.WrapErrorf(err, "Finding VM Failed '%s'", cid)
	}

	if inst != nil {
		return true, nil
	} else {
		return false, nil
	}
}
