package command

type Command interface {
	Work()
}

type CmdStirFry struct {
	Chef *Chef
}

func (cmd *CmdStirFry) Work() {
	cmd.Chef.StirFry()
}

type CmdSoupSimmering struct {
	Chef *Chef
}

func (cmd *CmdSoupSimmering) Work() {
	cmd.Chef.SoupSimmering()
}
