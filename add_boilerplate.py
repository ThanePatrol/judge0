#!/usr/bin/env python3

import sys

# file path as first argument
# second argument is line number to be added with boilerplate code
# fourth agument is the line number with language name
# third argument is the boilerplate code
# fifth argument is the language name
file = sys.argv[1]
code_line = int(sys.argv[2])
boilerplate = sys.argv[3]
lang_line = int(sys.argv[4])
lang_name = sys.argv[5]
lines = []

with open(file, 'r') as f:
    lines = f.readlines()

lines[code_line].replace('\n', '')
lines[code_line] += boilerplate + '\n'
lines[lang_line].replace('\n', '')
lines[lang_line] += lang_name + '\n'

with open("tmp.js", 'w') as f:
    f.writelines(lines)

