package __1


func InsertionSort(source []int)(desc []int){
	if len(source)==0||len(source)==1{
		desc = source
		return
	}
	for j:=1;j<len(source)-1;j++{
		key:=source[j]
		i:=j-1
		for i>=0 &&source[i]>key{
			source[i+1] = source[i]
			i--
		}
		source[i+1] = key
	}
	desc = source
	return
}



