package cmd

import "fmt"

func completionScript(shell string) (string, error) {
	switch shell {
	case "bash":
		return bashCompletionScript(), nil
	case "zsh":
		return zshCompletionScript(), nil
	case "fish":
		return fishCompletionScript(), nil
	case "powershell":
		return powerShellCompletionScript(), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}
}

func bashCompletionScript() string {
	return `#!/usr/bin/env bash

_aim_google_complete() {
  local IFS=$'\n'
  local completions
  completions=$(aim-google __complete --cword "$COMP_CWORD" -- "${COMP_WORDS[@]}")
  COMPREPLY=()
  if [[ -n "$completions" ]]; then
    COMPREPLY=( $completions )
  fi
}

complete -F _aim_google_complete aim-google
`
}

func zshCompletionScript() string {
	return `#compdef aim-google

_aim_google() {
  local -a completions
  completions=("${(@f)$(aim-google __complete --cword "$((CURRENT - 1))" -- "${words[@]}")}")
  _describe 'values' completions
}

compdef _aim_google aim-google
`
}

func fishCompletionScript() string {
	return `function __aim_google_complete
  set -l words (commandline -opc)
  set -l cur (commandline -ct)

  # Include the current token (partial word being typed) to match bash behavior.
  set words $words $cur

  # cword points to the last word (the one being completed).
  set -l cword (math (count $words) - 1)
  aim-google __complete --cword $cword -- $words
end

complete -c aim-google -f -a "(__aim_google_complete)"
`
}

func powerShellCompletionScript() string {
	return `Register-ArgumentCompleter -CommandName aim-google -ScriptBlock {
  param($commandName, $wordToComplete, $cursorPosition, $commandAst, $fakeBoundParameter)
  $elements = $commandAst.CommandElements | ForEach-Object { $_.ToString() }
  $cword = $elements.Count - 1
  $completions = aim-google __complete --cword $cword -- $elements
  foreach ($completion in $completions) {
    [System.Management.Automation.CompletionResult]::new($completion, $completion, 'ParameterValue', $completion)
  }
}
`
}
