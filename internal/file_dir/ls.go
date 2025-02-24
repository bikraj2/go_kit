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
	Ascending       bool
	Help            bool
}

var lsOptions = []string{"l", "r", "s", "t", "n", "r", "help"}

func Help() {
	fmt.Println("\nLS COMMAND - CUSTOM IMPLEMENTATION")
	fmt.Println("List directory contents with optional sorting, filtering, and detailed information.")
	fmt.Println("\nUSAGE:")
	fmt.Println("  ls [options]\n")

	fmt.Println("OPTIONS:")
	fmt.Println("  -l    Show detailed information (permissions, modification time, file size).")
	fmt.Println("  -a    Include hidden files (files starting with '.').")
	fmt.Println("  -n    Sort by name in ascending order (default).")
	fmt.Println("  -s    Sort by file size in ascending order.")
	fmt.Println("  -t    Sort by last modification time in ascending order.")
	fmt.Println("  -r    Reverse the sorting order (works with -n, -s, -t).")
	fmt.Println("\nSORTING BEHAVIOR:")
	fmt.Println("  - By default, files are displayed in ascending order by name.")
	fmt.Println("  - If multiple sorting flags (-n, -s, -t) are used, the last one takes precedence.")
	fmt.Println("  - The -r flag reverses the sorting order.")
	fmt.Println("\nEXAMPLES:")
	fmt.Println("  ls -l           # Show detailed file information.")
	fmt.Println("  ls -a           # List all files, including hidden ones.")
	fmt.Println("  ls -s           # Sort files by size in ascending order.")
	fmt.Println("  ls -t -r        # Sort files by modification time in descending order.")
	fmt.Println("  ls -a -l -t     # Show hidden files with details, sorted by modification time.")
	fmt.Println("\nERROR HANDLING:")
	fmt.Println("  - If an invalid option is used, an error message is displayed.")
	fmt.Println("  - The command does not accept additional arguments beyond options.")
	fmt.Println("\nNOTES:")
	fmt.Println("  - Directories are displayed with 📂 or 📁 icons depending on their contents.")
	fmt.Println("  - Color coding is used for better readability.")
	fmt.Println("\n")
}
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
	if l.SortBy != "" {
		dirs, err = SortDirEntries(dirs, l.SortBy, l.Ascending)
		if err != nil {
			panic(err)
		}
	}
	if l.Help {
		Help()
		return nil
	}
	for _, dir := range dirs {
		file_info, err := dir.Info()
		if err != nil {
			panic(err)
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
		} else {
			return fmt.Errorf("ls doesnot accept any argument")
		}
	}
	return nil
}
func (lOpt *LsOptions) setOption(opt string) error {
	if lOpt.flagSet() {
		return ErrFlagCollision
	}
	switch strings.ToLower(opt) {
	case "help":
		lOpt.Help = true
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
		lOpt.SortBy = "time"
	case "r":
		lOpt.Ascending = false
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
	return lOpt.Help
}

func (l *Ls) resetFlags() {
	l.MoreInfo = false
	l.ShowHiddenFiles = false
	l.SortBy = ""
	l.Ascending = true
	l.Help = false
}

// SortDirEntries sorts the entries by a given field: "name", "size", "modtime"
func SortDirEntries(entries []fs.DirEntry, field string, asc bool) ([]fs.DirEntry, error) {
	sort.Slice(entries, func(i, j int) bool {
		switch field {
		case "name":
			return asc && (entries[i].Name() > entries[j].Name())
		case "size":
			infoI, errI := entries[i].Info()
			infoJ, errJ := entries[j].Info()
			if errI != nil || errJ != nil {
				return false
			}
			return asc && (infoI.Size() > infoJ.Size())
		case "time":
			infoI, errI := entries[i].Info()
			infoJ, errJ := entries[j].Info()
			if errI != nil || errJ != nil {
				return false
			}
			return asc && (infoI.ModTime().After(infoJ.ModTime()))
		default:
			return false
		}
	})
	return entries, nil
}

// func ()
