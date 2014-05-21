  $ uvarint-encode 42 | hexdump -x
  0000000    002a                                                        
  0000001

  $ uvarint-encode -- -42
  uvarint-encode: cannot parse number: strconv.ParseUint: parsing "-42": invalid syntax
  [1]

  $ varint-encode 42 | hexdump -x
  0000000    0054                                                        
  0000001

  $ varint-encode -- -42 | hexdump -x
  0000000    0053                                                        
  0000001
