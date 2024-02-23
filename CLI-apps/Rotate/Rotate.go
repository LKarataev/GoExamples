package Rotate

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	archive_dir := flag.String("a", "", "Archive directory")
	flag.Parse()
	logfile_paths := flag.Args()

	if len(logfile_paths) == 0 {
		log.Fatal("No log files specified")
	}

	if *archive_dir == "" {
		current_dir, err := os.Getwd()
		if err != nil {
			log.Fatal("Failed to get current working directory:", err)
		}
		archive_dir = &current_dir
	}

	err := os.MkdirAll(*archive_dir, os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create archive directory:", err)
	}

	var wg sync.WaitGroup

	logfiles_count := 0
	for _, logfile_path := range logfile_paths {
		if filepath.Ext(logfile_path) != ".log" {
			log.Printf("Erorr: not log file: ", logfile_path)
		} else {
			wg.Add(1)
			go InitRotate(logfile_path, &wg, *archive_dir, &logfiles_count)
		}
	}

	wg.Wait()
	fmt.Printf("Rotated %d log files to %s\n", logfiles_count, *archive_dir)
}

func InitRotate(logfile_path string, wg *sync.WaitGroup, archive_dir string, count *int) {
	defer wg.Done()
	err := RotateLog(logfile_path, archive_dir)
	if err != nil {
		log.Printf("Failed to rotate log file %s: %v\n", logfile_path, err)
	} else {
		*count++
	}
}

func RotateLog(logfile_path string, archive_dir string) error {
	logfile, err := os.Open(logfile_path)
	if err != nil {
		return err
	}
	defer logfile.Close()

	file_info, err := logfile.Stat()
	if err != nil {
		return err
	}
	timestamp := file_info.ModTime().Unix()

	base_name := strings.TrimSuffix(filepath.Base(logfile_path), ".log")
	archive_file_name := fmt.Sprintf("%s_%d.tar.gz", base_name, timestamp)
	archive_file_path := filepath.Join(archive_dir, archive_file_name)
	archive_file, err := os.Create(archive_file_path)
	if err != nil {
		return err
	}
	defer archive_file.Close()

	gzip_writer := gzip.NewWriter(archive_file)
	defer gzip_writer.Close()

	tar_writer := tar.NewWriter(gzip_writer)
	defer tar_writer.Close()

	logfile_info, err := os.Stat(logfile_path)
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    filepath.Base(logfile_path),
		Size:    logfile_info.Size(),
		Mode:    int64(logfile_info.Mode()),
		ModTime: logfile_info.ModTime(),
	}
	err = tar_writer.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tar_writer, logfile)
	if err != nil {
		return err
	}

	err = os.Remove(logfile_path)
	if err != nil {
		return err
	}

	return nil
}
