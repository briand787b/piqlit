package plog_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/briand787b/piqlit/core/plog"
	"github.com/briand787b/piqlit/core/plog/plogtest"
	"github.com/briand787b/piqlit/core/util"

	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

func TestPLoggerError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		spanID            *string
		TraceID           *string
		msg               string
		args              []interface{}
		uuidGenSet        []string
		expFormattedStmts [][]byte
	}{
		{
			"full_msg_no_args_found",
			util.StrPtr("42"),
			util.StrPtr("69"),
			"banana",
			nil,
			nil,
			[][]byte{
				[]byte(`{
					"level": "ERROR",
					"message": "banana",
					"span_id": "42",
					"trace_id": "69"
				}`),
			},
		},
		{
			"full_msg_no_args_found_no_span_id",
			nil,
			util.StrPtr("69"),
			"orange",
			nil,
			[]string{"blah"},
			[][]byte{
				[]byte(`{
					"level": "ERROR",
					"message": "spanID is empty string",
					"span_id": "blah",
					"trace_id": "69",
					"data": {
						"new_span_id": "blah"
					}
				}`),
				[]byte(`{
					"level": "ERROR",
					"message": "orange",
					"span_id": "blah",
					"trace_id": "69"
				}`),
			},
		},
		{
			"full_msg_no_args_found_no_trace_id",
			util.StrPtr("42"),
			nil,
			"strawberry",
			nil,
			[]string{"blah"},
			[][]byte{
				[]byte(`{
					"level": "ERROR",
					"message": "traceID is empty string",
					"span_id": "42",
					"trace_id": "blah",
					"data": {
						"new_trace_id": "blah"
					}
				}`),
				[]byte(`{
					"level": "ERROR",
					"message": "strawberry",
					"span_id": "42",
					"trace_id": "blah"
				}`),
			},
		},
		{
			"full_msg_2_args_found",
			util.StrPtr("42"),
			util.StrPtr("69"),
			"banana",
			[]interface{}{"grape", "kiwi"},
			nil,
			[][]byte{
				[]byte(`{
					"level": "ERROR",
					"message": "banana",
					"span_id": "42",
					"trace_id": "69",
					"data": {
						"grape": "kiwi"
					}
				}`),
			},
		},
		{
			"full_msg_4_args_found",
			util.StrPtr("42"),
			util.StrPtr("69"),
			"banana",
			[]interface{}{"grape", "kiwi", "apple", "pear"},
			nil,
			[][]byte{
				[]byte(`{
						"level": "ERROR",
						"message": "banana",
						"span_id": "42",
						"trace_id": "69",
						"data": {
							"grape": "kiwi",
							"apple": "pear"
						}
					}`),
			},
		},
		{
			"full_msg_1_arg_found",
			util.StrPtr("42"),
			util.StrPtr("69"),
			"banana",
			[]interface{}{"passion_fruit"},
			nil,
			[][]byte{
				[]byte(`{
					"level": "ERROR",
					"message": "uneven number of args provided to PLogger",
					"span_id": "42",
					"trace_id": "69",
					"data": {
					  "number_of_args_provided": 1
					}
				}`),
				[]byte(`{
					"level": "ERROR",
					"message": "banana",
					"span_id": "42",
					"trace_id": "69",
					"data": {
						"args": ["passion_fruit"],
						"state": "invalid"
					}
				}`),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			if tt.spanID != nil {
				ctx = plog.StoreSpanID(ctx, *tt.spanID)
			}

			if tt.TraceID != nil {
				ctx = plog.StoreTraceID(ctx, *tt.TraceID)
			}

			slw := plogtest.SpyLogWriter{}
			ug := plogtest.MockUUIDGen{StringRetStrings: tt.uuidGenSet}
			plog.NewPLogger(&slw, &ug).Error(ctx, tt.msg, tt.args...)

			if len(tt.expFormattedStmts) != len(slw.PrintlnArgs) {
				t.Fatalf("expected arg count to be %v, was %v", len(tt.expFormattedStmts), len(slw.PrintlnArgs))
			}

			for i, args := range slw.PrintlnArgs {
				act, ok := args[0].(string)
				if !ok {
					t.Fatal("first arg to Printf is not a string")
				}

				d, err := diff.New().Compare(tt.expFormattedStmts[i], []byte(act))
				if err != nil {
					t.Fatal("error comparing log prints: ", err)
				}

				if d.Modified() {
					config := formatter.AsciiFormatterConfig{
						ShowArrayIndex: true,
						Coloring:       true,
					}

					var expMap map[string]interface{}
					if err := json.Unmarshal(tt.expFormattedStmts[i], &expMap); err != nil {
						t.Fatal(err)
					}

					formatter := formatter.NewAsciiFormatter(expMap, config)
					diffString, err := formatter.Format(d)
					if err != nil {
						t.Log("error getting diff string: ", err)
					}
					t.Fatalf("expected JSON and actual JSON are different: %s\n", diffString)
				}
			}
		})
	}
}
