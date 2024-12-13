package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// RelativizePath checks if the targetPath is absolute.
// If it is, it constructs the relative path from basePath to targetPath.
func RelativizePath(basePath, targetPath string) (string, error) {
	// Ensure both basePath and targetPath are absolute paths
	if !filepath.IsAbs(basePath) {
		return "", fmt.Errorf("basePath must be an absolute path")
	}
	if !filepath.IsAbs(targetPath) {
		return targetPath, nil
	}

	// Construct the relative path
	relativePath, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return "", fmt.Errorf("failed to construct relative path: %w", err)
	}

	return relativePath, nil
}
