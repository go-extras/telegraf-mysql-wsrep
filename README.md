# Telegraf MySQL Input External Plugin with Extended Wsrep Support

This is an extension to the original mysql plugin to support proper types for wsrep. The only difference is that this plugin version defines type conversion rules for `wsrep_*` variables according to MariaDB/Galera documentation. In addition as it is an external plugin, it runs as a separate process (ran and monitored by Telegraf's execd plugin).

For more information refer to the original [documentation for the mysql plugin](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/mysql) and the original [documentation for the execd plugin](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/execd).

# Compilation

Download the repo somewhere

    $ git clone git@github.com:go-extras/telegraf-mysql-wsrep.git

build the "mysql_wsrep" binary

    $ go build -o mysql_wsrep cmd/main.go
    
 (if you're using windows, you'll want to give it an .exe extension)
 
    go build -o mysql_wsrep.exe cmd/main.go

# Usage

You should be able to call this from telegraf now using execd:

`telegraf.conf`:
```toml
[[inputs.execd]]
  command = ["/path/to/mysql_wsrep_binary", "-c", "/path/to/mysql_wsrep.conf"]
  signal = "none"
```

`mysql_wsrep.conf` (refer to the built-in mysql plugin for more details):

```toml
    [[inputs.mysql_wsrep]]
      servers = ["user:password@tcp(127.0.0.1:3306)/"]
      metric_version = 2
      #fielddrop = []
      #gather_table_schema = true
      gather_process_list = true
      #gather_user_statistics = true
      #gather_info_schema_auto_inc = true
      gather_innodb_metrics = true
      #gather_slave_status = true
      #gather_binary_logs = true
      gather_global_variables = true
      #gather_table_io_waits = true
      #gather_table_lock_waits = true
      #gather_index_io_waits = true
      #gather_event_waits = true
      #gather_file_events_stats = true
      #gather_perf_events_statements = false
      ## the limits for metrics form perf_events_statements
      # perf_events_statements_digest_text_limit = 120
      # perf_events_statements_limit = 250
      # perf_events_statements_time_limit = 86400
      ## Some queries we may want to run less often (such as SHOW GLOBAL VARIABLES)
      ##   example: interval_slow = "30m"
      # interval_slow = ""
```
