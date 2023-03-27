package goroutinedata

import "fmt"

func DummyData1() {

	dataSet := DataTest{
		User: &User{
			Email:    "exampleUser1@gmail.com",
			Username: "user1",
		},
		Order: &Order{
			OrderId:    "001",
			TotalOrder: 2,
		},
	}
	result := map[string]interface{}{
		"UserInfo":  *dataSet.User,
		"OrderInfo": *dataSet.Order,
	}
	fmt.Println(result)
}

func DummyData2() {

	dataSet := DataTest{
		User: &User{
			Email:    "exampleUser2@gmail.com",
			Username: "user2",
		},
		Order: &Order{
			OrderId:    "002",
			TotalOrder: 100,
		},
	}

	result := map[string]interface{}{
		"UserInfo":  *dataSet.User,
		"OrderInfo": *dataSet.Order,
	}
	fmt.Println(result)

}
