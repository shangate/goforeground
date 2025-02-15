package goforeground

// Activate attempts to set a window foreground by its process id
// Linux and Windows are implemented; rest is unimplemented
// and will return nil
func Activate(pid int) error {
	return activate(pid)
}

func IsForegroundWindow(pid int) (yes bool, err error) {
	_, e := findWindow(pid)
	if e != nil {
		return false, e
	}
	return true, nil
}
