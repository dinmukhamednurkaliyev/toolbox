package command

import "io"

type registration struct {
	definition Definition
	execute    func([]string, io.Writer) error
}

func registeredCommands(program Program) []registration {
	return []registration{
		{
			definition: Initialize,
			execute:    runInitialize,
		},
		{
			definition: Help,
			execute: func(arguments []string, output io.Writer) error {
				return runHelp(program, arguments, output)
			},
		},
	}
}
