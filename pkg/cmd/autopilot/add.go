package main

import (
	"autopilot/pkg/core"
	"autopilot/pkg/library"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// DefaultLibraryFile is the default library file path
	DefaultLibraryFile = "~/.autopilot_library.json"
	// DefaultEditor is the default editor to use to create the item
	DefaultEditor = "vi"
)

var addItemCmd = &cobra.Command{
	Use:   "add-item [flags]",
	Short: "Add a new item to the library",
	Long: heredoc.Doc(`
		Add a new item to the library.
		The add-item command will launch the editor to define the item.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		libraryFile := viper.GetString("library")
		// Get the description from the flag
		description, _ := cmd.Flags().GetString("description")
		// Get the command to add from stdin
		var command []byte
		// check if there is somethinig to read on STDIN
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				command = append(command, scanner.Bytes()...)
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}

		addCommand(libraryFile, description, string(command))
	},
}

var addCmd = &cobra.Command{
	Use:   "add command",
	Short: "Add a new command to the library",
	Long: heredoc.Doc(`
		Add a new command to the library.
		It takes the command as an argument.
		The add command will launch the editor to edit the item before adding it to the library.
		It will write the item to the default library file (~/.autopilot_library.json)
		or the file specified by the environment variable AUTOPILOT_LIBRARY.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		libraryFile := viper.GetString("library")
		// No description - it will launch the editor to add the description
		description := ""
		// Get all arguments string as the command
		// If argument contains spaces, it will be quoted
		for i, arg := range args {
			if strings.Contains(arg, " ") {
				args[i] = fmt.Sprintf("%q", arg)
			}
		}
		command := strings.Join(args, " ")

		addCommand(libraryFile, description, command)
	},
}

func addCommand(libraryPath, description, command string) {
	// Create a new library
	lib := library.NewLibrary()
	// Load the library from the file before adding the item
	lib.Load(libraryPath)
	item := library.Item{}

	position := 0
	template := "# This is a YAML template for the new item\n\n"
	template += "# Command Description\n"
	if description != "" {
		item.Description = description
		template += fmt.Sprintf("description: %s\n\n", description)
	} else {
		template += "description: \n\n"
		position = len(template) - 1
	}

	template += "# Command to save\n"
	template += "command: |\n"
	template += "     " // indent the cursor
	if command != "" {
		item.Command = command
		template += fmt.Sprintf("%s\n\n", command)
	} else {
		if position == 0 {
			position = len(template) - 1
		}
	}

	// Launch the editor to edit the item
	if item.Description == "" || item.Command == "" {
		editor := viper.GetString("EDITOR")

		// Launch the editor to create the item
		command, err := core.LaunchEditor(editor, template, position)
		if err != nil {
			log.Fatal(err)
		}
		err = core.ParseContent(command, &item)
		if err != nil {
			log.Fatal(err)
		}
	}

	// validate the item before adding it to the library
	if err := item.Validate(); err != nil {
		log.Fatal(err)
	}

	// Add the item to the library
	if err := lib.Add(item); err != nil {
		log.Fatal(err)
	}

	// Save the library
	if err := lib.Save(libraryPath); err != nil {
		log.Fatal(err)
	}

	log.Println("Item added to the library")
}

// Helper function to expand the path
func expand(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, path[1:]), nil
}

func init() {
	expandedPath, err := expand(DefaultLibraryFile)
	if err != nil {
		log.Fatal(err)
	}

	addItemCmd.Flags().StringP("editor", "e", DefaultEditor, "Editor to use to create the item. Environment variable: EDITOR")
	addItemCmd.Flags().StringP("library", "l", expandedPath, "Library file. Environment variable: AUTOPILOT_LIBRARY")
	addItemCmd.Flags().StringP("description", "d", "", "Item description")

	// Bind the environment variables to the flags
	flags := addItemCmd.Flags()

	// Bind Editor flag
	if err := viper.BindPFlag("editor", flags.Lookup("editor")); err != nil {
		log.Fatal(err)
	}
	// Bind the environment variables
	viper.BindEnv("editor", "EDITOR")

	// Bind Library flag
	if err := viper.BindPFlag("library", flags.Lookup("library")); err != nil {
		log.Fatal(err)
	}
	// Bind the environment variables
	viper.BindEnv("library", "AUTOPILOT_LIBRARY")

	rootCmd.AddCommand(addItemCmd)
	addCmd.DisableFlagParsing = true
	rootCmd.AddCommand(addCmd)
}
