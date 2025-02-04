package main

import (
	"autopilot/pkg/cmd/autopilot/templates"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate Zsh key bindings",
	Long: heredoc.Doc(`
		Generate Zsh key bindings.
		The zsh command will generate Zsh key bindings for the autopilot command.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		tmpl, err := template.New("zsh").Parse(templates.ZshTemplate)
		if err != nil {
			log.Fatal(err)
		}
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		commandPath := path.Join(cwd, path.Base(os.Args[0]))
		tmpl.Execute(os.Stdout, map[string]string{
			"cmd": commandPath,
		})
	},
}

func init() {
	rootCmd.AddCommand(zshCmd)
}
