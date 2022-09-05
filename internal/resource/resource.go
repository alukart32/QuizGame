package resource

// RuleReader is an interface that declares methods for reading the rules of the game.
type RuleReader interface {
	// ReadAll reads records from a URL until an error and returns them.
	// A successful call returns err == nil. Parameter url can be a filepath
	// or website URL.
	ReadAll(url string) (records [][]string, err error)
}
