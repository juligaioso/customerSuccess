package csbalancing

import (
	"sort"
)

// Entity ...
type Entity struct {
	ID      int
	Score   int
	Counter int
}

// CustomerSuccessBalancing ...
func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessS[]int) int {
	// Write your solution here

	availableCustomerSuccess := []Entity{}
	for i := len(customerSuccess) - 1; i >= 0; i-- {
		containsCustomerSuccess := ContainsCustomerSuccess(customerSuccessS, customerSuccess[i])
		if !containsCustomerSuccess {
			availableCustomerSuccess = append(availableCustomerSuccess, customerSuccess[i])
		}
	}

	sort.SliceStable(availableCustomerSuccess, func(i, j int) bool {
		return availableCustomerSuccess[i].Score < availableCustomerSuccess[j].Score
	})

	var customerSuccessVerify = availableCustomerSuccess[len(availableCustomerSuccess)-1]
	var customerLast = customers[len(customers)-1]
	if customerSuccessVerify.Score == customerLast.Score {
		return customerSuccessVerify.ID
	}

	var maxCounter int = 0
	var sum int = 1
	for _, customer := range customers {
		for i := 0; i < len(availableCustomerSuccess); i++ {
			if availableCustomerSuccess[i].Score >= customer.Score {
				availableCustomerSuccess[i].Counter += sum
				if availableCustomerSuccess[i].Counter > maxCounter {
					maxCounter = availableCustomerSuccess[i].Counter
				}

				break
			}
		}
	}

	var customerSuccessFoundByCounter = FindCustomerSuccessByCounter(availableCustomerSuccess, maxCounter)
	var sizeCompare int = 1
	if len(customerSuccessFoundByCounter) == sizeCompare {
		return customerSuccessFoundByCounter[0].ID
	}

	return 0
}

func ContainsCustomerSuccess(customerSuccessId []int, css Entity) bool {
	for _, id := range customerSuccessId {
		if id == css.ID {
			return true
		}
	}
	return false
}

func FindCustomerSuccessByCounter(customerSuccess []Entity, maxCounter int) []Entity {
	customerSuccessFound := []Entity{}
	for _, css := range customerSuccess {
		if css.Counter == maxCounter {
			customerSuccessFound = append(customerSuccessFound, css)
		}
	}
	return customerSuccessFound
}
