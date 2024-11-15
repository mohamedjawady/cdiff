package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func InitRepo() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the project name: ")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	fmt.Print("Enter the author's name: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("Enter the license type (e.g., MIT, GPL-3.0): ")
	license, _ := reader.ReadString('\n')
	license = strings.TrimSpace(license)

	repoPath := ".cdiff"
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		if err := os.Mkdir(repoPath, 0755); err != nil {
			return fmt.Errorf("could not create .cdiff directory: %v", err)
		}

		subDirs := []string{"objects", "refs/heads"}
		for _, dir := range subDirs {
			if err := os.MkdirAll(filepath.Join(repoPath, dir), 0755); err != nil {
				return fmt.Errorf("could not create directory %s: %v", dir, err)
			}
		}

		configFile, err := os.Create(filepath.Join(repoPath, "config"))
		if err != nil {
			return fmt.Errorf("could not create config file: %v", err)
		}
		defer configFile.Close()

		_, err = configFile.WriteString(fmt.Sprintf("[project]\n\tname = %s\n\tauthor = %s\n\tlicense = %s\n", projectName, author, license))
		if err != nil {
			return fmt.Errorf("could not write to config file: %v", err)
		}

		_, err = configFile.WriteString("[user]\n\temail = you@example.com\n\tname = Your Name\n")
		if err != nil {
			return fmt.Errorf("could not write user info to config file: %v", err)
		}

		fmt.Println("Initialized an empty cdiff repository in .cdiff")
	} else {
		return fmt.Errorf(".cdiff directory already exists")
	}

	return nil
}
