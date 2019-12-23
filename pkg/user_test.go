package pkg_test

import (
	. "SatTV/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Recharge", func() {
When("recharge successful", func() {
	user:=User{
		UserName:     "sree",
		EmailID:      "sree@gmail.com",
		PhoneNumber:  963301456,
		Balance:      100,
		Subscription: Subscription{},
	}
	It("should match the response", func() {
		user.Recharge(50)
		Expect(user.Balance).To(Equal(float32(150)))
	})

})
})
var _=Describe( "PackageSubscription", func() {
	When("successfully subscribed silver package  for 3 month", func() {
		user:=User{
			UserName:     "sree",
			EmailID:      "sree@gmail.com",
			PhoneNumber:  963301456,
			Balance:      300,
			Subscription: Subscription{
			},
		}
		It("should match responses subscription more than 3 months", func() {
			real,discount,final:=user.PackSubscription(SilverPack,3)
			Expect(real).To(Equal(float32(150)))
			Expect(discount).To(Equal(float32(15)))
			Expect(final).To(Equal(float32(135)))
			Expect(user.Subscription.Package.PackageName).To(Equal(SilverPack.PackageName))
		})
		It("should match response for subscription less than 3 months", func() {
			real,discount,final:=user.PackSubscription(SilverPack,2)
			Expect(real).To(Equal(float32(100)))
			Expect(discount).To(Equal(float32(0)))
			Expect(final).To(Equal(float32(100)))
			Expect(user.Subscription.Package.PackageName).To(Equal(SilverPack.PackageName))
		})
	})
})
var _=Describe("Withdrawal", func() {
	When("amount is successfully withdrawn", func() {
		user:=User{
			UserName:     "sree",
			EmailID:      "sree@gmail.com",
			PhoneNumber:  963301456,
			Balance:      300,
			Subscription: Subscription{
			},
		}
		It("should match the response with no error", func() {
			err:=user.Withdrawal(100)
			Expect(err).To(BeNil())
			Expect(user.Balance).To(Equal(float32(200)))
		})
	})
	When("amount is failed to  withdraw", func() {
		user:=User{
			UserName:     "sree",
			EmailID:      "sree@gmail.com",
			PhoneNumber:  963301456,
			Balance:      300,
			Subscription: Subscription{
			},
		}
		It("should match the response with no error", func() {
			err:=user.Withdrawal(400)
			Expect(err).NotTo(BeNil())
			Expect(user.Balance).To(Equal(float32(300)))
		})
	})
})
var _=Describe("Add channels", func() {
	When("Successfully added channels for 1 month subscription", func() {
		user:=User{
			UserName:     "sree",
			EmailID:      "sree@gmail.com",
			PhoneNumber:  963301456,
			Balance:      300,
			Subscription: Subscription{},
		}
		It("should add channel for specified months", func() {
			channels:="Zee,Sony,Star Plus,Discovery,NatGeo"
			amount,err:=user.AddChannels(channels,2)
			Expect(amount).To(Equal(float32(150)))
			Expect(err).To(BeNil())
		})
	})
	When("failed to added channels for 1 month subscription", func() {
		user:=User{
			UserName:     "sree",
			EmailID:      "sree@gmail.com",
			PhoneNumber:  963301456,
			Balance:      300,
			Subscription: Subscription{},
		}
		It("should add channel for specified months", func() {
			channels:="zee,Sony,Star Plus,Discovery,NatGeo"
			amount,err:=user.AddChannels(channels,2)
			Expect(amount).To(Equal(float32(0)))
			Expect(err).NotTo(BeNil())
		})
	})
})
var _=Describe("SubscriptionService", func() {
	When("service subscribed successful", func() {
		user:=User{
			UserName:     "sree",
			EmailID:      "sree@gmail.com",
			PhoneNumber:  963301456,
			Balance:      300,
			Subscription: Subscription{},
		}
		It("should match the response value", func() {
			amount:=user.SubscriptionService(LearnEnglish,1)
			Expect(amount).To(Equal(float32(200)))
		})
	})
})
var _=Describe("viewBalance", func() {
	user:=User{
		UserName:     "sree",
		EmailID:      "sree@gmail.com",
		PhoneNumber:  963301456,
		Balance:      300,
		Subscription: Subscription{},
	}
	It("should print response", func() {
		user.ViewBalance()
	})
})
var _=Describe("CreateUser", func() {
	user:=User{
		UserName:     "",
		EmailID:      "",
		PhoneNumber:  0,
		Balance:      0,
		Subscription: Subscription{},
	}
	When("successfully created the client", func() {
		It("should return no error", func() {
			username:="sree"
			email:="sree@gmail.com"
			ph:=9896325574
			err:=user.CreateUser(username,email,ph)
			Expect(err).To(BeNil())
		})

	})
})