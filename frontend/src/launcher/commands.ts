import {main} from "../../wailsjs/go/models"

export interface OpenCommand {
    type: "open",
    params: {filepath: string}
}

export function isOpenCommand(cmd: main.Command): cmd is OpenCommand {
    return cmd.type  == "open"
}

export interface CopyToClipboardCommand {
    type: "copy-to-clipboard",
    params: {content: string}
}

export function isCopyToClipboardCommand(cmd: main.Command): cmd is CopyToClipboardCommand {
    return cmd.type  == "copy-to-clipbard"
}
