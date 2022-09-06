import React from "react";
import "./raycast.scss";
import { Command } from "cmdk";
import { TerminalIcon, RaycastLightIcon } from "./icons";
import * as App from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";
import * as runtime from "../../wailsjs/runtime/runtime";
import * as Popover from "@radix-ui/react-popover";
import { ReactSVG } from "react-svg";

// window.onblur = () => {
//   runtime.WindowHide();
// };

export function RaycastCMDK() {
  const inputRef = React.useRef<HTMLInputElement | null>(null);
  const listRef = React.useRef(null);
  const [query, setQuery] = React.useState("");
  const [focusedValue, setFocusedValue] = React.useState("");
  const [pages, setPages] = React.useState<
    {
      items: Record<string, main.SearchItem>;
      query?: string;
      focusedValue?: string;
    }[]
  >([]);

  const currentPage = pages[pages.length - 1];
  const focusedItem = currentPage.items && currentPage.items[focusedValue];

  function buildKey(searchItem: main.SearchItem) {
    return [searchItem.title, ...(searchItem.keywords || [])]
      .join(" ")
      .toLowerCase();
  }

  // Add root items on first load
  React.useEffect(() => {
    App.RootItems().then((items) => {
      const itemMap = Object.fromEntries(
        items.map((item) => [buildKey(item), item])
      );
      setPages([{ items: itemMap }]);
    });
  }, []);

  // restore the query when the current page change
  React.useEffect(() => {
    setQuery(currentPage.query || "");
    setFocusedValue(currentPage.focusedValue || "");
  }, [currentPage]);

  function handleAction(action: main.Action) {
    runtime.LogDebug(`Handling Action: ${JSON.stringify(action)}`);
    switch (action.type) {
      case "open":
        App.OpenFile(action.path);
        return;
      case "copy-to-clipboard":
        App.CopyToClipboard(action.content);
        return;
      case "run-script":
        App.RunScript(action.path, []);
        return;
      case "run-command":
        App.RunListCommand(action.path).then((items) => {
          const itemMap = Object.fromEntries(
            items.map((item) => [buildKey(item), item])
          );
          currentPage.query = query;
          currentPage.focusedValue = focusedValue;
          setPages([...pages, { items: itemMap }]);
          return;
        });
    }
  }

  // Listen for commands
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
    inputRef?.current?.focus();
  }, []);

  return (
    <div className="raycast">
      <Command
        shouldFilter={true}
        value={focusedValue}
        onKeyDown={(e) => {
          if (e.key === "Escape") {
            setPages(pages.slice(0, -1));
          } else if (e.key === "Tab") {
            if (focusedItem?.fill) {
              setQuery(focusedItem.fill);
            }
          }
        }}
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
          <Command.Empty>No Result Found.</Command.Empty>
          <Command.Group heading="Results">
            {Object.entries(currentPage.items || {}).map(([key, item]) => (
              <Item
                item={item}
                key={key}
                value={key}
                onSelect={() => {
                  handleAction(item.actions[0]);
                }}
              />
            ))}
          </Command.Group>
        </Command.List>

        <div cmdk-raycast-footer="">
          <div cmd-raycast-submenu="">
            <RaycastMenu listRef={listRef} inputRef={inputRef} />
          </div>

          {focusedItem ? (
            <>
              <button
                cmdk-raycast-subcommand-trigger=""
                onClick={() => {
                  handleAction(focusedItem.actions[0]);
                }}
              >
                {focusedItem.actions[0]?.title || ""}
                <kbd>↵</kbd>
              </button>
              <hr />
              <SubCommand
                listRef={listRef}
                focusedItem={focusedItem}
                inputRef={inputRef}
                onAction={(action) => handleAction(action)}
              />
            </>
          ) : null}
        </div>
      </Command>
    </div>
  );
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
    <Command.Item
      value={value}
      onSelect={onSelect}
      onPointerMove={(e) => e.preventDefault()}
    >
      {item.icon_src ? (
        <img width={24} height={24} src={item.icon_src} />
      ) : (
        <TerminalIcon />
      )}
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
              <Command.Item className="test">
                <RaycastIcon src="/raycast/icon-question-mark-circle-16.svg" />
                Manual
              </Command.Item>
              <Command.Item>
                <RaycastIcon src="/raycast/icon-moon-16.svg" />
                Switch to Dark Mode
              </Command.Item>

              <Command.Item>
                <RaycastIcon src="/raycast/icon-raycast-logo-neg-16.svg" />
                About Raycast
              </Command.Item>
              <Command.Item>
                <RaycastIcon src="/raycast/icon-cog-16.svg" />
                Preferences
              </Command.Item>

              <Command.Item
                onSelect={() => {
                  runtime.Quit();
                }}
              >
                <RaycastIcon src="/raycast/icon-logout-16.svg" />
                Quit Raycast
              </Command.Item>
            </Command.Group>
          </Command.List>
          <Command.Input placeholder="Search for actions..." />
        </Command>
      </Popover.Content>
    </Popover.Root>
  );
}

function RaycastIcon({ src }: { src: string }) {
  return (
    <ReactSVG
      src={src}
      beforeInjection={(svg) => {
        const rect = svg.querySelector("rect");
        rect?.setAttribute("fill", "none");
      }}
    />
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
                  <RaycastIcon src={action.icon} />
                  {action.title}
                  {/* <div cmdk-raycast-submenu-shortcuts="">
                    {action.shortcut.key ? (
                      <Shortcut shortcut={action.shortcut} />
                    ) : null}
                  </div> */}
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

// function Shortcut({ shortcut }: { shortcut: main.Shortcut }) {
//   return (
//     <div>
//       {shortcut.super ? <kbd key="⌘">⌘</kbd> : null}
//       {shortcut.ctrl ? <kbd key="⌃">⌃</kbd> : null}
//       {shortcut.alt ? <kbd key="⌥">⌥</kbd> : null}
//       {shortcut.shift ? <kbd key="⇧">⇧</kbd> : null}
//       <kbd key={shortcut.key}>{shortcut.key.toUpperCase()}</kbd>
//     </div>
//   );
// }
