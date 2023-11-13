#!/usr/bin/env bash

# does builds for the various different languages
# each language needs boilerplate

# python
# create working copy of file
cp editor.js tmp.js
boilerplate="\`class Solution(object):\n    def solve(self, A):\n        return 0\n\`,"
sed "9s|$| $boilerplate|" tmp.js > t.js && mv t.js tmp.js
sed "13s|$| python|" tmp.js > t.js && mv t.js tmp.js
#awk -v n=9 -v str="$boilerplate" 'NR == n {print $0 str; next} {print}' tmp.js
node_modules/.bin/rollup tmp.js -f iife -o resources/editor.python.js -p @rollup/plugin-node-resolve
#rm tmp.js
