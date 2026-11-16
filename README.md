# GUID

Generates a GUID and automatically copies it to your system clipboard.

## Compile

```
go build
```

## Usage

```
guid [OPTIONS] [<number>]

Options:
-n <number>      The number of UUIDs to generate. (default 1).
                 This is the equivalent of 'guid <number>'.
-no-copy         Don't copy output to clipboard.
-no-print        Don't print output to stdout.
-v <version>     Determine the UUID version to use. 4 or 7 accepted. (default 4)
-h               Show this help message and exit.
```
