package fileops

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

var trackedFilePattern = regexp.MustCompile(`^\d{2}_`) // e.g., 00_file.txt, 01_config.json

// Sync merges src into dst. Both must be of the same type:
// - If both are directories, contents are merged (with tracked-file cleanup).
// - If both are files, dst is created if missing, or replaced if "tracked".
// Rules:
//   - Non-tracked existing files/dirs are preserved untouched.
//   - Tracked files are copied if missing, replaced if present in src, and deleted if missing in src.
func Sync(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat source: %w", err)
	}

	dstInfo, err := os.Stat(dst)
	if err != nil {
		if os.IsNotExist(err) {
			// Destination does not exist → create
			if srcInfo.IsDir() {
				if err := os.MkdirAll(dst, 0o755); err != nil {
					return fmt.Errorf("failed to create destination directory: %w", err)
				}
				return syncDir(src, dst)
			}
			if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
				return fmt.Errorf("failed to create destination parent directory: %w", err)
			}
			return copyFile(src, dst)
		}
		return fmt.Errorf("failed to stat destination: %w", err)
	}

	// Type check
	if srcInfo.IsDir() != dstInfo.IsDir() {
		return fmt.Errorf("type mismatch: source is dir=%t, dest is dir=%t", srcInfo.IsDir(), dstInfo.IsDir())
	}

	if srcInfo.IsDir() {
		return syncDir(src, dst)
	}
	return syncFile(src, dst)
}

// syncDir merges src directory into dst directory and deletes tracked files missing from src.
func syncDir(src, dst string) error {
	srcEntries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read source directory: %w", err)
	}
	dstEntries, err := os.ReadDir(dst)
	if err != nil {
		return fmt.Errorf("failed to read destination directory: %w", err)
	}

	srcSet := make(map[string]bool)
	for _, entry := range srcEntries {
		srcSet[entry.Name()] = true
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if _, err := os.Stat(dstPath); os.IsNotExist(err) {
				if err := os.MkdirAll(dstPath, 0o755); err != nil {
					return fmt.Errorf("failed to create directory: %w", err)
				}
			}
			if err := syncDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := syncFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	// Cleanup tracked files in dst missing from src
	for _, entry := range dstEntries {
		if !srcSet[entry.Name()] && trackedFilePattern.MatchString(entry.Name()) {
			dstPath := filepath.Join(dst, entry.Name())
			if err := os.RemoveAll(dstPath); err != nil {
				return fmt.Errorf("failed to delete tracked file: %w", err)
			}
		}
	}

	return nil
}

// syncFile copies or preserves files based on tracking rules.
func syncFile(src, dst string) error {
	if _, err := os.Stat(dst); err == nil {
		// dst exists, only replace if tracked
		if !trackedFilePattern.MatchString(filepath.Base(src)) {
			return nil
		}
	}
	return copyFile(src, dst)
}

// copyFile copies src file to dst, overwriting if needed.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("failed to copy file contents: %w", err)
	}

	if info, err := os.Stat(src); err == nil {
		if err := os.Chmod(dst, info.Mode()); err != nil {
			return fmt.Errorf("failed to set file permissions: %w", err)
		}
	}

	return nil
}
