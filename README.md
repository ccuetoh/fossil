<p align="center">
  <img src="https://camiloh.com/fossil.png" width="200"/>
<p></p>

# Fossil
Fossil is a pure Go wrapper library for the Pterodactyl API and it's descendants. It allows client-wise and administrator-wise operations for easy and automated server, user, node, egg and location management.

## Installation
Fossil can be installed using the integrated Go package manager:

``go get github.com/camilohernandez/fossil``

## Examples
### Client
A Client gets access to all the functionalities a user might find on their server control panel. A Client Token is requiered for the Creation of  a Client. To start a new client use the ```NewClient()``` function:

```go
import "github.com/camilohernandez/fossil"

client := fossil.NewClient("https://example.com","NRVW42TME45WW3B3E5VTWOZ3MFZWIYLTMRQXGZDBMFZWIYLT")
```

#### Fetch all servers
```go
servers, err := client.GetServers()
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}

for _, s := range servers {
    fmt.Printf("ID: %s\n", s.ID)
    fmt.Printf("Name: %s\n", s.Name)
    fmt.Printf("Description: %s\n", s.Description)
    fmt.Printf("IP: %s\n\n", s.Allocation.IP)
}
```

#### Fetch a specific server
```go
server, err := client.GetServer("6a185444")
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}

fmt.Printf("ID: %s\n", server.ID)
fmt.Printf("Name: %s\n", server.Name)
fmt.Printf("Description: %s\n", server.Description)
fmt.Printf("IP: %s\n", server.Allocation.IP)
```

#### Turn server on
```go
err := client.SetPowerState("6a185444", fossil.ON)
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```

#### Execute a command on a server
```go
err := client.ExecuteCommand("6a185444", "say Hello!")
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```

## Disclaimer
Fossil is partially based on the [Crocgodyl](www.github.com/parkervcp/crocgodyl) library. Though no code of the before mentioned library is used, it does draw ideas from it. All the respective kudos to the maintainers. 

This library is licensed under the MIT Licence, please refer to the LICENCE file for more information.
