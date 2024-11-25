package handlers

import (
	"fmt"
	"goreplace/api"
	"log"
	"os"
	"path/filepath"
	"archive/zip"
	"io"
	"strings"
)

func	getCurrentPath() (string, error) {
	path, err := os.Getwd();
	if (err != nil) {log.Println("Error While getting the Current directory"); os.Exit(1)}
	return path, nil
}

func unzipFile(path string) error {
	zipFile, err := zip.OpenReader(path + "/MLX.zip")
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer zipFile.Close()
	outputDir := path + "/MLX"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}
	for _, file := range zipFile.File {
		if strings.HasPrefix(file.Name, "__MACOSX") || strings.HasPrefix(file.Name, "._") {
			continue
		}
		destPath := filepath.Join(outputDir, file.Name)
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(destPath, file.Mode()); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", destPath, err)
			}
			continue
		}
		destFile, err := os.Create(destPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", destPath, err)
		}
		defer destFile.Close()
		zipEntry, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open zip entry %s: %w", file.Name, err)
		}
		defer zipEntry.Close()
		_, err = io.Copy(destFile, zipEntry)
		if err != nil {
			return fmt.Errorf("failed to copy content to %s: %w", destPath, err)
		}
	}
	return nil
}

func	MlxHandler() (int) {
	path, _ := getCurrentPath()
	fmt.Println("Installing MLX in :" + path + "/	path...")
	err := api.InstallMlx(path)
	if (err != nil) {
		fmt.Println(err)
		return (-1)
	}
	fmt.Println("Unziping MLX.zip file...")
	err = unzipFile(path)
	if (err != nil) {
		log.Println(err); os.Exit(1)
	}
	fmt.Println("Should be Good. check " + path + " and Start building something good")
	return (0)
}
