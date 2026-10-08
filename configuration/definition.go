package configuration

type Definition struct {
	Actions map[string]Action `json:"actions"`
}

type Action struct {
	Description string            `json:"description"`
	Run         string            `json:"run"`
	Targets     map[string]Target `json:"targets"`
}

type Target struct {
	Description string `json:"description"`
	Run         string `json:"run"`
}
