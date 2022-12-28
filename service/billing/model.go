package billing

import "github.com/volcengine/volc-sdk-golang/base"

// see: https://www.volcengine.com/docs/6269/130259

type Bill struct {
	BillPeriod             string
	PayerID                string
	PayerUserName          string
	PayerCustomerName      string
	OwnerID                string
	OwnerUserName          string
	OwnerCustomerName      string
	Product                string
	ProductZh              string
	BusinessMode           string
	BillingMode            string
	ExpenseBeginTime       string
	ExpenseEndTime         string
	TradeTime              string
	BillID                 string
	BillCategoryParent     string
	OriginalBillAmount     float64
	PreferentialBillAmount float64
	RoundBillAmount        float64
	DiscountBillAmount     float64
	CouponAmount           float64
	PayableAmount          float64
	PaidAmount             float64
	UnpaidAmount           float64
	Currency               string
	PayStatus              string
}

type BillList struct {
	List   []*Bill
	Total  int
	Limit  int
	Offset int
}

type BillListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *BillList `json:",omitempty"`
}

type BillDetail struct {
	BillPeriod             string
	ExpenseDate            string
	PayerID                string
	PayerUserName          string
	PayerCustomerName      string
	OwnerID                string
	OwnerUserName          string
	OwnerCustomerName      string
	BusinessMode           string
	Product                string
	ProductZh              string
	BillingMode            string
	ExpenseBeginTime       string
	ExpenseEndTime         string
	UseDuration            string
	UseDurationUnit        string
	TradeTime              string
	BillID                 string
	BillCategory           string
	InstanceNo             string
	InstanceName           string
	ConfigName             string
	Element                string
	Region                 string
	Zone                   string
	Factor                 string
	ExpandField            string
	Price                  string
	PriceUnit              string
	Count                  string
	Unit                   string
	DeductionCount         string
	OriginalBillAmount     float64
	PreferentialBillAmount float64
	DiscountBillAmount     float64
	CouponAmount           float64
	PayableAmount          float64
	PaidAmount             float64
	UnpaidAmount           float64
	Currency               string
}

type BillDetailList struct {
	List   []*BillDetail
	Total  int
	Limit  int
	Offset int
}

type BillDetailListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *BillDetailList `json:",omitempty"`
}

type BillOverviewByProd struct {
	BillPeriod             string
	PayerID                string
	PayerUserName          string
	PayerCustomerName      string
	OwnerID                string
	OwnerUserName          string
	OwnerCustomerName      string
	BusinessMode           string
	Product                string
	ProductZh              string
	BillingMode            string
	BillCategoryParent     string
	OriginalBillAmount     float64
	PreferentialBillAmount float64
	RoundBillAmount        float64
	DiscountBillAmount     float64
	CouponAmount           float64
	PayableAmount          float64
	PaidAmount             float64
	UnpaidAmount           float64
}

type BillOverviewByProdList struct {
	List   []*BillOverviewByProd
	Total  int
	Limit  int
	Offset int
}

type BillOverviewByProdListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *BillOverviewByProdList `json:",omitempty"`
}
