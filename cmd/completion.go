/*
 * Copyright 2020 Netflix, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long:  `Generate shell completion script for Bash, Zsh, Fish, and Powershell.`,
}

var bashCompletionCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generate Bash completion",
	Long: `To load completions:

$ source <(weep completion bash)

To load completions for each session, execute once:

Linux:

  $ weep completion bash > /etc/bash_completion.d/weep

MacOS:

  $ weep completion bash > /usr/local/etc/bash_completion.d/weep
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Root().GenBashCompletion(os.Stdout); err != nil {
			log.Fatal(err)
		}
	},
}

var zshCompletionCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate Zsh completion",
	Long: `To load completions:

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for each session, execute once:

$ weep completion zsh > "${fpath[1]}/_weep"

You will need to start a new shell for this setup to take effect.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Root().GenZshCompletion(os.Stdout); err != nil {
			log.Fatal(err)
		}
	},
}

var fishCompletionCmd = &cobra.Command{
	Use:   "fish",
	Short: "Generate Fish completion",
	Long: `To load completions:

$ weep completion fish | source

To load completions for each session, execute once:

$ weep completion fish > ~/.config/fish/completions/weep.fish
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Root().GenFishCompletion(os.Stdout, true); err != nil {
			log.Fatal(err)
		}
	},
}

var powershellCompletionCmd = &cobra.Command{
	Use:   "powershell",
	Short: "Generate Powershell completion",
	Long: `To load completions:

We don't really know how Powershell works. 🙈

This doc has a little more context: https://github.com/spf13/cobra/blob/master/powershell_completions.md

Want to improve this help text? Check out this issue: https://github.com/Netflix/weep/issues/17
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Root().GenPowerShellCompletion(os.Stdout); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	completionCmd.AddCommand(bashCompletionCmd)
	completionCmd.AddCommand(zshCompletionCmd)
	completionCmd.AddCommand(fishCompletionCmd)
	completionCmd.AddCommand(powershellCompletionCmd)
	rootCmd.AddCommand(completionCmd)
}