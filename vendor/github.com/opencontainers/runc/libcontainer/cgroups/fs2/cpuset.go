// +build linux

package fs2

import (
	"github.com/opencontainers/runc/libcontainer/cgroups"
	"github.com/opencontainers/runc/libcontainer/configs"
)

func isCpusetSet(r *configs.Resources) bool {
	return r.CpusetCpus != "" || r.CpusetMems != ""
}

func setCpuset(dirPath string, r *configs.Resources) error {
	if !isCpusetSet(r) {
		return nil
	}

<<<<<<< HEAD
	if r.CpusetCpus != "" {
		if err := cgroups.WriteFile(dirPath, "cpuset.cpus", r.CpusetCpus); err != nil {
||||||| 5e58841cce7
	if cgroup.Resources.CpusetCpus != "" {
		if err := fscommon.WriteFile(dirPath, "cpuset.cpus", cgroup.Resources.CpusetCpus); err != nil {
=======
	if r.CpusetCpus != "" {
		if err := fscommon.WriteFile(dirPath, "cpuset.cpus", r.CpusetCpus); err != nil {
>>>>>>> v1.21.4
			return err
		}
	}
<<<<<<< HEAD
	if r.CpusetMems != "" {
		if err := cgroups.WriteFile(dirPath, "cpuset.mems", r.CpusetMems); err != nil {
||||||| 5e58841cce7
	if cgroup.Resources.CpusetMems != "" {
		if err := fscommon.WriteFile(dirPath, "cpuset.mems", cgroup.Resources.CpusetMems); err != nil {
=======
	if r.CpusetMems != "" {
		if err := fscommon.WriteFile(dirPath, "cpuset.mems", r.CpusetMems); err != nil {
>>>>>>> v1.21.4
			return err
		}
	}
	return nil
}
