programs: {
	service1: {
		path: "/path/to/service1"
		description: """
			service1 is a special service
			for special things
			"""
		args: ["hello", "world"]
		directory:    "/tmp"
		ignoreErrors: true
	}
	service2: {
		path: "/path/to/service2"
		description: """
			service2 is a special service
			for special things
			"""
		args: ["hello", "world"]
		directory: "/home/cueckoo"
	}
	service3: {
		path: "/path/to/service3"
		description: """
			service3 is a special service
			for special things
			"""
		directory: "/home/cueckoo"
	}
}
