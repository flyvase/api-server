package controllers

import (
	"harvest/entities"
	"harvest/interfaces"
)

func StartProfiler(pi interfaces.Profiler, c entities.ProfilerConfig) error {
	if err := pi.Start(c); err != nil {
		return err
	}

	return nil
}
