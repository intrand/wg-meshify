package main

import (
	"github.com/alecthomas/kingpin/v2"
)

var (
	// global
	app      = kingpin.New("wg-meshify", "Generates configuration files for all members of a wireguard mesh").Author("intrand")
	cmd_conf = app.Flag("conf", "Path to your_mesh_config.yml").Envar("wg_meshify_conf").Default("mesh.yml").String()

	cmd_version = app.Command("version", "prints version and exits")

	// mesh
	cmd_mesh = app.Command("mesh", "mesh stuff")

	cmd_mesh_create      = cmd_mesh.Command("create", "adds mesh to conf")
	cmd_mesh_create_name = cmd_mesh_create.Flag("name", "name of mesh to create").Short('n').Envar("wg_meshify_mesh_create_name").Required().String()

	cmd_mesh_rm      = cmd_mesh.Command("rm", "add/update a member of the mesh")                                                                                    // future
	cmd_mesh_rm_name = cmd_mesh_rm.Flag("name", "name of mesh to remove").Short('n').Envar("wg_meshify_mesh_rm_name").Required().String()                           // future
	cmd_mesh_rm_yes  = cmd_mesh_rm.Flag("yes", "remove mesh from conf without asking questions").Short('y').Envar("wg_meshify_mesh_rm_yes").Default("false").Bool() // future

	// show
	cmd_show        = app.Command("show", "show configuration data")                                                                                            // future
	cmd_show_output = cmd_show.Flag("output", "output format").Short('o').Envar("wg_meshify_show_output").Default("table").String()                             // future
	cmd_show_mesh   = cmd_show.Flag("mesh", "The name of the mesh you wish to operate upon.").Short('m').Envar("wg_meshify_show_mesh").Default("mesh").String() // future

	// genconf
	cmd_genconf      = app.Command("genconf", "generate mesh configuration data for each peer")
	cmd_genconf_dir  = cmd_genconf.Flag("dir", "directory in which to write configuration data").Short('d').Envar("wg_meshify_genconf_dir").Default(".").String() // future
	cmd_genconf_mesh = cmd_genconf.Flag("mesh", "The name of the mesh you wish to operate upon.").Short('m').Envar("wg_meshify_mesh").Default("mesh").String()    // future

	// peer
	cmd_peer      = app.Command("peer", "manage a member of the mesh")
	cmd_peer_name = cmd_peer.Flag("name", "name of peer; only shown by this tool").Short('n').Envar("wg_meshify_peer_name").Required().String()
	cmd_peer_mesh = cmd_peer.Flag("mesh", "The name of the mesh you wish to operate upon.").Short('m').Envar("wg_meshify_peer_mesh").Default("mesh").String()

	cmd_peer_set         = cmd_peer.Command("set", "add/update a member of the mesh")
	cmd_peer_description = cmd_peer_set.Flag("description", "description of peer; only shown by this tool").Short('d').Envar("wg_meshify_peer_set_description").String()
	cmd_peer_address     = cmd_peer_set.Flag("address", "IP address of this peer internal to the mesh").Short('a').Envar("wg_meshify_peer_set_address").String()
	cmd_peer_entrypoint  = cmd_peer_set.Flag("entrypoint", "external host.domain.tld or IP address of this peer").Short('e').Envar("wg_meshify_peer_set_address").String()
	cmd_peer_port        = cmd_peer_set.Flag("port", "UDP port of this peer on the mesh").Short('p').Envar("wg_meshify_set_peer_port").Uint16()
	cmd_peer_key         = cmd_peer_set.Flag("key", "base64-encoded private key").Short('k').Envar("wg_meshify_peer_set_key").String()
	cmd_peer_pub         = cmd_peer_set.Flag("pub", "base64-encoded public key").Short('b').Envar("wg_meshify_peer_set_pub").String()

	cmd_peer_rm     = cmd_peer.Command("rm", "remove a member from the mesh")
	cmd_peer_rm_yes = cmd_peer.Flag("yes", "remove peer from mesh without asking questions").Default("true").Hidden().Bool() // commands need something under them to work
)
