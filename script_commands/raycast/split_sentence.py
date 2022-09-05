#!/usr/bin/env python3

# @raycast.schemaVersion 1
# @raycast.title Split Sentence
# @raycast.description description
# @raycast.packageName package
# @raycast.mode search

import json
import sys

query = sys.stdin.read()

for word in query.split():
    print(
        json.dumps(
            {
                "title": word,
                "actions": [
                    {"title": "word", "type": "copy-to-clibpboard", "params": {"content": "word"}}
                ],
            }
        )
    )
