package resource

// RuleReader is an interface that declares methods for reading the rules of the game.
type RuleReader interface {
	// ReadAll reads records from a filepath until an error and returns them.
	// A successful call returns err == nil.
	ReadAll(filepath string) (records [][]string, err error)
}
