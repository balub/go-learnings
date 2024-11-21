package main

func getUserByID(userList []User, userID int) User {
	var user User

	for _, u := range userList {
		if u.Id == userID {
			user = u
		}
	}

	return user
}

func UpdateUserDetails(userList *[]User, userID int, newName string) {
	for i := range *userList {
		if (*userList)[i].Id == userID {
			(*userList)[i].Name = newName
			return
		}
	}
}
