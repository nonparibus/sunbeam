import {main} from "../../wailsjs/go/models"

export interface OpenAction extends main.Action {
    type: "open",
    params: {filepath: string}
}

export function isOpenAction(cmd: main.Action): cmd is OpenAction {
    return cmd.type  == "open-file"
}

export interface CopyToClipboardAction extends main.Action {
    type: "copy-to-clipboard",
    params: {content: string}
}

export function isCopyToClipboardAction(cmd: main.Action): cmd is CopyToClipboardAction {
    return cmd.type  == "copy-to-clipbard"
}

export interface RunScriptAction extends main.Action {
    type: "run-script",
    params: {scriptpath: string, cwd: string}
}

export function isRunScriptAction(cmd: main.Action): cmd is RunScriptAction {
    return cmd.type === "script"
}

export interface PushListAction extends main.Action {
    type: "push-list",
    params: Generator
}

export interface Generator {
    path: string, mode: "filter" | "search"
}

export function IsPushListAction(cmd: main.Action): cmd is PushListAction {
    return cmd.type === "push-list"
}
