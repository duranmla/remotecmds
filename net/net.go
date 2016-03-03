package net

import (
  "bytes"
  "io/ioutil"
  "os/user"
  "golang.org/x/crypto/ssh"
)

// this code can be found explained at: https://godoc.org/golang.org/x/crypto/ssh#Dial

func getKeyFile() (key ssh.Signer, err error){
  usr, _ := user.Current()
  file := usr.HomeDir + "/.ssh/id_rsa"
  buf, err := ioutil.ReadFile(file)

  if err != nil {
    panic(err)
  }

  key, err = ssh.ParsePrivateKey(buf)
  if err != nil {
    return
  }
  return
}

func ConnectToMachine(ip, username string) (*ssh.Session, error){
  key, err := getKeyFile();
  if err !=nil {
    panic(err)
  }

  config := &ssh.ClientConfig{
    User: username,
    Auth: []ssh.AuthMethod{
      ssh.PublicKeys(key),
    },
  }

  client, err := ssh.Dial("tcp", ip, config)
  if err != nil {
    panic("Failed to " + err.Error())
  }

  session, err := client.NewSession()
  if err != nil {
    panic("Failed to create session: " + err.Error())
  }

  var b bytes.Buffer
  session.Stdout = &b

  return session, err
}
