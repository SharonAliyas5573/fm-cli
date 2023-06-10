package main
import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	
)
func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage:fm [command] [source] [destination]")
		fmt.Println("Available commands:")
		fmt.Println("-cp: Copy files/directories")
		fmt.Println("-mv: Move files/directories")
		fmt.Println("-rn: Rename files")
		fmt.Println("-rm: Remove files/folders to trash")
		return
	}
	
	command := os.Args[1]
	source := os.Args[2]
	

	switch command {
	case "-cp":
		destination := os.Args[3]
		err := copyFile(source, destination)
		if err != nil {
			fmt.Printf("Error copying : %s\n", err)
		} else {
			fmt.Println("Copied successfully!")
		}
	case "-mv":
		destination := os.Args[3]
		err := moveFile(source, destination)
		if err != nil {
			fmt.Printf("Error moving : %s\n", err)
		} else {
			fmt.Println(" Moved successfully!")
		}
	case "-rn":
		if len(os.Args) ==0 {
			fmt.Println("Usage: fm -rn [current_name] [new_name]")
			return
		}
		currentName := os.Args[2]
		newName := os.Args[3]
		err := renameFile(currentName, newName)
		if err != nil {
			fmt.Printf("Error renaming : %s\n", err)
		} else {
			fmt.Println(" Renamed successfully!")
		}
	case "-rm":
		if len(os.Args) < 2 {
			fmt.Println("Usage:\tfm -rm [file/directory name] \\\\Move to trash\n \tfm -rm -f [file/directory name] \\\\Delete without trash")
		} else if len(os.Args) >= 3 && os.Args[2] == "-f" {
			if len(os.Args) < 4 {
				fmt.Println("File name is required after -f flag.")
				return
			}
			isForce := true
			fileName := os.Args[3]
			removeFile(fileName, isForce)
		} else {
			if len(os.Args) < 3 {
				fmt.Println("File name is required.")
				return
			}
			fileName := os.Args[2]
			removeFile(fileName, false)
		}
	
	default:
		fmt.Println("Invalid command. Available commands:")
		fmt.Println("-cp: Copy files/directories")
		fmt.Println("-mv: Move files/directories")
		fmt.Println("-rn: Rename files")
		fmt.Println("-rm: Remove files/folders to trash")
	}
}


func copyFile(source, destination string) error {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if sourceInfo.Mode().IsRegular() {
		// Copying a file
		sourceFile, err := os.Open(source)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		destFile, err := os.Create(destination)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, sourceFile)
		if err != nil {
			return err
		}

		err = destFile.Sync()
		if err != nil {
			return err
		}
	} else if sourceInfo.IsDir() {
		// Copying a directory
		err := copyDir(source, destination)
		if err != nil {
			return err
		}
	}
	
	return nil
}
func copyDir(sourceDir, destinationDir string) error {
	_, err := os.Stat(destinationDir)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		err = os.MkdirAll(destinationDir, 0755)
		if err != nil {
			return err
		}
	}

	sourceBase := filepath.Base(sourceDir)
	destinationPath := filepath.Join(destinationDir, sourceBase)

	err = os.MkdirAll(destinationPath, 0755)
	if err != nil {
		return err
	}

	directory, err := os.Open(sourceDir)
	if err != nil {
		return err
	}
	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		sourcePath := filepath.Join(sourceDir, obj.Name())
		destPath := filepath.Join(destinationPath, obj.Name())

		if obj.IsDir() {
			err = copyDir(sourcePath, destPath)
			if err != nil {
				fmt.Printf("Error copying directory: %s\n", err)
			}
		} else {
			err = copyFile(sourcePath, destPath)
			if err != nil {
				fmt.Printf("Error copying file: %s\n", err)
			}
		}
	}
	
	return nil
}


func moveFile(source, destination string) error {
	destPath := filepath.Join(destination, filepath.Base(source))
	err := os.Rename(source, destPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("source file does not exist: %s", source)
		}
		return err
	} 
	
	return nil
}


func renameFile(currentName string, newName string) error {
	files, err := filepath.Glob(currentName)
	if err != nil {
		return err
	}

	for _, file := range files {
		newPath := filepath.Join(filepath.Dir(file), newName)
		err := os.Rename(file, newPath)
		if err != nil {
			return err
		}
		
	}
	
	return nil
}

func removeFile(source string, isForce bool) error {
	if !isForce  {
		trashDir, err := getTrashDirectory()
		if err != nil {
			fmt.Println(err)
		}

		destPath := filepath.Join(trashDir, filepath.Base(source))
		err = os.Rename(source, destPath) 
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("source file does not exist: %s", source)
			}
			return err
		} 
		fmt.Printf("Moved '%s' to trash successfully\n", source)
	}else if isForce{ 
		err := os.Remove(source)
		if err != nil{
			return fmt.Errorf("error :%s", err)
		}
		fmt.Println("Deleted successfully")
	}
	return nil
}


func getTrashDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	trashDir := filepath.Join(homeDir, ".local", "share", "Trash")
	return trashDir, nil
}


// func hasFlag(flag string) bool {
// 	for _, arg := range os.Args {
// 		if arg == flag {
// 			return true
// 		}
// 	}
// 	return false
// }
