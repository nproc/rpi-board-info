package psutil

import (
	"time"

	"github.com/nproc/rpi-board-info/info"
	"github.com/shirou/gopsutil/host"
)

func NewBoardInformationProvider() *BoardInformationProvider {
	return &BoardInformationProvider{}
}

type BoardInformationProvider struct {
}

func (p *BoardInformationProvider) Board() (info.Board, error) {
	stat, err := host.HostInfo()
	if err != nil {
		return nil, err
	}
	return info.NewBoard(
		stat.Hostname,
		stat.OS,
		stat.Platform,
		stat.PlatformFamily,
		stat.PlatformVersion,
		time.Unix(int64(stat.BootTime), 0),
	), nil
}
