package opsfile

import (
	"io"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// CreateEmptyFile создает пустой файл
func CreateEmptyFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	return file.Close()
}

// CreateFileWithContent создает файл с указанным содержимым
func CreateFileWithContent(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// AppendToFile добавляет содержимое в конец файла
func AppendToFile(path, content string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// ClearFile очищает содержимое файла
func ClearFile(path string) error {
	return os.Truncate(path, 0)
}

// DeleteFile удаляет файл
func DeleteFile(path string) error {
	return os.Remove(path)
}

// RenameFile переименовывает файл
func RenameFile(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// MoveFile перемещает файл
func MoveFile(srcPath, dstPath string) error {
	// Создаем директорию назначения, если не существует
	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		return err
	}
	return os.Rename(srcPath, dstPath)
}

// CopyFile копирует файл
func CopyFile(srcPath, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	// Создаем директорию назначения, если не существует
	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		return err
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	// Копируем права доступа
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return err
	}
	return os.Chmod(dstPath, srcInfo.Mode())
}

// FileInfo содержит информацию о файле
type FileInfo struct {
	Path    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
	Owner   uint32
	Group   uint32
}

// ListFiles возвращает список файлов в директории (без рекурсии)
func ListFiles(dir string) ([]string, error) {
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		files = append(files, filepath.Join(dir, entry.Name()))
	}
	return files, nil
}

// ListFilesRecursive возвращает рекурсивный список всех файлов
func ListFilesRecursive(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// GetFileInfo возвращает информацию о файле
func GetFileInfo(path string) (*FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	sys, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, os.ErrInvalid
	}

	return &FileInfo{
		Path:    path,
		Size:    info.Size(),
		Mode:    info.Mode(),
		ModTime: info.ModTime(),
		IsDir:   info.IsDir(),
		Owner:   sys.Uid,
		Group:   sys.Gid,
	}, nil
}

// GetFileModTime возвращает время последней модификации файла
func GetFileModTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

// FileExists проверяет существование файла
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// SetFilePermissions устанавливает права доступа к файлу
func SetFilePermissions(path string, mode os.FileMode) error {
	return os.Chmod(path, mode)
}

// SetFileOwnerGroup устанавливает владельца и группу для файла
func SetFileOwnerGroup(path string, uid, gid int) error {
	return os.Chown(path, uid, gid)
}
