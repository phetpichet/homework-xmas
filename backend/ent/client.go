// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Piichet/app/ent/migrate"

	"github.com/Piichet/app/ent/department"
	"github.com/Piichet/app/ent/doctor"
	"github.com/Piichet/app/ent/office"
	"github.com/Piichet/app/ent/specialist"
	"github.com/Piichet/app/ent/working"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Department is the client for interacting with the Department builders.
	Department *DepartmentClient
	// Doctor is the client for interacting with the Doctor builders.
	Doctor *DoctorClient
	// Office is the client for interacting with the Office builders.
	Office *OfficeClient
	// Specialist is the client for interacting with the Specialist builders.
	Specialist *SpecialistClient
	// Working is the client for interacting with the Working builders.
	Working *WorkingClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Department = NewDepartmentClient(c.config)
	c.Doctor = NewDoctorClient(c.config)
	c.Office = NewOfficeClient(c.config)
	c.Specialist = NewSpecialistClient(c.config)
	c.Working = NewWorkingClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Department: NewDepartmentClient(cfg),
		Doctor:     NewDoctorClient(cfg),
		Office:     NewOfficeClient(cfg),
		Specialist: NewSpecialistClient(cfg),
		Working:    NewWorkingClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:     cfg,
		Department: NewDepartmentClient(cfg),
		Doctor:     NewDoctorClient(cfg),
		Office:     NewOfficeClient(cfg),
		Specialist: NewSpecialistClient(cfg),
		Working:    NewWorkingClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Department.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Department.Use(hooks...)
	c.Doctor.Use(hooks...)
	c.Office.Use(hooks...)
	c.Specialist.Use(hooks...)
	c.Working.Use(hooks...)
}

// DepartmentClient is a client for the Department schema.
type DepartmentClient struct {
	config
}

// NewDepartmentClient returns a client for the Department from the given config.
func NewDepartmentClient(c config) *DepartmentClient {
	return &DepartmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `department.Hooks(f(g(h())))`.
func (c *DepartmentClient) Use(hooks ...Hook) {
	c.hooks.Department = append(c.hooks.Department, hooks...)
}

// Create returns a create builder for Department.
func (c *DepartmentClient) Create() *DepartmentCreate {
	mutation := newDepartmentMutation(c.config, OpCreate)
	return &DepartmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Department.
func (c *DepartmentClient) Update() *DepartmentUpdate {
	mutation := newDepartmentMutation(c.config, OpUpdate)
	return &DepartmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DepartmentClient) UpdateOne(d *Department) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartment(d))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DepartmentClient) UpdateOneID(id int) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartmentID(id))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Department.
func (c *DepartmentClient) Delete() *DepartmentDelete {
	mutation := newDepartmentMutation(c.config, OpDelete)
	return &DepartmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DepartmentClient) DeleteOne(d *Department) *DepartmentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DepartmentClient) DeleteOneID(id int) *DepartmentDeleteOne {
	builder := c.Delete().Where(department.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DepartmentDeleteOne{builder}
}

// Create returns a query builder for Department.
func (c *DepartmentClient) Query() *DepartmentQuery {
	return &DepartmentQuery{config: c.config}
}

// Get returns a Department entity by its id.
func (c *DepartmentClient) Get(ctx context.Context, id int) (*Department, error) {
	return c.Query().Where(department.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DepartmentClient) GetX(ctx context.Context, id int) *Department {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// Hooks returns the client hooks.
func (c *DepartmentClient) Hooks() []Hook {
	return c.hooks.Department
}

// DoctorClient is a client for the Doctor schema.
type DoctorClient struct {
	config
}

// NewDoctorClient returns a client for the Doctor from the given config.
func NewDoctorClient(c config) *DoctorClient {
	return &DoctorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `doctor.Hooks(f(g(h())))`.
func (c *DoctorClient) Use(hooks ...Hook) {
	c.hooks.Doctor = append(c.hooks.Doctor, hooks...)
}

// Create returns a create builder for Doctor.
func (c *DoctorClient) Create() *DoctorCreate {
	mutation := newDoctorMutation(c.config, OpCreate)
	return &DoctorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Doctor.
func (c *DoctorClient) Update() *DoctorUpdate {
	mutation := newDoctorMutation(c.config, OpUpdate)
	return &DoctorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DoctorClient) UpdateOne(d *Doctor) *DoctorUpdateOne {
	mutation := newDoctorMutation(c.config, OpUpdateOne, withDoctor(d))
	return &DoctorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DoctorClient) UpdateOneID(id int) *DoctorUpdateOne {
	mutation := newDoctorMutation(c.config, OpUpdateOne, withDoctorID(id))
	return &DoctorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Doctor.
func (c *DoctorClient) Delete() *DoctorDelete {
	mutation := newDoctorMutation(c.config, OpDelete)
	return &DoctorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DoctorClient) DeleteOne(d *Doctor) *DoctorDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DoctorClient) DeleteOneID(id int) *DoctorDeleteOne {
	builder := c.Delete().Where(doctor.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DoctorDeleteOne{builder}
}

// Create returns a query builder for Doctor.
func (c *DoctorClient) Query() *DoctorQuery {
	return &DoctorQuery{config: c.config}
}

// Get returns a Doctor entity by its id.
func (c *DoctorClient) Get(ctx context.Context, id int) (*Doctor, error) {
	return c.Query().Where(doctor.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DoctorClient) GetX(ctx context.Context, id int) *Doctor {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// Hooks returns the client hooks.
func (c *DoctorClient) Hooks() []Hook {
	return c.hooks.Doctor
}

// OfficeClient is a client for the Office schema.
type OfficeClient struct {
	config
}

// NewOfficeClient returns a client for the Office from the given config.
func NewOfficeClient(c config) *OfficeClient {
	return &OfficeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `office.Hooks(f(g(h())))`.
func (c *OfficeClient) Use(hooks ...Hook) {
	c.hooks.Office = append(c.hooks.Office, hooks...)
}

// Create returns a create builder for Office.
func (c *OfficeClient) Create() *OfficeCreate {
	mutation := newOfficeMutation(c.config, OpCreate)
	return &OfficeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Office.
func (c *OfficeClient) Update() *OfficeUpdate {
	mutation := newOfficeMutation(c.config, OpUpdate)
	return &OfficeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OfficeClient) UpdateOne(o *Office) *OfficeUpdateOne {
	mutation := newOfficeMutation(c.config, OpUpdateOne, withOffice(o))
	return &OfficeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OfficeClient) UpdateOneID(id int) *OfficeUpdateOne {
	mutation := newOfficeMutation(c.config, OpUpdateOne, withOfficeID(id))
	return &OfficeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Office.
func (c *OfficeClient) Delete() *OfficeDelete {
	mutation := newOfficeMutation(c.config, OpDelete)
	return &OfficeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OfficeClient) DeleteOne(o *Office) *OfficeDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OfficeClient) DeleteOneID(id int) *OfficeDeleteOne {
	builder := c.Delete().Where(office.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OfficeDeleteOne{builder}
}

// Create returns a query builder for Office.
func (c *OfficeClient) Query() *OfficeQuery {
	return &OfficeQuery{config: c.config}
}

// Get returns a Office entity by its id.
func (c *OfficeClient) Get(ctx context.Context, id int) (*Office, error) {
	return c.Query().Where(office.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OfficeClient) GetX(ctx context.Context, id int) *Office {
	o, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return o
}

// Hooks returns the client hooks.
func (c *OfficeClient) Hooks() []Hook {
	return c.hooks.Office
}

// SpecialistClient is a client for the Specialist schema.
type SpecialistClient struct {
	config
}

// NewSpecialistClient returns a client for the Specialist from the given config.
func NewSpecialistClient(c config) *SpecialistClient {
	return &SpecialistClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `specialist.Hooks(f(g(h())))`.
func (c *SpecialistClient) Use(hooks ...Hook) {
	c.hooks.Specialist = append(c.hooks.Specialist, hooks...)
}

// Create returns a create builder for Specialist.
func (c *SpecialistClient) Create() *SpecialistCreate {
	mutation := newSpecialistMutation(c.config, OpCreate)
	return &SpecialistCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Specialist.
func (c *SpecialistClient) Update() *SpecialistUpdate {
	mutation := newSpecialistMutation(c.config, OpUpdate)
	return &SpecialistUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SpecialistClient) UpdateOne(s *Specialist) *SpecialistUpdateOne {
	mutation := newSpecialistMutation(c.config, OpUpdateOne, withSpecialist(s))
	return &SpecialistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SpecialistClient) UpdateOneID(id int) *SpecialistUpdateOne {
	mutation := newSpecialistMutation(c.config, OpUpdateOne, withSpecialistID(id))
	return &SpecialistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Specialist.
func (c *SpecialistClient) Delete() *SpecialistDelete {
	mutation := newSpecialistMutation(c.config, OpDelete)
	return &SpecialistDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SpecialistClient) DeleteOne(s *Specialist) *SpecialistDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SpecialistClient) DeleteOneID(id int) *SpecialistDeleteOne {
	builder := c.Delete().Where(specialist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SpecialistDeleteOne{builder}
}

// Create returns a query builder for Specialist.
func (c *SpecialistClient) Query() *SpecialistQuery {
	return &SpecialistQuery{config: c.config}
}

// Get returns a Specialist entity by its id.
func (c *SpecialistClient) Get(ctx context.Context, id int) (*Specialist, error) {
	return c.Query().Where(specialist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SpecialistClient) GetX(ctx context.Context, id int) *Specialist {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// Hooks returns the client hooks.
func (c *SpecialistClient) Hooks() []Hook {
	return c.hooks.Specialist
}

// WorkingClient is a client for the Working schema.
type WorkingClient struct {
	config
}

// NewWorkingClient returns a client for the Working from the given config.
func NewWorkingClient(c config) *WorkingClient {
	return &WorkingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `working.Hooks(f(g(h())))`.
func (c *WorkingClient) Use(hooks ...Hook) {
	c.hooks.Working = append(c.hooks.Working, hooks...)
}

// Create returns a create builder for Working.
func (c *WorkingClient) Create() *WorkingCreate {
	mutation := newWorkingMutation(c.config, OpCreate)
	return &WorkingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Working.
func (c *WorkingClient) Update() *WorkingUpdate {
	mutation := newWorkingMutation(c.config, OpUpdate)
	return &WorkingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkingClient) UpdateOne(w *Working) *WorkingUpdateOne {
	mutation := newWorkingMutation(c.config, OpUpdateOne, withWorking(w))
	return &WorkingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkingClient) UpdateOneID(id int) *WorkingUpdateOne {
	mutation := newWorkingMutation(c.config, OpUpdateOne, withWorkingID(id))
	return &WorkingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Working.
func (c *WorkingClient) Delete() *WorkingDelete {
	mutation := newWorkingMutation(c.config, OpDelete)
	return &WorkingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WorkingClient) DeleteOne(w *Working) *WorkingDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WorkingClient) DeleteOneID(id int) *WorkingDeleteOne {
	builder := c.Delete().Where(working.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkingDeleteOne{builder}
}

// Create returns a query builder for Working.
func (c *WorkingClient) Query() *WorkingQuery {
	return &WorkingQuery{config: c.config}
}

// Get returns a Working entity by its id.
func (c *WorkingClient) Get(ctx context.Context, id int) (*Working, error) {
	return c.Query().Where(working.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkingClient) GetX(ctx context.Context, id int) *Working {
	w, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return w
}

// Hooks returns the client hooks.
func (c *WorkingClient) Hooks() []Hook {
	return c.hooks.Working
}
