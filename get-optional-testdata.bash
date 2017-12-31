#!/bin/bash

START=$(pwd)
cd testdata
curl -L -O "https://fountain.io/_downloads/Big%20Fish.fountain"
curl -L -O "https://fountain.io/_downloads/Big%20Fish.fdx"
curl -L -O "https://fountain.io/_downloads/Big%20Fish.pdf"
curl -L -O "https://fountain.io/_downloads/Brick%20&%20Steel.fountain"
curl -L -O "https://fountain.io/_downloads/Brick%20&%20Steel.fdx"
curl -L -O "https://fountain.io/_downloads/Brick%20&%20Steel.pdf"
curl -L -O "https://fountain.io/_downloads/The%20Last%20Birthday%20Card.fountain"
curl -L -O "https://fountain.io/_downloads/The%20Last%20Birthday%20Card.fdx"
curl -L -O "https://fountain.io/_downloads/The%20Last%20Birthday%20Card.pdf"
cd "$START"
