# goprotein

WIP tools for efficiently encoding DNA and translating into proteins.

## Build

```
make
```

## Usage

### Pack

Packs a file containing plain text DNA into a compact binary serialised format.

```
./pack --input-file=<input-dna-file.txt> --output-file=<output-file.dna>
```

* `input-file` is required.
* `output-file` is optional. Path and name of the input-file will default to value of `input-file`.

### Translate

Translates a file containing packed DNA into proteins. This program is a WIP and currently just prints the proteins to STDOUT.

```
./translate --input-file=<input-dna-file.dna>
```

* `input-file` is required.
