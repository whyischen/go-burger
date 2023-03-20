package 面试题

import (
	"math/rand"
	"time"
)

/***
*流体网络*
公司有n个组,>=5，每组人数相同，>=2人，需要进行随机组队吃饭。

要求：
1. 两两一队或三人一队，不能落单
2. 两人队、三人队各自的队伍数均不得少于2
3. 一个人只出现一次
4. 队伍中所有人不能来自相同组
5. 随机组队，重复执行程序得到的结果不一样，总队伍数也不能一样
6. 注释注释注释
注：要同时满足条件1-6,
举例：
GroupList = [ # 小组列表
['小名', '小红', '小马', '小丽', '小强'],
['大壮', '大力', '大1', '大2', '大3'],
['阿花', '阿朵', '阿蓝', '阿紫', '阿红'],
['A', 'B', 'C', 'D', 'E'],
['一', '二', '三', '四', '五'],
['建国', '建军', '建民', '建超', '建跃'],
['爱民', '爱军', '爱国', '爱辉', '爱月']
]
输入：GroupList
示例输出:
[小强 大3] [阿红 E] [五 建跃] [爱月 小名] [大壮 阿花] [A 一] [建国 爱民] [小红 大力] [阿朵 B]
[二 建军] [爱军 小马] [大1 阿蓝] [C 三] [建民 爱国 小丽] [大2 阿紫 D] [四 建超 爱辉]
*/

func randomGroup(groupList [][]string) ([][]string, error) {
	//groupNum := len(groupList)
	//groupSize := len(groupList[0])

	return nil, nil
}

// 没解决【两人队、三人队各自的队伍数均不得少于2】
// funcGroupList函数接收一个二维字符串数组GroupList作为参数，返回一个二维字符串数组result
func funcGroupList(GroupList [][]string) [][]string {
	// 初始化result为一个空的二维字符串数组
	result := [][]string{[]string{}}
	// 初始化bit_index为0
	bit_index := 0
	// 获取GroupList的行数
	grouplist_y := len(GroupList)
	// 获取GroupList的列数
	grouplist_x := len(GroupList[0])
	// 初始化last_index为-1
	last_index := -1
	// 初始化随机数生成器r
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 当bit_index小于grouplist_y时，循环执行以下操作
	for bit_index < grouplist_y {
		// 生成一个随机数rand_y，表示在GroupList中随机选择一行
		rand_y := r.Intn(grouplist_y)
		// 如果上一次选择的行和这一次选择的行相同，则跳过本次循环
		if last_index == rand_y {
			continue
		}
		// 生成一个随机数rand_x，表示在GroupList[rand_y]中随机选择一个元素
		rand_x := r.Intn(grouplist_x)
		// 如果GroupList[rand_y]为空或者rand_x超出了GroupList[rand_y]的索引范围，则跳过本次循环
		if len(GroupList[rand_y]) == 0 || len(GroupList[rand_y]) <= rand_x {
			continue
		}
		// 获取GroupList[rand_y][rand_x]的值
		val := GroupList[rand_y][rand_x]
		// 从GroupList[rand_y]中删除GroupList[rand_y][rand_x]
		GroupList[rand_y] = append(GroupList[rand_y][:rand_x], GroupList[rand_y][rand_x+1:]...)
		// 将GroupList[rand_y][rand_x]的值添加到result的最后一个元素中
		result[len(result)-1] = append(result[len(result)-1], val)
		// 将last_index设置为rand_y
		last_index = rand_y
		if len(result[len(result)-1]) >= 2 {
			last_index = -1
			result = append(result, []string{})
		}
		if len(GroupList[rand_y]) == 0 {
			bit_index++
		}
	}
	if len(result[len(result)-1]) == 1 {
		temp := result[len(result)-1]
		result = result[:len(result)-1]
		result[len(result)-1] = append(result[len(result)-1], temp...)
	}
	if len(result[len(result)-1]) == 0 {
		return result[:len(result)-1]
	}
	return result
}
