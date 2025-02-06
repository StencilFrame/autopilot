package main

import (
	"autopilot/pkg/library"
	"fmt"
	"log"
	"os"

	"github.com/MakeNowJust/heredoc"
	fzf "github.com/junegunn/fzf/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fzfCmd = &cobra.Command{
	Use:   "fzf",
	Short: "Fuzzy find an item in the library",
	Long: heredoc.Doc(`
		Fuzzy find an item in the library.
		The fzf command will launch a fuzzy finder to select an item from the library.
		The matched item will be printed to the standard output.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		lib := library.NewLibrary()

		libFile := viper.GetString("library")

		// Load the library
		if err := lib.Load(libFile); err != nil {
			log.Fatalf("Error: %v", err)
		}

		inputChan := make(chan string)
		go func() {
			for _, s := range lib.Items(nil) {
				inputChan <- s
			}
			close(inputChan)
		}()

		outputChan := make(chan string)
		selectedItem := ""
		go func() {
			for s := range outputChan {
				selectedItem = s
			}
		}()

		exit := func(code int, err error) {
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			os.Exit(code)
		}

		// Build fzf.Options
		options, err := fzf.ParseOptions(
			true, // whether to load defaults ($FZF_DEFAULT_OPTS_FILE and $FZF_DEFAULT_OPTS)
			[]string{"--multi", "--reverse", "--border", "--height=40%"},
		)
		if err != nil {
			exit(fzf.ExitError, err)
		}

		// Set up input and output channels
		options.Input = inputChan
		options.Output = outputChan

		// Run fzf
		code, err := fzf.Run(options)
		if err != nil {
			exit(code, err)
		}

		// Link back the selected item to the library item
		if selectedItem != "" {
			// Extract the command from the selected item
			item, err := lib.GetItemByCommand(selectedItem)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			// Print the command in a way that it can be evaluated by the shell
			fmt.Println(item.Command)
		}
	},
}

func init() {
	expandedPath, err := expand(DefaultLibraryFile)
	if err != nil {
		log.Fatal(err)
	}

	fzfCmd.Flags().StringP("library", "l", expandedPath, "Library file. Environment variable: AUTOPILOT_LIBRARY")

	// Bind the environment variables to the flags
	flags := addItemCmd.Flags()

	// Bind Library flag
	if err := viper.BindPFlag("library", flags.Lookup("library")); err != nil {
		log.Fatal(err)
	}
	// Bind the environment variables
	viper.BindEnv("library", "AUTOPILOT_LIBRARY")

	rootCmd.AddCommand(fzfCmd)
}
