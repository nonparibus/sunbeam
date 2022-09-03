.PHONY: pop-launcher
pop-launcher:
	just --justfile pop-launcher/justfile
	HOME=$(PWD)/AppDir/home just --justfile pop-launcher/justfile install

.PHONY: raycast
raycast:
	wails build
	mkdir -p AppDir/usr/bin/
	cp build/bin/raycast AppDir/usr/bin/raycast

.PHONY: appimage
appimage: pop-launcher raycast
	appimagetool AppDir
	chmod +x Raycast-x86_64.AppImage
