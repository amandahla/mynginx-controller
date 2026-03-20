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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	webappv1 "github.com/amandahla/mynginx-controller/api/v1"
)

var _ = Describe("MyNginx Controller", func() {
	Context("When reconciling a resource", func() {
		var namespace *corev1.Namespace
		const resourceName = "test-resource"
		ctx := context.Background()
		typeNamespacedName := types.NamespacedName{}

		BeforeEach(func() {
			By("creating the custom resource for the Kind MyNginx")
			namespace = &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{GenerateName: "test-"}}
			Expect(k8sClient.Create(context.Background(), namespace)).To(Succeed())
			typeNamespacedName = types.NamespacedName{Name: resourceName, Namespace: namespace.Name}

			resource := &webappv1.MyNginx{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceName,
					Namespace: namespace.Name,
				},
				Spec: webappv1.MyNginxSpec{
					Replicas: 2,
				},
			}
			Expect(k8sClient.Create(ctx, resource)).To(Succeed())

		})
		AfterEach(func() {
			resource := &webappv1.MyNginx{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			if err == nil {
				resource.Finalizers = nil
				_ = k8sClient.Update(ctx, resource)
				_ = k8sClient.Delete(ctx, resource)
			}
			Expect(k8sClient.Delete(ctx, namespace)).To(Succeed())

		})
		It("should successfully add a finalizer if not found", func() {
			By("Reconciling the created resource")
			updated := &webappv1.MyNginx{}
			Eventually(func(g Gomega) {
				g.Expect(k8sClient.Get(ctx, typeNamespacedName, updated)).To(Succeed())
				g.Expect(updated.Finalizers).To(ContainElement(MyNginxFinalizer))
			}, "10s", "1s").Should(Succeed())
		})
		It("should create a Deployment with the correct number of replicas and condition is ready", func() {
			By("Reconciling the created resource")
			deployment := &appsv1.Deployment{}
			deploymentName := types.NamespacedName{
				Name:      resourceName,
				Namespace: namespace.Name,
			}
			Eventually(func() error {
				return k8sClient.Get(ctx, deploymentName, deployment)
			}).Should(Succeed())
			Expect(*deployment.Spec.Replicas).To(Equal(int32(2)))
			resource := &webappv1.MyNginx{}
			Eventually(func() error {
				return k8sClient.Get(ctx, typeNamespacedName, resource)
			}).Should(Succeed())
			var conditions []metav1.Condition
			Expect(resource.Status.Conditions).To(ContainElement(
				HaveField("Type", Equal("Ready")), &conditions))
			Expect(conditions[0].Status).To(Equal(metav1.ConditionTrue), "condition %s", "Ready")
			Expect(conditions[0].Reason).To(Equal(MyNginxReadyReason), "condition %s", "Ready")
		})
		It("should delete the Deployment when the resource is deleted", func() {
			By("Reconciling the deleted resource")

			mynginx := &webappv1.MyNginx{}
			Expect(k8sClient.Get(ctx, typeNamespacedName, mynginx)).To(Succeed())
			deployment := &appsv1.Deployment{}
			deploymentName := types.NamespacedName{
				Name:      resourceName,
				Namespace: namespace.Name,
			}
			Eventually(func() error {
				return k8sClient.Get(ctx, deploymentName, deployment)
			}).Should(Succeed())

			Expect(k8sClient.Delete(ctx, mynginx)).To(Succeed())

			Eventually(func() bool {
				err := k8sClient.Get(ctx, deploymentName, deployment)
				return errors.IsNotFound(err)
			}).Should(BeTrue())
		})
	})
})
