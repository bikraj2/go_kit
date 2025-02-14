package filedir

import (
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"go_kit.com/internal/color"
)

type Ls struct {
	CurrDir string
	LsOptions
}
type LsOptions struct {
	MoreInfo        bool
	ShowHiddenFiles bool
	SortBy          string
}

var lsOptions = []string{"l", "a", "s", "t", "n"}

func (l *Ls) ProcessCommand(args []string) error {
	defer l.resetFlags()
	err := l.processFlags(args)
	if err != nil {
		return err
	}
	dirs, err := list_file(l.CurrDir)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		file_info, err := dir.Info()
		if err != nil {
			panic(err)
		}
		if l.SortBy != "" {
			switch l.SortBy {
			case "size":
				dirs, err = SortDirEntries(dirs, "size")
			case "time":

				dirs, err = SortDirEntries(dirs, "time")
			default:
				dirs, err = SortDirEntries(dirs, "name")
			}
			if err != nil {
				panic(err)
			}
		}
		if strings.HasPrefix(file_info.Name(), ".") && !l.ShowHiddenFiles {
			continue
		}
		if l.MoreInfo {
			fmt.Printf("%v%-18s %v%-14s %v%-14v", color.Colors["gold"], file_info.Mode().String(), color.Colors["LightSeaGreen"], file_info.ModTime().Format("02 Jan 2006"), color.Colors["AntiqueWhite"], file_info.Size())
		}
		if dir.Type().IsDir() {

			// Additional Information on directory.
			// Icon selection
			if file_info.Size() == 0 {
				fmt.Print("📂")
			} else {
				fmt.Print("📁")
			}
			fmt.Printf("%v%v%v", color.Colors["cyan2"], dir.Name(), color.Colors["reset"])
		} else {
			fmt.Printf("%v", dir.Name())
		}
		fmt.Println()
	}
	return nil
}

func (l *LsOptions) processFlags(args []string) error {

	for i, arg := range args {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(arg, "-") {
			flag := strings.TrimPrefix(arg, "-")
			valid := isValidOptions(flag, lsOptions)
			if !valid {
				return fmt.Errorf("%v is not a valid flag", flag)
			}
			err := l.setOption(flag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (lOpt *LsOptions) setOption(opt string) error {
	if lOpt.flagSet() {
		return ErrFlagCollision
	}
	switch strings.ToLower(opt) {
	case "l":

		lOpt.MoreInfo = true
	case "a":
		lOpt.ShowHiddenFiles = true
	case "n":
		// fmt.Println("here")
		lOpt.SortBy = "name"
	case "s":
		lOpt.SortBy = "size"
	case "t":
		lOpt.SortBy = "t"
	default:

		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}

func (lOpt *LsOptions) flagSet() bool {
	// if lOpt.MoreInfo {
	// 	return true
	// }
	// return false
	return false
}

func (l *Ls) resetFlags() {
	l.MoreInfo = false
	l.ShowHiddenFiles = false
	l.SortBy = ""
}

// SortDirEntries sorts the entries by a given field: "name", "size", "modtime"
func SortDirEntries(entries []fs.DirEntry, field string) ([]fs.DirEntry, error) {
	sort.Slice(entries, func(i, j int) bool {
		switch field {
		case "name":
			// fmt.Println("here")
			return entries[i].Name() < entries[j].Name()

		case "size":
			infoI, errI := entries[i].Info()
			infoJ, errJ := entries[j].Info()
			if errI != nil || errJ != nil {
				return false
			}
			return infoI.Size() < infoJ.Size()

		case "modtime":
			infoI, errI := entries[i].Info()
			infoJ, errJ := entries[j].Info()
			if errI != nil || errJ != nil {
				return false
			}
			return infoI.ModTime().Before(infoJ.ModTime())

		default:
			return false
		}
	})
	return entries, nil
}

// func ()
