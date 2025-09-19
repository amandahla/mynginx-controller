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
	log := logf.FromContext(ctx)

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

	log.Info("handling finalizer")
	stop, err := r.handleFinalizer(ctx, myNginx, myDeployment)
	if err != nil {
		return ctrl.Result{}, err
	}

	if stop {
		return ctrl.Result{}, nil
	}

	log.Info("ensuring deployment")
	err = r.ensureDeployment(ctx, myNginx, ns)
	if err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *MyNginxReconciler) handleFinalizer(ctx context.Context, myNginx *webappv1.MyNginx, myDeployment *appsv1.Deployment) (bool, error) {
	myFinalizerName := "webapp.mynginx.amandahla.xyz/finalizer"
	// examine DeletionTimestamp to determine if object is under deletion
	if myNginx.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then let's add the finalizer and update the object. This is equivalent
		// to registering our finalizer.
		if !controllerutil.ContainsFinalizer(myNginx, myFinalizerName) {
			controllerutil.AddFinalizer(myNginx, myFinalizerName)
			if err := r.Update(ctx, myNginx); err != nil {
				return true, err
			}
		}
	} else {
		// The object is being deleted
		if controllerutil.ContainsFinalizer(myNginx, myFinalizerName) {
			// our finalizer is present, so let's handle any external dependency
			if err := r.Delete(ctx, myDeployment); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried.
				return true, err
			}

			// remove our finalizer from the list and update it.
			controllerutil.RemoveFinalizer(myNginx, myFinalizerName)
			if err := r.Update(ctx, myNginx); err != nil {
				return true, err
			}
		}
		// Stop reconciliation as the item is being deleted
		return true, nil
	}
	return false, nil
}

func (r *MyNginxReconciler) ensureDeployment(ctx context.Context, myNginx *webappv1.MyNginx, ns string) error {
	myDeployment := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Namespace: ns, Name: myNginx.Name}, myDeployment)
	if kerrors.IsNotFound(err) {
		labels := map[string]string{
			"app": myNginx.Name,
		}
		newDeployment := &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      myNginx.Name,
				Namespace: myNginx.Namespace,
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: ptr.To(int32(myNginx.Spec.Replicas)),
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

		err := r.Create(ctx, newDeployment)
		if err != nil {
			return err
		}
	}

	if err == nil {
		// myDeployment was found and we need to check if replicas are the same
		if *myDeployment.Spec.Replicas != int32(myNginx.Spec.Replicas) {
			myDeployment.Spec.Replicas = ptr.To(int32(myNginx.Spec.Replicas))
			r.Update(ctx, myDeployment) // change to patch
		}
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MyNginxReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.MyNginx{}).
		Named("mynginx").
		Complete(r)
}
