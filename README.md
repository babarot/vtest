vtest
=====

Unix `test` command for Version comparison

## Installation

Download the binary from [GitHub Releases][release] and drop it in your `$PATH`.

- [Darwin / Mac][release]
- [Linux][release]

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

[MIT][license]

## Author

[b4b4r07][author]

[author]:  https://tellme.tokyo
[release]: https://github.com/b4b4r07/vtest/releases/latest
[license]: https://b4b4r07.mit-license.org
