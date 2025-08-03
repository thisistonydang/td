package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/thisistonydang/td/internal/td"
)

var (
	SkipEnv     bool
	ProjectType string
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start development environment",
	Long: `Start a development environment based on the current directory's project type. 

By default, if a .env file is present, it is loaded before starting the development environment.
This can be skipped with the --skip-env flag.

Project types are auto-detected by the presence of the following files:
- Go - go.mod
- Elixir (Phoenix) - mix.exs
- JavaScript (pnpm) - package.json

To force a project type, rather than auto-detecting, use the --type flag. (e.g. --type go)

The following commands are run for each project type:
- Go - air
- Elixir (Phoenix) - iex -S mix phx.server
- JavaScript (pnpm) - pnpm dev
`,
	Example: `  td dev
  td dev --skip-env
  td dev --type go`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		startDevEnvironment()
	},
}

func init() {
	devCmd.Flags().BoolVarP(&SkipEnv, "skip-env", "s", false, "skip loading .env file before starting the dev environment")
	devCmd.Flags().StringVarP(&ProjectType, "type", "t", "", "force project type (e.g. go, ex, js)")
	rootCmd.AddCommand(devCmd)
}

func startDevEnvironment() {
	var argv []string
	goArgs := []string{"air"}
	exArgs := []string{"iex", "-S", "mix", "phx.server"}
	jsArgs := []string{"pnpm", "dev"}

	// Force project type if specified
	switch ProjectType {
	case "go", "golang", "air":
		argv = goArgs
	case "ex", "phx", "elixir", "phoenix", "iex":
		argv = exArgs
	case "js", "ts", "javascript", "typescript", "node", "nodejs", "npm", "pnpm":
		argv = jsArgs
	case "":
