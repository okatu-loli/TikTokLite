package logic

import (
	"context"
	"douyinFavoriteList_4/api/utool"
	"fmt"

	"douyinFavoriteList_4/api/internal/svc"
	"douyinFavoriteList_4/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteListLogic {
	return &GetFavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func action(req *types.FavoriteListReq) (resp *types.FavoriteListRes, err error) {
	// Go连接Mysql示例
	fmt.Println("\n\n>>>" + req.UserID)
	var f FavoriteList
	fmt.Println("\n\n>>>GetFavoriteListByUid")
	f.GetFavoriteListByUid(utool.S2I64(req.UserID))
	//f.PrintModel("FavoriteList")
	fmt.Printf("1>>>%#v", f.FollowData)
	f.PutMs2Api()
	fmt.Printf("2>>>OK")
	return &f.FavouriteListRes, nil
}
func (l *GetFavoriteListLogic) GetFavoriteList(req *types.FavoriteListReq) (resp *types.FavoriteListRes, err error) {
	// todo: add your logic here and delete this line

	return action(req)
}
