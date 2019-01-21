package model

type Cluster struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Addresses []string `json:"addresses"`
}