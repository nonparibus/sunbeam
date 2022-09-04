import {main} from "../../wailsjs/go/models"

export interface OpenCommand {
    type: "open",
    params: {filepath: string}
}

export function isOpenCommand(cmd: main.Command): cmd is OpenCommand {
    return cmd.type  == "open-file"
}

export interface CopyToClipboardCommand {
    type: "copy-to-clipboard",
    params: {content: string}
}

export function isCopyToClipboardCommand(cmd: main.Command): cmd is CopyToClipboardCommand {
    return cmd.type  == "copy-to-clipbard"
}

export interface RunScriptCommand {
    type: "run-script",
    params: {scriptpath: string, cwd: string}
}

export function isRunScriptCommand(cmd: main.Command): cmd is RunScriptCommand {
    return cmd.type === "script"
}

export interface PushListCommand {
    type: "push-list",
    params: {scriptpath: string, cwd: string}
}

export function IsPushListCommand(cmd: main.Command): cmd is PushListCommand {
    return cmd.type === "list"
}
