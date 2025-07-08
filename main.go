package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// check if the sizes of the arrays are the same
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	// create a map
	userMap := make(map[string]user)

	// iterate over the names and phone numbers
	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{names[i], phoneNumbers[i]}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
