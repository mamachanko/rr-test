package controllers_test

import (
	"github.com/mamachanko/rr-test/api/v1alpha1"
	"github.com/mamachanko/rr-test/controllers"
	"github.com/vmware-labs/reconciler-runtime/reconcilers"
	rtesting "github.com/vmware-labs/reconciler-runtime/testing"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
)

func TestApplyMySecret(t *testing.T) {
	rts := rtesting.SubReconcilerTestSuite{
		{
			Name: "Creates my secret",
			Resource: &v1alpha1.Thing{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-parent",
					Namespace: "test-ns",
				},
			},
			ExpectCreates: []client.Object{
				&v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-parent",
						Namespace: "test-ns",
						OwnerReferences: []metav1.OwnerReference{
							{
								APIVersion:         "things.mamachanko.com/v1alpha1",
								Kind:               "Thing",
								Name:               "test-parent",
								Controller:         boolP(true),
								BlockOwnerDeletion: boolP(true),
							},
						},
					},
					StringData: map[string]string{"key": "value"},
				},
			},
			ExpectEvents: []rtesting.Event{
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "Thing",
						APIVersion: "things.mamachanko.com/v1alpha1",
					},
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-parent",
					},
					Type:    v1.EventTypeNormal,
					Reason:  "Created",
					Message: `Created Secret "test-parent"`,
				},
			},
		},
		{
			Name: "Does not update my secret",
			Resource: &v1alpha1.Thing{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-parent",
					Namespace: "test-ns",
				},
			},
			GivenObjects: []client.Object{
				&v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-parent",
						Namespace: "test-ns",
						OwnerReferences: []metav1.OwnerReference{
							{
								APIVersion:         "things.mamachanko.com/v1alpha1",
								Kind:               "Thing",
								Name:               "test-parent",
								Controller:         boolP(true),
								BlockOwnerDeletion: boolP(true),
							},
						},
					},
					StringData: map[string]string{"key": "value"},
				},
			},
		},
	}

	scheme := runtime.NewScheme()
	if err := clientgoscheme.AddToScheme(scheme); err != nil {
		t.Errorf("could not add scheme %v", scheme)
	}
	if err := v1alpha1.AddToScheme(scheme); err != nil {
		t.Errorf("could not add scheme %v", scheme)
	}

	rts.Run(t, scheme, func(t *testing.T, rtc *rtesting.SubReconcilerTestCase, c reconcilers.Config) reconcilers.SubReconciler {
		return controllers.ApplyMySecret()
	})
}

func boolP(b bool) *bool {
	return &b
}
