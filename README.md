# chanme

Small Blockchain PoC Written in Go for Educational Purpose

- https://github.com/lhartikk/naivechain
- https://www.igvita.com/2014/05/05/minimum-viable-block-chain/
- https://github.com/izqui/blockchain
- https://github.com/bitcoinbook/bitcoinbook/tree/second_edition
- http://bitcoin.peryaudo.org/index.html
- https://github.com/jashmenn/bitcoin-reading-list
- https://www.machu.jp/diary/20080302.html#p01
- http://www.righto.com/2014/02/bitcoins-hard-way-using-raw-bitcoin.html


#### run server
```
wbs -c wbs.toml
```

#### get blocks
```
http http://localhost:8511/blocks
```

#### create block
```
http POST http://localhost:8511/blocks/add data=ahodfaisdfasd
```
