/*
Copyright 2019 The Knative Authors

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

package seccompprofile

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	seccompprofilereconciler "github.com/imjasonh/seccomp-profile/pkg/apis/injection/reconciler/seccomp/v1alpha1/seccompprofile"
	v1alpha1 "github.com/imjasonh/seccomp-profile/pkg/apis/seccomp/v1alpha1"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/reconciler"
)

const path = "/profiles"

// Reconciler implements seccompprofilereconciler.Interface for
// SeccompProfile resources.
type Reconciler struct {
}

// Check that our Reconciler implements Interface
var _ seccompprofilereconciler.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, p *v1alpha1.SeccompProfile) reconciler.Event {
	logger := logging.FromContext(ctx)
	logger.Infof("reconciling %s", p.Name)

	// Validate again just to be sure.
	if err := p.Validate(ctx); err != nil {
		return err
	}

	// Write policy contents to localhost.
	fn := fmt.Sprintf("%s/%s.json", path, p.Name)
	logger.Infof("writing %s", fn)
	f, err := os.Create(fn)
	if err != nil {
		return fmt.Errorf("error creating %s: %w", fn, err)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(&p.Spec.Contents); err != nil {
		return fmt.Errorf("error writing %s: %w", fn, err)
	}
	logger.Infof("wrote %s", fn)

	if err := listFiles(ctx); err != nil {
		return fmt.Errorf("error listing files after write: %w", err)
	}

	// TODO: Detect deleted policies and delete files.
	return nil
}
