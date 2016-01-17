package info

type Disk interface {
	Device() string
	MountPoint() string
	Type() string
	Total() uint64
	Used() uint64
	Available() uint64
}

type DiskInformationProvider interface {
	Disks() ([]Disk, error)
}

func NewDisk(device, mountPoint, fsType string, total, used, available uint64) Disk {
	return &disk{
		device:     device,
		mountPoint: mountPoint,
		fsType:     fsType,
		total:      total,
		used:       used,
		available:  available,
	}
}

type disk struct {
	device     string
	mountPoint string
	fsType     string
	total      uint64
	used       uint64
	available  uint64
}

func (d *disk) Device() string {
	return d.device
}

func (d *disk) MountPoint() string {
	return d.mountPoint
}

func (d *disk) Type() string {
	return d.fsType
}

func (d *disk) Total() uint64 {
	return d.total
}

func (d *disk) Used() uint64 {
	return d.used
}

func (d *disk) Available() uint64 {
	return d.available
}
