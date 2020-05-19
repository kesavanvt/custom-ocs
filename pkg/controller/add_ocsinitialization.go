package controller

import (
	"github.com/openshift/ocs-operator/pkg/controller/ocsinitialization"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, ocsinitialization.Add)
}
