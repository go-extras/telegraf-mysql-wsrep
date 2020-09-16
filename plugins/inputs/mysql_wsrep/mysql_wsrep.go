package mysql_wsrep

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/influxdata/telegraf/plugins/inputs/mysql"
	v2 "github.com/influxdata/telegraf/plugins/inputs/mysql/v2"
)

const (
	defaultTimeout                             = 5 * time.Second
	defaultPerfEventsStatementsDigestTextLimit = 120
	defaultPerfEventsStatementsLimit           = 250
	defaultPerfEventsStatementsTimeLimit       = 86400
	defaultGatherGlobalVars                    = true
)

func ParseFloat(value sql.RawBytes) (interface{}, error) {
	if val, err := strconv.ParseFloat(string(value), 64); err == nil {
		return val, nil
	}

	return nil, fmt.Errorf("unconvertible value: %q", string(value))
}

func ParseString(value sql.RawBytes) (interface{}, error) {
	if len(string(value)) > 0 {
		return string(value), nil
	}

	return nil, fmt.Errorf("unconvertible value: %q", string(value))
}

func init() {
	v2.GlobalStatusConversions["wsrep_applier_thread_count"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_apply_oooe"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_apply_oool"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_apply_window"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_causal_reads"] = v2.ParseBoolAsInteger
	v2.GlobalStatusConversions["wsrep_cert_deps_distance"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_cert_index_size"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_cert_interval"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_cluster_conf_id"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_cluster_size"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_cluster_state_uuid"] = ParseString
	v2.GlobalStatusConversions["wsrep_cluster_status"] = ParseString
	v2.GlobalStatusConversions["wsrep_cluster_weight"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_commit_oooe"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_commit_oool"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_commit_window"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_connected"] = v2.ParseBoolAsInteger
	v2.GlobalStatusConversions["wsrep_desync_count"] = v2.ParseInt

	// This status variable provides figures for the replication latency on group communication. It measures latency
	// from the time point when a message is sent out to the time point when a message is received. As replication is
	// a group operation, this essentially gives you the slowest ACK and longest RTT in the cluster.
	//
	// The value is a string of the following values: Minimum / Average / Maximum / Standard Deviation / Sample Size
	// Example: 0.00243433/0.144033/0.581963/0.215724/13
	v2.GlobalStatusConversions["wsrep_evs_repl_latency"] = ParseString

	v2.GlobalStatusConversions["wsrep_evs_state"] = ParseString
	v2.GlobalStatusConversions["wsrep_flow_control_paused"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_flow_control_paused_ns"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_flow_control_recv"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_flow_control_sent"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_gcomm_uuid"] = ParseString
	v2.GlobalStatusConversions["wsrep_incoming_addresses"] = ParseString
	v2.GlobalStatusConversions["wsrep_last_committed"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_bf_aborts"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_cached_downto"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_cert_failures"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_commits"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_index"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_recv_queue"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_recv_queue_avg"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_local_recv_queue_max"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_recv_queue_min"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_replays"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_send_queue"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_send_queue_avg"] = ParseFloat
	v2.GlobalStatusConversions["wsrep_local_send_queue_max"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_send_queue_min"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_state"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_local_state_comment"] = ParseString
	v2.GlobalStatusConversions["wsrep_local_state_uuid"] = ParseString
	v2.GlobalStatusConversions["wsrep_open_connections"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_open_transactions"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_protocol_version"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_provider_capabilities"] = ParseString
	v2.GlobalStatusConversions["wsrep_provider_name"] = ParseString
	v2.GlobalStatusConversions["wsrep_provider_vendor"] = ParseString
	v2.GlobalStatusConversions["wsrep_provider_version"] = ParseString
	v2.GlobalStatusConversions["wsrep_ready"] = v2.ParseBoolAsInteger
	v2.GlobalStatusConversions["wsrep_received"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_received_bytes"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_repl_data_bytes"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_repl_keys"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_repl_keys_bytes"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_repl_other_bytes"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_replicated"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_replicated_bytes"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_rollbacker_thread_count"] = v2.ParseInt
	v2.GlobalStatusConversions["wsrep_thread_count"] = v2.ParseInt

	inputs.Add("mysql_wsrep", func() telegraf.Input {
		return &mysql.Mysql{
			PerfEventsStatementsDigestTextLimit: defaultPerfEventsStatementsDigestTextLimit,
			PerfEventsStatementsLimit:           defaultPerfEventsStatementsLimit,
			PerfEventsStatementsTimeLimit:       defaultPerfEventsStatementsTimeLimit,
			GatherGlobalVars:                    defaultGatherGlobalVars,
		}
	})
}
