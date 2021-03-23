package main

import (
	"fmt"
	"log"
        "io/ioutil"
	"golang.org/x/crypto/ssh"
        "os"
)

func main() {


        var user = "root";
        var host = "192.168.1.150:22";
        var cmd = "ls";
	client, session, err := connectToHost(user, host)
	if err != nil {
		panic(err)
	}
	out, err := session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	client.Close()
}

func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {

        dir,err := os.UserHomeDir();

        sshpath := dir + "/.ssh/id_rsa";
        fmt.Printf("Path = %s\n",sshpath);
        key, err := ioutil.ReadFile(sshpath)

        // Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
                Auth: []ssh.AuthMethod{
			// Add in password check here for moar security.
			ssh.PublicKeys(signer),
                        },
	}
        sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()


	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}

