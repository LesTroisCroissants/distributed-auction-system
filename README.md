# distributed-auction-system

Handin 5 of the introductory Distributed Systems course of ITU.

# Guide — How to run the program

To run the auction system, first three servers must be started, then an arbitrary number of clients can be started. When everything has been started, the system can tolerate one crash failure, either the leading server may crash or one of the backups.

## Starting the servers

Three servers must be started with ID 1, 2 and 3 respectively. By default a server uses ID 1. The default server ID can be overwritten using the `-id` flag, thus servers with ID 2 and 3 can be started this way.

Starting all three servers might look like the following, when run from the root directory:

Server 1: `go run server/server.go`

Server 2: `go run server/server.go -id 2`

Server 3: `go run server/server.go -id 3`

## Starting the client(s)

Starting a client can be done by running the following command from the root directory:

`go run client/client.go`
