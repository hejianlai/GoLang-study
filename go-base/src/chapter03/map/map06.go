package main

import `fmt`

func modifyUser(users map[string]map[string]string,name string)  {

	// 判断用户是否存在
	// 用户存在
	if users[name] != nil {
		users[name]["pwd"] = "99999"
	}else {
		// 用户不存在
		users[name] = make(map[string]string,2)
		users[name]["pwd"] = "99999"
		users[name]["nickname"] = "昵称~" + name
	}

}
func main()  {
	users := make(map[string]map[string]string,10)
	users["smith"] = make(map[string]string,2)
	users["smith"]["pwd"] = "8888"
	users["smith"]["nickname"] = "笑话吗"
	modifyUser(users,"tom")
	modifyUser(users,"mary")
	modifyUser(users,"smith")
	fmt.Println(users)
	for  k,v := range users{
		fmt.Println(k,v)
		for k2,v2 := range v{
			fmt.Printf("\t%v=%v\n",k2,v2)
		}
	}
}
