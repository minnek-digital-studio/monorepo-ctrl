package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	mnkShellExec "github.com/minnek-digital-studio/mnk-shell-exec"
)

// Configuration struct to hold the mnk-config.json content
type ConfigurationConfig struct {
	Workspaces []string `json:"workspaces"`
	Extensions []string `json:"extensions"`
}

type Configuration struct {
	MonorepoCtrl struct {
		Global ConfigurationConfig `json:"global"`
		Configs []struct {
			Name     string   `json:"name"`
			Commands []string `json:"commands"`
		} `json:"configs"`
	} `json:"monorepo-ctrl"`
}

func findModifiedPackages(files []string, workspaces []string, extensions []string) map[string]bool {
	packages := make(map[string]bool)

	for _, file := range files {
		for _, workspace := range workspaces {
			if strings.HasPrefix(file, workspace+"/") {
				extension := filepath.Ext(file)
				if contains(extensions, extension) {
					packageDir := strings.Split(file, "/")[0:2]
					packagePath := strings.Join(packageDir, "/")
					packages[packagePath] = true
				}
			}
		}
	}

	return packages
}

func contains(slice []string, item string) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}

var checkFunc = `check_command() {
  COMMAND="$1"
  COMMAND_ARGUMENT=$(echo "$1" | cut -d " " -f 2)
  PACKAGE_FILE="./package.json"

  if grep -q "\"$COMMAND_ARGUMENT\":" "$PACKAGE_FILE"; then
    echo "true"
  else
    echo "false"
  fi
}`

// Function to check if a command exists
func checkCommand(name string) string {
	command := checkFunc + "; check_command \"" + name + "\""
	out, errout, err := mnkShellExec.Out(command)

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		os.Exit(1)
	}

	if strings.Contains(out, "true") {
		fmt.Printf("Running %s\n", name)
		_, _, err = mnkShellExec.OutLive(name)
		if err != nil {
			os.Exit(1)
		}
		return out
	} else {
		fmt.Printf("Skipping %s\n", name)
		return ""
	}
}

// Read and parse the mnk-config.json file
func readConfig(_configFile string) Configuration {
    configFile, err := os.Open(_configFile)
    if err != nil {
        fmt.Println("Error: Could not open " + _configFile)
        os.Exit(1)
    }
    defer configFile.Close()

    byteValue, _ := io.ReadAll(configFile)
    var config Configuration
    json.Unmarshal(byteValue, &config)

    return config
}

// Get the commands for the given command name
func getCommands(config Configuration, command string) ([]string, bool) {
    var commands []string
    commandFound := false

    for _, cfg := range config.MonorepoCtrl.Configs {
        if cfg.Name == command {
            commands = cfg.Commands
            commandFound = true
            break
        }
    }

    return commands, commandFound
}

// Get a list of the modified files
func getModifiedFiles() []string {
    output, err := exec.Command("git", "diff", "--cached", "--name-only", "--diff-filter=ACMR").Output()
    if err != nil {
        fmt.Println("Error: Could not get list of modified files.")
        os.Exit(1)
    }

    files := strings.Split(strings.TrimSpace(string(output)), "\n")
    return files
}

func Init(command string, configFile string) {
    config := readConfig(configFile)

    workspaces := config.MonorepoCtrl.Global.Workspaces
    extensions := config.MonorepoCtrl.Global.Extensions
    commands, commandFound := getCommands(config, command)

    if !commandFound {
        fmt.Printf("Error: Scope %s not found.\n", command)
        os.Exit(1)
    }

    files := getModifiedFiles()

    if len(files) == 0 {
        fmt.Println("No modified files. Skipping pre-commit hooks.")
        os.Exit(0)
    }

    packages := findModifiedPackages(files, workspaces, extensions)

    for packagePath := range packages {
        fmt.Printf("Running pre-commit hooks for %s\n", packagePath)
        err := os.Chdir(packagePath)
        if err != nil {
            fmt.Printf("Error: Package directory %s not found.\n", packagePath)
            continue
        }

        for _, command := range commands {
            checkCommand(command)
        }
    }

    err := os.Chdir("..")
    if err != nil {
        fmt.Println("Error: Could not change back to root directory.")
        os.Exit(1)
    }
}
