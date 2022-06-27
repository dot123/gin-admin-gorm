package service

import (
	"GameAdmin/pkg/helper"
	"GameAdmin/pkg/logger"
	"context"
	"github.com/google/wire"
)

var SystemSet = wire.NewSet(wire.Struct(new(SystemSrv), "*"))

type SystemSrv struct{}

func (a *SystemSrv) GetServerInfo(ctx context.Context) (*helper.Server, error) {
	var err error
	var s helper.Server

	s.Os = helper.GetOSInfo()
	if s.Cpu, err = helper.GetCpuInfo(); err != nil {
		logger.WithContext(ctx).Errorf("GetCpuInfo error: %v", err)
		return nil, err
	}

	if s.Rrm, err = helper.GetMemInfo(); err != nil {
		logger.WithContext(ctx).Errorf("GetMemInfo error: %v", err)
		return nil, err
	}

	if s.Disk, err = helper.GetDiskInfo(); err != nil {
		logger.WithContext(ctx).Errorf("GetDiskInfo error: %v", err)
		return nil, err
	}

	return &s, nil
}
