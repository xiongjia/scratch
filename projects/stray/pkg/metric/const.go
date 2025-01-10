package metric

import (
	"time"
)

const (
	LOG_COMPONENT_KEY = "component"

	COMPONENT_DISCOVERY = "service_discovery"
	COMPONENT_SCRAPE    = "scrape"
	COMPONENT_PMQL      = "pmql"

	QL_MAX_SAMPLES = 1000
	QL_TIMEOUT     = 20 * time.Second
)
