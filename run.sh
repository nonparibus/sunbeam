#!/bin/bash

DIRNAME=$( dirname "$(readlink -f -- "$0")" )

XDG_DATA_HOME=$DIRNAME/assets wails dev


