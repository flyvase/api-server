package interfaces

import "harvest/entities"

type Profiler interface {
	Start(entities.ProfilerConfig) error
}
