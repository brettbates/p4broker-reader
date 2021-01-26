A reader for the Perforce/helix broker filter output.

For example if you have a p4broker.conf stanza like so:

```
command: access
{
    action=filter;
    execute=/path/to/go/program;
}
```

You can use this p4broker-reader.Read() within the receiving go program to get the output dictionary as a map[string]string.

And example dictionary is like so:

```
brokerListenPort: 1998
brokerTargetPort: localhost:1999
command: access
clientProg: p4
clientVersion: 2020.2/LINUX26X86_64/2057778
clientProtocol: 89
apiProtocol: 99999
workspace: NP-B-BATES
user: brett.bates
clientIp: 127.0.0.1
clientHost: ahost
cwd: /home/brett/scratch/broker
clientPort: 1998
argCount: 2
Arg0: read
Arg1: //path/to/file
```

Example usage:
```go
// Pass in an io.Reader, though you probably want os.Stdin to get input from p4broker
res, err := Read(os.Stdin)
if err != nil {
	log.Fatalf("%v", err)
}
log.Printf("%v", res["Arg0"])
```
