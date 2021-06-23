package builder

import (
	"time"
)

// 抽象的产品生成器
// 可以理解为computer 这个产品中某一类型产品的生成器
// 抽象生成器即包含了产品(computer)的所有配置，也继承了 builder 公共生成器的所有生成步骤
type MacComputerBuilder struct {
	c *Computer
}

func NewMacComputerBuilder() builder {
	return &MacComputerBuilder{
		c: &Computer{name: "mac"},
	}
}

func (mc *MacComputerBuilder) getComputer() *Computer {
	return mc.c
}

// 设置 CPU型号
// 设置配置的时候要判断，如果客户端已经配置了，那么跳过
// 这块是因为 director 会在最后编译的时候统一整体执行一遍，防止客户端漏掉配置，走默认配置
func (mc *MacComputerBuilder) setCpuModel(m CPU_MODEL) {
	// demo
	if mc.c.cpuModel != "" {
		return
	}
	if price, ok := partsCpuPriceMap[m]; ok {
		mc.c.cpuModel = m
		mc.c.price += price
	} else {
		mc.c.cpuModel = MAC_CPU_I5 // 此为 mac 电脑默认 cpu 配置
		mc.c.price += partsCpuPriceMap[MAC_CPU_I5]
	}
}

// 设置 GPU型号
func (mc *MacComputerBuilder) setGpuModel(m GPU_MODEL) {
	// demo
	if mc.c.gpuModel != "" {
		return
	}
	if price, ok := partsGpuPriceMap[m]; ok {
		mc.c.gpuModel = m
		mc.c.price += price
	} else {
		mc.c.gpuModel = MAC_GPU_NVIDIA // 此为 mac 电脑默认 gpu 配置
		mc.c.price += partsGpuPriceMap[MAC_GPU_NVIDIA]
	}
}

// 设置内存大小
func (mc *MacComputerBuilder) setMemorySize(s MEM_SIZE) {
	// demo
	if mc.c.memorySize != "" {
		return
	}
	if price, ok := partsMemPriceMap[s]; ok {
		mc.c.memorySize = s
		mc.c.price += price
	} else {
		mc.c.memorySize = MAC_MEM_8G // 此为 mac 电脑默认 内存 配置
		mc.c.price += partsMemPriceMap[MAC_MEM_8G]
	}
}

// 设置 磁盘大小
func (mc *MacComputerBuilder) setDiskSize(s DISK_SIZE) {
	// demo
	if mc.c.diskSize != "" {
		return
	}
	if price, ok := partsDiskPriceMap[s]; ok {
		mc.c.diskSize = s
		mc.c.price += price
	} else {
		mc.c.diskSize = MAC_DISK_500G // 此为 mac 电脑默认 磁盘 配置
		mc.c.price += partsDiskPriceMap[MAC_DISK_500G]
	}
}

// 设置折扣
// 此操作为内置操作，不需要外部设置
func (mc *MacComputerBuilder) setDiscount() {
	// 2021-06-24 00:17:33
	// 如果大于这个时间，那么 mac 电脑整体打5折
	// 否则 整体打8折
	if time.Now().Unix() > 1624465043 {
		mc.c.discount = 0.5
	} else {
		mc.c.discount = 0.8
	}
}

// 计数价格
// 注意看，这块就是需要时序的地方，需要先setDiscount 才能进行报价
// 所以 需要通过 指挥者来统一进行构建，保证各个行为执行顺序
func (mc *MacComputerBuilder) calculatePrice() {
	mc.c.price = (mc.c.price * mc.c.discount)
}
