# Go-lang
Go with the flow.


Steps to create a module in workspace

## eg: steps
1. dir: /GO-LANG/modules/01namaste
2. cmd: ``` $ go mod init namaste```

    Result :

        go: creating new go.mod: module namaste
        go: to add module requirements and sums:
                go mod tidy
        
3. dir: /GO-LANG
4. cmd: ```$ go work init ```
5. cmd: ```$ go work use modules/01namaste/```
6. cmd: ```$ go work use modules/02variables/```

..
& so on !


## Working Dir Practice for 3 modules as example.

#### /GO-LANG/modules/
``` bash tree
├── 01namaste
│   ├── go.mod
│   ├── namaste.go
│   └── ...other files
├── 02variables
│   ├── go.mod
│   ├── variables.go
│   └── ...other files
└── nextmodule
    ├── go.mod
    ├── file.go
    └── ...otherfiles
```