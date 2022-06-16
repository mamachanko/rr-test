package controllers

import (
	"context"
	"github.com/mamachanko/rr-test/api/v1alpha1"
	"github.com/vmware-labs/reconciler-runtime/reconcilers"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ApplyMySecret() reconcilers.SubReconciler {
	return &reconcilers.ChildReconciler{
		Name:          "ApplyMySecret",
		ChildType:     &v1.Secret{},
		ChildListType: &v1.SecretList{},
		DesiredChild: func(ctx context.Context, parent *v1alpha1.Thing) (*v1.Secret, error) {
			return &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      parent.Name,
					Namespace: parent.Namespace,
				},
				StringData: map[string]string{"key": "value"},
			}, nil
		},
		HarmonizeImmutableFields: func(current, desired *v1.Secret) {
			desired.Data = current.Data
		},
		ReflectChildStatusOnParent: func(parent *v1alpha1.Thing, child *v1.Secret, err error) {
			// No op.
		},
		MergeBeforeUpdate: func(actual, desired *v1.Secret) {
			actual.Labels = desired.Labels
		},
		SemanticEquals: func(s1, s2 *v1.Secret) bool {
			return equality.Semantic.DeepEqual(s1.Data, s2.Data) &&
				equality.Semantic.DeepEqual(s1.Labels, s2.Labels)
		},
	}
}
