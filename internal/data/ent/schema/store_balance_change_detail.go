// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type StoreBalanceChangeDetail struct {
	ent.Schema
}

func (StoreBalanceChangeDetail) Fields() []ent.Field {
	return []ent.Field{field.Uint32("id").Comment("主键自增"), field.Int32("account_id").Comment("store_balance_accounts表主键"), field.String("store_code").Comment("门店编号"), field.Int8("type").Comment("1订货支付-|2提现-|3退款+|4线上充值+|5线下充值+|6虚拟充值+"), field.String("in_batch_no").Comment("内部流水号(订单编号)"), field.String("third_pay_type").Comment("第三方支付类型(付款银行)"), field.String("third_pay_organ").Comment("付款机构"), field.String("third_pay_no").Comment("第三方付款账号"), field.Int32("bc_bind_acc_id").Comment("bank_card_bind_accounts主键"), field.String("cash_account_no").Comment("收款账号"), field.String("cash_principal").Comment("收款主体"), field.String("cash_bank").Comment("收款银行名称"), field.String("transaction_no").Comment("交易流水号"), field.Time("transaction_at").Comment("交易时间"), field.Float("before_fee").Comment("变动前金额"), field.Float("change_fee").Comment("变动金额"), field.Float("after_fee").Comment("变动后金额"), field.String("operator_no").Comment("操作人编号"), field.String("flow_no").Comment("流程单号"), field.Int8("check_state").Comment("-1复合驳回 0复核 1复核通过"), field.Time("check_at").Comment("检查时间"), field.Int8("is_deleted").Comment("0未删除|1已删除"), field.Time("updated_at").Comment("更新时间"), field.Time("created_at").Comment("创建时间")}
}
func (StoreBalanceChangeDetail) Edges() []ent.Edge {
	return nil
}
func (StoreBalanceChangeDetail) Annotations() []schema.Annotation {
	return nil
}
