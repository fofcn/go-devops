package sshclient

import (
	"log"

	"golang.org/x/crypto/ssh"
)

type SSHConnectionInfo struct {
	Address   string
	Proto     string
	Username  string
	Password  string
	PublicKey string
}

type SSHSession struct {
	Client  *ssh.Client
	Session *ssh.Session
}

func CreateNewSesion(sessionInfo SSHConnectionInfo) (*SSHSession, error) {
	var sshClient *ssh.Client

	if len(sessionInfo.Username) != 0 {
		client, err := ssh.Dial(sessionInfo.Proto, sessionInfo.Address, &ssh.ClientConfig{
			User:            sessionInfo.Username,
			Auth:            []ssh.AuthMethod{ssh.Password(sessionInfo.Password)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})
		if err != nil {
			log.Fatalf("New session error: %s", err.Error())
			return nil, err
		}
		sshClient = client
	} else {
		// todo
		client, err := ssh.Dial(sessionInfo.Proto, sessionInfo.Address, &ssh.ClientConfig{
			User:            sessionInfo.Username,
			Auth:            []ssh.AuthMethod{ssh.Password(sessionInfo.Password)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})
		if err != nil {
			log.Fatalf("New session error: %s", err.Error())
			return nil, err
		}
		sshClient = client
	}

	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalf("New session error: %s", err.Error())
		return nil, err
	}

	var nodeSession SSHSession
	nodeSession.Client = sshClient
	nodeSession.Session = session

	return &nodeSession, nil
}

func CloseSession(nodeSession *SSHSession) {
	if nodeSession.Session != nil {
		defer nodeSession.Session.Close()

	}

	if nodeSession.Client != nil {
		defer nodeSession.Client.Close()
	}
}

func CloseClient(sshclient *ssh.Client) {
	if sshclient != nil {
		defer sshclient.Close()
	}
}
