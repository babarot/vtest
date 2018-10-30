vtest
=====

Unix `test` command for Version comparison

## Installation

```console
$ go get github.com/b4b4r07/vtest
```

## Usage

```bash
#!/bin/bash

if vtest "0.12.40 > ${version}"; then
    : ok
else
    echo "need to be upgraded" >&2
    exit 1
fi
```

## License

MIT

## Author

b4b4r07
