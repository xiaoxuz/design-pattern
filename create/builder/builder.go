package builder

// 生成器接口： 产品的生成器接口
// 可以理解为这个产品生成 必须要有哪些具体的步骤和行为, 后面每一个抽象的产品生成对象都要继承这个生成器接口
type builder interface {
	setCpuModel(CPU_MODEL)  // 设置 cpu 型号
	setGpuModel(GPU_MODEL)  // 设置 gpu 型号
	setMemorySize(MEM_SIZE) // 设置 内存型号
	setDiskSize(DISK_SIZE)  // 设置磁盘型号

	setDiscount()    // 设置折扣粒度, 这个折扣粒度是系统内置的，不需要客户端设置也就是说此功能不是给前台用户询价时自定义的。
	calculatePrice() // 计算报价

	getComputer() *Computer
}
