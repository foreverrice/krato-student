// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"finance/internal/data/ent/predicate"
	"finance/internal/data/ent/storebalancechangedetail"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreBalanceChangeDetailUpdate is the builder for updating StoreBalanceChangeDetail entities.
type StoreBalanceChangeDetailUpdate struct {
	config
	hooks    []Hook
	mutation *StoreBalanceChangeDetailMutation
}

// Where appends a list predicates to the StoreBalanceChangeDetailUpdate builder.
func (sbcdu *StoreBalanceChangeDetailUpdate) Where(ps ...predicate.StoreBalanceChangeDetail) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.Where(ps...)
	return sbcdu
}

// SetAccountID sets the "account_id" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetAccountID(i int32) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.ResetAccountID()
	sbcdu.mutation.SetAccountID(i)
	return sbcdu
}

// AddAccountID adds i to the "account_id" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) AddAccountID(i int32) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.AddAccountID(i)
	return sbcdu
}

// SetStoreCode sets the "store_code" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetStoreCode(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetStoreCode(s)
	return sbcdu
}

// SetType sets the "type" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetType(i int8) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.ResetType()
	sbcdu.mutation.SetType(i)
	return sbcdu
}

// AddType adds i to the "type" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) AddType(i int8) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.AddType(i)
	return sbcdu
}

// SetInBatchNo sets the "in_batch_no" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetInBatchNo(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetInBatchNo(s)
	return sbcdu
}

// SetThirdPayType sets the "third_pay_type" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetThirdPayType(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetThirdPayType(s)
	return sbcdu
}

// SetThirdPayOrgan sets the "third_pay_organ" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetThirdPayOrgan(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetThirdPayOrgan(s)
	return sbcdu
}

// SetThirdPayNo sets the "third_pay_no" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetThirdPayNo(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetThirdPayNo(s)
	return sbcdu
}

// SetBcBindAccID sets the "bc_bind_acc_id" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetBcBindAccID(i int32) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.ResetBcBindAccID()
	sbcdu.mutation.SetBcBindAccID(i)
	return sbcdu
}

// AddBcBindAccID adds i to the "bc_bind_acc_id" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) AddBcBindAccID(i int32) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.AddBcBindAccID(i)
	return sbcdu
}

// SetCashAccountNo sets the "cash_account_no" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetCashAccountNo(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetCashAccountNo(s)
	return sbcdu
}

// SetCashPrincipal sets the "cash_principal" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetCashPrincipal(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetCashPrincipal(s)
	return sbcdu
}

// SetCashBank sets the "cash_bank" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetCashBank(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetCashBank(s)
	return sbcdu
}

// SetTransactionNo sets the "transaction_no" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetTransactionNo(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetTransactionNo(s)
	return sbcdu
}

// SetTransactionAt sets the "transaction_at" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetTransactionAt(t time.Time) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetTransactionAt(t)
	return sbcdu
}

// SetBeforeFee sets the "before_fee" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetBeforeFee(f float64) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.ResetBeforeFee()
	sbcdu.mutation.SetBeforeFee(f)
	return sbcdu
}

// AddBeforeFee adds f to the "before_fee" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) AddBeforeFee(f float64) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.AddBeforeFee(f)
	return sbcdu
}

// SetChangeFee sets the "change_fee" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetChangeFee(f float64) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.ResetChangeFee()
	sbcdu.mutation.SetChangeFee(f)
	return sbcdu
}

// AddChangeFee adds f to the "change_fee" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) AddChangeFee(f float64) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.AddChangeFee(f)
	return sbcdu
}

// SetAfterFee sets the "after_fee" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetAfterFee(f float64) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.ResetAfterFee()
	sbcdu.mutation.SetAfterFee(f)
	return sbcdu
}

// AddAfterFee adds f to the "after_fee" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) AddAfterFee(f float64) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.AddAfterFee(f)
	return sbcdu
}

// SetOperatorNo sets the "operator_no" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetOperatorNo(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetOperatorNo(s)
	return sbcdu
}

// SetFlowNo sets the "flow_no" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetFlowNo(s string) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetFlowNo(s)
	return sbcdu
}

// SetCheckState sets the "check_state" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetCheckState(i int8) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.ResetCheckState()
	sbcdu.mutation.SetCheckState(i)
	return sbcdu
}

// AddCheckState adds i to the "check_state" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) AddCheckState(i int8) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.AddCheckState(i)
	return sbcdu
}

// SetCheckAt sets the "check_at" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetCheckAt(t time.Time) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetCheckAt(t)
	return sbcdu
}

// SetIsDeleted sets the "is_deleted" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetIsDeleted(i int8) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.ResetIsDeleted()
	sbcdu.mutation.SetIsDeleted(i)
	return sbcdu
}

// AddIsDeleted adds i to the "is_deleted" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) AddIsDeleted(i int8) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.AddIsDeleted(i)
	return sbcdu
}

// SetUpdatedAt sets the "updated_at" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetUpdatedAt(t time.Time) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetUpdatedAt(t)
	return sbcdu
}

// SetCreatedAt sets the "created_at" field.
func (sbcdu *StoreBalanceChangeDetailUpdate) SetCreatedAt(t time.Time) *StoreBalanceChangeDetailUpdate {
	sbcdu.mutation.SetCreatedAt(t)
	return sbcdu
}

// Mutation returns the StoreBalanceChangeDetailMutation object of the builder.
func (sbcdu *StoreBalanceChangeDetailUpdate) Mutation() *StoreBalanceChangeDetailMutation {
	return sbcdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sbcdu *StoreBalanceChangeDetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sbcdu.hooks) == 0 {
		affected, err = sbcdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StoreBalanceChangeDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sbcdu.mutation = mutation
			affected, err = sbcdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sbcdu.hooks) - 1; i >= 0; i-- {
			if sbcdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sbcdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sbcdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sbcdu *StoreBalanceChangeDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := sbcdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sbcdu *StoreBalanceChangeDetailUpdate) Exec(ctx context.Context) error {
	_, err := sbcdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sbcdu *StoreBalanceChangeDetailUpdate) ExecX(ctx context.Context) {
	if err := sbcdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sbcdu *StoreBalanceChangeDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   storebalancechangedetail.Table,
			Columns: storebalancechangedetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: storebalancechangedetail.FieldID,
			},
		},
	}
	if ps := sbcdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sbcdu.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: storebalancechangedetail.FieldAccountID,
		})
	}
	if value, ok := sbcdu.mutation.AddedAccountID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: storebalancechangedetail.FieldAccountID,
		})
	}
	if value, ok := sbcdu.mutation.StoreCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldStoreCode,
		})
	}
	if value, ok := sbcdu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldType,
		})
	}
	if value, ok := sbcdu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldType,
		})
	}
	if value, ok := sbcdu.mutation.InBatchNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldInBatchNo,
		})
	}
	if value, ok := sbcdu.mutation.ThirdPayType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldThirdPayType,
		})
	}
	if value, ok := sbcdu.mutation.ThirdPayOrgan(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldThirdPayOrgan,
		})
	}
	if value, ok := sbcdu.mutation.ThirdPayNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldThirdPayNo,
		})
	}
	if value, ok := sbcdu.mutation.BcBindAccID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: storebalancechangedetail.FieldBcBindAccID,
		})
	}
	if value, ok := sbcdu.mutation.AddedBcBindAccID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: storebalancechangedetail.FieldBcBindAccID,
		})
	}
	if value, ok := sbcdu.mutation.CashAccountNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldCashAccountNo,
		})
	}
	if value, ok := sbcdu.mutation.CashPrincipal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldCashPrincipal,
		})
	}
	if value, ok := sbcdu.mutation.CashBank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldCashBank,
		})
	}
	if value, ok := sbcdu.mutation.TransactionNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldTransactionNo,
		})
	}
	if value, ok := sbcdu.mutation.TransactionAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalancechangedetail.FieldTransactionAt,
		})
	}
	if value, ok := sbcdu.mutation.BeforeFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldBeforeFee,
		})
	}
	if value, ok := sbcdu.mutation.AddedBeforeFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldBeforeFee,
		})
	}
	if value, ok := sbcdu.mutation.ChangeFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldChangeFee,
		})
	}
	if value, ok := sbcdu.mutation.AddedChangeFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldChangeFee,
		})
	}
	if value, ok := sbcdu.mutation.AfterFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldAfterFee,
		})
	}
	if value, ok := sbcdu.mutation.AddedAfterFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldAfterFee,
		})
	}
	if value, ok := sbcdu.mutation.OperatorNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldOperatorNo,
		})
	}
	if value, ok := sbcdu.mutation.FlowNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldFlowNo,
		})
	}
	if value, ok := sbcdu.mutation.CheckState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldCheckState,
		})
	}
	if value, ok := sbcdu.mutation.AddedCheckState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldCheckState,
		})
	}
	if value, ok := sbcdu.mutation.CheckAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalancechangedetail.FieldCheckAt,
		})
	}
	if value, ok := sbcdu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldIsDeleted,
		})
	}
	if value, ok := sbcdu.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldIsDeleted,
		})
	}
	if value, ok := sbcdu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalancechangedetail.FieldUpdatedAt,
		})
	}
	if value, ok := sbcdu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalancechangedetail.FieldCreatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sbcdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storebalancechangedetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// StoreBalanceChangeDetailUpdateOne is the builder for updating a single StoreBalanceChangeDetail entity.
type StoreBalanceChangeDetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StoreBalanceChangeDetailMutation
}

// SetAccountID sets the "account_id" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetAccountID(i int32) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.ResetAccountID()
	sbcduo.mutation.SetAccountID(i)
	return sbcduo
}

// AddAccountID adds i to the "account_id" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) AddAccountID(i int32) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.AddAccountID(i)
	return sbcduo
}

// SetStoreCode sets the "store_code" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetStoreCode(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetStoreCode(s)
	return sbcduo
}

// SetType sets the "type" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetType(i int8) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.ResetType()
	sbcduo.mutation.SetType(i)
	return sbcduo
}

// AddType adds i to the "type" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) AddType(i int8) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.AddType(i)
	return sbcduo
}

// SetInBatchNo sets the "in_batch_no" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetInBatchNo(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetInBatchNo(s)
	return sbcduo
}

// SetThirdPayType sets the "third_pay_type" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetThirdPayType(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetThirdPayType(s)
	return sbcduo
}

// SetThirdPayOrgan sets the "third_pay_organ" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetThirdPayOrgan(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetThirdPayOrgan(s)
	return sbcduo
}

// SetThirdPayNo sets the "third_pay_no" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetThirdPayNo(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetThirdPayNo(s)
	return sbcduo
}

// SetBcBindAccID sets the "bc_bind_acc_id" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetBcBindAccID(i int32) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.ResetBcBindAccID()
	sbcduo.mutation.SetBcBindAccID(i)
	return sbcduo
}

// AddBcBindAccID adds i to the "bc_bind_acc_id" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) AddBcBindAccID(i int32) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.AddBcBindAccID(i)
	return sbcduo
}

// SetCashAccountNo sets the "cash_account_no" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetCashAccountNo(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetCashAccountNo(s)
	return sbcduo
}

// SetCashPrincipal sets the "cash_principal" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetCashPrincipal(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetCashPrincipal(s)
	return sbcduo
}

// SetCashBank sets the "cash_bank" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetCashBank(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetCashBank(s)
	return sbcduo
}

// SetTransactionNo sets the "transaction_no" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetTransactionNo(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetTransactionNo(s)
	return sbcduo
}

// SetTransactionAt sets the "transaction_at" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetTransactionAt(t time.Time) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetTransactionAt(t)
	return sbcduo
}

// SetBeforeFee sets the "before_fee" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetBeforeFee(f float64) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.ResetBeforeFee()
	sbcduo.mutation.SetBeforeFee(f)
	return sbcduo
}

// AddBeforeFee adds f to the "before_fee" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) AddBeforeFee(f float64) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.AddBeforeFee(f)
	return sbcduo
}

// SetChangeFee sets the "change_fee" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetChangeFee(f float64) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.ResetChangeFee()
	sbcduo.mutation.SetChangeFee(f)
	return sbcduo
}

// AddChangeFee adds f to the "change_fee" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) AddChangeFee(f float64) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.AddChangeFee(f)
	return sbcduo
}

// SetAfterFee sets the "after_fee" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetAfterFee(f float64) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.ResetAfterFee()
	sbcduo.mutation.SetAfterFee(f)
	return sbcduo
}

// AddAfterFee adds f to the "after_fee" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) AddAfterFee(f float64) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.AddAfterFee(f)
	return sbcduo
}

// SetOperatorNo sets the "operator_no" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetOperatorNo(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetOperatorNo(s)
	return sbcduo
}

// SetFlowNo sets the "flow_no" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetFlowNo(s string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetFlowNo(s)
	return sbcduo
}

// SetCheckState sets the "check_state" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetCheckState(i int8) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.ResetCheckState()
	sbcduo.mutation.SetCheckState(i)
	return sbcduo
}

// AddCheckState adds i to the "check_state" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) AddCheckState(i int8) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.AddCheckState(i)
	return sbcduo
}

// SetCheckAt sets the "check_at" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetCheckAt(t time.Time) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetCheckAt(t)
	return sbcduo
}

// SetIsDeleted sets the "is_deleted" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetIsDeleted(i int8) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.ResetIsDeleted()
	sbcduo.mutation.SetIsDeleted(i)
	return sbcduo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) AddIsDeleted(i int8) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.AddIsDeleted(i)
	return sbcduo
}

// SetUpdatedAt sets the "updated_at" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetUpdatedAt(t time.Time) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetUpdatedAt(t)
	return sbcduo
}

// SetCreatedAt sets the "created_at" field.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SetCreatedAt(t time.Time) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.mutation.SetCreatedAt(t)
	return sbcduo
}

// Mutation returns the StoreBalanceChangeDetailMutation object of the builder.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) Mutation() *StoreBalanceChangeDetailMutation {
	return sbcduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) Select(field string, fields ...string) *StoreBalanceChangeDetailUpdateOne {
	sbcduo.fields = append([]string{field}, fields...)
	return sbcduo
}

// Save executes the query and returns the updated StoreBalanceChangeDetail entity.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) Save(ctx context.Context) (*StoreBalanceChangeDetail, error) {
	var (
		err  error
		node *StoreBalanceChangeDetail
	)
	if len(sbcduo.hooks) == 0 {
		node, err = sbcduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StoreBalanceChangeDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sbcduo.mutation = mutation
			node, err = sbcduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sbcduo.hooks) - 1; i >= 0; i-- {
			if sbcduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sbcduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sbcduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*StoreBalanceChangeDetail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StoreBalanceChangeDetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) SaveX(ctx context.Context) *StoreBalanceChangeDetail {
	node, err := sbcduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := sbcduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sbcduo *StoreBalanceChangeDetailUpdateOne) ExecX(ctx context.Context) {
	if err := sbcduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sbcduo *StoreBalanceChangeDetailUpdateOne) sqlSave(ctx context.Context) (_node *StoreBalanceChangeDetail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   storebalancechangedetail.Table,
			Columns: storebalancechangedetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: storebalancechangedetail.FieldID,
			},
		},
	}
	id, ok := sbcduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StoreBalanceChangeDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sbcduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storebalancechangedetail.FieldID)
		for _, f := range fields {
			if !storebalancechangedetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != storebalancechangedetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sbcduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sbcduo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: storebalancechangedetail.FieldAccountID,
		})
	}
	if value, ok := sbcduo.mutation.AddedAccountID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: storebalancechangedetail.FieldAccountID,
		})
	}
	if value, ok := sbcduo.mutation.StoreCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldStoreCode,
		})
	}
	if value, ok := sbcduo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldType,
		})
	}
	if value, ok := sbcduo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldType,
		})
	}
	if value, ok := sbcduo.mutation.InBatchNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldInBatchNo,
		})
	}
	if value, ok := sbcduo.mutation.ThirdPayType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldThirdPayType,
		})
	}
	if value, ok := sbcduo.mutation.ThirdPayOrgan(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldThirdPayOrgan,
		})
	}
	if value, ok := sbcduo.mutation.ThirdPayNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldThirdPayNo,
		})
	}
	if value, ok := sbcduo.mutation.BcBindAccID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: storebalancechangedetail.FieldBcBindAccID,
		})
	}
	if value, ok := sbcduo.mutation.AddedBcBindAccID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: storebalancechangedetail.FieldBcBindAccID,
		})
	}
	if value, ok := sbcduo.mutation.CashAccountNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldCashAccountNo,
		})
	}
	if value, ok := sbcduo.mutation.CashPrincipal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldCashPrincipal,
		})
	}
	if value, ok := sbcduo.mutation.CashBank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldCashBank,
		})
	}
	if value, ok := sbcduo.mutation.TransactionNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldTransactionNo,
		})
	}
	if value, ok := sbcduo.mutation.TransactionAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalancechangedetail.FieldTransactionAt,
		})
	}
	if value, ok := sbcduo.mutation.BeforeFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldBeforeFee,
		})
	}
	if value, ok := sbcduo.mutation.AddedBeforeFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldBeforeFee,
		})
	}
	if value, ok := sbcduo.mutation.ChangeFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldChangeFee,
		})
	}
	if value, ok := sbcduo.mutation.AddedChangeFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldChangeFee,
		})
	}
	if value, ok := sbcduo.mutation.AfterFee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldAfterFee,
		})
	}
	if value, ok := sbcduo.mutation.AddedAfterFee(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: storebalancechangedetail.FieldAfterFee,
		})
	}
	if value, ok := sbcduo.mutation.OperatorNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldOperatorNo,
		})
	}
	if value, ok := sbcduo.mutation.FlowNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storebalancechangedetail.FieldFlowNo,
		})
	}
	if value, ok := sbcduo.mutation.CheckState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldCheckState,
		})
	}
	if value, ok := sbcduo.mutation.AddedCheckState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldCheckState,
		})
	}
	if value, ok := sbcduo.mutation.CheckAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalancechangedetail.FieldCheckAt,
		})
	}
	if value, ok := sbcduo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldIsDeleted,
		})
	}
	if value, ok := sbcduo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: storebalancechangedetail.FieldIsDeleted,
		})
	}
	if value, ok := sbcduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalancechangedetail.FieldUpdatedAt,
		})
	}
	if value, ok := sbcduo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storebalancechangedetail.FieldCreatedAt,
		})
	}
	_node = &StoreBalanceChangeDetail{config: sbcduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sbcduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storebalancechangedetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
