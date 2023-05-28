package model

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	CreatedBackups = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_count_of_created_backups",
		Help: "The total number of created backups",
	})
	ErrorDuringBackups = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_count_of_error_backups",
		Help: "The total number of errors during backup creation",
	})
)
