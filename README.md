<p align="center">
  <img src="https://camiloh.com/fossil.png" width="200"/>
<p></p>

# Fossil
[![Go Report Card](https://goreportcard.com/badge/github.com/CamiloHernandez/fossil)](https://goreportcard.com/report/github.com/CamiloHernandez/fossil)
[![Build Status](https://travis-ci.org/CamiloHernandez/fossil.svg?branch=master)](https://travis-ci.org/CamiloHernandez/fossil)
![Coverage](https://img.shields.io/badge/Go%20Coverage-72%25-brightgreen.svg?longCache=true&style=flat)
[![Licence](https://img.shields.io/github/license/CamiloHernandez/fossil)](https://github.com/CamiloHernandez/fossil/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/CamiloHernandez/fossil?status.svg)](https://godoc.org/github.com/CamiloHernandez/fossil)

Fossil is a pure Go wrapper library for the Pterodactyl API and it's descendants. It allows client-wise and administrator-wise operations for easy and automated server, user, node, egg and location management.

## Contents:
- [Installation](#installation)
- [Examples](#examples)
    - [Client API](#client-api)
        - [Servers](#client-servers)
            - [Fetch](#client-servers-fetch)
            - [Status and usage](#client-servers-status)
        - [Server Actions](#client-serveractions)
            - [Turn on/off](#client-serveractions-onoff)
            - [Execute commands](#client-serveractions-command)
    - [Application API](#app-api)
        - [Servers](#app-servers)
            - [Fetch](#app-servers-fetch)
            - [Create](#app-servers-create)
            - [Modify](#app-servers-modify)
            - [Suspend/Unsuspend](#app-servers-suspend)
            - [Reinstall or rebuild](#app-servers-re)
        - [Databases](#app-db)
            - [Fetch](#app-db-fetch)
            - [Create](#app-db-create)
            - [Delete](#app-db-delete)
        - [Users](#app-users)
            - [Fetch](#app-users-fetch)
            - [Modify](#app-users-modify)
            - [Create](#app-users-create)
            - [Delete](#app-users-delete)
        - [Nests and eggs](#app-nests)
            - [Fetch nests](#app-nests-fetch)
            - [Fetch eggs](#app-nests-fetch-eggs)
- [Disclaimer](#disclaimer)
- [Licence](#licence)

<a name="installation"></a>
## Installation
Fossil can be installed using the integrated Go package manager:

``go get github.com/camilohernandez/fossil``

<a name="examples"></a>
## Examples
<a name="client-api"></a>
### Client API
A Client gets access to all the functionalities a user might find on their server control panel. A Client Token is requiered for the Creation of a Client. To start a new Client use the ```NewClient()``` function:

```go
import "github.com/camilohernandez/fossil"

client := fossil.NewClient("https://example.com","NRVW42TME45WW3B3E5VTWOZ3MFZWIYLTMRQXGZDBMFZWIYLT") // Panel URL and Client Token
```
<a name="client-servers"></a>
#### Servers
<a name="client-servers-fetch"></a>
##### Fetch all servers
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
}
```

##### Fetch a specific server
```go
server, err := client.GetServer("6a185444")
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}

fmt.Printf("ID: %s\n", server.ID)
fmt.Printf("Name: %s\n", server.Name)
fmt.Printf("Description: %s\n", server.Description)
```

<a name="client-servers-status"></a>
##### Get a server's status and usage
```go
status, err := client.GetServerStatus("6a185444")
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}

fmt.Printf("State: %s\n", status.State)
fmt.Printf("CPU Current: %f\n", server.CPU.Current)
fmt.Printf("Memory used: %d\n", server.Memory.Used)
fmt.Printf("Disk used: %s\n", server.Disk.Used)
```

<a name="client-serveractions"></a>
#### Server Actions
<a name="client-serveractions-onoff"></a>
##### Turn server on
```go
err := client.SetPowerState("6a185444", fossil.ON)
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```
<a name="client-serveractions-command"></a>
##### Execute a command on a server
```go
err := client.ExecuteCommand("6a185444", "say Hello!")
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```

<a name="app-api"></a>
### Application API
An Application connection allows full access to server, user, location, nest and egg management. With Application calls you have full administrator-level access to the creation of users and servers. An Application Token (also called "API Token") is required to create an Application object. To start a new Application use the ```NewApplication()``` function:
```go
import "github.com/camilohernandez/fossil"

app := fossil.NewApplication("https://example.com","OF3WK4LXMVYXOZLYMRQXQWTYMFZWIYLTMRQXGZDBOM") // Panel URl and Application Token

```
<a name="app-servers"></a>
#### Servers
When interacting servers with an Application Token you get extra information, and more functions.
<a name="app-servers-fetch"></a>
##### Fetch all servers
```go
servers, err := app.GetServers()
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}

for _, s := range servers {
    fmt.Printf("ID: %d\n", s.ID)
    fmt.Printf("Name: %s\n", s.Name)
    fmt.Printf("Description: %s\n", s.Description)
    fmt.Printf("Node: %d\n", s.Node)
    fmt.Printf("Nest: %d\n\n", s.Nest)
}
```

##### Get a server with its Internal ID
```go
server, err := app.GetServer(17)
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```

##### Get a server with its External ID
```go
server, err := app.GetServerExternal("test_server1")
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```
<a name="app-servers-create"></a>
##### Create a new server
```go
server := &fossil.ApplicationServer{
    Name:        "Test Server",
    Description: "A server to test the Fossil library",
    ExternalID:  "test_server1",
    User:        12,
    Node:        2,
    Nest:        5,
    Egg:         20,
    Container: fossil.Container{
        StartupCommand: "java -Xms128M -Xmx 1024M -jar server.jar",
        Image:          "quay.io/pterodactyl/core:java-glibc", // Minecraft
        Environment: map[string]string{ // WISP Requieres some aditional enviroment variables for SRCDS images:
            "DL_VERSION": "1.12.2",     // SRCDS_APPID, STEAM_ACC, SRCDS_MAP
        },
    },
    Limits: fossil.Limits{
        Memory:    512,
        Swap:      512,
        Disk:      1024,
        IO:        500,
        CPU:       100,
        Databases: 1,
    },
    Allocation: 9000,
}

err := app.CreateServer(server)
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```
<a name="app-servers-modify"></a>
##### Change a server's name, user, external ID or description
```go
sv, err := app.GetServer(17)
if err != nil {
    fmt.Println(err.Error())
    return
}

sv.Name = "Fossil Test"

err = app.UpdateDetails(sv)
if err != nil {
    fmt.Println(err.Error())
}
```

##### Change a server's limits or allocations
```go
sv, err := app.GetServer(17)
if err != nil {
    fmt.Println(err.Error())
    return
}

sv.Limits = fossil.Limits{
    Memory:      1024,
    Databases:   2,
    Allocations: 2,
}

err = app.UpdateBuild(sv, []int{9001, 9002},[]int{9000}) // Add allocations 9001 and 9002, delete 9000
if err != nil {
    fmt.Println(err.Error())
    return
}
```

##### Change a server's image settings
```go
sv, err := app.GetServer(17)
if err != nil {
    fmt.Println(err.Error())
    return
}

sv.Container = fossil.Container{
    StartupCommand: "java -Xms128M -Xmx 1024M -jar server.jar",
    Image:          "quay.io/pterodactyl/core:java-glibc", // Minecraft
    Environment: map[string]string{
        "DL_VERSION": "1.12.2",
    },
}

err = app.UpdateStartup(sv)
if err != nil {
    fmt.Println(err.Error())
    return
}
```
<a name="app-servers-suspend"></a>
##### Suspend or unsuspend a server
```go
err := app.SuspendServer(17) // Suspending server with Internal ID 17 
if err != nil {
    fmt.Println(err.Error())
    return
}

err = app.UnsuspendServer(17) // Unsuspending server with Internal ID 17 
if err != nil {
    fmt.Println(err.Error())
    return
}
```
<a name="app-servers-re"></a>
##### Rebuild or reinstall a server
```go
err := app.RebuildServer(17) // Rebuilding server with Internal ID 17
if err != nil {
    fmt.Println(err.Error())
    return
}

err = app.ReinstallServer(17) // Reinstalling server with Internal ID 17
if err != nil {
    fmt.Println(err.Error())
    return
}
```
<a name="app-servers-delete"></a>
##### Delete a server
```go
err := app.DeleteServer(17) // Deleting server with Internal ID 17
if err != nil {
    fmt.Println(err.Error())
    return
}

err = app.ForceDeleteServer(17) // DeleteServer should allays be preferred, as ForceDeleteServer
if err != nil {                 // is an unsafe operation.
    fmt.Println(err.Error())
    return
}
```
<a name="app-db"></a>
#### Databases
<a name="app-db-fetch"></a>
##### Fetch all databases
```go
dbs, err := app.GetDatabases()
if err != nil {
    fmt.Println(err.Error())
    return
}

for _, db := range dbs{
    fmt.Println(db.Database)
}
```

##### Fetch a specific database
```go
db, err := app.GetDatabase(17, 1) // Get database 1 in server 17
if err != nil {
    fmt.Println(err.Error())
    return
}
```
<a name="app-db-create"></a>
##### Create a database
```go
db := &fossil.Database{
    Database: "fossil_test",
    Host:      15, // The server 15 will be the host server
    Remote:    "%", // IPs and wildcards work
}

err := app.CreateDatabase(17, db)
if err != nil {
    fmt.Println(err.Error())
    return
}
```
<a name="app-db-delete"></a>
##### Delete a database
```go
err := app.DeleteDatabase(17, 1) // Delete database 1 in server 17
if err != nil {
    fmt.Println(err.Error())
    return
}
```
<a name="app-users"></a>
### Users
<a name="app-users-fetch"></a>
##### Fetch all users
```go
users, err := app.GetUsers()
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}

for _, u := range users {
    fmt.Printf("ID: %d\n", u.ID)
    fmt.Printf("First Name: %s\n", u.FirstName)
    fmt.Printf("Last Name: %s\n\n", u.LastName)
}
```

##### Fetch a specific user
```go
user, err := app.GetUser(4) // Get user of Internal ID 4
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```
<a name="app-users-create"></a>
##### Create a user
```go
user := &fossil.User{
    ExternalID:              "test_user1",
    Username:                "TestUser",
    Email:                   "example@example.com",
    FirstName:               "John",
    LastName:                "Doe",
}

err := app.CreateUser(user, "password123")
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```
<a name="app-users-modify"></a>
##### Update a user
```go
user, err := app.GetUser(17)
if err != nil {
    return
}

user.FirstName = "John"

err = app.UpdateUser(user, "password123") // Password is optional
if err != nil {
    fmt.Println("ERROR: " + err.Error())
    return
}
```
<a name="app-users-delete"></a>
##### Delete a user
```go
err := app.DeleteUser(17)
if err != nil {
    return
}
```

<a name="app-nests"></a>
### Nests and eggs
<a name="app-nests-fetch"></a>
##### Fetch nests
```go
nests, err := app.GetNests()
if err != nil {
    fmt.Println(err.Error())
    return
}

for _, n := range nests{
    fmt.Println(n.Name)
    fmt.Println(n.Description)
}
```

##### Fetch specific nest
```go
nest, err := app.GetNest(2) // Get nest with ID 2 (Minecraft)
if err != nil {
    fmt.Println(err.Error())
    return
}
```

<a name="app-nests-fetch-eggs"></a>
##### Fetch eggs inside a nest
```go
eggs, err := app.GetEggs(2) // Get all the eggs for nest 2 (Minecraft)
if err != nil {
    fmt.Println(err.Error())
    return
}

for _, e := range eggs{
    fmt.Println(e.ID)
    fmt.Println(e.Description)
}
```

##### Fetch specific egg inside a nest
```go
egg, err := app.GetEgg(2, 2)
if err != nil {
    fmt.Println(err.Error())
    return
}
```

<a name="disclaimer"></a>
## Disclaimer
Fossil is partially based on the [Crocgodyl](https://www.github.com/parkervcp/crocgodyl) library. All the respective kudos to the author. 
<a name="licence"></a>
## Licence
This library is licensed under the MIT Licence, please refer to the LICENCE file for more information. The logo used in this project is provided by OpenClipart-Vectors, and is licenced under the [Pixbay Licence](https://pixabay.com/service/license/).
