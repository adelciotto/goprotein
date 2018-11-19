# goprotein

WIP tools for efficiently encoding DNA and translating into proteins.

## dnapack

Packs a plain text DNA sequence into a more compact binary serialisation format.

### Build

```
make
```

### Usage

```
./pack --input-file=<input-dna-file.txt> --output-file=<output-file.dna>
```

* `input-file` is required.
* `output-file` is optional. Path and name of the input-file will default to value of `input-file`.
