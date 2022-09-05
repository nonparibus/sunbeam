package main

type SearchItem struct {
	Icon           string   `json:"icon"`
	Title          string   `json:"title" validate:"required"`
	Subtitle       string   `json:"subtitle"`
	AccessoryTitle string   `json:"accessory_title"`
	Keywords       []string `json:"keywords"`
	Actions        []Action `json:"actions" validate:"required,dive"`
}

type Action struct {
	Icon     string   `json:"icon"`
	Title    string   `json:"title" validate:"required"`
	Command  Command  `json:"command" validate:"required"`
	Shortcut Shortcut `json:"shortcut"`
}

type Command struct {
	Type   string                 `json:"type" validate:"required"`
	Params map[string]interface{} `json:"params" validate:"required"`
}

type Shortcut struct {
	Ctrl  bool   `json:"ctrl"`
	Shift bool   `json:"shift"`
	Alt   bool   `json:"alt"`
	Super bool   `json:"super"`
	Key   string `json:"key" validate:"required"`
}

type CommandType int

const (
	OpenFile CommandType = iota + 1
	OpenUrl
	CopyToClipboard
	RunScript
	PushList
)

func (c CommandType) String() string {
	switch c {
	case OpenFile:
		return "open-file"
	case OpenUrl:
		return "open-url"
	case CopyToClipboard:
		return "copy-to-clipboard"
	case RunScript:
		return "run-script"
	case PushList:
		return "push-list"
	default:
		return "unknown"
	}
}

func NewOpenCommand(filepath string) Command {
	return Command{Type: OpenFile.String(), Params: map[string]interface{}{
		"filepath": filepath,
	}}
}

func NewOpenInBrowserCommand(url string) Command {
	return Command{Type: OpenUrl.String(), Params: map[string]interface{}{
		"url": url,
	}}
}

func NewCopyToClipboardCommand(content string) Command {
	return Command{Type: CopyToClipboard.String(), Params: map[string]interface{}{
		"content": content,
	}}
}

func NewRunScriptCommand(scriptPath string, args []string, mode string) Command {
	return Command{Type: RunScript.String(), Params: map[string]interface{}{
		"scriptpath": scriptPath,
		"args":       args,
	}}
}

func NewPushListCommand(scriptPath string, args []string, mode string) Command {
	return Command{Type: RunScript.String(), Params: map[string]interface{}{
		"scriptpath": scriptPath, "mode": mode,
	}}
}
