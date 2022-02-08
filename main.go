package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "" // to be filled in by goreleaser
	commit  = "" // to be filled in by goreleaser
	date    = "" // to be filled in by goreleaser
	builtBy = "" // to be filled in by goreleaser
	cmdname = filepath.Base(os.Args[0])
)

const filePerms = 0600

func main() {
	// parse args
	args := kingpin.MustParse(app.Parse(os.Args[1:]))

	// main decision tree
	switch args { // look for operations at the root of the command

	// version
	case cmd_version.FullCommand():
		fmt.Println(
			"{\"version\":\"" + version + "\",\"commit\":\"" + commit + "\",\"date\":\"" + date + "\",\"built_by\":\"" + builtBy + "\"}")

	// mesh
	case cmd_mesh_create.FullCommand(): // handle removing a peer
		all_meshes, _ := getMeshesFromConf()

		all_meshes[*cmd_mesh_create_name] = mesh{
			Peers: map[string]peer{},
		}

		writeConf(all_meshes)

	// peer
	case cmd_peer_rm.FullCommand(): // handle removing a peer
		all_meshes, err := getMeshesFromConf()
		if err != nil {
			log.Fatalln(err)
		}

		delete(all_meshes[*cmd_peer_mesh].Peers, *cmd_peer_name)
		writeConf(all_meshes)
		os.Exit(0)

	// peer
	case cmd_peer_set.FullCommand(): // handle adding or updating a peer
		all_meshes, err := getMeshesFromConf()
		if err != nil {
			log.Fatalln(err)
		}

		// create empty conf if mesh doesn't exist already
		_, found := all_meshes[*cmd_peer_mesh]
		if !found {
			all_meshes[*cmd_peer_mesh] = mesh{
				Peers: map[string]peer{},
			}
		}

		_mesh, err := updatePeer(all_meshes[*cmd_peer_mesh], *cmd_peer_name)
		if err != nil {
			log.Fatalln(err)
		}

		// write our new config
		all_meshes[*cmd_peer_mesh] = _mesh
		err = writeConf(all_meshes)
		if err != nil {
			log.Fatalln(err)
		}

	// genconf
	case cmd_genconf.FullCommand(): // handle generating config data
		all_meshes, err := getMeshesFromConf()
		if err != nil {
			log.Fatalln(err)
		}
		genConf(all_meshes)
	}
}
