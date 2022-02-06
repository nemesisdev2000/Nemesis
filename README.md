# Writing Custom C2

## Creating the server

### Required features (v0.001)

- The main C2 server to be created
- It should be able to send and receive messages from the agents
- It should be able to send commands and recieve the output of these commands
- A test agent might be used for this one. Later on this test agent created in golang will change to some other language. But for testing purposes this will be the one we use



### v0.002 

- Using Multiple agents
- Solve problem with freeing up close listener
	- ~~After the connection to an agent is closed, the server cannot reuse the same port~~
	- ~~The port needs to be closed separately~~
- Use multiple clients working on the team server
- HTTP communication


## Installation :

**NOTE: May vary with time**

- Install custom modules

```bash
cd lib/agentPool
go install

cd ../handleAgent
go install

cd ../../listeners/TcpListener
go install
```

- Install go-reuseport module from `github.com/libp2p/go-reuseport`

```bash
go get github.com/libp2p/go-reuseport
```


