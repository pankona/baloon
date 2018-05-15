package scene

type commandtype int

const (
	commandSpawn commandtype = iota
	commandSpawned
)

type command struct {
	ct      commandtype
	payload interface{}
}

func newCommand(ct commandtype, payload interface{}) *command {
	return &command{ct, payload}
}
