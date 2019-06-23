package commands

type CreateUser struct {
	BaseCommand
	Username string
	Email    string
	Password string
}
