#!/usr/bin/env python3
"""Post-build fix for sky-log: removes multi-return FFI wrappers and their references."""
import re
import os

os.chdir(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'sky-out'))

wrapper_files = [f for f in os.listdir('.') if f.startswith('sky_ffi_') and f.endswith('.go')]

# Step 1: Remove multi-return wrapper functions from FFI wrapper files
for wf in wrapper_files:
    with open(wf, 'r') as f:
        content = f.read()

    lines = content.split('\n')
    output = []
    i = 0
    removed = 0

    while i < len(lines):
        line = lines[i]
        m = re.match(r'^func (Sky_\w+)\([^)]*\)\s*\(', line)
        if m:
            func_end = i
            brace_count = line.count('{') - line.count('}')
            while brace_count > 0 and func_end < len(lines) - 1:
                func_end += 1
                brace_count += lines[func_end].count('{') - lines[func_end].count('}')
            removed += 1
            i = func_end + 1
            while i < len(lines) and lines[i].strip() == '':
                i += 1
            continue
        output.append(line)
        i += 1

    with open(wf, 'w') as f:
        f.write('\n'.join(output))

    if removed > 0:
        print(f'  {wf}: removed {removed} multi-return functions')

# Step 2: Iteratively remove undefined references from main.go
for iteration in range(5):
    # Collect all defined functions and vars from all Go files
    defined_symbols = set()
    for f in ['main.go'] + wrapper_files:
        try:
            with open(f) as fh:
                for line in fh:
                    m = re.match(r'^func (\w+)\(', line)
                    if m:
                        defined_symbols.add(m.group(1))
                    m = re.match(r'^var (\w+)\b', line)
                    if m:
                        defined_symbols.add(m.group(1))
        except:
            pass

    with open('main.go', 'r') as f:
        lines = f.read().split('\n')

    output = []
    i = 0
    removed_total = 0

    while i < len(lines):
        line = lines[i]

        # Check for func declaration
        m = re.match(r'^func (\w+)\(', line)
        if m:
            func_start = i
            func_end = i
            brace_count = line.count('{') - line.count('}')
            while brace_count > 0 and func_end < len(lines) - 1:
                func_end += 1
                brace_count += lines[func_end].count('{') - lines[func_end].count('}')

            body = '\n'.join(lines[func_start:func_end+1])
            # Find references to Sky_ or other symbols in function call/return position
            refs = set(re.findall(r'\b(Sky_\w+)\b', body))
            missing = [ref for ref in refs if ref not in defined_symbols]

            if missing:
                removed_total += 1
                i = func_end + 1
                while i < len(lines) and lines[i].strip() == '':
                    i += 1
                continue

        # Check for var declarations
        vm = re.match(r'^var \w+ = (\w+)\b', line.rstrip())
        if vm:
            ref = vm.group(1)
            if ref not in defined_symbols:
                removed_total += 1
                i += 1
                continue

        output.append(line)
        i += 1

    with open('main.go', 'w') as f:
        f.write('\n'.join(output))

    if removed_total == 0:
        break
    print(f'  main.go pass {iteration+1}: removed {removed_total} declarations')

print('Post-build fix complete.')
