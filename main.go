/**
  author: kevin
 */

package main

import (
	"fmt"
	"github.com/test.com/test/sdkInit"
	serveric "github.com/test.com/test/service"
	"os"
)

func main()  {
	const ConfigPath="config.yaml"
	sdk,err:=sdkInit.SetupSDK(ConfigPath,false)
	if err!=nil{
		fmt.Printf(err.Error())
	}
	initInfo := &sdkInit.InitInfo{

		ChannelID: "mychannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/test.com/test/fixtures/channel-artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"orgmai",
		OrdererOrgName: "orderer.test.com",
		UserName:"User1",

		ChaincodeID:"mycc",
		ChaincodePath:"github.com/test.com/test/chaincode",
		ChaincodeGoPath:os.Getenv("GOPATH"),
	}
	err=sdkInit.CreateChannel(sdk,initInfo)
	if err!=nil{
		fmt.Println(err.Error())
	}
	channelClient,err:=sdkInit.InstallAndInstantiateCC(sdk,initInfo)
	if err!=nil{
		fmt.Println(err.Error())
	}
	setup:=serveric.ServiceSetup{Client:channelClient,ChaincodeID:initInfo.ChaincodeID}
	transactionID,err:=setup.SetInfo("test1","mynameislinyesheng")
	if err!=nil{
		fmt.Println("链码调用set失败：",err)
	}
	fmt.Println("交易提案成功！,交易ID:",transactionID)
	msg,err:=setup.GetInfo("test1")
	if err!=nil{
		fmt.Println("链码调用“get”失败:",err)
	}
	fmt.Println("查询成功内容：",msg)
}