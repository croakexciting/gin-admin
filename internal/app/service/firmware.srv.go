package service

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/google/wire"

	"dishes-admin-mod/internal/app/config"
	"dishes-admin-mod/internal/app/dao"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
	"dishes-admin-mod/pkg/util/snowflake"
)

var FirmwareSet = wire.NewSet(wire.Struct(new(FirmwareSrv), "*"))

type FirmwareSrv struct {
	TransRepo    *dao.TransRepo
	FirmwareRepo *dao.FirmwareRepo
}

func (a *FirmwareSrv) Query(ctx context.Context, params schema.FirmwareQueryParam, opts ...schema.FirmwareQueryOptions) (*schema.FirmwareQueryResult, error) {
	result, err := a.FirmwareRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *FirmwareSrv) Get(ctx context.Context, id uint64, opts ...schema.FirmwareQueryOptions) (*schema.Firmware, error) {
	item, err := a.FirmwareRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *FirmwareSrv) Create(ctx context.Context, item schema.Firmware) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.FirmwareRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *FirmwareSrv) Update(ctx context.Context, id uint64, item schema.Firmware) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.FirmwareRepo.Update(ctx, id, item)
	})
}

func (a *FirmwareSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.FirmwareRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.FirmwareRepo.Delete(ctx, id)
	})
}

func (a *FirmwareSrv) UploadFile(ctx context.Context, file multipart.File, filename string, productID uint64, version string) error {
	d := config.C.FileServer.Directory + "/" + strconv.FormatUint(productID, 10) + "/" + version + "/"

	exist, err := a.pathExists(d)
	if err != nil {
		return err
	}
	if !exist {
		os.MkdirAll(d, os.ModePerm)
	}

	out, err := os.Create(d + filename)
	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}

func (a *FirmwareSrv) pathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
