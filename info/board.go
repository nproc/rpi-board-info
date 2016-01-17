package info

import "time"

type Board interface {
	Hostname() string
	OS() string
	Platform() string
	Family() string
	Version() string
	BootTime() time.Time
}

type BoardInformationProvider interface {
	Board() (Board, error)
}

func NewBoard(hostname, os, platform, family, version string, bootTime time.Time) Board {
	return &board{
		hostname: hostname,
		os:       os,
		platform: platform,
		family:   family,
		version:  version,
		bootTime: bootTime,
	}
}

type board struct {
	hostname string
	os       string
	platform string
	family   string
	version  string
	bootTime time.Time
}

func (b *board) Hostname() string {
	return b.hostname
}

func (b *board) OS() string {
	return b.os
}

func (b *board) Platform() string {
	return b.platform
}

func (b *board) Family() string {
	return b.family
}

func (b *board) Version() string {
	return b.version
}

func (b *board) BootTime() time.Time {
	return b.bootTime
}
