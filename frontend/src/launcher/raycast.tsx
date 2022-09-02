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
import { IconSource, isDesktopEntryResponse, isFillResponse, isMimeIcon, isNamedIcon as isSystemIcon, isUpdateResponse, SearchResult } from "./response";

export function RaycastCMDK() {
  const inputRef = React.useRef<HTMLInputElement | null>(null);
  const listRef = React.useRef(null);
  const [focusedItem, setFocusedItem] = React.useState<SearchResult>();
  const [query, setQuery] = React.useState("")
  const [items, setItems] = React.useState<SearchResult[]>([])

  React.useEffect(() => {
    function listener(e: KeyboardEvent) {
      if (!focusedItem || e.key !== "Tab") {
        return;
      }
      runtime.LogDebug(`Complete: ${focusedItem.name}`)
      e.preventDefault()
      App.Complete(focusedItem?.id)
    }

    document.addEventListener("keydown", listener);

    return () => {
      document.removeEventListener("keydown", listener);
    };
  }, [focusedItem]);

  useEffect(() => {
    runtime.EventsOn("update", (update) => {
      if (isUpdateResponse(update)) {
        const items = update.Update
        setItems(items)
        if (items.length > 0) {
          setFocusedItem(items[0])
        }
        return;
      }
      if (isDesktopEntryResponse(update)) {
        App.OpenApp(update.DesktopEntry.path)
        return
      }
      if (isFillResponse(update)) {
        setQuery(update.Fill)
        return;
      }
      return runtime.LogError(`No handler for : ${JSON.stringify(update)}`)
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
      <Command shouldFilter={false} value={focusedItem?.id.toString() || "-1"} onValueChange={(v) => {
        const index = parseInt(v)
        setFocusedItem(items[index])
      }}>
        <div cmdk-raycast-top-shine="" />
        <Command.Input
          ref={inputRef}
          value={query}
          onValueChange={setQuery}
          autoFocus
          placeholder="Search for apps and commands..."
        />
        <hr cmdk-raycast-loader="" />
        <Command.List ref={listRef} >
          <Command.Empty>{query ? "No Result Found." : "Type Something !"}</Command.Empty>
          {items.length > 0 ?
          <Command.Group heading="Results">
            {items.map(item => 
              <Item key={item.name} item={item} onSelect={() => {
                App.Activate(item.id)
              }}>
                <Logo>
                  {item.icon ? <ItemIcon icon={item.icon}/> : <TerminalIcon/>}
                </Logo>{item.name}
              </Item>
            )}
          </Command.Group>
          : null}
        </Command.List>

        <div cmdk-raycast-footer="">
          <RaycastLightIcon />

          <button cmdk-raycast-open-trigger="">
            Activate
            <kbd>↵</kbd>
          </button>

        </div>
      </Command>
    </div>
  );
}

function ItemIcon({icon}: {icon: IconSource}) {
  if (isMimeIcon(icon)) {
    return <img width={24} height={24} src={`/${icon.Mime}.svg?type=mime`}/>
  }
  if (isSystemIcon(icon)) {
    return <img width={24} height={24} src={`/${icon.Name}.png?type=system`}/>
  }
  return <TerminalIcon/>
}

function Item({
  children,
  item,
  onSelect}: {
  children: React.ReactNode;
  onSelect: () => void;
  item: SearchResult;
}) {
  return (
    <Command.Item value={item.id.toString()} onSelect={onSelect}>
      {children}
      <span cmdk-raycast-meta="">{item.description}</span>
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
