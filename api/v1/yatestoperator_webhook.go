/*
Copyright 2022, YasuhiroABE <yasu@yasundial.org>

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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var yatestoperatorlog = logf.Log.WithName("yatestoperator-resource")

func (r *YATestOperator) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-view-yasuhiroabe-github-io-v1-yatestoperator,mutating=true,failurePolicy=fail,sideEffects=None,groups=view.yasuhiroabe.github.io,resources=yatestoperators,verbs=create;update,versions=v1,name=myatestoperator.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &YATestOperator{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *YATestOperator) Default() {
	yatestoperatorlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-view-yasuhiroabe-github-io-v1-yatestoperator,mutating=false,failurePolicy=fail,sideEffects=None,groups=view.yasuhiroabe.github.io,resources=yatestoperators,verbs=create;update,versions=v1,name=vyatestoperator.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &YATestOperator{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *YATestOperator) ValidateCreate() error {
	yatestoperatorlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *YATestOperator) ValidateUpdate(old runtime.Object) error {
	yatestoperatorlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *YATestOperator) ValidateDelete() error {
	yatestoperatorlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
