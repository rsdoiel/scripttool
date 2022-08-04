#!/bin/bash
EXT=".osf"
NAME=$(basename "$1" "$EXT")
cp "$1" document.xml
zip "${NAME}.fadein" document.xml
rm document.xml
