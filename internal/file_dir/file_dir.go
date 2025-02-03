package filedir

type FileDir struct {
	Ls
}

func (fd *FileDir) ProcessCommand(args []string) error {
	fd.Ls.processCommand(args)
	return nil
}
