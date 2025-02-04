package main

import (
	"autopilot/pkg/cmd/autopilot/templates"
	"log"
	"os"
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
		tmpl.Execute(os.Stdout, map[string]string{
			"cmd": os.Args[0],
		})
	},
}

func init() {
	rootCmd.AddCommand(zshCmd)
}
