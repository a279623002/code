package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// OrderStatus 订单状态
type OrderStatus int32

const (
	OrderStatusPending   OrderStatus = 1 // 待支付
	OrderStatusPaid      OrderStatus = 2 // 已支付
	OrderStatusShipped   OrderStatus = 3 // 已发货
	OrderStatusCompleted OrderStatus = 4 // 已完成
	OrderStatusCanceled  OrderStatus = 5 // 已取消
)

// Order 订单模型
type Order struct {
	Id          int64       `gorm:"column:id;primaryKey;autoIncrement"`
	OrderNo     string      `gorm:"column:order_no;uniqueIndex;size:64;not null;comment:订单编号"`
	UserId      int64       `gorm:"column:user_id;index;not null;comment:用户ID"`
	TotalAmount float64     `gorm:"column:total_amount;type:decimal(10,2);not null;comment:订单总金额"`
	Status      OrderStatus `gorm:"column:status;default:1;comment:订单状态"`
	Remark      string      `gorm:"column:remark;size:512;comment:备注"`
	CreatedAt   time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;autoUpdateTime"`
}

// TableName 表名
func (Order) TableName() string {
	return "orders"
}

// OrderModel 订单数据访问层
type OrderModel struct {
	db *gorm.DB
}

// NewOrderModel 创建订单模型
func NewOrderModel(db *gorm.DB) *OrderModel {
	return &OrderModel{db: db}
}

// AutoMigrate 自动建表
func (m *OrderModel) AutoMigrate() error {
	return m.db.AutoMigrate(&Order{})
}

// Create 创建订单
func (m *OrderModel) Create(ctx context.Context, order *Order) error {
	return m.db.WithContext(ctx).Create(order).Error
}

// FindOne 根据 ID 查询订单
func (m *OrderModel) FindOne(ctx context.Context, id int64) (*Order, error) {
	var order Order
	if err := m.db.WithContext(ctx).First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}
	return &order, nil
}

// FindByOrderNo 根据订单号查询
func (m *OrderModel) FindByOrderNo(ctx context.Context, orderNo string) (*Order, error) {
	var order Order
	if err := m.db.WithContext(ctx).Where("order_no = ?", orderNo).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}
	return &order, nil
}

// List 分页查询订单
func (m *OrderModel) List(ctx context.Context, userId int64, page, pageSize int32) ([]*Order, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	db := m.db.WithContext(ctx).Model(&Order{})
	if userId > 0 {
		db = db.Where("user_id = ?", userId)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []*Order
	offset := (page - 1) * pageSize
	if err := db.Order("created_at DESC").Limit(int(pageSize)).Offset(int(offset)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// UpdateStatus 更新订单状态
func (m *OrderModel) UpdateStatus(ctx context.Context, id int64, status OrderStatus) error {
	result := m.db.WithContext(ctx).Model(&Order{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("order not found")
	}
	return nil
}
