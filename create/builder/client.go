package builder

import "fmt"

// 客户端询问报价
// 即用户在前台页面选择了 mac 电脑
// CPU i7
// GPU xxx
func getPrice() {
	// 先实例化抽象生成器对象，即 mac 电脑
	mcb := NewMacComputerBuilder()
	// 设置我想询问的配置
	mcb.setCpuModel(MAC_CPU_I7)
	mcb.setGpuModel(MAC_GPU_NVIDIA)
	mcb.setMemorySize(MAC_MEM_16G)
	// 磁盘我不选了，用默认的
	//mcb.setDiskSize()

	// 然后实例化一个主管，来准备生成报价
	d := NewDirector(mcb)
	// 执行编译，生成最终产品
	product := d.buildComputer()

	// ok 搞定了，我们可以看看最终这个产品的配置和报价
	fmt.Printf("current computer name: %s\n", product.name)
	fmt.Printf("choose config cpuModel: %s\n", product.cpuModel)
	fmt.Printf("choose config gpuModel: %s\n", product.gpuModel)
	fmt.Printf("choose config memorySize: %s\n", product.memorySize)
	fmt.Printf("choose config diskSize: %s\n", product.diskSize)

	fmt.Printf("give you discount: %f\n", product.discount)
	fmt.Printf("final offer: %f\n", product.price)

	fmt.Printf("---------------询问下一个电脑---------------\n")
	// 下面 我们再生成一个华为的电脑报价
	hwcb := NewHuaweiComputerBuilder()
	hwcb.setCpuModel(HW_CPU_I7)
	hwcb.setGpuModel(HW_GPU_ATI)
	hwcb.setMemorySize(HW_MEM_16G)
	hwcb.setDiskSize(HW_DISK_1T)
	d.resetBuilder(hwcb)
	// 执行编译，生成最终产品
	product2 := d.buildComputer()

	// ok 搞定了，我们可以看看最终这个产品的配置和报价
	fmt.Printf("current computer name: %s\n", product2.name)
	fmt.Printf("choose config cpuModel: %s\n", product2.cpuModel)
	fmt.Printf("choose config gpuModel: %s\n", product2.gpuModel)
	fmt.Printf("choose config memorySize: %s\n", product2.memorySize)
	fmt.Printf("choose config diskSize: %s\n", product2.diskSize)

	fmt.Printf("give you discount: %f\n", product2.discount)
	fmt.Printf("final offer: %f\n", product2.price)

}
