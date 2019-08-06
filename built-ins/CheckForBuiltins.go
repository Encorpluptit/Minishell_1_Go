package builtins

func Check(exitStatus *uint8, name string, argv []string) bool {
	// à transformer en mmap ?? (plus facile pour alias)
	switch name {
	case "exit":
		return exitCmd(exitStatus, name, argv)
	case "cd":
		return cd(exitStatus, name, argv)
	case "env":
		return env(exitStatus, name, argv)
	case "setenv":
		return setEnv(exitStatus, name, argv)
	case "unsetenv":
		return unSetEnv(exitStatus, name, argv)
	//case "echo":
	//	echo(exitStatus, name, argv)
	default:
		return false
	}
}
