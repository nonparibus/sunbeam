import React from "react";
import "./raycast.scss";
import { Command } from "cmdk";
import { TerminalIcon, RaycastLightIcon } from "./icons";
import * as App from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";
import * as runtime from "../../wailsjs/runtime/runtime";
import { isCopyToClipboardCommand, isOpenCommand, isRunCommand } from "./commands";
import * as Popover from "@radix-ui/react-popover";

// window.onblur = () => {
//   runtime.WindowHide();
// };

export function RaycastCMDK() {
  const inputRef = React.useRef<HTMLInputElement | null>(null);
  const listRef = React.useRef(null);
  const [focusedValue, setFocusedValue] = React.useState("");
  const [query, setQuery] = React.useState("");
  const [items, setItems] = React.useState<Record<string, main.SearchItem>>();

  const focusedItem = items ? items[focusedValue] : undefined;

  function handleResponse(response: main.Response) {
    switch (response.type) {
      case "filter": {
        const items = Object.fromEntries(
          response.items.map((item) => [item.title.trim().toLowerCase(), item])
        );
        setItems(items);
        break;
      }
      default: {
        runtime.LogError(
          `No handler for response: ${JSON.stringify(response)}`
        );
      }
    }
  }

  function handleCommand(command: main.Command) {
    runtime.LogDebug(`Handling Command: ${JSON.stringify(command)}`)
    if (isCopyToClipboardCommand(command)) {
      App.CopyToClipboard(command.params.content);
      return;
    }
    if (isOpenCommand(command)) {
      App.OpenFile(command.params.filepath);
      return;
    }
    if (isRunCommand(command)) {
      App.RunScript(command.params.scriptpath, [])
      return;
    }
  }

  React.useEffect(() => {
    App.RootItems().then(handleResponse);
  }, []);

  React.useEffect(() => {
    runtime.EventsOn("toggle", () => {
      runtime.WindowCenter();
      runtime.WindowShow();
    });
    return () => {
      runtime.EventsOff("toggle");
    };
  }, []);

  React.useEffect(() => {
    function listener(e: KeyboardEvent) {
      if (e.key !== "Escape") {
        return;
      }
      runtime.WindowHide();
      setQuery("");
    }

    document.addEventListener("keydown", listener);

    return () => {
      document.removeEventListener("keydown", listener);
    };
  }, []);

  React.useEffect(() => {
    inputRef?.current?.focus();
  }, []);

  return (
    <div className="raycast">
      <Command
        shouldFilter={true}
        value={focusedValue}
        onValueChange={setFocusedValue}
      >
        <div cmdk-raycast-top-shine="" />
        <Command.Input
          ref={inputRef}
          value={query}
          onValueChange={setQuery}
          autoFocus
          placeholder="Search for apps and commands..."
        />
        <hr cmdk-raycast-loader="" />
        <Command.List ref={listRef}>
          <Command.Empty>
            {query ? "No Result Found." : "Type Something !"}
          </Command.Empty>
          <Command.Group heading="Results">
            {Object.entries(items || {}).map(([key, item]) => (
              <Item
                item={item}
                key={key}
                value={key}
                onSelect={() => {
                  handleCommand(item.actions[0].command);
                }}
              />
            ))}
          </Command.Group>
        </Command.List>

        <div cmdk-raycast-footer="">
          <div cmd-raycast-submenu="">
            <RaycastMenu
              listRef={listRef}
              inputRef={inputRef}
            />
          </div>

          {focusedItem ? (
            <>
              <button cmdk-raycast-subcommand-trigger=""onClick={() => {handleCommand(focusedItem.actions[0]?.command)}}>
                {focusedItem.actions[0]?.title || ""}
                <kbd>↵</kbd>
              </button>
              <hr />
              <SubCommand
                listRef={listRef}
                focusedItem={focusedItem}
                inputRef={inputRef}
                onAction={(action) => handleCommand(action.command)}
              />
            </>
          ) : null}
        </div>
      </Command>
    </div>
  );
}

function ItemIcon({ icon }: { icon: string }) {
  return <img width={24} height={24} src={`/${icon}`} />;
}

function Item({
  item,
  onSelect,
  value,
}: {
  item: main.SearchItem;
  value: string;
  onSelect: () => void;
}) {
  return (
    <Command.Item value={value} onSelect={onSelect}>
      {item.icon ? <ItemIcon icon={item.icon} /> : <TerminalIcon />}
      {item.title}
      <span cmdk-raycast-subtitle="">{item.subtitle}</span>
      <span cmdk-raycast-accessory-title="">{item.accessory_title}</span>
    </Command.Item>
  );
}

function RaycastMenu({
  inputRef,
  listRef,
}: {
  inputRef: React.RefObject<HTMLInputElement>;
  listRef: React.RefObject<HTMLElement>;
}) {
  const [open, setOpen] = React.useState(false);

  React.useEffect(() => {
    const el = listRef.current;

    if (!el) return;

    if (open) {
      el.style.overflow = "hidden";
    } else {
      el.style.overflow = "";
    }
  }, [open, listRef]);

  return (
    <Popover.Root open={open} onOpenChange={setOpen} modal>
      <Popover.Trigger
        cmdk-raycast-subcommand-trigger=""
        onClick={() => setOpen(true)}
        aria-expanded={open}
      >
      <RaycastLightIcon />
      </Popover.Trigger>
      <Popover.Content
        side="top"
        align="start"
        className="raycast-submenu"
        sideOffset={16}
        alignOffset={0}
        onCloseAutoFocus={(e) => {
          e.preventDefault();
          inputRef?.current?.focus();
        }}
      >
        <Command>
          <Command.List>
            <Command.Group heading="Raycast">
              <Command.Item>Manual</Command.Item>
              <Command.Item>Switch to Dark Mode</Command.Item>
              <Command.Item onSelect={() => {runtime.Quit()}}>Quit Raycast</Command.Item>
            </Command.Group>
          </Command.List>
          <Command.Input placeholder="Search for actions..." />
        </Command>
      </Popover.Content>
    </Popover.Root>
  );
}

function SubCommand({
  inputRef,
  listRef,
  focusedItem,
  onAction,
}: {
  inputRef: React.RefObject<HTMLInputElement>;
  listRef: React.RefObject<HTMLElement>;
  focusedItem?: main.SearchItem;
  onAction: (action: main.Action) => void;
}) {
  const [open, setOpen] = React.useState(false);

  React.useEffect(() => {
    function listener(e: KeyboardEvent) {
      if (e.key === "k" && e.ctrlKey) {
        e.preventDefault();
        setOpen((o) => !o);
      }
    }

    document.addEventListener("keydown", listener);

    return () => {
      document.removeEventListener("keydown", listener);
    };
  }, []);

  React.useEffect(() => {
    const el = listRef.current;

    if (!el) return;

    if (open) {
      el.style.overflow = "hidden";
    } else {
      el.style.overflow = "";
    }
  }, [open, listRef]);

  return (
    <Popover.Root open={open} onOpenChange={setOpen} modal>
      <Popover.Trigger
        cmdk-raycast-subcommand-trigger=""
        onClick={() => setOpen(true)}
        aria-expanded={open}
      >
        Actions
        <kbd>⌃</kbd>
        <kbd>K</kbd>
      </Popover.Trigger>
      <Popover.Content
        side="top"
        align="end"
        className="raycast-submenu"
        sideOffset={16}
        alignOffset={0}
        onCloseAutoFocus={(e) => {
          e.preventDefault();
          inputRef?.current?.focus();
        }}
      >
        <Command>
          <Command.List>
            <Command.Group heading={focusedItem?.title}>
              {focusedItem?.actions.map((action) => (
                <Command.Item onSelect={() => onAction(action)}>
                  {action.title}
                  <div cmdk-raycast-submenu-shortcuts="">
                    {action.shortcut.key ? <Shortcut shortcut={action.shortcut} /> : null}
                  </div>
                </Command.Item>
              ))}
            </Command.Group>
          </Command.List>
          <Command.Input placeholder="Search for actions..." />
        </Command>
      </Popover.Content>
    </Popover.Root>
  );
}

function Shortcut({ shortcut }: { shortcut: main.Shortcut }) {
  return (
    <div>
      {shortcut.super ? <kbd key="⌘">⌘</kbd> : null}
      {shortcut.ctrl ? <kbd key="⌃">⌃</kbd> : null}
      {shortcut.alt ? <kbd key="⌥">⌥</kbd> : null}
      {shortcut.shift ? <kbd key="⇧">⇧</kbd> : null}
      <kbd key={shortcut.key}>{shortcut.key.toUpperCase()}</kbd>
    </div>
  );
}
