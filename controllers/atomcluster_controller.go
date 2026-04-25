package controllers

import (
	"context"
	"fmt"

	atomv1alpha1 "github.com/mahaasur13-sys/ATOMFederationOS/atom-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ATOMClusterReconciler reconciles ATOMCluster CRD.
type ATOMClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *ATOMClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling ATOMCluster", "name", req.Name, "namespace", req.Namespace)

	var cluster atomv1alpha1.ATOMCluster
	if err := r.Get(ctx, req.NamespacedName, &cluster); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Update status conditions
	ready := len(cluster.Spec.Nodes)
	if cluster.Status.ReadyNodes != ready {
		cluster.Status.ReadyNodes = ready
		cluster.Status.Phase = "Running"
		if err := r.Status().Update(ctx, &cluster); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to update status: %w", err)
		}
	}
	logger.Info("ATOMCluster reconciled", "clusterID", cluster.Spec.ClusterID, "nodes", ready)
	return ctrl.Result{}, nil
}

func (r *ATOMClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&atomv1alpha1.ATOMCluster{}).
		Complete(r)
}
