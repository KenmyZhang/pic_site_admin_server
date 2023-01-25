package order

const (
	PayStatusUnpay      = 0 // 未支付
	PayStatusPaySuccess = 1 // 已支付
	PayStatusPayFail    = 2 // 支付失败
)

// 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)'
