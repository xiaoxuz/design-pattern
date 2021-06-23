package builder

// 兜底校验入参
const DIRECTOR_CHECK_PARAMS = "nil"

// CPU
type CPU_MODEL string

const (
	MAC_CPU_I5 CPU_MODEL = "maci5"
	MAC_CPU_I7 CPU_MODEL = "maci7"

	HW_CPU_I5 CPU_MODEL = "hwi5"
	HW_CPU_I7 CPU_MODEL = "hwi7"
)

var partsCpuPriceMap = map[CPU_MODEL]float64{
	MAC_CPU_I5: 100,
	MAC_CPU_I7: 500,
	HW_CPU_I5:  200,
	HW_CPU_I7:  700,
}

// GPU
type GPU_MODEL string

const (
	MAC_GPU_NVIDIA GPU_MODEL = "mac-NVIDIA"
	MAC_GPU_ATI    GPU_MODEL = "mac-ATI"

	HW_GPU_NVIDIA GPU_MODEL = "hw-NVIDIA"
	HW_GPU_ATI    GPU_MODEL = "hw-ATI"
)

var partsGpuPriceMap = map[GPU_MODEL]float64{
	MAC_GPU_NVIDIA: 100,
	MAC_GPU_ATI:    500,
	HW_GPU_NVIDIA:  200,
	HW_GPU_ATI:     700,
}

// Memory
type MEM_SIZE string

const (
	MAC_MEM_8G  MEM_SIZE = "mac-8g"
	MAC_MEM_16G MEM_SIZE = "mac-16g"

	HW_MEM_8G  MEM_SIZE = "hw-8g"
	HW_MEM_16G MEM_SIZE = "hw-16g"
)

var partsMemPriceMap = map[MEM_SIZE]float64{
	MAC_MEM_8G:  100,
	MAC_MEM_16G: 500,
	HW_MEM_8G:   200,
	HW_MEM_16G:  700,
}

// Disk
type DISK_SIZE string

const (
	MAC_DISK_500G DISK_SIZE = "mac-500g"
	MAC_DISK_1T   DISK_SIZE = "mac-1t"
	HW_DISK_500G  DISK_SIZE = "hw-500g"
	HW_DISK_1T    DISK_SIZE = "hw-1t"
)

var partsDiskPriceMap = map[DISK_SIZE]float64{
	MAC_DISK_500G: 100,
	MAC_DISK_1T:   500,
	HW_DISK_500G:  200,
	HW_DISK_1T:    700,
}
