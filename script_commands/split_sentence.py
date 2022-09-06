#!/usr/bin/env python3

# @raycast.schemaVersion 1
# @raycast.title Split Sentence
# @raycast.description description
# @raycast.packageName package
# @raycast.mode command

import json
import sys

query = sys.stdin.read()

for word in query.split():
    print(
        json.dumps(
            {
                "title": word,
                "subtitle": "truc",
                "accessory_title": "test",
                "icon": "https://...",
                "keywords": [],
                "fill": "test",
                "actions": [
                    {"title": "word", "type": "copy-to-clibpboard", "content": "test"}
                ],
            }
        )
    )
