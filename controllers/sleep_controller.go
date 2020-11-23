/*
Copyright 2020 reoring.

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
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pridev1beta1 "github.com/inflion/pride/api/v1beta1"
	sleeperv1beta1 "github.com/inflion/sleeper/api/v1beta1"
)

// SleepReconciler reconciles a Sleep object
type SleepReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=inflion.inflion.com,resources=sleeps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=inflion.inflion.com,resources=sleeps/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=inflion.inflion.com,resources=prides,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=inflion.inflion.com,resources=prides/status,verbs=update;patch

func (r *SleepReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sleep", req.NamespacedName)

	var sleep sleeperv1beta1.Sleep
	if err := r.Get(ctx, req.NamespacedName, &sleep); err != nil {
		log.Error(err, "unable to get sleep object")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if _, err := ctrl.CreateOrUpdate(ctx, r.Client, &sleep, func() error {
		st := sleepTime{now: time.Now(), bedtimeAt: sleep.Spec.Bedtime, wakeupAt: sleep.Spec.Wakeup}
		if st.isSleepTime() {
			log.Info("Good night")
			sleep.Status.Sleeping = true
		} else {
			log.Info("Good morning")
			sleep.Status.Sleeping = false
		}
		return nil
	}); err != nil {
		log.Error(err, "unable to update sleep status")
		return ctrl.Result{}, err
	}

	var pride pridev1beta1.Pride
	var prideNamespacedName = client.ObjectKey{Namespace: req.Namespace, Name: sleep.Spec.PrideName}

	if err := r.Get(ctx, prideNamespacedName, &pride); err != nil {
		log.Error(err, "unable to get pride")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log.Info("pride fetched")

	if _, err := ctrl.CreateOrUpdate(ctx, r.Client, &pride, func() error {
		pride.Status.Sleeping = sleep.Status.Sleeping
		return nil
	}); err != nil {
		log.Error(err, "unable to update pride status")
		return ctrl.Result{}, err
	}
	log.Info("pride status updated")

	//if err := r.Status().Update(ctx, &pride); err != nil {
	//	log.Error(err, "unable to update pride status")
	//	return ctrl.Result{}, err
	//}

	return ctrl.Result{}, nil
}

func (r *SleepReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sleeperv1beta1.Sleep{}).
		Complete(r)
}
