# Stepping Up from Bash in DevOps: Python, Perl, Lua, or Go?

In DevOps, our journey often starts at the command line, where tasks are executed with precision, much like C-style logic, finally. Let's explore this with a simple task: reading and writing files line by line, a fundamental operation in CI/CD processes. Below, we examine how this task translates across different languages, moving from Bash to more advanced scripting environments.

![DevOpsNextStep](DevOpsNextStep.png)

## Starting Point: Bash Efficiency

We often leverage Bash for its efficiency in handling Unix/Linux system and network tasks—perfect for quick scripting and initial automation. But what comes next?

## Python: The Popular Choice

```python
def main():
    filename = '/etc/os-release'

    try:
        with open(filename, 'r') as file:
            line_number = 1
            for line in file:
                line = line.strip()
                print(f"{line_number}: {line}")
                line_number += 1
    except FileNotFoundError:
        print(f"Unable to open the file: {filename}")
    except Exception as e:
        print(f"An error occurred: {str(e)}")

if __name__ == "__main__":
    main()
```

Many opt for Python for its vast libraries and readability, which excel in complex automation and data applications. However, should we limit our tools to what's familiar?

## Consider Perl

```perl
use strict;
use warnings;

sub main {
    my $filename = '/etc/os-release';
    open(my $file, '<', $filename) or die "Unable to open the file '$filename': $!";

    my $line_number = 1;
    while (my $line = <$file>) {
        chomp $line;
        print "$line_number: $line\n";
        $line_number++;
    }

    close $file;
}

main();
```

Perl's syntax is close to C and excels in text processing—ideal for those bridging scripting with system-level programming.

## Check Lua

```lua
function main()
    local filename = "/etc/os-release"
    local file, err = io.open(filename, "r")
    if not file then
        error("Unable to open the file: " .. err)
    end

    local line_number = 1
    for line in file:lines() do
        line = line:gsub("\n$", "")
        print(line_number .. ": " .. line)
        line_number = line_number + 1
    end

    file:close()
end

main()
```

Lua offers lightweight, flexible scripting with minimal memory needs, perfect for developing domain-specific languages and configurations across platforms.

## Then There's Go

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    filename := "/etc/os-release"

    file, err := os.Open(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to open the file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineNumber := 1

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("%d: %s\n", lineNumber, line)
        lineNumber++
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading the file: %v\n", err)
        os.Exit(1)
    }
}
```

Go's compiled nature and concurrency handling make it unmatched for high-performance backends. Plus, there's no need to ship a runtime environment or build a Docker image.

## Let's remember finally for comparison, the C

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    FILE *file;
    char buffer[1024];
    int line_number = 1;

    file = fopen("/etc/os-release", "r");
    if (file == NULL) {
        perror("Unable to open the file");
        exit(EXIT_FAILURE);
    }

    while (fgets(buffer, sizeof(buffer), file)) {
        printf("%d: %s", line_number++, buffer);
    }

    fclose(file);
    
    return 0;
}
```

Ultimately, this is where we'll end up (either at Go or Rust, but it doesn't really matter in this case).

## Ranking - Language Similarity to C

1. Perl
2. Go
3. Lua
4. Python
5. Bash

This ranking reflects not just syntactical similarities but also each language's design philosophy, balancing control with abstraction.

## Final Thoughts

Choosing the right tool goes beyond familiarity. Consider where you want your tech stack to grow. Perl, Lua, or Go might be your next steps if you're aiming beyond Python's ease.

> **What has been your experience with these languages? Do you find one more advantageous for DevOps tasks over the others?**

## Links

- [C](https://www.open-std.org/jtc1/sc22/wg14/)
- [Perl](https://www.perl.org)
- [Go](https://go.dev)
- [Lua](https://www.lua.org)
- [Python](https://www.python.org)
- [Bash](https://www.gnu.org/software/bash/)

[#DevOps](#DevOps) [#Programming](#Programming)  [#TechnologyChoices](#TechnologyChoices) [#CareerDevelopment](#CareerDevelopment)
