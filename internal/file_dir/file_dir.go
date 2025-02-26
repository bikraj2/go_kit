package filedir

type FileDir struct {
	Ls
	Cd
	MkDir
	Pwd
}

// func (fd *FileDir) ProcessCommand(args []string) error {
// 	fd.Ls.processCommand(args)
// }
// 	return nil
