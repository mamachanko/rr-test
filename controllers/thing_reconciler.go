package controllers

import (
	"context"
	"github.com/mamachanko/rr-test/api/v1alpha1"
	contourv1 "github.com/projectcontour/contour/apis/projectcontour/v1"
	"github.com/vmware-labs/reconciler-runtime/reconcilers"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ApplyMyHTTPProxy() reconcilers.SubReconciler {
	return &reconcilers.ChildReconciler{
		Name:          "ApplyMyHTTPProxy",
		ChildType:     &contourv1.HTTPProxy{},
		ChildListType: &contourv1.HTTPProxyList{},
		DesiredChild: func(ctx context.Context, parent *v1alpha1.Thing) (*contourv1.HTTPProxy, error) {
			virtualHost := &contourv1.VirtualHost{
				Fqdn: "test-fqdn.example.com",
			}

			if parent.Spec.TLSEnabled {
				virtualHost.TLS = &contourv1.TLS{
					SecretName: "test-tls",
				}
			}

			return &contourv1.HTTPProxy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      parent.Name,
					Namespace: parent.Namespace,
				},
				Spec: contourv1.HTTPProxySpec{
					VirtualHost: virtualHost,
				},
			}, nil
		},
		HarmonizeImmutableFields: func(current, desired *contourv1.HTTPProxy) {
			// Noop
		},
		ReflectChildStatusOnParent: func(parent *v1alpha1.Thing, child *contourv1.HTTPProxy, err error) {
			// Noop
		},
		MergeBeforeUpdate: func(actual, desired *contourv1.HTTPProxy) {
			actual.Labels = desired.Labels
		},
		SemanticEquals: func(s1, s2 *contourv1.HTTPProxy) bool {
			return equality.Semantic.DeepEqual(s1.Spec, s2.Spec) &&
				equality.Semantic.DeepEqual(s1.Labels, s2.Labels)
		},
	}
}
