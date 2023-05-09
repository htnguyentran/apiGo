package models

type UpdateOrderStatusReq struct {
	OrderNumber string `json:"OrderNumber"`
	// TrackingNumber string  `json:"TrackingNumber" binding:"required"`
	// UpdateTime     int64   `json:"UpdateTime" binding:"required"`
	// ActualWeight   *int64  `json:"ActualWeight"`
	// Status         string  `json:"Status" binding:"required"`
	// Note           string  `json:"Note"`
	// Name           string  `json:"Name"`
	// CityId         int     `json:"CityId"`
	// CityName       string  `json:"CityName"`
	// DistrictId     int     `json:"DistrictId"`
	// DistrictName   string  `json:"DistrictName"`
	// WardId         int     `json:"WardId"`
	// WardName       string  `json:"WardName"`
	// Address        string  `json:"Address"`
	// Lat            float64 `json:"Lat"`
	// Lng            float64 `json:"Lng"`
	// IsPrinted      bool    `json:"IsPrinted"`
}
