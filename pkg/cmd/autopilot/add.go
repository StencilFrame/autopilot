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

var addCmd = &cobra.Command{
	Use:   "add [flags]",
	Short: "Add a new item to the library",
	Long: heredoc.Doc(`
		Add a new item to the library.
		The add command will launch the editor to define the item.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		libraryFile, _ := cmd.Flags().GetString("library")
		if libraryFile == "" {
			log.Fatal("Library file is required")
		}
		expandedPath, err := expand(libraryFile)
		if err != nil {
			log.Fatal(err)
		}

		// Create a new library
		lib := library.NewLibrary()
		// Load the library from the file before adding the item
		lib.Load(expandedPath)
		item := library.Item{}

		template := "# This is a YAML template for the new item\n\n"

		template += "# Command Description\n"
		if description, _ := cmd.Flags().GetString("description"); description != "" {
			item.Description = description
			template += fmt.Sprintf("# Description: %s\n\n", description)
		} else {
			template += "description: \n\n"
		}

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
			item.Command = string(command)
		} else {
			editor := viper.GetString("EDITOR")
			// Launch the editor to create the item
			template += "# Command to save\n"
			template += "command: |\n"
			template += "     " // indent the cursor
			command, err := core.LaunchEditor(editor, template, len(template))
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
		if err := lib.Save(expandedPath); err != nil {
			log.Fatal(err)
		}

		log.Println("Item added to the library")
	},
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

	addCmd.Flags().StringP("editor", "e", DefaultEditor, "Editor to use to create the item. Environment variable: EDITOR")
	addCmd.Flags().StringP("library", "l", expandedPath, "Library file. Environment variable: AUTOPILOT_LIBRARY")
	addCmd.Flags().StringP("description", "d", "", "Item description")

	// Bind the environment variables to the flags
	flags := addCmd.Flags()

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

	rootCmd.AddCommand(addCmd)
}
