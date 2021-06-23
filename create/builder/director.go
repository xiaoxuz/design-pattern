package builder

// director 主管，负责整体 build 执行
// 可以理解为总指挥，他来负责计算报价
type director struct {
	builder builder
}

// 实例化一个主管
func NewDirector(b builder) *director {
	return &director{
		builder: b,
	}
}

// 手动重置主管，方便进行多次不同产品生成构建
func (d *director) resetBuilder(b builder) {
	d.builder = b
}

// 执行编译生成，这块就是要严格统一管理编译的步骤和顺序
// 当前这个 demo , 因为时计算报价的例子而不是生成电脑配置的例子，所以前置的那些 setXXX 都在客户端自定义执行了
// 但是有可能前台用户没有选择某些配置，所以需要主管统一兜底
// 1. 兜底每个电脑配置
// 2. 根据当前时间选择折扣粒度
// 3. 计算报价
func (d *director) buildComputer() *Computer {
	// 第一步，兜底每一个电脑配置
	d.builder.setCpuModel(DIRECTOR_CHECK_PARAMS)
	d.builder.setGpuModel(DIRECTOR_CHECK_PARAMS)
	d.builder.setMemorySize(DIRECTOR_CHECK_PARAMS)
	d.builder.setDiskSize(DIRECTOR_CHECK_PARAMS)

	// 第二步设置折扣
	d.builder.setDiscount()
	// 第三步 计算报价
	d.builder.calculatePrice()

	// 返回产品对象
	return d.builder.getComputer()
}
