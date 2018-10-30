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

# Check ${version} is greater than the threshold or not.
if vtest "${version} < 0.12.40"; then
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
