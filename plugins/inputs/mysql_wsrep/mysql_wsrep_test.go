package mysql_wsrep_test

import (
	"database/sql"
	"fmt"
	"testing"

	. "github.com/go-extras/telegraf-mysql-wsrep/plugins/inputs/mysql_wsrep"

	v2 "github.com/influxdata/telegraf/plugins/inputs/mysql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFloat(t *testing.T) {
	testCases := []struct {
		rawByte sql.RawBytes
		output  interface{}
		err     error
	}{
		{sql.RawBytes("123"), float64(123), nil},
		{sql.RawBytes("123.12"), float64(123.12), nil},
		{sql.RawBytes(".12"), float64(0.12), nil},
		{sql.RawBytes("+123.12"), float64(123.12), nil},
		{sql.RawBytes("-123.12"), float64(-123.12), nil},
		{sql.RawBytes("1.7976931348623157e+308"), float64(1.7976931348623157e+308), nil},
		{sql.RawBytes("1111111111111111111111111111111111111.11111111"), float64(1111111111111111111111111111111111111.11111111), nil},
		{sql.RawBytes("1.7976931348623157e+309"), nil, fmt.Errorf("unconvertible value: %q", "1.7976931348623157e+309")},
		{sql.RawBytes("invalid"), nil, fmt.Errorf("unconvertible value: %q", "invalid")},
		{sql.RawBytes(""), nil, fmt.Errorf("unconvertible value: %q", "")},
		{nil, nil, fmt.Errorf("unconvertible value: %q", "")},
		{sql.RawBytes("\x00aaa"), nil, fmt.Errorf("unconvertible value: %q", "\x00aaa")},
	}
	for _, cas := range testCases {
		val, err := ParseFloat(cas.rawByte)
		require.Equal(t, cas.err, err)
		assert.Equal(t, cas.output, val)
	}
}

func TestParseString(t *testing.T) {
	testCases := []struct {
		rawByte sql.RawBytes
		output  interface{}
		err     error
	}{
		{sql.RawBytes("123"), "123", nil},
		{sql.RawBytes(""), nil, fmt.Errorf("unconvertible value: %q", "")},
		{nil, nil, fmt.Errorf("unconvertible value: %q", "")},
	}
	for _, cas := range testCases {
		val, err := ParseString(cas.rawByte)
		require.Equal(t, cas.err, err)
		assert.Equal(t, cas.output, val)
	}
}

func TestConvertGlobalStatus(t *testing.T) {
	testCases := []struct {
		key     string
		rawByte sql.RawBytes
		output  interface{}
	}{
		{"wsrep_applier_thread_count", sql.RawBytes("1"), int64(1)},
		{"wsrep_apply_oooe", sql.RawBytes("0"), float64(0)},
		{"non_registered", sql.RawBytes("0"), int64(0)}, // comparing with previous, we get int64 here
		{"wsrep_apply_oooe", sql.RawBytes("0.1"), float64(0.1)},
		{"wsrep_ready", sql.RawBytes("ON"), int64(1)},
	}
	for _, cas := range testCases {
		val, err := v2.ConvertGlobalStatus(cas.key, cas.rawByte)
		require.NoError(t, err)
		assert.Equal(t, cas.output, val)
	}
}
