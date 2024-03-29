/*
Copyright 2023.

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

package controller

import (
	"context"
	"github.com/retinacodeworks/mikrotik-operator/internal/routeros"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	routerosv1alpha1 "github.com/retinacodeworks/mikrotik-operator/api/v1alpha1"
)

// Ipv4AddressReconciler reconciles a Ipv4Address object
type Ipv4AddressReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Sdk    *routeros.RouterOS
}

//+kubebuilder:rbac:groups=routeros.mikrotik.com,resources=ipv4addresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=routeros.mikrotik.com,resources=ipv4addresses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=routeros.mikrotik.com,resources=ipv4addresses/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Ipv4Address object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *Ipv4AddressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Get our IP address resource
	var address routerosv1alpha1.Ipv4Address
	if err := r.Get(ctx, req.NamespacedName, &address); err != nil {
		if errors.IsNotFound(err) {
			logger.Error(err, "Ipv4Address resource did not exist")
			return ctrl.Result{}, nil
		} else {
			logger.Error(err, "failed finding Ipv4Address resource")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Ipv4AddressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&routerosv1alpha1.Ipv4Address{}).
		Complete(r)
}
