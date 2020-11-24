// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/privacy/ent/task"
	"github.com/facebook/ent/entc/integration/privacy/ent/user"
	"github.com/google/uuid"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Status holds the value of the "status" field.
	Status task.Status `json:"status,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges     TaskEdges `json:"edges"`
	UserTasks *int
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Teams holds the value of the teams edge.
	Teams []*Team
	// Owner holds the value of the owner edge.
	Owner *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TeamsOrErr returns the Teams value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) TeamsOrErr() ([]*Team, error) {
	if e.loadedTypes[0] {
		return e.Teams, nil
	}
	return nil, &NotLoadedError{edge: "teams"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // title
		&sql.NullString{}, // description
		&sql.NullString{}, // status
		&uuid.UUID{},      // uuid
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Task) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // UserTasks
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(values ...interface{}) error {
	if m, n := len(values), len(task.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[0])
	} else if value.Valid {
		t.Title = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[1])
	} else if value.Valid {
		t.Description = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[2])
	} else if value.Valid {
		t.Status = task.Status(value.String)
	}
	if value, ok := values[3].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field uuid", values[3])
	} else if value != nil {
		t.UUID = *value
	}
	values = values[4:]
	if len(values) == len(task.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field UserTasks", value)
		} else if value.Valid {
			t.UserTasks = new(int)
			*t.UserTasks = int(value.Int64)
		}
	}
	return nil
}

// QueryTeams queries the teams edge of the Task.
func (t *Task) QueryTeams() *TeamQuery {
	return (&TaskClient{config: t.config}).QueryTeams(t)
}

// QueryOwner queries the owner edge of the Task.
func (t *Task) QueryOwner() *UserQuery {
	return (&TaskClient{config: t.config}).QueryOwner(t)
}

// Update returns a builder for updating this Task.
// Note that, you need to call Task.Unwrap() before calling this method, if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return (&TaskClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", title=")
	builder.WriteString(t.Title)
	builder.WriteString(", description=")
	builder.WriteString(t.Description)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", t.UUID))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

func (t Tasks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
