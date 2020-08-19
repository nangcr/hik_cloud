package hik_cloud

import (
	"github.com/Insua/hik_cloud/access_token"
	"github.com/Insua/hik_cloud/building"
	"github.com/Insua/hik_cloud/community"
	"github.com/Insua/hik_cloud/config"
	"github.com/Insua/hik_cloud/context"
	"github.com/Insua/hik_cloud/credential"
	"github.com/Insua/hik_cloud/device"
	"github.com/Insua/hik_cloud/http"
	"github.com/Insua/hik_cloud/permission"
	"github.com/Insua/hik_cloud/person"
)

type HikCloud struct {
	ctx *context.Context
}

func NewHikCloud(cfg *config.Config) *HikCloud {
	akHandle := credential.NewAccessToken(cfg.ClientId, cfg.ClientSecret, cfg.Redis)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: akHandle,
	}
	return &HikCloud{ctx: ctx}
}

func (hc *HikCloud) GetAccessToken() *access_token.AccessToken {
	return access_token.NewAccessToken(hc.ctx)
}

func (hc *HikCloud) GetCommunity() *community.Community {
	h := http.NewHttp(hc.ctx)
	return community.NewCommunity(h)
}

func (hc *HikCloud) GetBuilding() *building.Building {
	h := http.NewHttp(hc.ctx)
	return building.NewBuilding(h)
}

func (hc *HikCloud) GetPerson() *person.Person {
	h := http.NewHttp(hc.ctx)
	return person.NewPerson(h)
}

func (hc *HikCloud) GetDevice() *device.Device {
	h := http.NewHttp(hc.ctx)
	return device.NewDevice(h)
}

func (hc *HikCloud) GetPermission() *permission.Permission {
	h := http.NewHttp(hc.ctx)
	return permission.NewPermission(h)
}
