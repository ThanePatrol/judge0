#!/usr/bin/env bash

# does builds for the various different languages
# each language needs boilerplate

# python
# create working copy of file
cp editor.js tmp.js
boilerplate="\`class Solution(object):\n    def solve(self, A):\n        return 0\n\`,"
./add_boilerplate.py tmp.js 8 "$boilerplate" 12 "python"

./node_modules/.bin/rollup tmp.js -f iife -o resources/editor.python.js -p @rollup/plugin-node-resolve
#rm tmp.js
