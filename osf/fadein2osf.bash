#!/bin/bash
EXT=".fadein"
NAME=$(basename "$1" "$EXT")
unzip "$1" document.xml
mv document.xml "${NAME}.osf"
