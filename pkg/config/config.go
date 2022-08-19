/*
This file is part of REANA.
Copyright (C) 2022 CERN.

REANA is free software; you can redistribute it and/or modify it
under the terms of the MIT License; see LICENSE file for more details.
*/

// Package config gives constants and small functions that specify the REANA client configuration.
package config

// FilesBlacklist list of files to be ignored.
var FilesBlacklist = []string{".git/", "/.git/"}

// InteractiveSessionTypes list of supported types of interactive sessions.
var InteractiveSessionTypes = []string{"jupyter"}

// ReanaComputeBackends maps the backends' command line references to their real names.
var ReanaComputeBackends = map[string]string{
	"kubernetes": "Kubernetes",
	"htcondor":   "HTCondor",
	"slurm":      "Slurm",
}

// ReanaComputeBackendKeys valid options for compute backends, used in the command line.
// These keys are the same used in ReanaComputeBackends.
var ReanaComputeBackendKeys = []string{"kubernetes", "htcondor", "slurm"}

// LeadingMark prefix used when displaying headers or important messages.
var LeadingMark = "==>"

// GetRunStatuses provides a list of currently supported run statuses.
// Includes the deleted status if includeDeleted is set to true.
func GetRunStatuses(includeDeleted bool) []string {
	runStatuses := []string{
		"created",
		"running",
		"finished",
		"failed",
		"stopped",
		"queued",
		"pending",
	}
	if includeDeleted {
		runStatuses = append(runStatuses, "deleted")
	}
	return runStatuses
}

// DuMultiFilters available filters with multiple values in du command
var DuMultiFilters = []string{"size", "name"}

// ListMultiFilters available filters with multiple values in list command
var ListMultiFilters = []string{"name", "status"}

// LogsSingleFilters available filters with a single value in logs command
var LogsSingleFilters = []string{"compute_backend", "docker_img", "status"}

// LogsMultiFilters available filters with multiple values in logs command
var LogsMultiFilters = []string{"step"}