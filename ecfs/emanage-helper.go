package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/elastifile/emanage-go/pkg/emanage"
	"github.com/elastifile/errors"
)

type emanageClient struct {
	*emanage.Client
}

var emsConfig *config

// Connect to eManage
func newEmanageClient() (client *emanage.Client, err error) {

	if emsConfig == nil {
		glog.V(2).Infof("AAAAA GetClient - initializing new eManage client") // TODO: DELME
		emsConfig, err = newPluginConfig()
		if err != nil {
			return
			//panic(fmt.Sprintf("Failed to create new ECFS plugin config - %v", err))
		}
	}

	// TODO: Add retries

	glog.V(2).Infof("AAAAA newEmanageClient - config: %+v", emsConfig) // TODO: DELME
	glog.Infof("newEmanageClient on %s", emsConfig.EmanageURL)
	baseURL, err := url.Parse(strings.TrimSuffix(emsConfig.EmanageURL, "\n"))
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	glog.V(2).Info("AAAAA newEmanageClient - calling NewClient()") // TODO: DELME
	client = emanage.NewClient(baseURL)
	glog.V(2).Infof("AAAAA newEmanageClient - got new client: %+v", client) // TODO: DELME
	err = client.Sessions.Login(emsConfig.Username, emsConfig.Password)
	glog.V(2).Infof("AAAAA newEmanageClient - login err: %v", err) // TODO: DELME
	if err != nil {
		err = errors.WrapPrefix(err, fmt.Sprintf("Failed to log into eManage on %v", emsConfig.EmanageURL), 0)
		err = status.Error(codes.Unauthenticated, err.Error())
		return
	}
	return
}

func (ems *emanageClient) GetClient() *emanage.Client {
	var err error

	if ems.Client == nil {
		glog.V(2).Infof("AAAAA CreateVolume - creating eManage client") // TODO: DELME
		ems.Client, err = newEmanageClient()
		if err != nil {
			panic(fmt.Sprintf("Failed to create eManage client. err: %v", err))
		}
		glog.V(2).Infof("AAAAA GetClient - initialized new eManage client - %+v", ems.Client) // TODO: DELME
	}

	return ems.Client
}

func (ems *emanageClient) GetDcByName(dcName string) (*emanage.DataContainer, error) {
	glog.V(2).Infof("AAAAA GetDcByName - getting DCs from ems: %+v", ems) // TODO: DELME
	dcs, err := ems.GetClient().DataContainers.GetAll(nil)
	if err != nil {
		return nil, errors.WrapPrefix(err, "Failed to list Data Containers", 0)
	}
	for _, dc := range dcs {
		if dc.Name == dcName {
			return &dc, nil
		}
	}
	return nil, errors.Errorf("Container '%v' not found", dcName)
}

func (ems *emanageClient) GetDcExportByName(dcExportName string) (*emanage.DataContainer, *emanage.Export, error) {
	// Here we assume the Dc and the Export have the same name
	glog.V(2).Infof("AAAAA GetDcExportByName - enter. Looking for Dc & export named: %v", dcExportName) // TODO: DELME
	dc, err := ems.GetDcByName(dcExportName)
	if err != nil {
		return nil, nil, errors.Wrap(err, 0)
	}

	exports, err := ems.GetClient().Exports.GetAll(nil)
	if err != nil {
		return nil, nil, errors.WrapPrefix(err, "Failed to get exports from eManage", 0)
	}

	for _, export := range exports {
		if dc.Id == export.DataContainerId && export.Name == dcExportName {
			glog.V(2).Infof("AAAAA GetDcExportByName - success. Returning DC: %+v EXPORT: %+v", dc, export) // TODO: DELME
			return dc, &export, nil
		}
	}
	return nil, nil, errors.Errorf("Export not found by DataContainer&Export name", dcExportName)
}

func GetProvisionerSettings() (configMap map[string]string, secret map[string]string, err error) {
	// TODO: Implement fetching cluster configuration and credentials from configmap+secret - these are fake values
	glog.Warning("Config map and secrets are not yet supported - using hard-coded values!!!")

	configMap = make(map[string]string)
	secret = make(map[string]string)

	configMap[managementAddress] = "https://35.241.144.0"
	configMap[managementUserName] = "admin"
	configMap[nfsAddress] = "172.28.0.5"
	secret[managementPassword] = "changeme"

	return
}