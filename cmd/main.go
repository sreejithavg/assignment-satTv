package main

import (
	"SatTV/pkg"
	"fmt"
	"log"
)

func main() {
	// create a SatTV user
	var service,phoneNumber int
	var username,email string
	satUser := new(pkg.User)
	fmt.Println("Enter the UserName : ")
	_, err := fmt.Scanf("%s", &username)
	if err != nil {
		log.Fatal("Error occurred while fetching the username ",err)
	}
	fmt.Println("Enter the email_ID : ")
	_, err = fmt.Scanf("%s", &email)
	if err != nil {
		log.Fatal("Error occurred while fetching the EmailID ")

	}
	fmt.Println("Enter the Phone Number : ")
	_, err = fmt.Scanf("%d", &phoneNumber)
	if err != nil {
		log.Fatal("Error occurred while fetching the phone no: ")

	}
	err=satUser.CreateUser(username,email,phoneNumber)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Welcome to SatTV \n" +
		"1. View current balance in the account \n" +
		"2. Recharge Account \n" +
		"3. View available packs, channels and services\n" +
		"4. Subscribe to base packs\n" +
		"5. Add channels to an existing subscription\n" +
		"6. Subscribe to special services \n" +
		"7. View current subscription details\n" +
		"8. Update email and phone number for notifications\n" +
		"9. Exit ")

SWITCH:
	fmt.Println("Enter the Option : ")
	fmt.Scanf("%d", &service)
	switch service {
	case 1:
		satUser.ViewBalance()
		goto SWITCH
	case 2:
		var amount float32
		fmt.Println("Enter the amount to recharge : ")
		fmt.Scanf("%f", &amount)
		satUser.Recharge(amount)
		goto SWITCH
	case 3:
		ViewServices(pkg.Pack, pkg.Channels, pkg.Services)
		goto SWITCH
	case 4:
		var packChoice rune
		var packSub pkg.Package
		fmt.Println("Subscribe to channel packs \nEnter the Pack you wish to subscribe : (Silver: ‘S’, Gold: ‘G’) : ")
		fmt.Scanf("%c", &packChoice)
	MONTH:
		fmt.Println("Enter the months:")
		var m int
		fmt.Scan(&m)
		if m == 0 {
			fmt.Println("Please Enter the valid count")
			goto MONTH
		}
		if packChoice == 'S' {
			packSub = pkg.SilverPack
		} else {
			packSub = pkg.GoldPackage
		}
		amount, discount, finalPrice := satUser.PackSubscription(packSub, m)
		err := satUser.Withdrawal(finalPrice)
		if err != nil {
			fmt.Println("Subscription failed due to  ", err)
			goto SWITCH
		}
		fmt.Println("You have successfully subscribed the following packs - ", packSub.PackageName, "\nMonthly"+
			" price: ", packSub.Price, " Rs. No of months: ", m, "\nSubscription Amount: ", amount, " Rs. \nDiscount applied: ", discount, " Rs. "+
			"\nFinal Price after discount: ", finalPrice, " Rs. \nAccount balance: ", satUser.Balance, " Rs. "+
			"\nEmail notification sent successfully \nSMS notification sent successfully")
		goto SWITCH
	case 5:
	READ:
		fmt.Println("Add channels to existing subscription \nEnter channel names to add (separated by commas):")
		var channel string
		fmt.Scanf("%s", &channel)
		fmt.Println(channel)
	MONTHS:
		fmt.Println("Enter the months:")
		var m int
		fmt.Scan(&m)
		if m == 0 {
			fmt.Println("Please Enter the valid count")
			goto MONTHS
		}
		finalPrice, err := satUser.AddChannels(channel, m)
		if err != nil {
			log.Println(err)
			goto READ
		}
		err = satUser.Withdrawal(finalPrice)
		if err != nil {
			log.Println(err)
			goto SWITCH
		}
		fmt.Printf("Channels added successfully.\n Account balance: %f Rs.", satUser.Balance)
		goto SWITCH
	case 6:
		fmt.Println("Enter the service name:(LearnEnglish Service:'E' , LearnCooking Service:'C' )")
		var ser rune
		var servi pkg.Service
		fmt.Scanf("%c", &ser)
	MONTHSER:
		fmt.Println("Enter the months:")
		var m int
		fmt.Scanf("%d",&m)
		if m == 0 {
			fmt.Println("Please Enter the valid count")
			goto MONTHSER
		}
		if ser == 'E' {
			servi = pkg.LearnEnglish
		} else if ser == 'C' {
			servi = pkg.LearnCooking
		} else {
			log.Println("Invalid service Please check the service again")
			goto SWITCH
		}
		finalPrice := satUser.SubscriptionService(servi, m)
		err := satUser.Withdrawal(finalPrice)
		if err != nil {
			log.Println(err)
			goto SWITCH
		}
		fmt.Printf("\n Service subscribed successfully \nAccount balance: %f Rs. \nEmail notification sent successfully \nSMS notification sent successfull.", satUser.Balance)
		goto SWITCH
	case 7:
		fmt.Println("View current subscription details\nCurrently subscribed packs and channels: ", satUser.Subscription.Package.PackageName)
		for _, channels := range satUser.Subscription.Channel {
			fmt.Print("+", channels.ChannelName)
		}
		fmt.Println("Currently subscribed services:")
		for _, service := range satUser.Subscription.Service {
			fmt.Print(service.ServiceName, ",")
		}
		goto SWITCH
	case 8:
		var email string
		var ph int
		fmt.Println("Update email and phone number for notifications \nEnter the email: ")
		fmt.Scanf("%s", &email)
		fmt.Println("Enter phone:")
		fmt.Scanf("%d", &ph)
		satUser.EmailID = email
		satUser.PhoneNumber = ph
		fmt.Println(" Email and Phone updated successfully ")
		goto SWITCH
	case 9:
		break

	default:
		log.Println("invalid choice !")
		goto SWITCH
	}

}
func ViewServices(pack []pkg.Package, channel []pkg.Channel, service []pkg.Service) {
	fmt.Println("View available packs, channels and services Available packs for subscription")
	for _, v := range pack {
		fmt.Printf("%s :", v.PackageName)
		for _, c := range v.Channel {
			fmt.Printf("%s,", c.ChannelName)

		}
		fmt.Printf(": %f RS . \n", v.Price)
	}

	fmt.Println("Available channels for subscription")
	for _, c := range channel {
		fmt.Printf(" %s: %f Rs. \n", c.ChannelName, c.Price)
	}
	fmt.Println("Available services for subscription ")
	for _, s := range service {
		fmt.Printf(" %s : %f Rs. \n", s.ServiceName, s.Price)
	}
}
