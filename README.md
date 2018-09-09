# pipesum
Sums numbers from stdin

# Installation

`go get -u github.com/carlpett/pipesum`, or download binary from release page.

# Usage

```
usage: pipesum [<flags>]

Flags:
      --help             Show context-sensitive help (also try --help-long and --help-man).
  -n, --numeric          Numeric (default)
  -h, --human-readable   Human readable (prefixes like M, K, etc)
      --human-iec        Use IEC prefixes for human-readable format (Mi, Ki, etc)
      --human-unit=UNIT  Specify a unit to follow the prefix
```

# Examples

Standard numeric sum
```bash
$ echo -e '10\n20\n12' | pipesum
42
```

Human-readable
```bash
echo -e '10M\n200K' | pipesum -h
10.2M
```

Human-readable with IEC prefixes (base 2)
```bash
echo -e '10Mi\n200Ki' | pipesum -h --human-iec
10.2Mi
```

Human-readable with unit
```bash
echo -e '10MB\n200KB' | pipesum -h --human-unit B
10.2MB
```
