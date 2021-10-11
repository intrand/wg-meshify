package main

import (
	"log"
	"os"
	"strconv"
)

// very possibly the ugliest Golang I've
// written to date. wow. go-ini can't
// repeat sections, so here we are...
func genConf(meshes map[string]mesh) {
	for _mesh, _peers := range meshes {
		for _peer, conf := range _peers.Peers {
			// [Interface] section for the current host
			out := "[Interface]\n" +
				"# " + _peer + "\n" +
				"Address = " + conf.Address + "\n" +
				"PrivateKey = " + conf.PrivateKey + "\n" +
				"ListenPort = " + strconv.Itoa(int(conf.Port)) + "\n"
			for _peer_again, conf_again := range _peers.Peers {
				if _peer != _peer_again {
					// one [Peer] section per other host
					out = out + "\n" +
						"[Peer]\n" +
						"# " + _peer_again + "\n" +
						"PublicKey = " + conf_again.PublicKey + "\n" +
						"Endpoint = " + conf_again.Entrypoint + ":" + strconv.Itoa(int(conf_again.Port)) + "\n" +
						"AllowedIPs = " + conf_again.Address + "\n"
				}
			}
			err := os.WriteFile(_mesh+"-"+_peer+".conf", []byte(out), 0600) // write it out
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
