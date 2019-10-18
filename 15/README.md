- https://learn.hashicorp.com/consul/day-0/containers-guide#configure-and-run-a-consul-server
```
docker run \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    --name=badger \
    consul:1.6.1 agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
```

- https://stackoverflow.com/questions/28802525/unable-to-generate-gpg-keys-in-linux
- https://easyengine.io/tutorials/linux/gpg-keys/
```
gpg-agent --daemon --use-standard-socket --pinentry-program /usr/bin/pinentry-curses
sudo rngd -r /dev/urandom
sudo gpg --gen-key
gpg --export -a -o mypublickey.txt user@replaceurmail.com
gpg --export-secret-key -a -o myprivatekey.txt user@replaceurmail.com
```
```
gpg-agent (GnuPG) 2.2.4
libgcrypt 1.8.1
Copyright (C) 2017 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
```
```
rngd 5
Copyright 2001-2004 Jeff Garzik
Copyright (c) 2001 by Philipp Rumpf
This is free software; see the source for copying conditions.  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
```
```
gpg (GnuPG) 2.2.4
libgcrypt 1.8.1
Copyright (C) 2017 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: /home/zzh/.gnupg
Supported algorithms:
Pubkey: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
Cipher: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
        CAMELLIA128, CAMELLIA192, CAMELLIA256
Hash: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
Compression: Uncompressed, ZIP, ZLIB, BZIP2
```

- crypt v0.0.1
```
https://github.com/xordataexchange/crypt/releases/download/v0.0.1/crypt-0.0.1-darwin-amd64
https://github.com/xordataexchange/crypt/releases/download/v0.0.1/crypt-0.0.1-linux-amd64
```
```
crypt set --backend consul --endpoint 127.0.0.1:8500 --keyring mypublickey.txt key config.json
crypt get --backend consul --endpoint 127.0.0.1:8500 --keyring myprivatekey.txt key
```