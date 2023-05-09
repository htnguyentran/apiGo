package statics

const (
	QueueName                     = "LOGISTIC_%s_JITT_UPDATE_ORDER_STATUS"
	QueueNameRetry                = "LOGISTIC_%s_JITT_UPDATE_ORDER_STATUS_RETRY"
	QueueDOPName                  = "LOGISTIC_FRTDOP_UPDATE_ORDER_STATUS"
	QueueDOPNameRetry             = "LOGISTIC_FRTDOP_UPDATE_ORDER_STATUS_RETRY"
	ShippmentStatusProcessing     = "processing"      // Đang xử lý
	ShippmentStatusPickingUp      = "picking_up"      // Đang lấy hàng
	ShippmentStatusSorting        = "sorting"         // Đang xếp hàng
	ShippmentStatusDelivering     = "delivering"      // Đang giao hàng
	ShippmentStatusPod            = "pod"             // Đã giao hàng
	ShippmentStatusShipmentReturn = "shipment_return" // Người bán đã nhận lại hàng
	ShippmentStatusRequestReturn  = "request_return"  // Trả hàng người bán
	ShippmentStatusChanged        = "changed"         // Đã đổi hàng
	ShippmentStatusReturning      = "returning"       // Đang trả hàng
	ShippmentStatusReturned       = "returned"        // Đổi trả thành công
	ShippmentStatusPickupSuccess  = "pick_up_success" // lấy hàng thành công
	ShippmentStatusHandedOver     = "handed_over"     //ban giao thanh cong lm
)
