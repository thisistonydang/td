package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rgCmd)
}

