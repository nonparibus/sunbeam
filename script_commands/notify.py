#!/usr/bin/env python3

# @raycast.schemaVersion 1
# @raycast.title Say Something!
# @raycast.packageName Why?
# @raycast.mode command
# @raycast.description This makes no sense...

import sys
import json
import subprocess

if len(sys.argv) > 1:
    sentence = sys.argv[1]
    print(
        {
            "type": "details",
            "markdown": sentence,
            "actions": [{"type": "copy-to-clipboard", "content": "sentence"}],
        }
    )
    sys.exit(0)

sentences = ["Hello World!", "Raycast is Awesome!"]

print(
    json.dumps(
        {
            "type": "list",
            "list_items": [
                {
                    "title": sentence,
                    "actions": [
                        {
                            "type": "callback",
                            "title": "Copy Sentence",
                            "args": [sentence],
                        },
                        {"type": "copy-to-clipboard", "title": "Copy to Clipboard", "content": "sentence"},
                    ],
                }
                for sentence in sentences
            ],
        }
    )
)
