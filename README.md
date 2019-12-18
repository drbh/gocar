# gocar

kinda like cargo for Rust - but a cheap knockoff for Golang. 

## Installation

```bash
go get -v github.com/drbh/gocar
go install github.com/drbh/gocar
```


## Usage

Start new project  
```bash
gocar new [project-name]
```

Change directory to the new project  
```bash
cd [project-name]
```

Run it with `gocar`  
```bash
gocar run
# "Hello, World!"
```


## Whats created

```
$ tree .
.
├── .gitignore
├── README.md
├── gocar.toml
└── src
    └── main.go
```

