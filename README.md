# GoRe — a Go regexp toolbelt

Written out of annoyance with the many different regexp syntaxes and escaping requirements in sed, grep, etc.

Intended to cover the most common use-cases without any magical flags.

## Install

* Install Go
* `go get nevill.io/gore`

## Commands

All commands traverse directories recursively.

All patterns use RE2 syntax (https://golang.org/s/re2syntax).

All patterns are applied line-by-line.

All replacements may reference matched groups with `$name` or `${name}` (https://golang.org/pkg/regexp/#Regexp.Expand)

```
find|list|ls <pattern> [<path...>]
  find file and directories with names matching pattern.  
search <pattern> [<path...>]
  list files with lines matching the pattern.
print|p|grep <pattern> [<path...>]
  print lines matching pattern.
replace|re|sub|s <pattern> <replacement> <path>
  replace occurrences of pattern with replacement.
delete|d <pattern> <path...>
  delete liens matching pattern.
rename <pattern> <replacement> <path>
  replace occurrences of pattern in file and directory names with replacement.
```

## License

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
