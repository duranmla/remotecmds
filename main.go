package main

import (
  "fmt"
  "os"
  "log"
  "net"
  "github.com/duranmla/remotecmds/cmdutil"
  "github.com/duranmla/remotecmds/commands"
  connect "github.com/duranmla/remotecmds/net"
)

var (
  Stdout        *os.File   = os.Stdout
)

func getLocalIP() net.IP {
  host, _ := os.Hostname()
  addrs, _ := net.LookupIP(host)
  var ip net.IP

  for _, addr := range addrs {
    if ipv4 := addr.To4(); ipv4 != nil {
      ip = ipv4
      break
    }
  }

  return ip
}

func main() {
  ip := getLocalIP().String()
  fmt.Fprint(Stdout, "Username: ")
  username := cmdutil.ReadLine()
  log.Printf("Connecting to %s as %s...", ip, username)
  session, err := connect.ConnectToMachine(ip + ":22", username)

  if err != nil {
    panic("Failed to run: " + err.Error())
  } else {
    log.Println("Conection sucess!! <CTRL + C> to Exit\n")
  }

  commands.ListenCommands(session)
}
