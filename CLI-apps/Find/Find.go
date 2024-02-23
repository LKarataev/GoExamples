package Find

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	ext           = flag.String("ext", "", "find only files with concrete extension")
	only_dirs     = flag.Bool("d", false, "print only directories")
	only_files    = flag.Bool("f", false, "print only regular files")
	only_symlinks = flag.Bool("sl", false, "print only symbolic links")
)

func main() {
	flag.Parse()
	CheckFlags()
	path := flag.Arg(0)

	err := filepath.Walk(path, WalkFunc)
	CheckError(err)
}

func WalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil && !os.IsPermission(err) {
		return err
	}
	if *only_files {
		if !info.Mode().IsRegular() {
			return nil
		}
		if *ext != "" && filepath.Ext(path) != "."+*ext {
			return nil
		}
		fmt.Println(path)
		return nil
	}
	if *only_dirs {
		if !info.Mode().IsDir() {
			return nil
		}
		fmt.Println(path)
		return nil
	}
	if *only_symlinks {
		if info.Mode()&os.ModeSymlink == 0 {
			return nil
		}
		link, err := filepath.EvalSymlinks(path)
		if err != nil {
			fmt.Printf("%s -> [broken]\n", path)
			return nil
		}
		fmt.Printf("%s -> %s\n", path, link)
		return nil
	}
	return nil
}

func CheckFlags() {
	if !(*only_dirs || *only_files || *only_symlinks) {
		*only_dirs = true
		*only_files = true
		*only_symlinks = true
	}
	if (*only_dirs || *only_symlinks) && *ext != "" {
		log.Fatal("You can use flag '-ext' only with flag '-f")
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
