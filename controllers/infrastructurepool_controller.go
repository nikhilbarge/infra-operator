/*


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

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	cachev1 "github.com/nikhilbarge/infra-operator/api/v1"
)
var thisCrName = "Unknown"

// InfrastructurePoolReconciler reconciles a InfrastructurePool object
type InfrastructurePoolReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=cache.example.com,resources=infrastructurepools,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.example.com,resources=infrastructurepools/status,verbs=get;update;patch

func (r *InfrastructurePoolReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("infrastructurepool", req.NamespacedName)
	thisCrName = req.Name
	// your logic here
	fmt.Println("infrastructurepool fcswitch at ", time.Now())
	return ctrl.Result{}, nil
}

func (r *InfrastructurePoolReconciler) SetupWithManager(mgr ctrl.Manager) error {
	mapFn := handler.ToRequestsFunc(
		func(a handler.MapObject) []reconcile.Request {
			instance := &cachev1.InfrastructurePool{}

			depKey := types.NamespacedName{
				Name:      thisCrName,
				Namespace: a.Meta.GetNamespace(),
			}

			err := mgr.GetClient().Get(context.TODO(), depKey, instance)
			if err != nil {
				return []reconcile.Request{}
			}

			return []reconcile.Request{
				{
					NamespacedName: types.NamespacedName{
						Name:      thisCrName,
						Namespace: a.Meta.GetNamespace(),
					},
				},
			}
		},
	)
	return ctrl.NewControllerManagedBy(mgr). 
		For(&cachev1.InfrastructurePool{}).
		WithEventFilter(BasicPredicate()).
		Watches(&source.Kind{Type: &cachev1.EthSwitch{}}, &handler.EnqueueRequestsFromMapFunc{
			ToRequests: mapFn,
		}).
		Watches(&source.Kind{Type: &cachev1.FcSwitch{}}, &handler.EnqueueRequestsFromMapFunc{
			ToRequests: mapFn,
		}).
		WithOptions(controller.Options{ 
			MaxConcurrentReconciles: 1,
		}).
		Complete(r)
}

func BasicPredicate() predicate.Predicate {
	return predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			// Ignore updates to CR status in which case metadata.Generation does not change
			GenerationIDAfterFirstReconcile := int64(1)
			if e.MetaNew.GetGeneration() <= GenerationIDAfterFirstReconcile {
				return false
			}
			return e.MetaOld.GetGeneration() != e.MetaNew.GetGeneration()
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			// Evaluates to false if the object has been confirmed deleted.
			return !e.DeleteStateUnknown
		},
	}
}
