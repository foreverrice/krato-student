// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"finance/internal/data/ent/predicate"
	"finance/internal/data/ent/storebalanceaccount"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreBalanceAccountUpdate is the builder for updating StoreBalanceAccount entities.
type StoreBalanceAccountUpdate struct {
	config
	hooks    []Hook
	mutation *StoreBalanceAccountMutation
}

// Where appends a list predicates to the StoreBalanceAccountUpdate builder.
func (sbau *StoreBalanceAccountUpdate) Where(ps ...predicate.StoreBalanceAccount) *StoreBalanceAccountUpdate {
	sbau.mutation.Where(ps...)
	return sbau
}

// SetAccountNo sets the "account_no" field.
func (sbau *StoreBalanceAccountUpdate) SetAccountNo(s string) *StoreBalanceAccountUpdate {
	sbau.mutation.SetAccountNo(s)
	return sbau
}

// SetNillableAccountNo sets the "account_no" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillableAccountNo(s *string) *StoreBalanceAccountUpdate {
	if s != nil {
		sbau.SetAccountNo(*s)
	}
	return sbau
}

// ClearAccountNo clears the value of the "account_no" field.
func (sbau *StoreBalanceAccountUpdate) ClearAccountNo() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearAccountNo()
	return sbau
}

// SetStoreCode sets the "store_code" field.
func (sbau *StoreBalanceAccountUpdate) SetStoreCode(s string) *StoreBalanceAccountUpdate {
	sbau.mutation.SetStoreCode(s)
	return sbau
}

// SetNillableStoreCode sets the "store_code" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillableStoreCode(s *string) *StoreBalanceAccountUpdate {
	if s != nil {
		sbau.SetStoreCode(*s)
	}
	return sbau
}

// ClearStoreCode clears the value of the "store_code" field.
func (sbau *StoreBalanceAccountUpdate) ClearStoreCode() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearStoreCode()
	return sbau
}

// SetUpperOrganNo sets the "upper_organ_no" field.
func (sbau *StoreBalanceAccountUpdate) SetUpperOrganNo(s string) *StoreBalanceAccountUpdate {
	sbau.mutation.SetUpperOrganNo(s)
	return sbau
}

// SetNillableUpperOrganNo sets the "upper_organ_no" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillableUpperOrganNo(s *string) *StoreBalanceAccountUpdate {
	if s != nil {
		sbau.SetUpperOrganNo(*s)
	}
	return sbau
}

// ClearUpperOrganNo clears the value of the "upper_organ_no" field.
func (sbau *StoreBalanceAccountUpdate) ClearUpperOrganNo() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearUpperOrganNo()
	return sbau
}

// SetPwd sets the "pwd" field.
func (sbau *StoreBalanceAccountUpdate) SetPwd(s string) *StoreBalanceAccountUpdate {
	sbau.mutation.SetPwd(s)
	return sbau
}

// SetNillablePwd sets the "pwd" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillablePwd(s *string) *StoreBalanceAccountUpdate {
	if s != nil {
		sbau.SetPwd(*s)
	}
	return sbau
}

// ClearPwd clears the value of the "pwd" field.
func (sbau *StoreBalanceAccountUpdate) ClearPwd() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearPwd()
	return sbau
}

// SetPwdSalt sets the "pwd_salt" field.
func (sbau *StoreBalanceAccountUpdate) SetPwdSalt(s string) *StoreBalanceAccountUpdate {
	sbau.mutation.SetPwdSalt(s)
	return sbau
}

// SetNillablePwdSalt sets the "pwd_salt" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillablePwdSalt(s *string) *StoreBalanceAccountUpdate {
	if s != nil {
		sbau.SetPwdSalt(*s)
	}
	return sbau
}

// ClearPwdSalt clears the value of the "pwd_salt" field.
func (sbau *StoreBalanceAccountUpdate) ClearPwdSalt() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearPwdSalt()
	return sbau
}

// SetBalanceFee sets the "balance_fee" field.
func (sbau *StoreBalanceAccountUpdate) SetBalanceFee(f float64) *StoreBalanceAccountUpdate {
	sbau.mutation.ResetBalanceFee()
	sbau.mutation.SetBalanceFee(f)
	return sbau
}

// SetNillableBalanceFee sets the "balance_fee" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillableBalanceFee(f *float64) *StoreBalanceAccountUpdate {
	if f != nil {
		sbau.SetBalanceFee(*f)
	}
	return sbau
}

// AddBalanceFee adds f to the "balance_fee" field.
func (sbau *StoreBalanceAccountUpdate) AddBalanceFee(f float64) *StoreBalanceAccountUpdate {
	sbau.mutation.AddBalanceFee(f)
	return sbau
}

// ClearBalanceFee clears the value of the "balance_fee" field.
func (sbau *StoreBalanceAccountUpdate) ClearBalanceFee() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearBalanceFee()
	return sbau
}

// SetTotalChargeFee sets the "total_charge_fee" field.
func (sbau *StoreBalanceAccountUpdate) SetTotalChargeFee(f float64) *StoreBalanceAccountUpdate {
	sbau.mutation.ResetTotalChargeFee()
	sbau.mutation.SetTotalChargeFee(f)
	return sbau
}

// SetNillableTotalChargeFee sets the "total_charge_fee" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillableTotalChargeFee(f *float64) *StoreBalanceAccountUpdate {
	if f != nil {
		sbau.SetTotalChargeFee(*f)
	}
	return sbau
}

// AddTotalChargeFee adds f to the "total_charge_fee" field.
func (sbau *StoreBalanceAccountUpdate) AddTotalChargeFee(f float64) *StoreBalanceAccountUpdate {
	sbau.mutation.AddTotalChargeFee(f)
	return sbau
}

// ClearTotalChargeFee clears the value of the "total_charge_fee" field.
func (sbau *StoreBalanceAccountUpdate) ClearTotalChargeFee() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearTotalChargeFee()
	return sbau
}

// SetIsDeleted sets the "is_deleted" field.
func (sbau *StoreBalanceAccountUpdate) SetIsDeleted(i int8) *StoreBalanceAccountUpdate {
	sbau.mutation.ResetIsDeleted()
	sbau.mutation.SetIsDeleted(i)
	return sbau
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillableIsDeleted(i *int8) *StoreBalanceAccountUpdate {
	if i != nil {
		sbau.SetIsDeleted(*i)
	}
	return sbau
}

// AddIsDeleted adds i to the "is_deleted" field.
func (sbau *StoreBalanceAccountUpdate) AddIsDeleted(i int8) *StoreBalanceAccountUpdate {
	sbau.mutation.AddIsDeleted(i)
	return sbau
}

// ClearIsDeleted clears the value of the "is_deleted" field.
func (sbau *StoreBalanceAccountUpdate) ClearIsDeleted() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearIsDeleted()
	return sbau
}

// SetUpdatedAt sets the "updated_at" field.
func (sbau *StoreBalanceAccountUpdate) SetUpdatedAt(t time.Time) *StoreBalanceAccountUpdate {
	sbau.mutation.SetUpdatedAt(t)
	return sbau
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillableUpdatedAt(t *time.Time) *StoreBalanceAccountUpdate {
	if t != nil {
		sbau.SetUpdatedAt(*t)
	}
	return sbau
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (sbau *StoreBalanceAccountUpdate) ClearUpdatedAt() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearUpdatedAt()
	return sbau
}

// SetCreatedAt sets the "created_at" field.
func (sbau *StoreBalanceAccountUpdate) SetCreatedAt(t time.Time) *StoreBalanceAccountUpdate {
	sbau.mutation.SetCreatedAt(t)
	return sbau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sbau *StoreBalanceAccountUpdate) SetNillableCreatedAt(t *time.Time) *StoreBalanceAccountUpdate {
	if t != nil {
		sbau.SetCreatedAt(*t)
	}
	return sbau
}

// ClearCreatedAt clears the value of the "created_at" field.
func (sbau *StoreBalanceAccountUpdate) ClearCreatedAt() *StoreBalanceAccountUpdate {
	sbau.mutation.ClearCreatedAt()
	return sbau
}

// Mutation returns the StoreBalanceAccountMutation object of the builder.
func (sbau *StoreBalanceAccountUpdate) Mutation() *StoreBalanceAccountMutation {
	return sbau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sbau *StoreBalanceAccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sbau.hooks) == 0 {
		affected, err = sbau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StoreBalanceAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sbau.mutation = mutation
			affected, err = sbau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sbau.hooks) - 1; i >= 0; i-- {
			if sbau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sbau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sbau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sbau *StoreBalanceAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := sbau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sbau *StoreBalanceAccountUpdate) Exec(ctx context.Context) error {
	_, err := sbau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sbau *StoreBalanceAccountUpdate) ExecX(ctx context.Context) {
	if err := sbau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sbau *StoreBalanceAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   storebalanceaccount.Table,
			Columns: storebalanceaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: storebalanceaccount.FieldID,
			},
		},
	}
	if ps := sbau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sbau.mutation.AccountNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldAccountNo,
		})
	}
	if sbau.mutation.AccountNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldAccountNo,
		})
	}
	if value, ok := sbau.mutation.StoreCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldStoreCode,
		})
	}
	if sbau.mutation.StoreCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldStoreCode,
		})
	}
	if value, ok := sbau.mutation.UpperOrganNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldUpperOrganNo,
		})
	}
	if sbau.mutation.UpperOrganNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldUpperOrganNo,
		})
	}
	if value, ok := sbau.mutation.Pwd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldPwd,
		})
	}
	if sbau.mutation.PwdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldPwd,
		})
	}
	if value, ok := sbau.mutation.PwdSalt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldPwdSalt,
		})
	}
	if sbau.mutation.PwdSaltCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldPwdSalt,
		})
	}
	if value, ok := sbau.mutation.BalanceFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalanceaccount.FieldBalanceFee,
		})
	}
	if value, ok := sbau.mutation.AddedBalanceFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalanceaccount.FieldBalanceFee,
		})
	}
	if sbau.mutation.BalanceFeeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: storebalanceaccount.FieldBalanceFee,
		})
	}
	if value, ok := sbau.mutation.TotalChargeFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalanceaccount.FieldTotalChargeFee,
		})
	}
	if value, ok := sbau.mutation.AddedTotalChargeFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalanceaccount.FieldTotalChargeFee,
		})
	}
	if sbau.mutation.TotalChargeFeeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: storebalanceaccount.FieldTotalChargeFee,
		})
	}
	if value, ok := sbau.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalanceaccount.FieldIsDeleted,
		})
	}
	if value, ok := sbau.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalanceaccount.FieldIsDeleted,
		})
	}
	if sbau.mutation.IsDeletedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Column: storebalanceaccount.FieldIsDeleted,
		})
	}
	if value, ok := sbau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalanceaccount.FieldUpdatedAt,
		})
	}
	if sbau.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: storebalanceaccount.FieldUpdatedAt,
		})
	}
	if value, ok := sbau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalanceaccount.FieldCreatedAt,
		})
	}
	if sbau.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: storebalanceaccount.FieldCreatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sbau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storebalanceaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// StoreBalanceAccountUpdateOne is the builder for updating a single StoreBalanceAccount entity.
type StoreBalanceAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StoreBalanceAccountMutation
}

// SetAccountNo sets the "account_no" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetAccountNo(s string) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.SetAccountNo(s)
	return sbauo
}

// SetNillableAccountNo sets the "account_no" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillableAccountNo(s *string) *StoreBalanceAccountUpdateOne {
	if s != nil {
		sbauo.SetAccountNo(*s)
	}
	return sbauo
}

// ClearAccountNo clears the value of the "account_no" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearAccountNo() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearAccountNo()
	return sbauo
}

// SetStoreCode sets the "store_code" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetStoreCode(s string) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.SetStoreCode(s)
	return sbauo
}

// SetNillableStoreCode sets the "store_code" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillableStoreCode(s *string) *StoreBalanceAccountUpdateOne {
	if s != nil {
		sbauo.SetStoreCode(*s)
	}
	return sbauo
}

// ClearStoreCode clears the value of the "store_code" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearStoreCode() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearStoreCode()
	return sbauo
}

// SetUpperOrganNo sets the "upper_organ_no" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetUpperOrganNo(s string) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.SetUpperOrganNo(s)
	return sbauo
}

// SetNillableUpperOrganNo sets the "upper_organ_no" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillableUpperOrganNo(s *string) *StoreBalanceAccountUpdateOne {
	if s != nil {
		sbauo.SetUpperOrganNo(*s)
	}
	return sbauo
}

// ClearUpperOrganNo clears the value of the "upper_organ_no" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearUpperOrganNo() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearUpperOrganNo()
	return sbauo
}

// SetPwd sets the "pwd" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetPwd(s string) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.SetPwd(s)
	return sbauo
}

// SetNillablePwd sets the "pwd" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillablePwd(s *string) *StoreBalanceAccountUpdateOne {
	if s != nil {
		sbauo.SetPwd(*s)
	}
	return sbauo
}

// ClearPwd clears the value of the "pwd" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearPwd() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearPwd()
	return sbauo
}

// SetPwdSalt sets the "pwd_salt" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetPwdSalt(s string) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.SetPwdSalt(s)
	return sbauo
}

// SetNillablePwdSalt sets the "pwd_salt" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillablePwdSalt(s *string) *StoreBalanceAccountUpdateOne {
	if s != nil {
		sbauo.SetPwdSalt(*s)
	}
	return sbauo
}

// ClearPwdSalt clears the value of the "pwd_salt" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearPwdSalt() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearPwdSalt()
	return sbauo
}

// SetBalanceFee sets the "balance_fee" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetBalanceFee(f float64) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ResetBalanceFee()
	sbauo.mutation.SetBalanceFee(f)
	return sbauo
}

// SetNillableBalanceFee sets the "balance_fee" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillableBalanceFee(f *float64) *StoreBalanceAccountUpdateOne {
	if f != nil {
		sbauo.SetBalanceFee(*f)
	}
	return sbauo
}

// AddBalanceFee adds f to the "balance_fee" field.
func (sbauo *StoreBalanceAccountUpdateOne) AddBalanceFee(f float64) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.AddBalanceFee(f)
	return sbauo
}

// ClearBalanceFee clears the value of the "balance_fee" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearBalanceFee() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearBalanceFee()
	return sbauo
}

// SetTotalChargeFee sets the "total_charge_fee" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetTotalChargeFee(f float64) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ResetTotalChargeFee()
	sbauo.mutation.SetTotalChargeFee(f)
	return sbauo
}

// SetNillableTotalChargeFee sets the "total_charge_fee" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillableTotalChargeFee(f *float64) *StoreBalanceAccountUpdateOne {
	if f != nil {
		sbauo.SetTotalChargeFee(*f)
	}
	return sbauo
}

// AddTotalChargeFee adds f to the "total_charge_fee" field.
func (sbauo *StoreBalanceAccountUpdateOne) AddTotalChargeFee(f float64) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.AddTotalChargeFee(f)
	return sbauo
}

// ClearTotalChargeFee clears the value of the "total_charge_fee" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearTotalChargeFee() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearTotalChargeFee()
	return sbauo
}

// SetIsDeleted sets the "is_deleted" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetIsDeleted(i int8) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ResetIsDeleted()
	sbauo.mutation.SetIsDeleted(i)
	return sbauo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillableIsDeleted(i *int8) *StoreBalanceAccountUpdateOne {
	if i != nil {
		sbauo.SetIsDeleted(*i)
	}
	return sbauo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (sbauo *StoreBalanceAccountUpdateOne) AddIsDeleted(i int8) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.AddIsDeleted(i)
	return sbauo
}

// ClearIsDeleted clears the value of the "is_deleted" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearIsDeleted() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearIsDeleted()
	return sbauo
}

// SetUpdatedAt sets the "updated_at" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetUpdatedAt(t time.Time) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.SetUpdatedAt(t)
	return sbauo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillableUpdatedAt(t *time.Time) *StoreBalanceAccountUpdateOne {
	if t != nil {
		sbauo.SetUpdatedAt(*t)
	}
	return sbauo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearUpdatedAt() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearUpdatedAt()
	return sbauo
}

// SetCreatedAt sets the "created_at" field.
func (sbauo *StoreBalanceAccountUpdateOne) SetCreatedAt(t time.Time) *StoreBalanceAccountUpdateOne {
	sbauo.mutation.SetCreatedAt(t)
	return sbauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sbauo *StoreBalanceAccountUpdateOne) SetNillableCreatedAt(t *time.Time) *StoreBalanceAccountUpdateOne {
	if t != nil {
		sbauo.SetCreatedAt(*t)
	}
	return sbauo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (sbauo *StoreBalanceAccountUpdateOne) ClearCreatedAt() *StoreBalanceAccountUpdateOne {
	sbauo.mutation.ClearCreatedAt()
	return sbauo
}

// Mutation returns the StoreBalanceAccountMutation object of the builder.
func (sbauo *StoreBalanceAccountUpdateOne) Mutation() *StoreBalanceAccountMutation {
	return sbauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sbauo *StoreBalanceAccountUpdateOne) Select(field string, fields ...string) *StoreBalanceAccountUpdateOne {
	sbauo.fields = append([]string{field}, fields...)
	return sbauo
}

// Save executes the query and returns the updated StoreBalanceAccount entity.
func (sbauo *StoreBalanceAccountUpdateOne) Save(ctx context.Context) (*StoreBalanceAccount, error) {
	var (
		err  error
		node *StoreBalanceAccount
	)
	if len(sbauo.hooks) == 0 {
		node, err = sbauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StoreBalanceAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sbauo.mutation = mutation
			node, err = sbauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sbauo.hooks) - 1; i >= 0; i-- {
			if sbauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sbauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sbauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*StoreBalanceAccount)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StoreBalanceAccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sbauo *StoreBalanceAccountUpdateOne) SaveX(ctx context.Context) *StoreBalanceAccount {
	node, err := sbauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sbauo *StoreBalanceAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := sbauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sbauo *StoreBalanceAccountUpdateOne) ExecX(ctx context.Context) {
	if err := sbauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sbauo *StoreBalanceAccountUpdateOne) sqlSave(ctx context.Context) (_node *StoreBalanceAccount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   storebalanceaccount.Table,
			Columns: storebalanceaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: storebalanceaccount.FieldID,
			},
		},
	}
	id, ok := sbauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StoreBalanceAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sbauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storebalanceaccount.FieldID)
		for _, f := range fields {
			if !storebalanceaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != storebalanceaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sbauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sbauo.mutation.AccountNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldAccountNo,
		})
	}
	if sbauo.mutation.AccountNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldAccountNo,
		})
	}
	if value, ok := sbauo.mutation.StoreCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldStoreCode,
		})
	}
	if sbauo.mutation.StoreCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldStoreCode,
		})
	}
	if value, ok := sbauo.mutation.UpperOrganNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldUpperOrganNo,
		})
	}
	if sbauo.mutation.UpperOrganNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldUpperOrganNo,
		})
	}
	if value, ok := sbauo.mutation.Pwd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldPwd,
		})
	}
	if sbauo.mutation.PwdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldPwd,
		})
	}
	if value, ok := sbauo.mutation.PwdSalt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalanceaccount.FieldPwdSalt,
		})
	}
	if sbauo.mutation.PwdSaltCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: storebalanceaccount.FieldPwdSalt,
		})
	}
	if value, ok := sbauo.mutation.BalanceFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalanceaccount.FieldBalanceFee,
		})
	}
	if value, ok := sbauo.mutation.AddedBalanceFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalanceaccount.FieldBalanceFee,
		})
	}
	if sbauo.mutation.BalanceFeeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: storebalanceaccount.FieldBalanceFee,
		})
	}
	if value, ok := sbauo.mutation.TotalChargeFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalanceaccount.FieldTotalChargeFee,
		})
	}
	if value, ok := sbauo.mutation.AddedTotalChargeFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalanceaccount.FieldTotalChargeFee,
		})
	}
	if sbauo.mutation.TotalChargeFeeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: storebalanceaccount.FieldTotalChargeFee,
		})
	}
	if value, ok := sbauo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalanceaccount.FieldIsDeleted,
		})
	}
	if value, ok := sbauo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalanceaccount.FieldIsDeleted,
		})
	}
	if sbauo.mutation.IsDeletedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Column: storebalanceaccount.FieldIsDeleted,
		})
	}
	if value, ok := sbauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalanceaccount.FieldUpdatedAt,
		})
	}
	if sbauo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: storebalanceaccount.FieldUpdatedAt,
		})
	}
	if value, ok := sbauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalanceaccount.FieldCreatedAt,
		})
	}
	if sbauo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: storebalanceaccount.FieldCreatedAt,
		})
	}
	_node = &StoreBalanceAccount{config: sbauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sbauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storebalanceaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
