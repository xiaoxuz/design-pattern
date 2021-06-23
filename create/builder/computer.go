package builder

// 产品： 这个是我们的目标，computer 要有这些配置
// computer 可以理解成我们要制作一个什么产品
// 结构体字段 可以理解为我们要做的产品都要哪些配置，对应上文 生成函数的 N 多个入参
type Computer struct {
	name       string    // 电脑类型 比如 mac/华为
	cpuModel   CPU_MODEL // cpu 型号
	gpuModel   GPU_MODEL // gpu 型号
	memorySize MEM_SIZE  // 内存大小
	diskSize   DISK_SIZE // 磁盘大小

	discount float64 // 折扣
	price    float64 // 整体报价
}
