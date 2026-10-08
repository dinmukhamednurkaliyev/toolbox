package configuration

import "github.com/dinmukhamednurkaliyev/toolbox/foundation"

const FileName = "toolbox.json"

const initialConfiguration = "{\n  \"actions\": {}\n}\n"

func Initialize() error {
	return foundation.CreateFile(FileName, []byte(initialConfiguration))
}
