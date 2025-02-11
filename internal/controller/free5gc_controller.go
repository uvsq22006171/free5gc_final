/*
Copyright 2025.

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
// +kubebuilder:rbac:groups=free5gc.free5gc.org,resources=free5gcs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=free5gc.free5gc.org,resources=free5gcs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=free5gc.free5gc.org,resources=free5gcs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Free5GC object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile

package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	free5gv1alpha1 "free5gc-operator/api/v1alpha1"
)

// Free5GCReconciler gère la ressource Free5GC
type Free5GCReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=core.free5gc,resources=free5gcs,verbs=get;list;watch;create;update;patch;delete

func (r *Free5GCReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Récupération de la ressource Free5GC
	var free5gc free5gv1alpha1.Free5GC
	if err := r.Get(ctx, req.NamespacedName, &free5gc); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Free5GC resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get Free5GC resource")
		return ctrl.Result{}, err
	}

	// Vérification de l'état actuel et mise à jour
	free5gc.Status.Ready = true
	if err := r.Status().Update(ctx, &free5gc); err != nil {
		logger.Error(err, "Failed to update Free5GC status")
		return ctrl.Result{}, err
	}

	logger.Info("Successfully reconciled Free5GC", "name", free5gc.Name)
	return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}

// SetupWithManager initialise le contrôleur
func (r *Free5GCReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&free5gv1alpha1.Free5GC{}).
		Complete(r)
}
