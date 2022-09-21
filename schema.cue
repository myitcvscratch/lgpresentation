package main

programs: [string]: #Program

#Program: {
	path: string
	args?: [...string]
	description:   string
	retries?:      int
	ignoreErrors?: bool
	directory?:    string
}
