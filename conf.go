package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func writeConf(meshes map[string]mesh) error {
	// if the file doesn't exist, create it
	_, err := os.OpenFile(*cmd_conf, os.O_CREATE, filePerms)
	if err != nil {
		log.Fatalln(err)
	}

	// golang -> yaml
	var output bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&output)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(&meshes)
	if err != nil {
		return err
	}

	// write to file
	err = ioutil.WriteFile(*cmd_conf, output.Bytes(), filePerms)
	if err != nil {
		return err
	}

	return nil
}

func getMeshesFromConf() (map[string]mesh, error) {
	all_meshes := map[string]mesh{}

	// read file to mem
	yamlFile, err := ioutil.ReadFile(*cmd_conf)
	if err != nil {
		return all_meshes, err
	}

	// parse yaml to golang
	err = yaml.Unmarshal(yamlFile, &all_meshes)
	if err != nil {
		return all_meshes, err
	}

	return all_meshes, nil
}
