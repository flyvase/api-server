package repositories

import (
	"cloud.google.com/go/profiler"

	"harvest/entities"
)

func toCloudProfilerConfig(c entities.ProfilerConfig) profiler.Config {
	return profiler.Config{
		NoCPUProfiling: c.NoCPUProfiling,
	}
}

type CloudProfiler struct{}

func (cp CloudProfiler) Start(c entities.ProfilerConfig) error {
	if err := profiler.Start(toCloudProfilerConfig(c)); err != nil {
		return err
	}

	return nil
}
