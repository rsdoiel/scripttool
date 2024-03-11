#!/bin/bash

START=$(pwd)
cd testdata
curl -L -O "https://fountain.io/_downloads/Big-Fish.fountain"
curl -L -O "https://fountain.io/_downloads/Big-Fish.fdx"
#curl -L -O "https://fountain.io/_downloads/Big-Fish.pdf"
curl -L -O "https://fountain.io/_downloads/Brick-&-Steel.fountain"
curl -L -O "https://fountain.io/_downloads/Brick-&-Steel.fdx"
#curl -L -O "https://fountain.io/_downloads/Brick-&-Steel.pdf"
curl -L -O "https://fountain.io/_downloads/The-Last-Birthday-Card.fountain"
curl -L -O "https://fountain.io/_downloads/The-Last-Birthday-Card.fdx"
#curl -L -O "https://fountain.io/_downloads/The-Last-Birthday-Card.pdf"
cd "$START"
