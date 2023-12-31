#! /bin/bash

BUILD_DIR=./build
BLURHASH_CLI=${BUILD_DIR}/blurhash_cli

HASH=$(${BLURHASH} encode --filepath ../imgs/salad.png --xcomp 4 --ycomp 3)
${BLURHASH_CLI} decode --hash "${HASH}" --width 1024 --height 800 --punch 3 --dest ${BUILD_DIR}/output.png