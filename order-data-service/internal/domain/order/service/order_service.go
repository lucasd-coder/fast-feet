package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/lucasd-coder/fast-feet/order-data-service/internal/domain/order"
	pkgErrors "github.com/lucasd-coder/fast-feet/order-data-service/internal/errors"
	"github.com/lucasd-coder/fast-feet/order-data-service/internal/provider/validator"
	"github.com/lucasd-coder/fast-feet/order-data-service/internal/shared"
	"github.com/lucasd-coder/fast-feet/order-data-service/pkg/pb"
	"github.com/lucasd-coder/fast-feet/pkg/logger"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	validate        shared.Validator
	orderRepository order.OrderRepository
}

func NewOrderService(
	validate *validator.Validation,
	orderRepo order.OrderRepository) *OrderService {
	return &OrderService{validate: validate, orderRepository: orderRepo}
}

func (s *OrderService) Save(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	log := logger.FromContext(ctx)

	slog.With("payload", req).Info("received request")

	pld := order.CreateOrder{
		DeliverymanID: req.GetDeliverymanId(),
		Product:       order.NewProduct(req.GetProduct().GetName()),
		Address:       s.newAddress(req),
	}

	if err := pld.Validate(s.validate); err != nil {
		return nil, pkgErrors.ValidationErrors(err)
	}

	order := order.NewOrder(pld)

	newOrder, err := s.orderRepository.Save(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("error when save: %w", err)
	}

	log.Infof("successfully created order with id: %s", newOrder.ID.Hex())

	return &pb.OrderResponse{
		Id:        newOrder.ID.Hex(),
		CreatedAt: order.GetCreatedAt(),
	}, nil
}

func (s *OrderService) newAddress(req *pb.OrderRequest) order.Address {
	return order.Address{
		Address:      req.GetAddresses().GetAddress(),
		Number:       req.GetAddresses().GetNumber(),
		PostalCode:   req.GetAddresses().GetPostalCode(),
		Neighborhood: req.GetAddresses().GetNeighborhood(),
		City:         req.GetAddresses().GetCity(),
		State:        req.GetAddresses().GetState(),
	}
}

func (s *OrderService) GetAllOrder(ctx context.Context, req *pb.GetAllOrderRequest) (*pb.GetAllOrderResponse, error) {
	log := logger.FromContext(ctx)

	slog.With("payload", req).Info("received request")

	pld := &order.GetAllOrderRequest{
		ID:            req.GetId(),
		DeliverymanID: req.GetDeliverymanId(),
		StartDate:     req.GetStartDate(),
		EndDate:       req.GetEndDate(),
		CreatedAt:     req.GetCreatedAt(),
		UpdatedAt:     req.GetUpdatedAt(),
		CanceledAt:    req.GetCanceledAt(),
		Limit:         req.GetLimit(),
		Offset:        req.GetOffset(),
		Product:       order.GetProduct{Name: req.GetProduct().GetName()},
		Address:       s.newGetAddress(req),
	}

	if err := pld.Validate(s.validate); err != nil {
		return nil, pkgErrors.ValidationErrors(err)
	}

	orders, err := s.orderRepository.FindAll(ctx, pld)
	if err != nil {
		return nil, fmt.Errorf("error when orderRepository findAll: %w", err)
	}

	log.Info("successfully return getAllOrder")

	return s.extractGetAllOrderResponse(pld, orders), nil
}

func (s *OrderService) newGetAddress(req *pb.GetAllOrderRequest) order.GetAddress {
	return order.GetAddress{
		Address:      req.GetAddresses().GetAddress(),
		Number:       req.GetAddresses().GetNumber(),
		PostalCode:   req.GetAddresses().GetPostalCode(),
		Neighborhood: req.GetAddresses().GetNeighborhood(),
		City:         req.GetAddresses().GetCity(),
		State:        req.GetAddresses().GetState(),
	}
}

func (s *OrderService) extractGetAllOrderResponse(pld *order.GetAllOrderRequest, orders []order.Order) *pb.GetAllOrderResponse {
	if len(orders) == 0 {
		return &pb.GetAllOrderResponse{}
	}

	pbOrders := []*pb.Order{}

	for _, order := range orders {
		pbOrders = append(pbOrders, s.extractPbOrder(order))
	}

	return &pb.GetAllOrderResponse{
		Total:  int32(len(pbOrders)),
		Offset: int32(pld.Offset),
		Limit:  int32(pld.Limit),
		Orders: pbOrders,
	}
}

func (s *OrderService) extractPbOrder(order order.Order) *pb.Order {
	return &pb.Order{
		Id:            order.ID.Hex(),
		DeliverymanId: order.DeliverymanID,
		StartDate:     order.GetStartDate(),
		EndDate:       order.GetStartDate(),
		Product:       &pb.Product{Name: order.Product.Name},
		Addresses: &pb.Address{
			Address:      order.Address.Address,
			PostalCode:   order.Address.PostalCode,
			Neighborhood: order.Address.Neighborhood,
			City:         order.Address.City,
			State:        order.Address.State,
			Number:       order.Address.Number,
		},
		CreatedAt:  order.GetCreatedAt(),
		UpdatedAt:  order.GetUpdatedAt(),
		CanceledAt: order.GetCanceledAt(),
	}
}
