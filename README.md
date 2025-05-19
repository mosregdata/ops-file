# OpsFile

Модуль предназначен для работы с файлами в Unix-системах. 
Он предоставляет базовые функции работы с файлами: создание пустого файла, создание файла с содержимым, дозапись, удаление, очистка, переименование, копирование, перемещение, получение списка файлов, проверка наличия, смена прав.

## Возможности
- Создание пустого файла (`CreateEmptyFile`).
- Создание файла с указанным содержимым (`CreateFileWithContent`).
- Добавление содержимого в конец файла (`AppendToFile`).
- Очистка содержимого файла (`ClearFile`).
- Удаление файла (`DeleteFile`).
- Переименование файла (`RenameFile`).
- Перемещение файла (`MoveFile`).
- Копирование файла (`CopyFile`).
- Получение списка файлов в директории (без рекурсии) (`ListFiles`).
- Получение рекурсивного списка всех файлов (`ListFilesRecursive`).
- Получение времени последней модификации файла (`GetFileModTime`).
- Проверка существования файла (`FileExists`).
- Установка прав доступа к файлу (`SetFilePermissions`).
- Установка владельца и группы для файла (`SetFileOwnerGroup`).

## Требования
- Go 1.20 или выше.
- Unix-подобная ОС (Linux, macOS и т.д.).

## Установка
Склонируйте репозиторий или добавьте модуль в ваш проект:
```shell
go get github.com/mosregdata/ops-file
```

## Использование
Пример использования модуля:
```go
package main

import (
    "fmt"
    opsfile "github.com/mosregdata/ops-file"
)

func main() {
    // Создать пустой файл
    err := opsfile.CreateEmptyFile("/tmp/test.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Создать файл с содержимым
    err = opsfile.CreateFileWithContent("/tmp/test2.txt", "Hello, World!")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Добавить содержимое
    err = opsfile.AppendToFile("/tmp/test2.txt", "\nNew line")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Очистить файл
    err = opsfile.ClearFile("/tmp/test2.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Переименовать файл
    err = opsfile.RenameFile("/tmp/test.txt", "/tmp/test_renamed.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Переместить файл
    err = opsfile.MoveFile("/tmp/test_renamed.txt", "/tmp/newdir/test_moved.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Скопировать файл
    err = opsfile.CopyFile("/tmp/test2.txt", "/tmp/test_copy.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Удалить файл
    err = opsfile.DeleteFile("/tmp/test2.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }

	// Получить список файлов
	files, err := opsfile.ListFiles("/tmp")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Files:", files)

	// Рекурсивный список
	recFiles, err := opsfile.ListFilesRecursive("/tmp")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Recursive files:", recFiles)

	// Информация о файле
	info, err := opsfile.GetFileInfo("/tmp/test_copy.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("File info: %+v\n", info)

	// Время модификации
	modTime, err := opsfile.GetFileModTime("/tmp/test_copy.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Mod time:", modTime)

	// Проверка существования
	exists, err := opsfile.FileExists("/tmp/test_copy.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File exists:", exists)

	// Установить права
	err = opsfile.SetFilePermissions("/tmp/test_copy.txt", 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Установить владельца и группу
	err = opsfile.SetFileOwnerGroup("/tmp/test_copy.txt", 1000, 1000)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
```
