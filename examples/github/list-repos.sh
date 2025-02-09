#!/usr/bin/env bash

set -eo pipefail

if [ -n "$1" ]; then
    ENDPOINT=/users/$1/repos
else
    ENDPOINT=/user/repos
fi

gh api "$ENDPOINT" --paginate --cache 3h --jq '.[] |
    {
        title: .name,
        subtitle: (.description // ""),
        accessories: [
            "\(.stargazers_count) *"
        ],
        actions: [
            {type: "open-url", url: .html_url},
            {
                type: "run-command",
                command: "list-prs",
                title: "List Pull Requests",
                shortcut: "ctrl+p",
                with: {repository: .full_name}
            },
            {
                type: "run-command",
                command: "view-readme",
                title: "View Readme",
                shortcut: "ctrl+r",
                with: {repository: .full_name}
            }
        ]
    }
' | sunbeam query --slurp '{
    type: "list",
    list: {
        items: .
    }
}'
