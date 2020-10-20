#! /bin/sh

SRC_DIR="assets/styles"
OUT_DIR="static/styles"

sass $SRC_DIR/main.scss $OUT_DIR/main.css
echo "Successfully compiled assets"
