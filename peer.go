package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func updatePeer(_mesh mesh, _peer string) (mesh, error) {
	if _mesh.Peers == nil {
		_mesh.Peers = map[string]peer{}
	}

	this_peer := _mesh.Peers[_peer]

	cmd := app.GetCommand("peer") // look at the user input
	cmd = cmd.GetCommand("set")

	// keys: setup some vars
	regen := false                                         // assume no keypair (re-)generation is necessary
	input_key := cmd.GetFlag("key").Model().Value.String() // get --key's parameter, if any
	input_pub := cmd.GetFlag("pub").Model().Value.String() // get --pub's parameter, if any

	// keys: check input
	if len(input_key) > 0 || len(input_pub) > 0 { // either key or pub had length
		if len(input_key) < 1 || len(input_pub) < 1 { // if using input, must supply both
			fmt.Println("You must supply both --key and --pub when you want to set just one of them.")
			os.Exit(1)
		} // input passed check
		this_peer.PrivateKey = *cmd_peer_key // use supplied input
		this_peer.PublicKey = *cmd_peer_pub
	} else { // check config - no input was supplied at all
		if len(this_peer.PrivateKey) > 0 || len(this_peer.PublicKey) > 0 { // either key or pub had length
			if len(this_peer.PrivateKey) < 1 || len(this_peer.PublicKey) < 1 {
				fmt.Println("Only one of key or pub had data. Odd. Regenerating new keypair for " + _peer + ".")
				regen = true // only fix is to create a new keypair
			} // config check passed, just use that (no need to do anything)
		} else { // there is no key or pub in the config
			regen = true // we must create them
		}
		if regen { // we hit a scenario in which we need to generate a keypair
			keypair, _ := wgtypes.GeneratePrivateKey()         // generate a keypair for them
			this_peer.PrivateKey = keypair.String()            // get just the key value
			this_peer.PublicKey = keypair.PublicKey().String() // get just the public key value
		}
	}

	// handle description
	input_description := cmd.GetFlag("description").Model().Value.String()
	if len(input_description) > 0 {
		this_peer.Description = *cmd_peer_description
	}

	// handle address
	input_address := cmd.GetFlag("address").Model().Value.String()
	if len(input_address) > 0 {
		this_peer.Address = *cmd_peer_address
	} else {
		if len(this_peer.Address) < 1 {
			fmt.Println("It looks like --address has an invalid value or is missing. Please consult --help and try again.")
			os.Exit(1)
		}
	}

	// handle entrypoint
	input_entrypoint := cmd.GetFlag("entrypoint").Model().Value.String()
	if len(input_entrypoint) > 0 {
		this_peer.Entrypoint = *cmd_peer_entrypoint
	} else {
		if len(this_peer.Entrypoint) < 1 {
			fmt.Println("It looks like --address has an invalid value or is missing. Please consult --help and try again.")
			os.Exit(1)
		}
	}

	// handle port
	input_port, err := strconv.Atoi(cmd.GetFlag("port").Model().Value.String())
	if err != nil {
		log.Fatalln(err)
	}
	if input_port > 0 {
		this_peer.Port = *cmd_peer_port
	} else {
		if this_peer.Port < 1 {
			fmt.Println("It looks like --port has an invalid value or is missing. Please consult --help and try again.")
			os.Exit(1)
		}
	}

	// we made it through all of the args we need to set up a new peer
	_mesh.Peers[_peer] = this_peer

	return _mesh, nil
}
