package psutil

import (
	"strings"

	"github.com/nproc/rpi-board-info/info"
	"github.com/shirou/gopsutil/disk"
)

var (
	ignoredMountPoints = []string{
		"/dev",
		"/run",
		"/sys",
		"/proc",
	}
)

func NewDiskInformationProvider() *DiskInformationProvider {
	return &DiskInformationProvider{}
}

type DiskInformationProvider struct {
}

func (p *DiskInformationProvider) Disks() ([]info.Disk, error) {
	stat, err := disk.DiskPartitions(false)
	if err != nil {
		return nil, err
	}
	disks := []info.Disk{}
	for _, part := range stat {
		usage, err := disk.DiskUsage(part.Mountpoint)
		if err != nil {
			return nil, err
		}
		if p.shouldIgnore(part, usage) {
			continue
		}
		disks = append(disks, info.NewDisk(
			part.Device,
			part.Mountpoint,
			part.Fstype,
			usage.Total,
			usage.Used,
			usage.Free,
		))
	}
	return disks, nil
}

func (p *DiskInformationProvider) shouldIgnore(part disk.DiskPartitionStat, usage *disk.DiskUsageStat) bool {
	for _, p := range ignoredMountPoints {
		if strings.HasPrefix(part.Mountpoint, p) {
			return true
		}
	}

	return false
}
