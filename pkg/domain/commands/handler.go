package commands

type CommandHandler interface {
	HandleCommand(command Command) error
}