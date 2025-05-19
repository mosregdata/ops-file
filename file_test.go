package opsfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileUtils(t *testing.T) {
	// Временная директория для тестов
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	testFile2 := filepath.Join(tmpDir, "test2.txt")

	t.Run("CreateEmptyFile", func(t *testing.T) {
		if err := CreateEmptyFile(testFile); err != nil {
			t.Errorf("CreateEmptyFile failed: %v", err)
		}
		if _, err := os.Stat(testFile); os.IsNotExist(err) {
			t.Error("File was not created")
		}
	})

	t.Run("CreateFileWithContent", func(t *testing.T) {
		content := "test content"
		if err := CreateFileWithContent(testFile2, content); err != nil {
			t.Errorf("CreateFileWithContent failed: %v", err)
		}
		data, err := os.ReadFile(testFile2)
		if err != nil || string(data) != content {
			t.Errorf("File content mismatch: got %q, want %q", string(data), content)
		}
	})

	t.Run("AppendToFile", func(t *testing.T) {
		appendContent := "\nappended"
		if err := AppendToFile(testFile2, appendContent); err != nil {
			t.Errorf("AppendToFile failed: %v", err)
		}
		data, err := os.ReadFile(testFile2)
		if err != nil || !contains(string(data), appendContent) {
			t.Errorf("Append failed: got %q, want containing %q", string(data), appendContent)
		}
	})

	t.Run("ClearFile", func(t *testing.T) {
		if err := ClearFile(testFile2); err != nil {
			t.Errorf("ClearFile failed: %v", err)
		}
		data, err := os.ReadFile(testFile2)
		if err != nil || len(data) != 0 {
			t.Errorf("File not cleared: got %d bytes", len(data))
		}
	})

	t.Run("RenameFile", func(t *testing.T) {
		newPath := filepath.Join(tmpDir, "renamed.txt")
		if err := RenameFile(testFile, newPath); err != nil {
			t.Errorf("RenameFile failed: %v", err)
		}
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			t.Error("File was not renamed")
		}
	})

	t.Run("MoveFile", func(t *testing.T) {
		newDir := filepath.Join(tmpDir, "newdir")
		movedPath := filepath.Join(newDir, "moved.txt")
		if err := MoveFile(testFile2, movedPath); err != nil {
			t.Errorf("MoveFile failed: %v", err)
		}
		if _, err := os.Stat(movedPath); os.IsNotExist(err) {
			t.Error("File was not moved")
		}
	})

	t.Run("CopyFile", func(t *testing.T) {
		copyPath := filepath.Join(tmpDir, "copy.txt")
		if err := CreateFileWithContent(testFile, "copy test"); err != nil {
			t.Fatal(err)
		}
		if err := CopyFile(testFile, copyPath); err != nil {
			t.Errorf("CopyFile failed: %v", err)
		}
		data, err := os.ReadFile(copyPath)
		if err != nil || string(data) != "copy test" {
			t.Errorf("Copied file content mismatch: got %q", string(data))
		}
	})

	t.Run("DeleteFile", func(t *testing.T) {
		if err := DeleteFile(testFile); err != nil {
			t.Errorf("DeleteFile failed: %v", err)
		}
		if _, err := os.Stat(testFile); !os.IsNotExist(err) {
			t.Error("File was not deleted")
		}
	})
}

// Вспомогательная функция для проверки содержимого
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[len(s)-len(substr):] == substr
}

func TestListFiles(t *testing.T) {
	tmpDir := t.TempDir()
	createTestFile(t, tmpDir, "file1.txt")
	createTestFile(t, tmpDir, "file2.txt")

	files, err := ListFiles(tmpDir)
	if err != nil {
		t.Fatalf("ListFiles failed: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
}

func TestListFilesRecursive(t *testing.T) {
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}
	createTestFile(t, tmpDir, "file1.txt")
	createTestFile(t, subDir, "file2.txt")

	files, err := ListFilesRecursive(tmpDir)
	if err != nil {
		t.Fatalf("ListFilesRecursive failed: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
}

func TestGetFileInfo(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := createTestFile(t, tmpDir, "test.txt")

	info, err := GetFileInfo(filePath)
	if err != nil {
		t.Fatalf("GetFileInfo failed: %v", err)
	}
	if info.Path != filePath {
		t.Errorf("Expected path %s, got %s", filePath, info.Path)
	}
	if info.Size != 0 {
		t.Errorf("Expected size 0, got %d", info.Size)
	}
}

func TestGetFileModTime(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := createTestFile(t, tmpDir, "test.txt")

	modTime, err := GetFileModTime(filePath)
	if err != nil {
		t.Fatalf("GetFileModTime failed: %v", err)
	}
	if modTime.IsZero() {
		t.Error("Expected non-zero mod time")
	}
}

func TestFileExists(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := createTestFile(t, tmpDir, "test.txt")

	exists, err := FileExists(filePath)
	if err != nil {
		t.Fatalf("FileExists failed: %v", err)
	}
	if !exists {
		t.Error("Expected file to exist")
	}

	exists, err = FileExists(filepath.Join(tmpDir, "nonexistent.txt"))
	if err != nil {
		t.Fatalf("FileExists failed: %v", err)
	}
	if exists {
		t.Error("Expected file to not exist")
	}
}

func TestSetFilePermissions(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := createTestFile(t, tmpDir, "test.txt")

	err := SetFilePermissions(filePath, 0600)
	if err != nil {
		t.Fatalf("SetFilePermissions failed: %v", err)
	}

	info, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf("Stat failed: %v", err)
	}
	if info.Mode().Perm() != 0600 {
		t.Errorf("Expected mode 0600, got %o", info.Mode().Perm())
	}
}

// Вспомогательная функция для создания тестового файла
func createTestFile(t *testing.T, dir, name string) string {
	path := filepath.Join(dir, name)
	err := os.WriteFile(path, []byte{}, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	return path
}
