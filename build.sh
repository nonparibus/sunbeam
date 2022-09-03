#!/bin/sh

set -ex

DIRNAME=$( dirname "$(readlink -f -- "$0")" )
APPDIR="${DIRNAME}/AppDir"

# build pop-launcher
cd "${DIRNAME}/pop-launcher"
just
HOME=${APPDIR}/home just install

# Build raycast
wails build
mkdir -p "${APPDIR}/usr/bin"
cp -r build/bin/raycast "${APPDIR}/usr/bin"

appimagetool "${APPDIR}"
