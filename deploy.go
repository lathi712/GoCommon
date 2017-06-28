package main

type Deploy struct {
	Depname         string `json:"depname"`
	Depid           string `json:"depid"`
	Projname        string `json:"projname"`
	Username        string `json:"username"`
	Gitcodeurl      string `json:"gitcodeurl"`
	Namespace       string `json:"namespace"`
	Nginxpath       string `json:"nginxpath"`
	Automataenv     string `json:"automataenv"`
	Subenv          string `json:"subenv"`
	Kubedepapplabel string `json:"kubedepapplabel"`
	Kubeservname    string `json:"kubeservname"`
	Clustername     string `json:"clustername"`
	Zone            string `json:"zone"`
	Createdat       string `json:"createdat"`
}
