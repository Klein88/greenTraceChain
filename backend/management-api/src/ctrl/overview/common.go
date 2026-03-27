package overview

import (
	"github.com/ethereum/go-ethereum/log"

	"management_backend/src/utils"
)

// CHAINSTARTTLS chain start tls
const CHAINSTARTTLS = 0

// LOCALHOST localhost
const LOCALHOST = "localhost"

func createCertFile(name string, cert string, priKey string, path string) error {
	err := utils.CreateAndRename(name+".crt", path+"/"+name+".crt", cert)
	if err != nil {
		log.Error("create and rename crt file err :", err.Error())
		return err
	}
	err = utils.CreateAndRename(name+".key", path+"/"+name+".key", priKey)
	if err != nil {
		log.Error("create and rename kry file err :", err.Error())
	}
	return err
}

func createPkNodeFile(name string, publicKey string, priKey string, nodeId string, path string) error {
	err := utils.CreateAndRename(name+".crt", path+"/"+name+".pem", publicKey)
	if err != nil {
		log.Error("create and rename crt file err :", err.Error())
		return err
	}
	err = utils.CreateAndRename(name+".key", path+"/"+name+".key", priKey)
	if err != nil {
		log.Error("create and rename kry file err :", err.Error())
	}
	err = utils.CreateAndRename(name+".id", path+"/"+name+".id", nodeId)
	if err != nil {
		log.Error("create and rename kry file err :", err.Error())
	}
	return err
}

func createPkAdminFile(name string, publicKey string, priKey string, addr string, path string) error {
	err := utils.CreateAndRename(name+".crt", path+"/"+name+".pem", publicKey)
	if err != nil {
		log.Error("create and rename crt file err :", err.Error())
		return err
	}
	err = utils.CreateAndRename(name+".key", path+"/"+name+".key", priKey)
	if err != nil {
		log.Error("create and rename kry file err :", err.Error())
	}
	err = utils.CreateAndRename(name+".addr", path+"/"+name+".addr", addr)
	if err != nil {
		log.Error("create and rename kry file err :", err.Error())
	}
	return err
}
