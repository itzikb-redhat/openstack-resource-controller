/*
Copyright 2024 The ORC Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package securitygroup

import (
	"context"

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"

	ctrlexport "github.com/k-orc/openstack-resource-controller/internal/controllers/export"
	"github.com/k-orc/openstack-resource-controller/internal/scope"
)

const (
	Finalizer = "openstack.k-orc.cloud/securitygroup"

	FieldOwner = "openstack.k-orc.cloud/securitygroupcontroller"
	// Field owner of the object finalizer.
	SSAFinalizerTxn = "finalizer"
	// Field owner of transient status.
	SSAStatusTxn = "status"
)

// ssaFieldOwner returns the field owner for a specific named SSA transaction.
func ssaFieldOwner(txn string) client.FieldOwner {
	return client.FieldOwner(FieldOwner + "/" + txn)
}

type securitygroupReconcilerConstructor struct {
	scopeFactory scope.Factory
}

func New(scopeFactory scope.Factory) ctrlexport.Controller {
	return securitygroupReconcilerConstructor{
		scopeFactory: scopeFactory,
	}
}

func (securitygroupReconcilerConstructor) GetName() string {
	return "securitygroup"
}

// orcSecurityGroupReconciler reconciles an ORC Subnet.
type orcSecurityGroupReconciler struct {
	client       client.Client
	recorder     record.EventRecorder
	scopeFactory scope.Factory
}

// SetupWithManager sets up the controller with the Manager.
func (c securitygroupReconcilerConstructor) SetupWithManager(_ context.Context, mgr ctrl.Manager, options controller.Options) error {
	reconciler := orcSecurityGroupReconciler{
		client:       mgr.GetClient(),
		recorder:     mgr.GetEventRecorderFor("orc-securitygroup-controller"),
		scopeFactory: c.scopeFactory,
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&orcv1alpha1.SecurityGroup{}).
		WithOptions(options).
		Complete(&reconciler)

}
