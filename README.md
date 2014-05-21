# {u,}varint-{en,de}code -- Variable integer encoding/decoding tools

This repository provides command-line tools to interact with the
varint/uvarint format, as used by
[Protocol Buffer](https://developers.google.com/protocol-buffers/docs/overview)
and Go's [encoding/binary](http://golang.org/pkg/encoding/binary/)
package.

The `{u,}varint-encode` tools take integers on the command line and
output binary. Use `--` to pass negative integers.

The `{u,}varint-decode` tools read standard input and output integers,
one per line.

Examples:

``` console
$ uvarint-encode 42 | hexdump -x
0000000    002a                                                        
0000001

$ varint-encode -- -42 | hexdump -x
0000000    0053                                                        
0000001

$ uvarint-encode 42 | uvarint-decode
42

$ varint-encode 42 | varint-decode
42

$ varint-encode -- -42 | varint-decode
-42
```
