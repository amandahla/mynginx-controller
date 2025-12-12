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

package controller

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	webappv1 "github.com/amandahla/mynginx-controller/api/v1"
	"k8s.io/utils/ptr"
)

// MyNginxFinalizer defines finalizer name
const MyNginxFinalizer = "webapp.mynginx.amandahla.xyz/finalizer"

// MyNginxReconciler reconciles a MyNginx object
type MyNginxReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=webapp.mynginx.amandahla.xyz,resources=mynginxes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=webapp.mynginx.amandahla.xyz,resources=mynginxes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=webapp.mynginx.amandahla.xyz,resources=mynginxes/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
func (r *MyNginxReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx).WithValues("name", req.Name, "namespace", req.Namespace)

	myNginx := &webappv1.MyNginx{}
	if err := r.Get(ctx, req.NamespacedName, myNginx); err != nil {
		log.Error(err, "unable to fetch myNginx")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	ns := req.NamespacedName.Namespace
	if ns == "" {
		ns = "default"
	}
	myDeployment := &appsv1.Deployment{}
	_ = r.Get(ctx, types.NamespacedName{Namespace: ns, Name: myNginx.Name}, myDeployment)

	log.Info("reconcile finalizer")
	if changed := controllerutil.AddFinalizer(myNginx, MyNginxFinalizer); changed {
		err := r.Update(ctx, myNginx) // TODO replace by patch
		return ctrl.Result{}, err
	}

	if !myNginx.ObjectMeta.DeletionTimestamp.IsZero() {
		log.Info("reconcile delete")
		return r.reconcileDelete(ctx, myNginx, myDeployment)
	}

	log.Info("reconcile deployment")
	return r.reconcileDeployment(ctx, myNginx, ns)
}

func (r *MyNginxReconciler) reconcileDelete(ctx context.Context, myNginx *webappv1.MyNginx, myDeployment *appsv1.Deployment) (ctrl.Result, error) {
	if myDeployment.Name == "" { // TODO watch deployments so if deleted, is recreated and we dont need to check this.
		controllerutil.RemoveFinalizer(myNginx, MyNginxFinalizer)
		err := r.Update(ctx, myNginx)
		return ctrl.Result{}, err
	}
	err := r.Delete(ctx, myDeployment)
	if err != nil && !kerrors.IsNotFound(err) {
		// if fail to delete the external dependency here, return with error
		// so that it can be retried.
		return ctrl.Result{}, err
	}

	// remove our finalizer from the list and update it.
	controllerutil.RemoveFinalizer(myNginx, MyNginxFinalizer)
	err = r.Update(ctx, myNginx)
	return ctrl.Result{}, err
}

func (r *MyNginxReconciler) reconcileDeployment(ctx context.Context, myNginx *webappv1.MyNginx, ns string) (ctrl.Result, error) {
	myDeployment := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Namespace: ns, Name: myNginx.Name}, myDeployment)
	if kerrors.IsNotFound(err) {
		return ctrl.Result{}, r.createDeployment(ctx, myNginx)
	}

	if err != nil || *myDeployment.Spec.Replicas == myNginx.Spec.Replicas {
		return ctrl.Result{}, err
	}

	patch := client.MergeFrom(myDeployment.DeepCopy())
	myDeployment.Spec.Replicas = ptr.To(myNginx.Spec.Replicas)
	return ctrl.Result{}, r.Patch(ctx, myDeployment, patch)
}

func (r *MyNginxReconciler) createDeployment(ctx context.Context, myNginx *webappv1.MyNginx) error {
	labels := map[string]string{
		"app": myNginx.Name,
	}
	newDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      myNginx.Name,
			Namespace: myNginx.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.To(myNginx.Spec.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx",
						},
					},
				},
			},
		},
	}

	if err := ctrl.SetControllerReference(myNginx, newDeployment, r.Scheme); err != nil {
		return err
	}

	return r.Create(ctx, newDeployment)
}

// SetupWithManager sets up the controller with the Manager.
func (r *MyNginxReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.MyNginx{}).
		Named("mynginx").
		Complete(r)
}
