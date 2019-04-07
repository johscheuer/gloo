// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionClusterIngressFunc func(original, desired *ClusterIngress) (bool, error)

type ClusterIngressReconciler interface {
	Reconcile(namespace string, desiredResources ClusterIngressList, transition TransitionClusterIngressFunc, opts clients.ListOpts) error
}

func clusterIngresssToResources(list ClusterIngressList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, clusterIngress := range list {
		resourceList = append(resourceList, clusterIngress)
	}
	return resourceList
}

func NewClusterIngressReconciler(client ClusterIngressClient) ClusterIngressReconciler {
	return &clusterIngressReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type clusterIngressReconciler struct {
	base reconcile.Reconciler
}

func (r *clusterIngressReconciler) Reconcile(namespace string, desiredResources ClusterIngressList, transition TransitionClusterIngressFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "clusterIngress_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*ClusterIngress), desired.(*ClusterIngress))
		}
	}
	return r.base.Reconcile(namespace, clusterIngresssToResources(desiredResources), transitionResources, opts)
}
