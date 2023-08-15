package logic

import (
	"douyinFavoriteList_4/api/internal/logic/model"
	"douyinFavoriteList_4/api/internal/types"
	"strconv"
)

type FavoriteList struct {
	Uid           int64
	f             model.Favourite
	FavouriteData []model.Favourite
	v             model.Video
	VideosData    []model.Video
	u             model.User
	UsersData     []model.User
	fo            model.Follow
	FollowData    []model.Follow

	Size             int64
	FavouriteListRes types.FavoriteListRes
}

func (fl *FavoriteList) GetFavoriteListByUid(uid int64) {
	fl.Size = 0
	fl.FavouriteData = fl.f.GetFavouritesByUid(uid)
	//fmt.Printf("%+v", favouriteData)

	for _, f := range fl.FavouriteData {
		fl.VideosData = append(fl.VideosData, fl.v.GetVideosByAid(f.Id)...)
		fl.Size++
	}
	//fmt.Printf("%+v", videosData)

	for _, f := range fl.VideosData {
		fl.UsersData = append(fl.UsersData, fl.u.GetUserById(f.Id))
	}
	//fmt.Printf("%+v", fl.usersData)

	for i, _ := range fl.VideosData {
		fl.FollowData = append(fl.FollowData, fl.fo.GetFollowsByUids(uid, fl.UsersData[i].Id))
	}

}

func (fl FavoriteList) PrintModel(text string) {
	println(text)
	println("vvvvvvvvvvvvvvvvvvvv\n")
	for i, f := range fl.FavouriteData {
		f.Print("	FavouriteData" + strconv.Itoa(i))
	}
	for i, f := range fl.VideosData {
		f.Print("	VideosData" + strconv.Itoa(i))
	}
	for i, f := range fl.UsersData {
		f.Print("	UsersData" + strconv.Itoa(i))
	}
	for i, f := range fl.FollowData {
		f.Print("	FollowData" + strconv.Itoa(i))
	}
	println("^^^^^^^^^^^^^^^^^^")
}

func (fl *FavoriteList) PutMs2Api() {
	var VideoList = make([]types.Video, fl.Size)
	var Author = make([]types.User, fl.Size)
	for i := 0; i < int(fl.Size); i++ {
		Author[i].Avatar = fl.UsersData[i].GetAvatar()                   //
		Author[i].BackgroundImage = fl.UsersData[i].GetBackgroundImage() //
		Author[i].FavoriteCount = fl.UsersData[i].GetUserFavoriteCount()
		Author[i].FollowCount = fl.UsersData[i].FollowCount
		Author[i].FollowerCount = fl.UsersData[i].FollowerCount
		Author[i].ID = fl.UsersData[i].Id
		Author[i].IsFollow = func() bool {
			if fl.FollowData[i].Status == 1 {
				return true
			}
			return false
		}()
		Author[i].Name = fl.UsersData[i].UserName
		Author[i].Signature = fl.UsersData[i].GetSignature() //
		Author[i].TotalFavorited = fl.UsersData[i].GetUserTotalFavorited()
		Author[i].WorkCount = fl.UsersData[i].GetWorkCount()

		VideoList[i].Author = Author[i]
		VideoList[i].CommentCount = fl.VideosData[i].CommentCount
		VideoList[i].CoverURL = fl.VideosData[i].CoverUrl
		VideoList[i].FavoriteCount = fl.VideosData[i].FavoriteCount
		VideoList[i].ID = fl.VideosData[i].Id
		VideoList[i].IsFavorite = func() bool {
			if fl.FavouriteData[i].Status == 1 {
				return true
			}
			return false
		}()
		VideoList[i].PlayURL = fl.VideosData[i].PlayUrl
		VideoList[i].Title = fl.VideosData[i].Title
	}
	fl.FavouriteListRes.VideoList = VideoList
	fl.FavouriteListRes.StatusCode = "0"
	s := "succ"
	fl.FavouriteListRes.StatusMsg = &s
}
