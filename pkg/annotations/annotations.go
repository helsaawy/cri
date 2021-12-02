/*
Copyright 2018 The Containerd Authors.

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

package annotations

// ContainerType values
// Following OCI annotations are used by katacontainers now.
// We'll switch to standard secure pod API after it is defined in CRI.
const (
	// ContainerTypeSandbox represents a pod sandbox container
	ContainerTypeSandbox = "sandbox"

	// ContainerTypeContainer represents a container running within a pod
	ContainerTypeContainer = "container"

	// ContainerType is the container type (sandbox or container) annotation
	ContainerType = "io.kubernetes.cri.container-type"

	// SandboxID is the sandbox ID annotation
	SandboxID = "io.kubernetes.cri.sandbox-id"

	// UntrustedWorkload is the sandbox annotation for untrusted workload. Untrusted
	// workload can only run on dedicated runtime for untrusted workload.
	UntrustedWorkload = "io.kubernetes.cri.untrusted-workload"

	// CopyExistingScratch uses the scratch space of another specified, pre-existing container as
	// its own scratch space. Other container must be stopped.
	CopyExistingScratch = "containerd.io/snapshot/io.microsoft.container.storage.copy-existing-scratch"
	// CopyScratchVhd specified the absolute path of a vhd file to copy as the source scratch space
	CopyScratchVhd = "containerd.io/snapshot/io.microsoft.container.storage.copy-scratch-vhd"

	// ReuseScratch enables containers to reuse the same scratch VHD as their pod
	ReuseScratch = "containerd.io/snapshot/io.microsoft.container.storage.reuse-scratch"

	// ScratchVhdSize specifies the size (in GB) for a containers scratch size, must be greater
	// than 0. Windows process and hypervisor isolated containers must specify a value greater
	// than the default scratch size of 20 (GB).
	ScratchVhdSize = "containerd.io/snapshot/io.microsoft.container.storage.rootfs.size-gb"
)
