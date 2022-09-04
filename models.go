package main

type SearchItem struct {
	Key            string   `json:"key"`
	Icon           string   `json:"icon"`
	Title          string   `json:"title"`
	Subtitle       string   `json:"subtitle"`
	AccessoryTitle string   `json:"accessory_title"`
	Actions        []Action `json:"actions"`
}

type Action struct {
	Icon     string   `json:"icon"`
	Title    string   `json:"title"`
	Command  Command  `json:"command"`
	Shortcut Shortcut `json:"shortcut"`
}

type Command struct {
	Type   CommandType            `json:"type"`
	Params map[string]interface{} `json:"params"`
}

type Shortcut struct {
	Ctrl  bool   `json:"ctrl"`
	Shift bool   `json:"shift"`
	Alt   bool   `json:"alt"`
	Super bool   `json:"super"`
	Key   string `json:"key"`
}

type CommandType string

const (
	OpenFile        = "open-file"
	OpenUrl         = "open-url"
	CopyToClipboard = "copy-to-clipboard"
	RunScript       = "run-script"
	PushList        = "push-list"
)

func NewOpenCommand(filepath string) Command {
	return Command{Type: OpenFile, Params: map[string]interface{}{
		"filepath": filepath,
	}}
}

func NewOpenInBrowserCommand(url string) Command {
	return Command{Type: OpenUrl, Params: map[string]interface{}{
		"url": url,
	}}
}

func NewCopyToClipboardCommand(content string) Command {
	return Command{Type: CopyToClipboard, Params: map[string]interface{}{
		"content": content,
	}}
}

func RunScriptCommand(scriptPath string, args ...string) Command {
	return Command{Type: RunScript, Params: map[string]interface{}{
		"scriptpath": scriptPath,
		"args":       args,
	}}
}

func PushListCommand(mode string, scriptPath string, args []string) Command {
	return Command{Type: RunScript, Params: map[string]interface{}{
		"scriptpath": scriptPath, "mode": mode,
	}}
}
