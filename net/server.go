package main
import (
        "bufio"
        //"context"
        //"crypto/rand"
        "flag"
        "fmt"
        //"io/ioutil"
        "log"
        //golog "github.com/ipfs/go-log"
        //libp2p "github.com/libp2p/go-libp2p"
        //crypto "github.com/libp2p/go-libp2p-crypto"
        //host "github.com/libp2p/go-libp2p-host"
        //gologging "github.com/whyrusleeping/go-logging"
        //ma "github.com/multiformats/go-multiaddr"
        net "github.com/libp2p/go-libp2p-net"
        //peer "github.com/libp2p/go-libp2p-peer"
        //pstore "github.com/libp2p/go-libp2p-peerstore"
)

func doEcho(s net.Stream) error {
        buf := bufio.NewReader(s)
        str, err := buf.ReadString('\n')
        if err != nil {
                return err
        }

        log.Printf("read: %s\n", str)
        _, err = s.Write([]byte(str))
        return err
}

func stream(s net.Stream){
	fmt.Println("123")
}

func main() {
        secio := flag.Bool("secio", false, "enable secio")
        seed := flag.Int64("seed", 0, "set random seed for id generation")
        flag.Parse()
        ha, err := makeBasicHost(9999, *secio, *seed)
        if err != nil {
                log.Fatal(err)
        }
        ha.SetStreamHandler("/echo/1.0.0", stream)
                select {} // hang forever
}


