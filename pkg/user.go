package pkg

import (
	"errors"
	"fmt"
	"github.com/dariubs/percent"
	"log"
	"strings"
)

func (u *User)CreateUser()  {
	fmt.Println("Enter the UserName : ")
	_, err := fmt.Scanf("%s", &u.UserName)
	if err != nil {
		log.Fatal("Error occurred while fetching the username ",err)
	}
	fmt.Println("Enter the email_ID : ")
	_, err = fmt.Scanf("%s", &u.EmailID)
	if err != nil {
		log.Fatal("Error occurred while fetching the EmailID ")
	}
	fmt.Println("Enter the Phone Number : ")
	_, err = fmt.Scanf("%d", &u.PhoneNumber)
	if err != nil {
		log.Fatal("Error occurred while fetching the phone no: ")
	}
	u.Balance=100.0
}
func (u *User)Recharge(amount float32)  {
	u.Balance=u.Balance+amount
	fmt.Println(u.Balance)
	fmt.Println("Recharge completed successfully. Current balance is : ",u.Balance)
}
func (u *User)ViewBalance()  {
	fmt.Printf("Current balance is %f Rs.  \n",u.Balance)
}
func (u *User)PackSubscription(pack Package, months int) (float32,float32,float32) {
	var discountAmount float64
	var amount float32
		u.Subscription.Package=pack
		u.Subscription.Package.Month=months
		amount=pack.Price*float32(months)
		if months<3{
			return amount,0,amount
		}
		discountAmount=percent.PercentFloat(10,float64(amount))
		finalAmount:=amount-float32(discountAmount)
		return amount,float32(discountAmount),finalAmount
}
func (u *User) Withdrawal(amount float32) error{
	if amount>u.Balance{
		err:=errors.New("there is not enough balance in your account.Please recharge! ")
		log.Println(err)
		return err
	}
	u.Balance-=amount
	return nil
}
func (u *User)AddChannels(channel string,months int) (float32,error) {
	chann:=strings.Split(channel,",")
	var amount float32
	for _,v:=range chann{
		switch v {
		case Zee.ChannelName:
			u.Subscription.Channel=append(u.Subscription.Channel,Zee)
			amount=amount+(Zee.Price*float32(months))
		case Sony.ChannelName:
			u.Subscription.Channel=append(u.Subscription.Channel,Sony)
			amount=amount+(Zee.Price*float32(months))
		case StarPlus.ChannelName:
			u.Subscription.Channel=append(u.Subscription.Channel,StarPlus)
			amount=amount+(Zee.Price*float32(months))
		case Discovery.ChannelName :
			u.Subscription.Channel=append(u.Subscription.Channel,StarPlus)
			amount=amount+(Zee.Price*float32(months))
		case NatGeo.ChannelName:
			u.Subscription.Channel=append(u.Subscription.Channel,StarPlus)
			amount=amount+(Zee.Price*float32(months))
		default:
			return amount,errors.New("invalid Channel. cannot subscribe! ")
		}
	}
	return amount,nil
}
func (u *User)SubscriptionService(service Service,month int) float32 {
	var amount float32
	u.Subscription.Service=append(u.Subscription.Service,service)
	amount=amount+(service.Price*float32(month))
	return amount
}
