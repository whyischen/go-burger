package question

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
