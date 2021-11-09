package couponallocated

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponallocated"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateCouponAllocated(info *npool.CouponAllocated) error {
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invlaid user id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invlaid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCouponID()); err != nil {
		return xerrors.Errorf("invlaid coupon id: %v", err)
	}
	return nil
}

func dbRowToCouponAllocated(row *ent.CouponAllocated) *npool.CouponAllocated {
	return &npool.CouponAllocated{
		ID:       row.ID.String(),
		UserID:   row.UserID.String(),
		AppID:    row.AppID.String(),
		Used:     row.Used,
		CouponID: row.CouponID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateCouponAllocatedRequest) (*npool.CreateCouponAllocatedResponse, error) {
	if err := validateCouponAllocated(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		CouponAllocated.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetCouponID(uuid.MustParse(in.GetInfo().GetCouponID())).
		SetUsed(false).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create coupon allocated: %v", err)
	}

	return &npool.CreateCouponAllocatedResponse{
		Info: dbRowToCouponAllocated(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateCouponAllocatedRequest) (*npool.UpdateCouponAllocatedResponse, error) {
	if err := validateCouponAllocated(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	info, err := db.Client().
		CouponAllocated.
		UpdateOneID(id).
		SetUsed(in.GetInfo().GetUsed()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update coupon allocated: %v", err)
	}

	return &npool.UpdateCouponAllocatedResponse{
		Info: dbRowToCouponAllocated(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetCouponAllocatedRequest) (*npool.GetCouponAllocatedResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		CouponAllocated.
		Query().
		Where(
			couponallocated.And(
				couponallocated.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coupon allocated: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty coupon allocated")
	}

	return &npool.GetCouponAllocatedResponse{
		Info: dbRowToCouponAllocated(infos[0]),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetCouponsAllocatedByAppRequest) (*npool.GetCouponsAllocatedByAppResponse, error) {
	return nil, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetCouponsAllocatedByAppUserRequest) (*npool.GetCouponsAllocatedByAppUserResponse, error) {
	return nil, nil
}