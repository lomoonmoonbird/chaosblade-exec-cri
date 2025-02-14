/*
 * Copyright 1999-2019 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package exec

import (
	"github.com/lomoonmoonbird/chaosblade-exec-cri/exec/container"
	"github.com/lomoonmoonbird/chaosblade-exec-cri/exec/container/containerd"
	"github.com/lomoonmoonbird/chaosblade-exec-cri/exec/container/docker"
	"github.com/lomoonmoonbird/chaosblade-spec-go/spec"
)

func GetClientByRuntime(expModel *spec.ExpModel) (container.Container, error) {
	switch expModel.ActionFlags[ContainerRuntime.Name] {
	case container.ContainerdRuntime:
		return containerd.NewClient(expModel.ActionFlags[EndpointFlag.Name], expModel.ActionFlags[ContainerNamespace.Name])
	default:
		return docker.NewClient(expModel.ActionFlags[EndpointFlag.Name])
		//default:
		//	return nil,errors.New(fmt.Sprintf("`%s`, the container runtime not support", expModel.ActionFlags[ContainerRuntime.Name]))
	}
}