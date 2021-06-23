package builder

import "C"

// 抽象的产品生成器
// 可以理解为computer 这个产品中某一类型产品的生成器
// 抽象生成器即包含了产品(computer)的所有配置，也继承了 builder 公共生成器的所有生成步骤
type HuaweiComputerBuilder struct {
	c *Computer
}

func NewHuaweiComputerBuilder() builder {
	return &HuaweiComputerBuilder{
		c: &Computer{name: "huawei"},
	}
}

func (hc *HuaweiComputerBuilder) getComputer() *Computer {
	return hc.c
}

// 设置 CPU型号
// 设置配置的时候要判断，如果客户端已经配置了，那么跳过
// 这块是因为 director 会在最后编译的时候统一整体执行一遍，防止客户端漏掉配置，走默认配置
func (hc *HuaweiComputerBuilder) setCpuModel(m CPU_MODEL) {
	// demo
	if hc.c.cpuModel != "" {
		return
	}

	if price, ok := partsCpuPriceMap[m]; ok {
		hc.c.cpuModel = m
		hc.c.price += price
	} else {
		hc.c.cpuModel = HW_CPU_I5 // 此为 mac 电脑默认 cpu 配置
		hc.c.price += partsCpuPriceMap[HW_CPU_I5]
	}
}

// 设置 GPU型号
func (hc *HuaweiComputerBuilder) setGpuModel(m GPU_MODEL) {
	// demo
	if hc.c.gpuModel != "" {
		return
	}
	if price, ok := partsGpuPriceMap[m]; ok {
		hc.c.gpuModel = m
		hc.c.price += price
	} else {
		hc.c.gpuModel = HW_GPU_NVIDIA // 此为 mac 电脑默认 gpu 配置
		hc.c.price += partsGpuPriceMap[HW_GPU_NVIDIA]
	}
}

// 设置内存大小
func (hc *HuaweiComputerBuilder) setMemorySize(s MEM_SIZE) {
	// demo
	if hc.c.memorySize != "" {
		return
	}
	if price, ok := partsMemPriceMap[s]; ok {
		hc.c.memorySize = s
		hc.c.price += price
	} else {
		hc.c.memorySize = HW_MEM_8G // 此为 mac 电脑默认 内存 配置
		hc.c.price += partsMemPriceMap[HW_MEM_8G]
	}
}

// 设置 磁盘大小
func (hc *HuaweiComputerBuilder) setDiskSize(s DISK_SIZE) {
	// demo
	if hc.c.diskSize != "" {
		return
	}
	if price, ok := partsDiskPriceMap[s]; ok {
		hc.c.diskSize = s
		hc.c.price += price
	} else {
		hc.c.diskSize = HW_DISK_500G // 此为 mac 电脑默认 磁盘 配置
		hc.c.price += partsDiskPriceMap[HW_DISK_500G]
	}
}

func (hc *HuaweiComputerBuilder) setDiscount() {
	// 华为机器不打折，国产赞。 这块就是和 mac 差异化的地方
	hc.c.discount = 1
}

// 既然华为不打折，那么直接输出就好了
func (hc *HuaweiComputerBuilder) calculatePrice() {
}
