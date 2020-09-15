# Telegraf execd mysql_wsrep input

This is an extension to the original mysql plugin to support proper types for wsrep.

# Install Instructions

Download the repo somewhere

    $ git clone git@github.com:go-extras/telegraf-mysql-wsrep.git

build the "mysql_wsrep" binary

    $ go build -o mysql_wsrep cmd/main.go
    
 (if you're using windows, you'll want to give it an .exe extension)
 
    go build -o mysql_wsrep.exe cmd/main.go

You should be able to call this from telegraf now using execd:

```
[[inputs.execd]]
  command = ["/path/to/mysql_wsrep_binary"]
  signal = "none"
  
# sample output: write metrics to stdout
[[outputs.file]]
  files = ["stdout"]
```
