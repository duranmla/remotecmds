package commands

import (
  "github.com/duranmla/remotecmds/cmdutil"
  "golang.org/x/crypto/ssh"
  "time"
  "os"
  "fmt"
  "strings"
)

var (
  Stdout        *os.File   = os.Stdout
)

func ListenCommands(session *ssh.Session) {
  commands := make(chan string)

  go func ()  {
    for {
      fmt.Fprint(Stdout, "-> ")
      command := cmdutil.ReadLine()
      commands <- command
    }
  }()

  for {
    select {
    case <-commands:
      executeCommand(<-commands, session)
    case <-time.After(10000 * time.Millisecond):
      fmt.Println("\n\nSession has been closed, due to inactivity...")
      session.Close()
      os.Exit(1)
    }
  }
}

func executeCommand(command string, session *ssh.Session) {
  method := strings.Split(command, " ")[0]

  switch method {
	case "say":
    // session.Start(command) <-- sadly this automatically close the session
		fmt.Printf("\nexecuting say\n<PRESS ENTER>")
	case "time":
    fmt.Printf("\ngetting Time\n<PRESS ENTER>")
	case "cpu":
		fmt.Printf("\nchecking CPU\n<PRESS ENTER>")
	case "ram":
		fmt.Printf("\nchecking RAM\n<PRESS ENTER>")
	case "screenshot":
		fmt.Printf("\ntaking screenshot\n<PRESS ENTER>")
	default:
    cpu := "cpu: to check CPU performance"
    ram := "ram: to check RAM performance"
    scr := "screenshot: to take an screenshot"
    say := "say: to make me say something"
    fmt.Printf("\nOf course! we can't allow you to run any command\nAvailable commands:\n- %s\n- %s\n- %s\n- %s\n\n <PRESS ENTER>", cpu, ram, scr, say)
	}
}
