//go:build windows
// +build windows

/*
Copyright 2017 The Kubernetes Authors.

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

package server

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/containerd/containerd/log"
	"github.com/containerd/containerd/plugin"
	"github.com/pkg/errors"

	"github.com/containerd/cri/pkg/constants"
	sandboxstore "github.com/containerd/cri/pkg/store/sandbox"
)

func (c *criService) doStopPodSandbox(ctx context.Context, id string, sandbox sandboxstore.Sandbox) error {
	// Teardown network for sandbox.
	if sandbox.NetNSPath != "" && sandbox.NetNS != nil {
		if teardownErr := c.teardownPod(id, sandbox.NetNSPath, sandbox.Config); teardownErr != nil {
			return errors.Wrapf(teardownErr, "failed to destroy network for sandbox %q", id)
		}
		/*TODO:It is still possible that containerd crashes after we teardown the network, but before we remove the network namespace.
		In that case, we'll not be able to remove the sandbox anymore. The chance is slim, but we should be aware of that.
		In the future, once TearDownPod is idempotent, this will be fixed.*/

		//Close the sandbox network namespace if it was created
		if err := sandbox.NetNS.Remove(); err != nil {
			return errors.Wrapf(err, "failed to remove network namespace for sandbox %q", id)
		}
	}

	log.G(ctx).Infof("TearDown network for sandbox %q successfully", id)

	// TODO: remove this when retry logic is added to Bundle.Delete in upstream containerd
	// There is a race condition if hcsshim does not exit early for the lock on the bundle path
	// (its working directory) to be release in time for contaienrd  to successfully delete the
	// bundle during
	// "github.com/containerd/containerd/containerd/runtime/v2/".shim.Delete()
	// Even if the shim exits before the delete opteration, it may take several milliseconds for
	// the OS to flush that change.
	// Delete the bundle if it still exits to ensure propper cleanup
	bp := filepath.Join(filepath.Dir(c.config.StateDir),
		fmt.Sprintf("%s.%s", plugin.RuntimePluginV2, "task"), // from "github.com/containerd/containerd/runtime/v2".TaskManager{}.ID()
		constants.K8sContainerdNamespace,
		id)
	if err := os.RemoveAll(bp); err != nil { // does not error if path is nonexistant
		log.G(ctx).WithError(err).Warning("could not remove bundle path")
	}

	return nil

}
