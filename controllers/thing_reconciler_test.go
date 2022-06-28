package controllers_test

import (
	"github.com/mamachanko/rr-test/api/v1alpha1"
	"github.com/mamachanko/rr-test/controllers"
	contourv1 "github.com/projectcontour/contour/apis/projectcontour/v1"
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

func TestApplyMyHTTPProxy(t *testing.T) {
	rts := rtesting.SubReconcilerTestSuite{
		{
			Name: "Creates HTTPProxy with TLS secret",
			Resource: &v1alpha1.Thing{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-parent",
					Namespace: "test-ns",
				},
				Spec: v1alpha1.ThingSpec{
					TLSEnabled: true,
				},
			},
			ExpectCreates: []client.Object{
				&contourv1.HTTPProxy{
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
					Spec: contourv1.HTTPProxySpec{
						VirtualHost: &contourv1.VirtualHost{
							Fqdn: "test-fqdn.example.com",
							TLS: &contourv1.TLS{
								SecretName: "test-tls",
							},
						},
					},
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
					Message: `Created HTTPProxy "test-parent"`,
				},
			},
		},
		{
			Name: "Updates HTTPProxy and removes TLS Secret",
			Resource: &v1alpha1.Thing{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-parent",
					Namespace: "test-ns",
				},
			},
			GivenObjects: []client.Object{
				&contourv1.HTTPProxy{
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
					Spec: contourv1.HTTPProxySpec{
						VirtualHost: &contourv1.VirtualHost{
							Fqdn: "test-fqdn.example.com",
							TLS: &contourv1.TLS{
								SecretName: "test-tls",
							},
						},
					},
				},
			},
			ExpectUpdates: []client.Object{
				&contourv1.HTTPProxy{
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
					Spec: contourv1.HTTPProxySpec{
						VirtualHost: &contourv1.VirtualHost{
							Fqdn: "test-fqdn.example.com",
						},
					},
				},
			},
			ExpectEvents: []rtesting.Event{
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "Thing",
						APIVersion: "things.mamachanko.com/v1alpha1",
					},
					NamespacedName: types.NamespacedName{
						Name:      "test-parent",
						Namespace: "test-ns",
					},
					Type:    "Normal",
					Reason:  "Updated",
					Message: "Updated HTTPProxy \"test-parent\"",
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
	if err := contourv1.AddToScheme(scheme); err != nil {
		t.Errorf("could not add scheme %v", scheme)
	}

	rts.Run(t, scheme, func(t *testing.T, rtc *rtesting.SubReconcilerTestCase, c reconcilers.Config) reconcilers.SubReconciler {
		return controllers.ApplyMyHTTPProxy()
	})
}

func boolP(b bool) *bool {
	return &b
}
