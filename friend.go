package yunxin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Friend struct {
	APPKEY    string
	APPSECRET string
}

type FriendInfo struct {
	CreateTime  int64  `json:"createtime"`
	Bidirection bool   `json:"bidirection"`
	Faccid      string `json:"faccid"`
	Alias       string `json:"alias"`
}

type FriendInfoResponse struct {
	BaseResp
	Size    int          `json:"size"`
	Friends []FriendInfo `json:"friends"`
}

type BlackAndMuteList struct {
	BaseResp
	MuteList  []string `json:"mutelist"`
	BlackList []string `json:"blacklist"`
}

// AddFriend ...
func (this *Friend) AddFriend(accid string, faccid string, friendType int, msg string) (*BaseResp, error) {
	data, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_FRIEND_ADD, url.Values{"accid": {accid}, "faccid": {faccid}, "type": {strconv.Itoa(friendType)}, "msg": {msg}})
	if err != nil {
		return nil, err
	}
	resp := &BaseResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateFriend ...
func (this *Friend) UpdateFriend(accid, faccid, alias, ex string) (*BaseResp, error) {
	data, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_FRIEND_UPDATE, url.Values{"accid": {accid}, "faccid": {faccid}, "alias": {alias}, "ex": {ex}})
	if err != nil {
		return nil, err
	}
	resp := &BaseResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// DeleeteFriend ...
func (this *Friend) DeleteFriend(accid, faccid string) (*BaseResp, error) {
	data, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_FRIEND_DELETE, url.Values{"accid": {accid}, "faccid": {faccid}})
	if err != nil {
		return nil, err
	}
	resp := &BaseResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// GetFriendList ...
func (this *Friend) GetFriendList(accid string, updateTime int64) (*FriendInfoResponse, error) {
	data, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_FRIEND_GET, url.Values{"accid": {accid}, "updatetime": {strconv.FormatInt(updateTime, 10)}})
	if err != nil {
		return nil, err
	}
	friendsInfo := &FriendInfoResponse{}
	err = json.Unmarshal(data, friendsInfo)
	if err != nil {
		return nil, err
	}
	return friendsInfo, nil

}

// SetSpecialRelation ...
func (this *Friend) SetSpecialRelation(accid string, targetAcc string, relationType int, value int) (*BaseResp, error) {
	data, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_USER_SET_SPECIAL_RELATION, url.Values{"accid": {accid}, "targetAcc": {targetAcc}, "relationType": {strconv.Itoa(relationType)}, "value": {strconv.Itoa(value)}})
	if err != nil {
		return nil, err
	}
	baseResp := &BaseResp{}
	err = json.Unmarshal(data, baseResp)
	if err != nil {
		return nil, err
	}
	return baseResp, nil

}

// ListBlackAndMuteList ...
func (this *Friend) ListBlackAndMuteList(accid string) (*BlackAndMuteList, error) {
	data, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_USER_LIST_BLACK_AND_MUTE_LIST, url.Values{"accid": {accid}})
	if err != nil {
		return nil, err
	}
	blackAndMuteList := &BlackAndMuteList{}

	err = json.Unmarshal(data, blackAndMuteList)
	if err != nil {
		return nil, err
	}
	return blackAndMuteList, nil

}
