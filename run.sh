#!/bin/bash

DIRNAME=$( dirname "$(readlink -f -- "$0")" )

export DESKTOP_SESSION=xfce
export XDG_DATA_DIRS=$DIRNAME/assets:$XDG_DATA_DIRS

exec wails dev
