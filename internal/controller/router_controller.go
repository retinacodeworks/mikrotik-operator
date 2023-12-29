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
	"k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	routerosv1alpha1 "github.com/launchboxio/mikrotik-operator/api/v1alpha1"
)

// RouterReconciler reconciles a Router object
type RouterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=routeros.mikrotik.com,resources=routers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=routeros.mikrotik.com,resources=routers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=routeros.mikrotik.com,resources=routers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Router object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *RouterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var router routerosv1alpha1.Router
	if err := r.Get(ctx, req.NamespacedName, &router); err != nil {
		if errors.IsNotFound(err) {
			logger.Error(err, "router resource did not exist")
			return ctrl.Result{}, nil
		} else {
			logger.Error(err, "failed finding router resource")
			return ctrl.Result{}, err
		}
	}

	routerClient, err := router.GetClient(r.Client)
	if err != nil {
		logger.Error(err, "failed getting routeros client")
		return ctrl.Result{}, err
	}

	info, err := routerClient.GetVersionInfo()
	if err != nil {
		logger.Error(err, "failed getting version info")
		return ctrl.Result{}, err
	}

	router.Status.BoardName = info.BoardName
	router.Status.Version = info.Version

	if err := r.Status().Update(ctx, &router); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RouterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&routerosv1alpha1.Router{}).
		Complete(r)
}
