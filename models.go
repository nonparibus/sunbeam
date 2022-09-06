package main

type SearchItem struct {
	IconSource     string   `json:"icon_src"`
	Title          string   `json:"title" validate:"required"`
	Subtitle       string   `json:"subtitle"`
	Fill           string   `json:"fill"`
	AccessoryTitle string   `json:"accessory_title"`
	Keywords       []string `json:"keywords"`
	Actions        []Action `json:"actions" validate:"required,gte=1,dive"`
}

type Action struct {
	Icon    string   `json:"icon"`
	Title   string   `json:"title" validate:"required"`
	Type    string   `json:"type" validate:"required"`
	Content string   `json:"content"`
	Path    string   `json:"path"`
	Url     string   `json:"url"`
	Args    []string `json:"args"`
}

type Shortcut struct {
	Ctrl  bool   `json:"ctrl"`
	Shift bool   `json:"shift"`
	Alt   bool   `json:"alt"`
	Super bool   `json:"super"`
	Key   string `json:"key"`
}

type CommandType int

const (
	Open CommandType = iota + 1
	OpenInBrowser
	CopyToClipboard
	RunScript
	RunCommand
)

func (c CommandType) String() string {
	switch c {
	case Open:
		return "open"
	case OpenInBrowser:
		return "open-in-browser"
	case CopyToClipboard:
		return "copy-to-clipboard"
	case RunScript:
		return "run-script"
	case RunCommand:
		return "run-command"
	default:
		return ""
	}
}

func NewOpenAction(title string, path string) Action {
	return Action{Title: title, Icon: "/raycast/icon-app-window-16.svg", Type: Open.String(), Path: path}
}

func NewOpenInBrowser(title string, icon string, url string) Action {
	return Action{Title: title, Icon: icon, Type: OpenInBrowser.String(), Url: url}
}

func NewCopyToClipboardAction(title string, content string) Action {
	return Action{Icon: "/raycast/icon-copy-clipboard-16.svg", Type: CopyToClipboard.String(), Content: content}
}

func NewRunScriptAction(title string, path string, args ...string) Action {
	return Action{Title: title, Icon: "raycast/icon-terminal-16.svg", Type: RunScript.String(), Path: path}
}

func NewRunCommandAction(title string, path string, args ...string) Action {
	return Action{Type: RunCommand.String(), Icon: "raycast/icon-window-list-16.svg", Args: args}
}
