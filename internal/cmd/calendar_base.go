package cmd

import "github.com/BrianV1981/aim-google/internal/googleapi"

var newCalendarService = googleapi.NewCalendar

const (
	scopeAll    = literalAll
	scopeSingle = "single"
	scopeFuture = "future"
)
