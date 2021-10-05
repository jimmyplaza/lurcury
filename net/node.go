package main

import (
        "context"
        "crypto/rand"
        "fmt"
        "io"
        "log"
        mrand "math/rand"
        libp2p "github.com/libp2p/go-libp2p"
        crypto "github.com/libp2p/go-libp2p-crypto"
        host "github.com/libp2p/go-libp2p-host"
        ma "github.com/multiformats/go-multiaddr"
)

func makeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {
        var r io.Reader
        if randseed == 0 {
                r = rand.Reader
        } else {
                r = mrand.New(mrand.NewSource(randseed))
        }
        priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
        if err != nil {
                return nil, err
        }
        opts := []libp2p.Option{
                libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
                libp2p.Identity(priv),
        }
        if !secio {
                opts = append(opts, libp2p.NoSecurity)
        }
        basicHost, err := libp2p.New(context.Background(), opts...)
        if err != nil {
                return nil, err
        }
        hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))
        addr := basicHost.Addrs()[0]
        fullAddr := addr.Encapsulate(hostAddr)
        log.Printf("peer info %s\n", fullAddr)
        if secio {
		                log.Printf("Now run \"./echo -l %d -d %s -secio\" on a different terminal\n", listenPort+1, fullAddr)
				        } else {
						                log.Printf("Now run \"./echo -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
								        }
        return basicHost, nil
}
