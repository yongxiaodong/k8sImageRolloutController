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

package controllers

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"strings"
)

// T1Reconciler reconciles a T1 object
type T1Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;update;watch;patch;list;get
//+kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=create;update;watch;patch;list;get
//+kubebuilder:rbac:groups=apps,resources=deployments/finalizers,verbs=create;update;watch;patch;list;get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the T1 object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *T1Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	if req.Namespace == "default" {
		l.Info(fmt.Sprintf("Project %s change", req.Name))
		// TODO(user): your logic here
		deployment := &appsv1.Deployment{}
		err := r.Get(ctx, req.NamespacedName, deployment)
		if err != nil {
			if errors.IsNotFound(err) {
				return reconcile.Result{}, nil
			}
			// Error reading the object - requeue the request.
			return reconcile.Result{}, err
		}
		//fmt.Println(deployment)
		var data ImageInfo
		data.DpName = req.Name
		data.NameSpace = req.Namespace
		c := &deployment.Spec.Template.Spec.Containers
		temp := make([]string, 0)
		for _, v := range *c {
			temp = append(temp, v.Image)
		}
		data.Image = strings.Join(temp, ",")
		MysqlWrite(data)
	}

	return ctrl.Result{}, nil
}

//
// SetupWithManager sets up the controller with the Manager.
func (r *T1Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Deployment{}).
		Complete(r)
}
