package main

import "github.com/glebnaz/go-platform/metrics"

var (
	timeHandleMetric = metrics.MustRegisterHistogramVec("time_handle", "rebrain", "",
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []string{})
)
