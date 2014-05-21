  $ uvarint-encode bad
  uvarint-encode: cannot parse number: strconv.ParseUint: parsing "bad": invalid syntax
  [1]

  $ varint-encode bad
  varint-encode: cannot parse number: strconv.ParseInt: parsing "bad": invalid syntax
  [1]

  $ printf '\377' | uvarint-decode
  uvarint-decode: truncated input
  [1]

  $ printf '\377' | varint-decode
  varint-decode: truncated input
  [1]
