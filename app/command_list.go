package main

type respCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]respCommand{
	"PING": {
		name:        "ping",
		description: "sends back a pong response",
		callback:    commandPing,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}
