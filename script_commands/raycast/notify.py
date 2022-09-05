#!/usr/bin/env python3

# @raycast.schemaVersion 1
# @raycast.title Say Something!
# @raycast.packageName Why?
# @raycast.mode filter
# @raycast.description This makes no sense...

import sys
import json
import subprocess

for sentence in ["Hello World!", "Raycast is Awesome!"]:
    print(
        json.dumps(
            {
                "title": sentence,
                "actions": [
                    {
                        "type": "callback",
                        "title": "Copy",
                        "params": {"content": sentence},
                    }
                ],
            }
        )
    )
