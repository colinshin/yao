package helper

import "github.com/yaoapp/gou"

// ProcessArrayPluck  xiang.helper.ArrayPluck 将多个数据记录集合，合并为一个数据记录集合
func ProcessArrayPluck(process *gou.Process) interface{} {
	process.ValidateArgNums(2)
	columns := process.ArgsStrings(0)
	pluck := process.ArgsMap(1)
	return ArrayPluck(columns, pluck)
}

// ProcessArraySplit  xiang.helper.ArraySplit 将多条数记录集合，分解为一个 columns:[]string 和 values: [][]interface{}
func ProcessArraySplit(process *gou.Process) interface{} {
	process.ValidateArgNums(1)
	records := process.ArgsRecords(0)
	columns, values := ArraySplit(records)
	return map[string]interface{}{
		"columns": columns,
		"values":  values,
	}
}

// ProcessArrayColumn  xiang.helper.ArrayColumn  返回多条数据记录，指定字段数值。
func ProcessArrayColumn(process *gou.Process) interface{} {
	process.ValidateArgNums(2)
	records := process.ArgsRecords(0)
	name := process.ArgsString(1)
	values := ArrayColumn(records, name)
	return values
}

// ProcessArrayKeep  xiang.helper.ArrayKeep  仅保留指定键名的数据
func ProcessArrayKeep(process *gou.Process) interface{} {
	process.ValidateArgNums(2)
	records := process.ArgsRecords(0)
	columns := process.ArgsStrings(1)
	return ArrayKeep(records, columns)
}

// ProcessArrayTree  xiang.helper.ArrayTree  转换为属性结构
func ProcessArrayTree(process *gou.Process) interface{} {
	process.ValidateArgNums(2)
	records := process.ArgsRecords(0)
	setting := process.ArgsMap(1)
	return ArrayTree(records, setting)
}