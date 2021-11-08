// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/agencysetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcouponsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/newuserrewardsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/purchaseinvitation"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/registrationinvitation"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userinvitationcode"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AgencySetting is the client for interacting with the AgencySetting builders.
	AgencySetting *AgencySettingClient
	// AppCouponSetting is the client for interacting with the AppCouponSetting builders.
	AppCouponSetting *AppCouponSettingClient
	// CouponAllocated is the client for interacting with the CouponAllocated builders.
	CouponAllocated *CouponAllocatedClient
	// CouponPool is the client for interacting with the CouponPool builders.
	CouponPool *CouponPoolClient
	// NewUserRewardSetting is the client for interacting with the NewUserRewardSetting builders.
	NewUserRewardSetting *NewUserRewardSettingClient
	// PurchaseInvitation is the client for interacting with the PurchaseInvitation builders.
	PurchaseInvitation *PurchaseInvitationClient
	// RegistrationInvitation is the client for interacting with the RegistrationInvitation builders.
	RegistrationInvitation *RegistrationInvitationClient
	// UserInvitationCode is the client for interacting with the UserInvitationCode builders.
	UserInvitationCode *UserInvitationCodeClient
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
	c.AgencySetting = NewAgencySettingClient(c.config)
	c.AppCouponSetting = NewAppCouponSettingClient(c.config)
	c.CouponAllocated = NewCouponAllocatedClient(c.config)
	c.CouponPool = NewCouponPoolClient(c.config)
	c.NewUserRewardSetting = NewNewUserRewardSettingClient(c.config)
	c.PurchaseInvitation = NewPurchaseInvitationClient(c.config)
	c.RegistrationInvitation = NewRegistrationInvitationClient(c.config)
	c.UserInvitationCode = NewUserInvitationCodeClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                    ctx,
		config:                 cfg,
		AgencySetting:          NewAgencySettingClient(cfg),
		AppCouponSetting:       NewAppCouponSettingClient(cfg),
		CouponAllocated:        NewCouponAllocatedClient(cfg),
		CouponPool:             NewCouponPoolClient(cfg),
		NewUserRewardSetting:   NewNewUserRewardSettingClient(cfg),
		PurchaseInvitation:     NewPurchaseInvitationClient(cfg),
		RegistrationInvitation: NewRegistrationInvitationClient(cfg),
		UserInvitationCode:     NewUserInvitationCodeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:                 cfg,
		AgencySetting:          NewAgencySettingClient(cfg),
		AppCouponSetting:       NewAppCouponSettingClient(cfg),
		CouponAllocated:        NewCouponAllocatedClient(cfg),
		CouponPool:             NewCouponPoolClient(cfg),
		NewUserRewardSetting:   NewNewUserRewardSettingClient(cfg),
		PurchaseInvitation:     NewPurchaseInvitationClient(cfg),
		RegistrationInvitation: NewRegistrationInvitationClient(cfg),
		UserInvitationCode:     NewUserInvitationCodeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AgencySetting.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.AgencySetting.Use(hooks...)
	c.AppCouponSetting.Use(hooks...)
	c.CouponAllocated.Use(hooks...)
	c.CouponPool.Use(hooks...)
	c.NewUserRewardSetting.Use(hooks...)
	c.PurchaseInvitation.Use(hooks...)
	c.RegistrationInvitation.Use(hooks...)
	c.UserInvitationCode.Use(hooks...)
}

// AgencySettingClient is a client for the AgencySetting schema.
type AgencySettingClient struct {
	config
}

// NewAgencySettingClient returns a client for the AgencySetting from the given config.
func NewAgencySettingClient(c config) *AgencySettingClient {
	return &AgencySettingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `agencysetting.Hooks(f(g(h())))`.
func (c *AgencySettingClient) Use(hooks ...Hook) {
	c.hooks.AgencySetting = append(c.hooks.AgencySetting, hooks...)
}

// Create returns a create builder for AgencySetting.
func (c *AgencySettingClient) Create() *AgencySettingCreate {
	mutation := newAgencySettingMutation(c.config, OpCreate)
	return &AgencySettingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AgencySetting entities.
func (c *AgencySettingClient) CreateBulk(builders ...*AgencySettingCreate) *AgencySettingCreateBulk {
	return &AgencySettingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AgencySetting.
func (c *AgencySettingClient) Update() *AgencySettingUpdate {
	mutation := newAgencySettingMutation(c.config, OpUpdate)
	return &AgencySettingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AgencySettingClient) UpdateOne(as *AgencySetting) *AgencySettingUpdateOne {
	mutation := newAgencySettingMutation(c.config, OpUpdateOne, withAgencySetting(as))
	return &AgencySettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AgencySettingClient) UpdateOneID(id int) *AgencySettingUpdateOne {
	mutation := newAgencySettingMutation(c.config, OpUpdateOne, withAgencySettingID(id))
	return &AgencySettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AgencySetting.
func (c *AgencySettingClient) Delete() *AgencySettingDelete {
	mutation := newAgencySettingMutation(c.config, OpDelete)
	return &AgencySettingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AgencySettingClient) DeleteOne(as *AgencySetting) *AgencySettingDeleteOne {
	return c.DeleteOneID(as.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AgencySettingClient) DeleteOneID(id int) *AgencySettingDeleteOne {
	builder := c.Delete().Where(agencysetting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AgencySettingDeleteOne{builder}
}

// Query returns a query builder for AgencySetting.
func (c *AgencySettingClient) Query() *AgencySettingQuery {
	return &AgencySettingQuery{
		config: c.config,
	}
}

// Get returns a AgencySetting entity by its id.
func (c *AgencySettingClient) Get(ctx context.Context, id int) (*AgencySetting, error) {
	return c.Query().Where(agencysetting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AgencySettingClient) GetX(ctx context.Context, id int) *AgencySetting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AgencySettingClient) Hooks() []Hook {
	return c.hooks.AgencySetting
}

// AppCouponSettingClient is a client for the AppCouponSetting schema.
type AppCouponSettingClient struct {
	config
}

// NewAppCouponSettingClient returns a client for the AppCouponSetting from the given config.
func NewAppCouponSettingClient(c config) *AppCouponSettingClient {
	return &AppCouponSettingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `appcouponsetting.Hooks(f(g(h())))`.
func (c *AppCouponSettingClient) Use(hooks ...Hook) {
	c.hooks.AppCouponSetting = append(c.hooks.AppCouponSetting, hooks...)
}

// Create returns a create builder for AppCouponSetting.
func (c *AppCouponSettingClient) Create() *AppCouponSettingCreate {
	mutation := newAppCouponSettingMutation(c.config, OpCreate)
	return &AppCouponSettingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AppCouponSetting entities.
func (c *AppCouponSettingClient) CreateBulk(builders ...*AppCouponSettingCreate) *AppCouponSettingCreateBulk {
	return &AppCouponSettingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AppCouponSetting.
func (c *AppCouponSettingClient) Update() *AppCouponSettingUpdate {
	mutation := newAppCouponSettingMutation(c.config, OpUpdate)
	return &AppCouponSettingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppCouponSettingClient) UpdateOne(acs *AppCouponSetting) *AppCouponSettingUpdateOne {
	mutation := newAppCouponSettingMutation(c.config, OpUpdateOne, withAppCouponSetting(acs))
	return &AppCouponSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppCouponSettingClient) UpdateOneID(id int) *AppCouponSettingUpdateOne {
	mutation := newAppCouponSettingMutation(c.config, OpUpdateOne, withAppCouponSettingID(id))
	return &AppCouponSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AppCouponSetting.
func (c *AppCouponSettingClient) Delete() *AppCouponSettingDelete {
	mutation := newAppCouponSettingMutation(c.config, OpDelete)
	return &AppCouponSettingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AppCouponSettingClient) DeleteOne(acs *AppCouponSetting) *AppCouponSettingDeleteOne {
	return c.DeleteOneID(acs.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AppCouponSettingClient) DeleteOneID(id int) *AppCouponSettingDeleteOne {
	builder := c.Delete().Where(appcouponsetting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppCouponSettingDeleteOne{builder}
}

// Query returns a query builder for AppCouponSetting.
func (c *AppCouponSettingClient) Query() *AppCouponSettingQuery {
	return &AppCouponSettingQuery{
		config: c.config,
	}
}

// Get returns a AppCouponSetting entity by its id.
func (c *AppCouponSettingClient) Get(ctx context.Context, id int) (*AppCouponSetting, error) {
	return c.Query().Where(appcouponsetting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppCouponSettingClient) GetX(ctx context.Context, id int) *AppCouponSetting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppCouponSettingClient) Hooks() []Hook {
	return c.hooks.AppCouponSetting
}

// CouponAllocatedClient is a client for the CouponAllocated schema.
type CouponAllocatedClient struct {
	config
}

// NewCouponAllocatedClient returns a client for the CouponAllocated from the given config.
func NewCouponAllocatedClient(c config) *CouponAllocatedClient {
	return &CouponAllocatedClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `couponallocated.Hooks(f(g(h())))`.
func (c *CouponAllocatedClient) Use(hooks ...Hook) {
	c.hooks.CouponAllocated = append(c.hooks.CouponAllocated, hooks...)
}

// Create returns a create builder for CouponAllocated.
func (c *CouponAllocatedClient) Create() *CouponAllocatedCreate {
	mutation := newCouponAllocatedMutation(c.config, OpCreate)
	return &CouponAllocatedCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CouponAllocated entities.
func (c *CouponAllocatedClient) CreateBulk(builders ...*CouponAllocatedCreate) *CouponAllocatedCreateBulk {
	return &CouponAllocatedCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CouponAllocated.
func (c *CouponAllocatedClient) Update() *CouponAllocatedUpdate {
	mutation := newCouponAllocatedMutation(c.config, OpUpdate)
	return &CouponAllocatedUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CouponAllocatedClient) UpdateOne(ca *CouponAllocated) *CouponAllocatedUpdateOne {
	mutation := newCouponAllocatedMutation(c.config, OpUpdateOne, withCouponAllocated(ca))
	return &CouponAllocatedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CouponAllocatedClient) UpdateOneID(id uuid.UUID) *CouponAllocatedUpdateOne {
	mutation := newCouponAllocatedMutation(c.config, OpUpdateOne, withCouponAllocatedID(id))
	return &CouponAllocatedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CouponAllocated.
func (c *CouponAllocatedClient) Delete() *CouponAllocatedDelete {
	mutation := newCouponAllocatedMutation(c.config, OpDelete)
	return &CouponAllocatedDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CouponAllocatedClient) DeleteOne(ca *CouponAllocated) *CouponAllocatedDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CouponAllocatedClient) DeleteOneID(id uuid.UUID) *CouponAllocatedDeleteOne {
	builder := c.Delete().Where(couponallocated.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CouponAllocatedDeleteOne{builder}
}

// Query returns a query builder for CouponAllocated.
func (c *CouponAllocatedClient) Query() *CouponAllocatedQuery {
	return &CouponAllocatedQuery{
		config: c.config,
	}
}

// Get returns a CouponAllocated entity by its id.
func (c *CouponAllocatedClient) Get(ctx context.Context, id uuid.UUID) (*CouponAllocated, error) {
	return c.Query().Where(couponallocated.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CouponAllocatedClient) GetX(ctx context.Context, id uuid.UUID) *CouponAllocated {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CouponAllocatedClient) Hooks() []Hook {
	return c.hooks.CouponAllocated
}

// CouponPoolClient is a client for the CouponPool schema.
type CouponPoolClient struct {
	config
}

// NewCouponPoolClient returns a client for the CouponPool from the given config.
func NewCouponPoolClient(c config) *CouponPoolClient {
	return &CouponPoolClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `couponpool.Hooks(f(g(h())))`.
func (c *CouponPoolClient) Use(hooks ...Hook) {
	c.hooks.CouponPool = append(c.hooks.CouponPool, hooks...)
}

// Create returns a create builder for CouponPool.
func (c *CouponPoolClient) Create() *CouponPoolCreate {
	mutation := newCouponPoolMutation(c.config, OpCreate)
	return &CouponPoolCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CouponPool entities.
func (c *CouponPoolClient) CreateBulk(builders ...*CouponPoolCreate) *CouponPoolCreateBulk {
	return &CouponPoolCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CouponPool.
func (c *CouponPoolClient) Update() *CouponPoolUpdate {
	mutation := newCouponPoolMutation(c.config, OpUpdate)
	return &CouponPoolUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CouponPoolClient) UpdateOne(cp *CouponPool) *CouponPoolUpdateOne {
	mutation := newCouponPoolMutation(c.config, OpUpdateOne, withCouponPool(cp))
	return &CouponPoolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CouponPoolClient) UpdateOneID(id uuid.UUID) *CouponPoolUpdateOne {
	mutation := newCouponPoolMutation(c.config, OpUpdateOne, withCouponPoolID(id))
	return &CouponPoolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CouponPool.
func (c *CouponPoolClient) Delete() *CouponPoolDelete {
	mutation := newCouponPoolMutation(c.config, OpDelete)
	return &CouponPoolDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CouponPoolClient) DeleteOne(cp *CouponPool) *CouponPoolDeleteOne {
	return c.DeleteOneID(cp.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CouponPoolClient) DeleteOneID(id uuid.UUID) *CouponPoolDeleteOne {
	builder := c.Delete().Where(couponpool.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CouponPoolDeleteOne{builder}
}

// Query returns a query builder for CouponPool.
func (c *CouponPoolClient) Query() *CouponPoolQuery {
	return &CouponPoolQuery{
		config: c.config,
	}
}

// Get returns a CouponPool entity by its id.
func (c *CouponPoolClient) Get(ctx context.Context, id uuid.UUID) (*CouponPool, error) {
	return c.Query().Where(couponpool.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CouponPoolClient) GetX(ctx context.Context, id uuid.UUID) *CouponPool {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CouponPoolClient) Hooks() []Hook {
	return c.hooks.CouponPool
}

// NewUserRewardSettingClient is a client for the NewUserRewardSetting schema.
type NewUserRewardSettingClient struct {
	config
}

// NewNewUserRewardSettingClient returns a client for the NewUserRewardSetting from the given config.
func NewNewUserRewardSettingClient(c config) *NewUserRewardSettingClient {
	return &NewUserRewardSettingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `newuserrewardsetting.Hooks(f(g(h())))`.
func (c *NewUserRewardSettingClient) Use(hooks ...Hook) {
	c.hooks.NewUserRewardSetting = append(c.hooks.NewUserRewardSetting, hooks...)
}

// Create returns a create builder for NewUserRewardSetting.
func (c *NewUserRewardSettingClient) Create() *NewUserRewardSettingCreate {
	mutation := newNewUserRewardSettingMutation(c.config, OpCreate)
	return &NewUserRewardSettingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of NewUserRewardSetting entities.
func (c *NewUserRewardSettingClient) CreateBulk(builders ...*NewUserRewardSettingCreate) *NewUserRewardSettingCreateBulk {
	return &NewUserRewardSettingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for NewUserRewardSetting.
func (c *NewUserRewardSettingClient) Update() *NewUserRewardSettingUpdate {
	mutation := newNewUserRewardSettingMutation(c.config, OpUpdate)
	return &NewUserRewardSettingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NewUserRewardSettingClient) UpdateOne(nurs *NewUserRewardSetting) *NewUserRewardSettingUpdateOne {
	mutation := newNewUserRewardSettingMutation(c.config, OpUpdateOne, withNewUserRewardSetting(nurs))
	return &NewUserRewardSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NewUserRewardSettingClient) UpdateOneID(id int) *NewUserRewardSettingUpdateOne {
	mutation := newNewUserRewardSettingMutation(c.config, OpUpdateOne, withNewUserRewardSettingID(id))
	return &NewUserRewardSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for NewUserRewardSetting.
func (c *NewUserRewardSettingClient) Delete() *NewUserRewardSettingDelete {
	mutation := newNewUserRewardSettingMutation(c.config, OpDelete)
	return &NewUserRewardSettingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NewUserRewardSettingClient) DeleteOne(nurs *NewUserRewardSetting) *NewUserRewardSettingDeleteOne {
	return c.DeleteOneID(nurs.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NewUserRewardSettingClient) DeleteOneID(id int) *NewUserRewardSettingDeleteOne {
	builder := c.Delete().Where(newuserrewardsetting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NewUserRewardSettingDeleteOne{builder}
}

// Query returns a query builder for NewUserRewardSetting.
func (c *NewUserRewardSettingClient) Query() *NewUserRewardSettingQuery {
	return &NewUserRewardSettingQuery{
		config: c.config,
	}
}

// Get returns a NewUserRewardSetting entity by its id.
func (c *NewUserRewardSettingClient) Get(ctx context.Context, id int) (*NewUserRewardSetting, error) {
	return c.Query().Where(newuserrewardsetting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NewUserRewardSettingClient) GetX(ctx context.Context, id int) *NewUserRewardSetting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NewUserRewardSettingClient) Hooks() []Hook {
	return c.hooks.NewUserRewardSetting
}

// PurchaseInvitationClient is a client for the PurchaseInvitation schema.
type PurchaseInvitationClient struct {
	config
}

// NewPurchaseInvitationClient returns a client for the PurchaseInvitation from the given config.
func NewPurchaseInvitationClient(c config) *PurchaseInvitationClient {
	return &PurchaseInvitationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `purchaseinvitation.Hooks(f(g(h())))`.
func (c *PurchaseInvitationClient) Use(hooks ...Hook) {
	c.hooks.PurchaseInvitation = append(c.hooks.PurchaseInvitation, hooks...)
}

// Create returns a create builder for PurchaseInvitation.
func (c *PurchaseInvitationClient) Create() *PurchaseInvitationCreate {
	mutation := newPurchaseInvitationMutation(c.config, OpCreate)
	return &PurchaseInvitationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PurchaseInvitation entities.
func (c *PurchaseInvitationClient) CreateBulk(builders ...*PurchaseInvitationCreate) *PurchaseInvitationCreateBulk {
	return &PurchaseInvitationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PurchaseInvitation.
func (c *PurchaseInvitationClient) Update() *PurchaseInvitationUpdate {
	mutation := newPurchaseInvitationMutation(c.config, OpUpdate)
	return &PurchaseInvitationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PurchaseInvitationClient) UpdateOne(pi *PurchaseInvitation) *PurchaseInvitationUpdateOne {
	mutation := newPurchaseInvitationMutation(c.config, OpUpdateOne, withPurchaseInvitation(pi))
	return &PurchaseInvitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PurchaseInvitationClient) UpdateOneID(id int) *PurchaseInvitationUpdateOne {
	mutation := newPurchaseInvitationMutation(c.config, OpUpdateOne, withPurchaseInvitationID(id))
	return &PurchaseInvitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PurchaseInvitation.
func (c *PurchaseInvitationClient) Delete() *PurchaseInvitationDelete {
	mutation := newPurchaseInvitationMutation(c.config, OpDelete)
	return &PurchaseInvitationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PurchaseInvitationClient) DeleteOne(pi *PurchaseInvitation) *PurchaseInvitationDeleteOne {
	return c.DeleteOneID(pi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PurchaseInvitationClient) DeleteOneID(id int) *PurchaseInvitationDeleteOne {
	builder := c.Delete().Where(purchaseinvitation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PurchaseInvitationDeleteOne{builder}
}

// Query returns a query builder for PurchaseInvitation.
func (c *PurchaseInvitationClient) Query() *PurchaseInvitationQuery {
	return &PurchaseInvitationQuery{
		config: c.config,
	}
}

// Get returns a PurchaseInvitation entity by its id.
func (c *PurchaseInvitationClient) Get(ctx context.Context, id int) (*PurchaseInvitation, error) {
	return c.Query().Where(purchaseinvitation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PurchaseInvitationClient) GetX(ctx context.Context, id int) *PurchaseInvitation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PurchaseInvitationClient) Hooks() []Hook {
	return c.hooks.PurchaseInvitation
}

// RegistrationInvitationClient is a client for the RegistrationInvitation schema.
type RegistrationInvitationClient struct {
	config
}

// NewRegistrationInvitationClient returns a client for the RegistrationInvitation from the given config.
func NewRegistrationInvitationClient(c config) *RegistrationInvitationClient {
	return &RegistrationInvitationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `registrationinvitation.Hooks(f(g(h())))`.
func (c *RegistrationInvitationClient) Use(hooks ...Hook) {
	c.hooks.RegistrationInvitation = append(c.hooks.RegistrationInvitation, hooks...)
}

// Create returns a create builder for RegistrationInvitation.
func (c *RegistrationInvitationClient) Create() *RegistrationInvitationCreate {
	mutation := newRegistrationInvitationMutation(c.config, OpCreate)
	return &RegistrationInvitationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RegistrationInvitation entities.
func (c *RegistrationInvitationClient) CreateBulk(builders ...*RegistrationInvitationCreate) *RegistrationInvitationCreateBulk {
	return &RegistrationInvitationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RegistrationInvitation.
func (c *RegistrationInvitationClient) Update() *RegistrationInvitationUpdate {
	mutation := newRegistrationInvitationMutation(c.config, OpUpdate)
	return &RegistrationInvitationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RegistrationInvitationClient) UpdateOne(ri *RegistrationInvitation) *RegistrationInvitationUpdateOne {
	mutation := newRegistrationInvitationMutation(c.config, OpUpdateOne, withRegistrationInvitation(ri))
	return &RegistrationInvitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RegistrationInvitationClient) UpdateOneID(id int) *RegistrationInvitationUpdateOne {
	mutation := newRegistrationInvitationMutation(c.config, OpUpdateOne, withRegistrationInvitationID(id))
	return &RegistrationInvitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RegistrationInvitation.
func (c *RegistrationInvitationClient) Delete() *RegistrationInvitationDelete {
	mutation := newRegistrationInvitationMutation(c.config, OpDelete)
	return &RegistrationInvitationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RegistrationInvitationClient) DeleteOne(ri *RegistrationInvitation) *RegistrationInvitationDeleteOne {
	return c.DeleteOneID(ri.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RegistrationInvitationClient) DeleteOneID(id int) *RegistrationInvitationDeleteOne {
	builder := c.Delete().Where(registrationinvitation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RegistrationInvitationDeleteOne{builder}
}

// Query returns a query builder for RegistrationInvitation.
func (c *RegistrationInvitationClient) Query() *RegistrationInvitationQuery {
	return &RegistrationInvitationQuery{
		config: c.config,
	}
}

// Get returns a RegistrationInvitation entity by its id.
func (c *RegistrationInvitationClient) Get(ctx context.Context, id int) (*RegistrationInvitation, error) {
	return c.Query().Where(registrationinvitation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RegistrationInvitationClient) GetX(ctx context.Context, id int) *RegistrationInvitation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RegistrationInvitationClient) Hooks() []Hook {
	return c.hooks.RegistrationInvitation
}

// UserInvitationCodeClient is a client for the UserInvitationCode schema.
type UserInvitationCodeClient struct {
	config
}

// NewUserInvitationCodeClient returns a client for the UserInvitationCode from the given config.
func NewUserInvitationCodeClient(c config) *UserInvitationCodeClient {
	return &UserInvitationCodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userinvitationcode.Hooks(f(g(h())))`.
func (c *UserInvitationCodeClient) Use(hooks ...Hook) {
	c.hooks.UserInvitationCode = append(c.hooks.UserInvitationCode, hooks...)
}

// Create returns a create builder for UserInvitationCode.
func (c *UserInvitationCodeClient) Create() *UserInvitationCodeCreate {
	mutation := newUserInvitationCodeMutation(c.config, OpCreate)
	return &UserInvitationCodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserInvitationCode entities.
func (c *UserInvitationCodeClient) CreateBulk(builders ...*UserInvitationCodeCreate) *UserInvitationCodeCreateBulk {
	return &UserInvitationCodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserInvitationCode.
func (c *UserInvitationCodeClient) Update() *UserInvitationCodeUpdate {
	mutation := newUserInvitationCodeMutation(c.config, OpUpdate)
	return &UserInvitationCodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserInvitationCodeClient) UpdateOne(uic *UserInvitationCode) *UserInvitationCodeUpdateOne {
	mutation := newUserInvitationCodeMutation(c.config, OpUpdateOne, withUserInvitationCode(uic))
	return &UserInvitationCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserInvitationCodeClient) UpdateOneID(id uuid.UUID) *UserInvitationCodeUpdateOne {
	mutation := newUserInvitationCodeMutation(c.config, OpUpdateOne, withUserInvitationCodeID(id))
	return &UserInvitationCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserInvitationCode.
func (c *UserInvitationCodeClient) Delete() *UserInvitationCodeDelete {
	mutation := newUserInvitationCodeMutation(c.config, OpDelete)
	return &UserInvitationCodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserInvitationCodeClient) DeleteOne(uic *UserInvitationCode) *UserInvitationCodeDeleteOne {
	return c.DeleteOneID(uic.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserInvitationCodeClient) DeleteOneID(id uuid.UUID) *UserInvitationCodeDeleteOne {
	builder := c.Delete().Where(userinvitationcode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserInvitationCodeDeleteOne{builder}
}

// Query returns a query builder for UserInvitationCode.
func (c *UserInvitationCodeClient) Query() *UserInvitationCodeQuery {
	return &UserInvitationCodeQuery{
		config: c.config,
	}
}

// Get returns a UserInvitationCode entity by its id.
func (c *UserInvitationCodeClient) Get(ctx context.Context, id uuid.UUID) (*UserInvitationCode, error) {
	return c.Query().Where(userinvitationcode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserInvitationCodeClient) GetX(ctx context.Context, id uuid.UUID) *UserInvitationCode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserInvitationCodeClient) Hooks() []Hook {
	return c.hooks.UserInvitationCode
}
