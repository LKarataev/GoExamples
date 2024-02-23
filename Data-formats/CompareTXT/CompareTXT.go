package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	old_filename := flag.String("old", "", "txt file for compare")
	new_filename := flag.String("new", "", "txt file for compare")
	flag.Parse()

	if IsFileTxt(*old_filename) && IsFileTxt(*new_filename) {
		Compare(*old_filename, *new_filename)
	} else {
		fmt.Printf("Error; unexpected extension (not .txt)")
	}
}

func Compare(old_filename string, new_filename string) {
	old_files, err := ReadFileList(old_filename)
	CheckError(err)
	new_files, err := os.Open(new_filename)
	CheckError(err)

	old_files_seen := make([]bool, len(old_files))

	scanner := bufio.NewScanner(new_files)

	for scanner.Scan() {
		line := scanner.Text()
		found := false
		for i, str := range old_files {
			if line == str {
				old_files_seen[i] = true
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("ADDED %s\n", line)
		}
	}

	for i, seen := range old_files_seen {
		if !seen {
			fmt.Printf("REMOVED %s\n", old_files[i])
		}
	}

	new_files.Close()
	err = scanner.Err()
	CheckError(err)
}

func ReadFileList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var file_list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		file_list = append(file_list, scanner.Text())
	}
	file.Close()
	return file_list, scanner.Err()
}

func IsFileTxt(filename string) bool {
	if filename == "" {
		return false
	}
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return false
	}
	return parts[len(parts)-1] == "txt"
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
