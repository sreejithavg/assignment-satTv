package pkg

type User struct {
	UserName     string
	EmailID      string
	PhoneNumber  int
	Balance      float32
	Subscription Subscription
}

type Subscription struct {
	Package Package
	Channel []Channel
	Service []Service
	pkgMonths	int
	channelMonths	int
	serviceMonths	int
}

type Package struct {
	PackageName string
	Price       float32
	Channel     []Channel
	Month		int
}

type Channel struct {
	ChannelName string
	Price       float32

}
type Service struct {
	ServiceName string
	Price       float32

}
var Zee= Channel{
	ChannelName:"Zee",
	Price:10,
}
var Sony =Channel{
	ChannelName: "Sony",
	Price:       15,
}
var StarPlus =Channel{
	ChannelName: "Star Plus",
	Price:       20,
}
var Discovery =Channel{
	ChannelName: "LearnCooking Service",
	Price:       10,
}
var NatGeo =Channel{
	ChannelName: "NatGeo",
	Price:       20,
}
var SilverPack = Package{
	PackageName: "Silver Pack",
	Price:       50.0,
	Channel:     []Channel{Zee,Sony,StarPlus},
}

var GoldPackage =Package{
	PackageName: "Gold Pack",
	Price:       100.0,
	Channel:      []Channel{Zee,Sony,StarPlus,Discovery,NatGeo},
}

var LearnEnglish =Service{
	ServiceName: "LearnEnglish Service",
	Price:       200.0,
}
var LearnCooking =Service{
	ServiceName: "LearnCooking Service",
	Price:       100.0,
}

var Pack = []Package{SilverPack,GoldPackage}
var Channels = []Channel{Zee,Sony,StarPlus,Discovery,NatGeo}
var Services=[]Service{LearnEnglish,LearnCooking}