  $ uvarint-encode 42 | uvarint-decode
  42

  $ varint-encode 42 | varint-decode
  42

  $ varint-encode -- -42 | varint-decode
  -42
