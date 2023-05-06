package models

type RouteStatus struct {
	Path   string `json:"path"`
	Status string `json:"status"`
}

type RouteInfo struct {
	Name   string      `json:"name"`
	Status RouteStatus `json:"status"`
}
