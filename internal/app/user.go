package app

import (
	"context"

	"github.com/LyricTian/gin-admin/v8/internal/app/schema"
	"github.com/LyricTian/gin-admin/v8/internal/app/service"
)

func InitRole(r *service.RoleSrv, m *service.MenuSrv) (*schema.IDResult, error) {
	ctx := context.Background()
	var params schema.MenuQueryParam
	params.Pagination = true
	mResult, err := m.Query(ctx, params, schema.MenuQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC)),
	})
	if err != nil {
		return nil, err
	}

	mData := mResult.Data

	var params2 schema.RoleQueryParam
	params2.Name = "worker"
	rResult, err := r.Query(ctx, params2)
	if err != nil {
		return nil, err
	}
	if len(rResult.Data) == 0 {
		_, err = r.Create(ctx, schema.Role{
			Name:     "worker",
			Sequence: 6,
			Status:   1,
		})
		if err != nil {
			return nil, err
		}
	}

	params2.Name = "admin"
	rResult, err = r.Query(ctx, params2)
	if err != nil {
		return nil, err
	}
	if len(rResult.Data) == 0 {
		return r.Create(ctx, schema.Role{
			Name:     "admin",
			Sequence: 9,
			Status:   1,
			RoleMenus: schema.RoleMenus{
				&schema.RoleMenu{
					MenuID:   mData[0].ID,
					ActionID: mData[0].Actions[0].ID,
				},
				&schema.RoleMenu{
					MenuID:   mData[0].ID,
					ActionID: mData[0].Actions[1].ID,
				},
				&schema.RoleMenu{
					MenuID:   mData[0].ID,
					ActionID: mData[0].Actions[2].ID,
				},
				&schema.RoleMenu{
					MenuID:   mData[0].ID,
					ActionID: mData[0].Actions[3].ID,
				},
				&schema.RoleMenu{
					MenuID:   mData[0].ID,
					ActionID: mData[0].Actions[4].ID,
				},
				&schema.RoleMenu{
					MenuID:   mData[0].ID,
					ActionID: mData[0].Actions[5].ID,
				},
			},
		})
	} else {
		return &schema.IDResult{ID: rResult.Data[0].ID}, nil
	}
}

func InitUser(u *service.UserSrv, r *service.RoleSrv, m *service.MenuSrv) error {
	ctx := context.Background()
	admin, err := InitRole(r, m)
	if err != nil {
		return err
	}

	param := schema.UserQueryParam{
		UserName: "admin",
	}
	result, err := u.Query(ctx, param)
	if err != nil {
		return err
	}

	if len(result.Data) > 0 {
		return nil
	} else {
		_, err = u.Create(ctx, schema.User{
			UserName: "admin",
			RealName: "admin",
			Password: "admin",
			Status:   1,
			UserRoles: schema.UserRoles{
				&schema.UserRole{
					RoleID: admin.ID,
				},
			},
		})
	}

	return err
}
