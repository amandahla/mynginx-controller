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

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	webappv1 "github.com/amandahla/mynginx-controller/api/v1"
)

var _ = Describe("MyNginx Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default", // TODO(user):Modify as needed
		}
		mynginx := &webappv1.MyNginx{}

		BeforeEach(func() {
			By("creating the custom resource for the Kind MyNginx")
			err := k8sClient.Get(ctx, typeNamespacedName, mynginx)
			if err != nil && errors.IsNotFound(err) {
				resource := &webappv1.MyNginx{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					Spec: webappv1.MyNginxSpec{
						Replicas: 2,
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})
		It("should successfully reconcile the resource", func() {
			By("Reconciling the created resource")
			controllerReconciler := &MyNginxReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())
		})
		It("should successfully add a finalizer if not found", func() {
			By("Reconciling the created resource")
			controllerReconciler := &MyNginxReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			updated := &webappv1.MyNginx{}
			Expect(k8sClient.Get(ctx, typeNamespacedName, updated)).To(Succeed())
			Expect(updated.Finalizers).To(ContainElement(MyNginxFinalizer))
		})
		It("should create a Deployment with the correct number of replicas", func() {
			By("Reconciling the created resource")

			controllerReconciler := &MyNginxReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			deployment := &appsv1.Deployment{}
			deploymentName := types.NamespacedName{
				Name:      resourceName,
				Namespace: "default",
			}
			Eventually(func() error {
				return k8sClient.Get(ctx, deploymentName, deployment)
			}).Should(Succeed())
			Expect(*deployment.Spec.Replicas).To(Equal(int32(2)))
		})
		It("should delete the Deployment when the resource is deleted", func() {
			By("Reconciling the deleted resource")

			mynginx := &webappv1.MyNginx{}
			Expect(k8sClient.Get(ctx, typeNamespacedName, mynginx)).To(Succeed())
			deployment := &appsv1.Deployment{}
			deploymentName := types.NamespacedName{
				Name:      resourceName,
				Namespace: "default",
			}
			Eventually(func() error {
				return k8sClient.Get(ctx, deploymentName, deployment)
			}).Should(Succeed())

			Expect(k8sClient.Delete(ctx, mynginx)).To(Succeed())

			controllerReconciler := &MyNginxReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}
			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				err := k8sClient.Get(ctx, deploymentName, deployment)
				return errors.IsNotFound(err)
			}).Should(BeTrue())
		})
	})
})
