package main


func Sum(arr1 []int, arr2 [] int)[]int{
	var res [] int

	for i := range arr1{
		res = append(res, arr1[i] + arr2[i])
	}

	return res

}


func main(){


}