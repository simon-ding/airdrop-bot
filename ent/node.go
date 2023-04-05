// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airdrop-bot/ent/node"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Node is the model entity for the Node schema.
type Node struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Region holds the value of the "region" field.
	Region string `json:"region,omitempty"`
	// DnsName holds the value of the "dnsName" field.
	DnsName string `json:"dnsName,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NodeQuery when eager-loading is set.
	Edges NodeEdges `json:"edges"`
}

// NodeEdges holds the relations/edges for other nodes in the graph.
type NodeEdges struct {
	// StepRuns holds the value of the step_runs edge.
	StepRuns []*StepRun `json:"step_runs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StepRunsOrErr returns the StepRuns value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) StepRunsOrErr() ([]*StepRun, error) {
	if e.loadedTypes[0] {
		return e.StepRuns, nil
	}
	return nil, &NotLoadedError{edge: "step_runs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Node) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case node.FieldID:
			values[i] = new(sql.NullInt64)
		case node.FieldName, node.FieldRegion, node.FieldDnsName, node.FieldAddress:
			values[i] = new(sql.NullString)
		case node.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Node", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Node fields.
func (n *Node) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case node.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case node.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				n.Name = value.String
			}
		case node.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				n.Region = value.String
			}
		case node.FieldDnsName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dnsName", values[i])
			} else if value.Valid {
				n.DnsName = value.String
			}
		case node.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				n.Address = value.String
			}
		case node.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryStepRuns queries the "step_runs" edge of the Node entity.
func (n *Node) QueryStepRuns() *StepRunQuery {
	return (&NodeClient{config: n.config}).QueryStepRuns(n)
}

// Update returns a builder for updating this Node.
// Note that you need to call Node.Unwrap() before calling this method if this Node
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Node) Update() *NodeUpdateOne {
	return (&NodeClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Node entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Node) Unwrap() *Node {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Node is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString("Node(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", name=")
	builder.WriteString(n.Name)
	builder.WriteString(", region=")
	builder.WriteString(n.Region)
	builder.WriteString(", dnsName=")
	builder.WriteString(n.DnsName)
	builder.WriteString(", address=")
	builder.WriteString(n.Address)
	builder.WriteString(", updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Nodes is a parsable slice of Node.
type Nodes []*Node

func (n Nodes) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
