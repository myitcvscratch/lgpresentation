import "strings"

programs: [_name=string]: {
	directory:   *runtime.workingDirectory | _
	path:        *"/path/to/\(_name)" | _
	description: strings.HasPrefix(_name)
}

programs: {
	service1: {
		description: """
			service1 is a special service
			for special things
			"""
		args: ["hello", "world"]
		directory:    "/tmp"
		ignoreErrors: true
	}

	service2: {
		description: """
			service2 is a special service
			for special things
			"""
		args: service1.args
	}

	service3: {
		description: """
			service3 is a special service
			for special things
			"""
	}
}

runtime: _
