import React, { useEffect } from "react";
import * as Popover from "@radix-ui/react-popover";
import "./raycast.scss";
import { Command } from "cmdk";
import {
  Logo,
  FinderIcon,
  StarIcon,
  WindowIcon,
  TerminalIcon,
} from "./icons";
import * as App from "../../wailsjs/go/main/App"
import * as runtime from "../../wailsjs/runtime/runtime"
import { isUpdateResponse, SearchResult } from "./response";

export function RaycastCMDK() {
  const inputRef = React.useRef<HTMLInputElement | null>(null);
  const listRef = React.useRef(null);
  const [value, setValue] = React.useState("linear");
  const [query, setQuery] = React.useState("")
  const [items, setItems] = React.useState<SearchResult[]>([])
  runtime.LogDebug(JSON.stringify(items))

  useEffect(() => {
    runtime.EventsOn("update", (event) => {
      if (isUpdateResponse(event)) {
        setItems(event.Update)
        return;
      }
      return runtime.LogError(`No handler for : ${JSON.stringify(event)}`)
    })
    return () => {
      runtime.EventsOff("update")
    }
  }, [])


  useEffect(() => {
    App.Search(query);
  }, [query]);

  React.useEffect(() => {
    inputRef?.current?.focus();
  }, []);

  return (
    <div className="raycast">
      <Command shouldFilter={false} value={value} onValueChange={(v) => setValue(v)}>
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
          <Command.Empty>{query ? "No Result Found." : "Type Something !"}</Command.Empty>
            {items.map(item => 
              <Item item={item} onSelect={() => {}}>
                <Logo>
                  <TerminalIcon/>
                </Logo>{item.name}
              </Item>
            )}
        </Command.List>

        <div cmdk-raycast-footer="">
          <RaycastLightIcon />

          <button cmdk-raycast-open-trigger="">
            Open Application
            <kbd>↵</kbd>
          </button>

          <hr />

          <SubCommand
            listRef={listRef}
            selectedValue={value}
            inputRef={inputRef}
          />
        </div>
      </Command>
    </div>
  );
}

function Item({
  children,
  item,
  onSelect,
  isCommand = false,
}: {
  children: React.ReactNode;
  onSelect: () => void;
  item: SearchResult;
  isCommand?: boolean;
}) {
  return (
    <Command.Item value={item.name} onSelect={onSelect}>
      {children}
      <span cmdk-raycast-meta="">{item.category_icon.Name}</span>
    </Command.Item>
  );
}

function SubCommand({
  inputRef,
  listRef,
  selectedValue,
}: {
  inputRef: React.RefObject<HTMLInputElement>;
  listRef: React.RefObject<HTMLElement>;
  selectedValue: string;
}) {
  const [open, setOpen] = React.useState(false);

  React.useEffect(() => {
    function listener(e: KeyboardEvent) {
      if (e.key === "k" && e.metaKey) {
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
        <kbd>⌘</kbd>
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
            <Command.Group heading={selectedValue}>
              <SubItem>
                <WindowIcon />
                Open Application
              </SubItem>
              <SubItem>
                <FinderIcon />
                Show in Finder
              </SubItem>
              <SubItem>
                <FinderIcon />
                Show Info in Finder
              </SubItem>
              <SubItem>
                <StarIcon />
                Add to Favorites
              </SubItem>
            </Command.Group>
          </Command.List>
          <Command.Input placeholder="Search for actions..." />
        </Command>
      </Popover.Content>
    </Popover.Root>
  );
}

function SubItem({
  children,
  shortcut,
}: {
  children: React.ReactNode;
  shortcut?: string;
}) {
  return (
    <Command.Item>
      {children}
      <div cmdk-raycast-submenu-shortcuts="">
        {shortcut?.split(" ").map((key) => {
          return <kbd key={key}>{key}</kbd>;
        })}
      </div>
    </Command.Item>
  );
}

function RaycastLightIcon() {
  return (
    <svg
      width="1024"
      height="1024"
      viewBox="0 0 1024 1024"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M934.302 511.971L890.259 556.017L723.156 388.902V300.754L934.302 511.971ZM511.897 89.5373L467.854 133.583L634.957 300.698H723.099L511.897 89.5373ZM417.334 184.275L373.235 228.377L445.776 300.923H533.918L417.334 184.275ZM723.099 490.061V578.209L795.641 650.755L839.74 606.652L723.099 490.061ZM697.868 653.965L723.099 628.732H395.313V300.754L370.081 325.987L322.772 278.675L278.56 322.833L325.869 370.146L300.638 395.379V446.071L228.097 373.525L183.997 417.627L300.638 534.275V634.871L133.59 467.925L89.4912 512.027L511.897 934.461L555.996 890.359L388.892 723.244H489.875L606.516 839.892L650.615 795.79L578.074 723.244H628.762L653.994 698.011L701.303 745.323L745.402 701.221L697.868 653.965Z"
        fill="#FF6363"
      />
    </svg>
  );
}

function RaycastDarkIcon() {
  return (
    <svg
      width="1024"
      height="1024"
      viewBox="0 0 1024 1024"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M301.144 634.799V722.856L90 511.712L134.244 467.804L301.144 634.799ZM389.201 722.856H301.144L512.288 934L556.34 889.996L389.201 722.856ZM889.996 555.956L934 511.904L512.096 90L468.092 134.052L634.799 300.952H534.026L417.657 184.679L373.605 228.683L446.065 301.144H395.631V628.561H723.048V577.934L795.509 650.395L839.561 606.391L723.048 489.878V389.105L889.996 555.956ZM323.17 278.926L279.166 322.978L326.385 370.198L370.39 326.145L323.17 278.926ZM697.855 653.61L653.994 697.615L701.214 744.834L745.218 700.782L697.855 653.61ZM228.731 373.413L184.679 417.465L301.144 533.93V445.826L228.731 373.413ZM578.174 722.856H490.07L606.535 839.321L650.587 795.269L578.174 722.856Z"
        fill="#FF6363"
      />
    </svg>
  );
}
