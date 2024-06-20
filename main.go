package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	xstrings "github.com/charmbracelet/x/exp/strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
)

var (
	cli_tools        []string
	aliases          []string
	code_editor      string
	zsh              bool
	zed_installed    bool
	vscode_installed bool
	current_os       string

	curr_step int = 0
)

func CheckOperatingSystem() {
	if runtime.GOOS == "windows" {
		current_os = "windows"
	}

	if runtime.GOOS == "darwin" {
		current_os = "darwin"
	}

	if runtime.GOOS == "linux" {
		current_os = "linux"
	}
}

func main() {

	CheckOperatingSystem()

	form := huh.NewForm(
		huh.NewGroup(huh.NewMultiSelect[string]().
			Title("CLI tools to install").
			Description("Select the tools you would like to install").
			Options(
				huh.NewOption("Eza - better LS", "eza"),
				huh.NewOption("FZF - fuzzy finder", "fzf"),
				huh.NewOption("Bat - better cat", "bat"),
				huh.NewOption("Ripgrep - better grep", "ripgrep"),
				huh.NewOption("Oh my zsh - ZSH theming", "oh-my-zsh"),
				huh.NewOption("TheFuck - CLI typo fixer", "thefuck"),
				huh.NewOption("lazygit - terminal git manage", "lazygit"),
				huh.NewOption("Skip step", "skip"),
			).
			Validate(func(s []string) error {
				if len(s) == 0 {
					return errors.New("please select at least one tool to install")
				}
				return nil
			}).
			Value(&cli_tools)),
		// Install handy dandy aliases
		huh.NewGroup(huh.NewMultiSelect[string]().
			Title("Aliases").
			Description("Select the aliases you would like to install").
			Options(
				huh.NewOption("git purge", "git-purge"),
				huh.NewOption("Skip step", "skip"),
			).
			Validate(func(s []string) error {
				if len(s) == 0 {
					return errors.New("please select at least one alias to install")
				}
				return nil
			}).
			Value(&aliases)),
		// Install text editor
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Code editor").
				Description("Select the code editor you want to install").
				Options(
					huh.NewOption("Visual Studio Code", "visual-studio-code"),
					huh.NewOption("Zed", "zed"),
					huh.NewOption("Skip step", "skip"),
				).
				Validate(func(s string) error {
					if s == "" {
						return errors.New("please select at least one code editor to install")
					}
					return nil
				}).
				Value(&code_editor)),

		// Final info about cli installs
		huh.NewGroup(
			huh.NewConfirm().
				Title("Is your default shell zsh?").
				Description("This is required for Oh my zsh to function.\n").
				Value(&zsh),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if !zsh {
		color.Red("Oh my zsh requires zsh as the default shell.\n")
		os.Exit(1)
	}

	install := func() {

		// install cli tools
		for _, tool := range cli_tools {
			switch tool {
			case "eza":
				_ = spinner.New().Title("Installing Eza...").Action(func() {
					curr_step += 1
					if current_os == "windows" {
						color.Red("can not install eza on this operating system\n")
						return
					}
					_, err := exec.Command("brew", "install", "eza").Output()
					if err != nil {
						color.Red("Error installing eza\n")
						os.Exit(1)
					}

					_, setupErr := exec.Command("/bin/zsh", "-c", "alias ls='eza --color=always --long --git --no-filesize --no-time --no-user --no-permission --tree --level=2'").Output()
					if setupErr != nil {
						color.Red("Can not set ls alias\n")
					}
				}).Run()
			case "fzf":
				_ = spinner.New().Title("Installing fzf...").Action(func() {
					curr_step += 1
					if current_os == "windows" {
						color.Red("can not install fzf on this operating system\n")
						return
					}
					_, err := exec.Command("brew", "install", "fzf").Output()
					if err != nil {
						color.Red("Error installing fzf\n")
						os.Exit(1)
					}
				}).Run()
			case "bat":
				_ = spinner.New().Title("Installing bat...").Action(func() {
					curr_step += 1
					if current_os == "windows" {
						color.Red("can not install bat on this operating system\n")
						return
					}
					_, err := exec.Command("brew", "install", "bat").Output()
					if err != nil {
						color.Red("Error installing bat\n")
						os.Exit(1)
					}
				}).Run()
			case "ripgrep":
				_ = spinner.New().Title("Installing Ripgrep...").Action(func() {
					curr_step += 1
					if current_os == "windows" {
						color.Red("can not install ripgrep on this operating system\n")
						return
					}
					_, err := exec.Command("brew", "install", "ripgrep").Output()
					if err != nil {
						color.Red("Error installing ripgrep\n")
						os.Exit(1)
					}
				}).Run()
			case "oh-my-zsh":
				_ = spinner.New().Title("Installing Oh my zsh...").Action(func() {
					curr_step += 1
					if current_os == "windows" {
						color.Red("can not install oh my zsh on this operating system\n")
						return
					}
					_, err := exec.Command("brew", "install", "oh-my-zsh").Output()
					if err != nil {
						color.Red("Error installing oh my zsh\n")
						os.Exit(1)
					}
				}).Run()
			case "thefuck":
				_ = spinner.New().Title("Installing thefuck...").Action(func() {
					curr_step += 1
					if current_os != "darwin" {
						color.Red("can not install thefuck on this operating system\n")
						return
					}

					_, err := exec.Command("brew", "install", "thefuck").Output()
					if err != nil {
						color.Red("Error installing thefuck\n")
						os.Exit(1)
					}

					_, setupErr := exec.Command("/bin/zsh", "-c", "eval $(thefuck --alias)").Output()
					if setupErr != nil {
						color.Red("Can not set thefuck alias\n")
					}
				}).Run()
			case "lazygit":
				_ = spinner.New().Title("Installing lazygit..").Action(func() {
					curr_step += 1
					if current_os == "windows" {
						color.Red("can not install lazygit on this operating system using brew\n")
						return
					}
					_, err := exec.Command("brew", "install", "lazygit").Output()
					if err != nil {
						color.Red("Error installing oh my lazygit\n")
						os.Exit(1)
					}
				}).Run()
			case "skip":
				continue
			}
		}

		for _, alias := range aliases {
			switch alias {
			case "git-purge":
				_ = spinner.New().Title("Setting git-purge alias").Action(func() {
					curr_step += 1
					if current_os == "windows" {
						color.Red("can not set git-purge alias on this operating system \n")
						return
					}

					_, err := exec.Command("/bin/zsh",
						"-c",
						"alias git-purge=\"git fetch -p && git branch --merged | grep -v '*' | grep -v 'master' | xargs git branch -d\"").
						Output()
					if err != nil {
						color.Red("Can not set git purge alias\n")
					}
				}).Run()
			case "skip":
				continue
			}
		}

		// install code editor
		if code_editor != "skip" {

			if current_os == "windows" {
				color.Red("can not set install via brew on this operating system\n")
				return
			}

			if current_os == "linux" || code_editor == "zed" {
				color.Red("can not install zed on linux yet...\n")
				code_editor = "skip"
				return
			} else {
				curr_step += 1

				_ = spinner.New().Title("Installing text editor...").Action(func() {
					_, err := exec.Command("brew", "install", "--cask", code_editor).Output()
					if err != nil {
						color.Red("Error installing " + code_editor)
						os.Exit(1)
					}
				})
			}
		}

	}

	_ = spinner.New().Title("").Action(install).Run()

	var sb strings.Builder
	keyword := func(s string) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
	}
	if cli_tools[0] == "skip" {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("209")).Render("CLI Tools skipped."))
	} else {
		fmt.Fprintf(&sb,
			"Following tools were installed \n%s\n",
			keyword(xstrings.SpokenLanguageJoin(cli_tools, xstrings.EN)),
		)
		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				Padding(1, 2).
				Render(sb.String()),
		)
	}

	var aliases_sb strings.Builder
	if aliases[0] == "skip" {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("209")).Render("Aliases skipped."))
	} else {
		fmt.Fprintf(&aliases_sb,
			"Following aliases have been set \n%s\n",
			keyword(xstrings.SpokenLanguageJoin(aliases, xstrings.EN)),
		)
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Render(aliases_sb.String()))
	}

	if code_editor == "skip" {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("209")).
			Render("Code editor install skipped."))
	} else {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("211")).
			Render("Code editor installed " + code_editor))
	}

}
