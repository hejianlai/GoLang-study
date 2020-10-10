package main

import `fmt`

func main()  {
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string,2)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	fmt.Println(studentMap["stu01"])
	studentMap["stu02"] = make(map[string]string,3)
	studentMap["stu02"]["name"] = "xx"
	studentMap["stu02"]["sex"] = "男"
	studentMap["stu02"]["addr"] = "guangxi"
	studentMap["stu02"]["addrs"] = "guangxi"
	studentMap["stu02"]["age"] = "23"
	fmt.Println(studentMap["stu02"])
	studentMap["stu01"]["sex"] = "女"
	fmt.Println(studentMap["stu01"])
	delete(studentMap,"stu01")
	fmt.Println(studentMap)
	fmt.Println(len(studentMap))
	fmt.Println("-------------")
	for k,v := range studentMap{
		fmt.Println(k,v)
		for k2,v2 := range v{
			fmt.Printf("\t k2=%v v2=%v\n",k2,v2)
		}
		fmt.Println()
	}

}
