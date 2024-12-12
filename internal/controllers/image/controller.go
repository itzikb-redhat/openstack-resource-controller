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

package image

import (
	"context"
	"time"

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"

	ctrlexport "github.com/k-orc/openstack-resource-controller/internal/controllers/export"
	"github.com/k-orc/openstack-resource-controller/internal/scope"
)

const (
	Finalizer = "openstack.k-orc.cloud/image"

	FieldOwner = "openstack.k-orc.cloud/imagecontroller"
	// Field owner of the object finalizer.
	SSAFinalizerTxn = "finalizer"
	// Field owner of transient status.
	SSAStatusTxn = "status"
)

// ssaFieldOwner returns the field owner for a specific named SSA transaction.
func ssaFieldOwner(txn string) client.FieldOwner {
	return client.FieldOwner(FieldOwner + "/" + txn)
}

const (
	// The time to wait before reconciling again when we are expecting OpenStack to finish some task and update status.
	externalUpdatePollingPeriod = 15 * time.Second

	// Size of the upload and download buffers.
	transferBufferSizeBytes = 64 * 1024

	// The maximum number of times we will attempt to download content before failing.
	maxDownloadAttempts = 5
)

type imageReconcilerConstructor struct {
	scopeFactory scope.Factory
}

func New(scopeFactory scope.Factory) ctrlexport.Controller {
	return imageReconcilerConstructor{scopeFactory: scopeFactory}
}

func (imageReconcilerConstructor) GetName() string {
	return "image"
}

// orcImageReconciler reconciles an ORC Image.
type orcImageReconciler struct {
	client       client.Client
	recorder     record.EventRecorder
	scopeFactory scope.Factory
}

// SetupWithManager sets up the controller with the Manager.
func (c imageReconcilerConstructor) SetupWithManager(_ context.Context, mgr ctrl.Manager, options controller.Options) error {
	reconciler := orcImageReconciler{
		client:       mgr.GetClient(),
		recorder:     mgr.GetEventRecorderFor("orc-image-controller"),
		scopeFactory: c.scopeFactory,
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&orcv1alpha1.Image{}).
		WithOptions(options).
		Complete(&reconciler)
}