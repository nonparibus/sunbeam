import React, { useEffect } from "react";
import "./raycast.scss";
import { Command } from "cmdk";
import {
  TerminalIcon,
  RaycastLightIcon
} from "./icons";
import * as App from "../../wailsjs/go/main/App"
import * as runtime from "../../wailsjs/runtime/runtime"
import { IconSource, isDesktopEntryResponse, isFillResponse, isMimeIcon, isNamedIcon as isSystemIcon, isUpdateResponse, SearchResult } from "./response";

window.onblur = () => {
  runtime.WindowHide()
}

export function RaycastCMDK() {
  const inputRef = React.useRef<HTMLInputElement | null>(null);
  const listRef = React.useRef(null);
  const [focusedItem, setFocusedItem] = React.useState<SearchResult>();
  const [query, setQuery] = React.useState("")
  const [items, setItems] = React.useState<SearchResult[]>([])

  React.useEffect(() => {
    runtime.EventsOn("toggle", () => {
      runtime.WindowCenter()
      runtime.WindowShow()
    })    
    return () => {
      runtime.EventsOff("toggle")
    }
  }, [])

  React.useEffect(() => {
    function listener(e: KeyboardEvent) {
      if (e.key !== "Escape") {
        return;
      }
      runtime.WindowHide()
      setQuery("")
    }

    document.addEventListener("keydown", listener);

    return () => {
      document.removeEventListener("keydown", listener);
    };
  }, []);


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
                <Item key={item.name} item={item} />


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

function ItemIcon({ icon }: { icon: IconSource }) {
  if (icon.Name?.endsWith(".ico")) {
    // return <img width={24} height={24} src={`/${icon.Name}?type=file`} />
    return <TerminalIcon />
  }
  if (isMimeIcon(icon)) {
    return <img width={24} height={24} src={`/${icon.Mime}?type=mime`} />
  }
  if (isSystemIcon(icon)) {
    return <img width={24} height={24} src={`/${icon.Name}?type=system`} />
  }
  return <TerminalIcon />
}

function Item({
  item
}: {
  item: SearchResult;
}) {
  return (
    <Command.Item value={item.id.toString()} onSelect={() => App.Activate(item.id)}>
      {item.icon ? <ItemIcon icon={item.icon} /> : <TerminalIcon />}
      {item.name}

      <span cmdk-raycast-meta="">{item.description}</span>
    </Command.Item>
  );
}
