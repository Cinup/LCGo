package main

import (
	"strings"
)

//数瓶子
//input := [100]int{0}
//result := [100]int{0}
//num := 0
//for {
//	tmp := 0
//	fmt.Scanf("%d", &tmp)
//	if tmp != 0 {
//		input[num] = tmp
//		num++
//	} else {
//		break
//	}
//}
//for i := 0; i < num; i++ {
//	num := input[i]
//	for num >= 3 {
//		result[i] += num / 3
//		num = num/3 + num%3
//		//fmt.Println("tmpResult", tmpResult)
//		if num == 2 {
//			result[i]++
//		}
//	}
//}
//for i := 0; i < num; i++ {
//	fmt.Println(result[i])
//}
//==================================================
//去重排序
//需要处理多组数据
//num := 0
//for {
//	n, _ := fmt.Scan(&num)
//	if n == 0 {
//		break
//	} else {
//		array := [1001]int{0}
//		for i := 0; i < num; i++ {
//			input := 0
//			fmt.Scanf("%d", &input)
//			if array[input] == 0 {
//				array[input] = 1
//			}
//		}
//		s := make([]int, 1001)
//		num = 0
//		for i := 0; i < 1001; i++ {
//			if array[i] != 0 {
//				s[num] = i
//				num++
//			}
//		}
//		sort.Ints(s)
//		s = s[len(s)-num:]
//		for i := 0; i < num; i++ {
//			fmt.Println(s[i])
//		}
//	}
//}
//===============================================
//16进制转10进制
//var s string
//for {
//	n, _ := fmt.Scanf("%s", &s)
//	if n == 0 {
//		break
//	} else {
//		result := 0
//		for i := 2; i < len(s); i++ {
//			result = result * 16
//			switch s[i] {
//			case 'A':
//				result = result + 10
//			case 'B':
//				result = result + 11
//			case 'C':
//				result = result + 12
//			case 'D':
//				result = result + 13
//			case 'E':
//				result = result + 14
//			case 'F':
//				result = result + 15
//			default:
//				result = result + int(s[i]) - 48
//			}
//		}
//		fmt.Println(result)
//	}
//}
//=============================================
//stuNum, opNum := 0, 0
//stu := make([]int, 30001)
//for {
//	n, _ := fmt.Scanf("%d %d", &stuNum, &opNum)
//	if n == 0 {
//		break
//	}
//	for i := 1; i <= stuNum; i++ {
//		fmt.Scanf("%d", &stu[i])
//	}
//	var op string
//	s, e := 0, 0
//	for i := 0; i < opNum; i++ {
//		fmt.Scanf("%s %d %d", &op, &s, &e)
//		switch op {
//		case "Q":
//			var tmp []int
//			//不保证Q A B 中 A > B 也可能是A < B
//			if s > e {
//				tmp = stu[e : s+1]
//			} else {
//				tmp = stu[s : e+1]
//			}
//			sortarrays := make([]int, 0)
//			sortarrays = append(sortarrays, tmp...)
//			sort.Ints(sortarrays)
//			fmt.Println(sortarrays[len(sortarrays)-1])
//		case "U":
//			stu[s] = e
//		}
//	}
//}
//=====================================
//type Record struct {
//	filename string
//	rank     int
//	num      int
//}
//type Records []Record
//
//func (rs Records) Len() int {
//	return len(rs)
//}
//func (rs Records) Less(i, j int) bool {
//	if rs[i].num == rs[j].num {
//		return rs[i].rank < rs[j].rank
//	} else {
//		return rs[i].num > rs[j].num
//	}
//}
//func (rs Records) Swap(i, j int) {
//	rs[i], rs[j] = rs[j], rs[i]
//}
//
//func main() {
//	rs := make(Records, 1000)
//	//键值为文件名加行数
//	input := make(map[string]int)
//	file := ""
//	line := 0
//
//	//记录Records数量
//	index := 0
//	//记录文件名出现的次数,当出现次数一样时,需要根据这个排序
//	findex := 0
//
//	for {
//		n, _ := fmt.Scanf("%s %d", &file, &line)
//		if n == 0 {
//			break
//		}
//		//分割字符得到文件名
//		file = file[strings.LastIndex(file, "\\")+1:]
//		//使用:将文件名和行数连接起来,后续容易分割
//		file = file + ":" + strconv.Itoa(line)
//		v, ok := input[file]
//		if ok {
//			rs[v].num++
//		} else {
//			input[file] = index
//			if index > 1000 {
//				rs = append(rs, Record{file, findex, 1})
//			} else {
//				rs[index] = Record{file, findex, 1}
//			}
//			index++
//		}
//		findex++
//	}
//	//排序
//	sort.Sort(rs)
//	//最多输出前八个
//	for i := 0; i < 8 && i < index; i++ {
//		//将文件名和行号分割,因为这会影响文件名字数的判断
//		fn := rs[i].filename
//		si := strings.LastIndex(fn, ":")
//		f := fn[:si]
//		num := fn[si+1:]
//		if len(f) > 16 {
//			//Println会自动加一个空格
//			fmt.Println(f[len(f)-16:], num, rs[i].num)
//		} else {
//			fmt.Println(f, num, rs[i].num)
//		}
//	}
//}
//=========================================
//func convertToTitle(n int) string {
//	return
//}
func shortestToChar(S string, C byte) []int {
	var distance = make([]int, 10000)
	firstIndex := strings.Index(S, string(C))
	for i := 0; i <= firstIndex; i++ {
		distance[i] = firstIndex - i
	}
	dis := 0
	for i := firstIndex + 1; i < len(S); i++ {
		if S[i] == C {
			dis = 0
		} else {
			dis++
		}
		distance[i] = dis

	}
	lastIndex := strings.LastIndex(S, string(C))
	for i := len(S) - 1; i >= lastIndex; i-- {
		distance[i] = i - lastIndex
	}
	dis = 0
	for i := lastIndex - 1; i >= 0; i-- {
		if S[i] == C {
			dis = 0
		} else {
			dis++
		}
		if dis < distance[i] {
			distance[i] = dis
		}

	}
	return distance[:len(S)+1]
}
func main() {
	S := "aaba"
	var C byte = 'b'
	shortestToChar(S, C)
}
