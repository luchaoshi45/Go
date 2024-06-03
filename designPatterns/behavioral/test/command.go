package test

import "Go/designPatterns/behavioral/command"

func Command() {
	chef := new(command.Chef)
	waiter := new(command.Waiter)

	cmdStirFry := command.CmdStirFry{Chef: chef}
	cmdSoupSimmering := command.CmdSoupSimmering{Chef: chef}

	waiter.CmdList = append(waiter.CmdList, &cmdStirFry)
	waiter.CmdList = append(waiter.CmdList, &cmdSoupSimmering)

	waiter.Notify()
}
