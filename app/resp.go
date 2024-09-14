package main

type command string

func (cmd *command) ToRESP() string {
	return string(*cmd)
}

func (cmd *command) FromRESP() string {
	return string(*cmd)
}
