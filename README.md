# b2sum

This simple program operates similar to command line tools like sha256sum. Prints a 64 byte blake2b checksum encoded in base64 to standard output. Works with files or standard input.

#### Operation

  - `b2sum data.bin`
  - `cat data.bin | b2sum`
