# Implementing interface from different package golang

https://www.golangprograms.com/implementing-interface-from-different-package-golang.html


This example aims to demonstrate the implementation of interfaces in Go and import your custom package. You will be able to define and declare an interface for an application in custom packages and implement that interface in your applications.
The following will be the directory structure of our application.

├── analysis
│   ├── go.mod
│   ├── main.go
│   ├── vehicle
│   │   └── vehicle.go
│   └── human
│       └── human.go
