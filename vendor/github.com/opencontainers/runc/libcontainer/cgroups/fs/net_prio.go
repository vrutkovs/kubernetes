// +build linux

package fs

import (
	"github.com/opencontainers/runc/libcontainer/cgroups"
	"github.com/opencontainers/runc/libcontainer/configs"
)

type NetPrioGroup struct{}

func (s *NetPrioGroup) Name() string {
	return "net_prio"
}

func (s *NetPrioGroup) Apply(path string, d *cgroupData) error {
	return join(path, d.pid)
}

<<<<<<< HEAD
func (s *NetPrioGroup) Set(path string, r *configs.Resources) error {
	for _, prioMap := range r.NetPrioIfpriomap {
		if err := cgroups.WriteFile(path, "net_prio.ifpriomap", prioMap.CgroupString()); err != nil {
||||||| 5e58841cce7
func (s *NetPrioGroup) Set(path string, cgroup *configs.Cgroup) error {
	for _, prioMap := range cgroup.Resources.NetPrioIfpriomap {
		if err := fscommon.WriteFile(path, "net_prio.ifpriomap", prioMap.CgroupString()); err != nil {
=======
func (s *NetPrioGroup) Set(path string, r *configs.Resources) error {
	for _, prioMap := range r.NetPrioIfpriomap {
		if err := fscommon.WriteFile(path, "net_prio.ifpriomap", prioMap.CgroupString()); err != nil {
>>>>>>> v1.21.4
			return err
		}
	}

	return nil
}

func (s *NetPrioGroup) GetStats(path string, stats *cgroups.Stats) error {
	return nil
}
