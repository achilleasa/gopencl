package cl

/*
#cgo darwin  LDFLAGS: -framework OpenCL
#cgo linux   LDFLAGS: -lCL
#cgo windows LDFLAGS: -lopencl32
#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif
extern void contextErrorCallback(char *, void *, size_t, void *);
extern void memObjectDestroyCallback(cl_mem,void*);
extern void programObjectBuildCompleteCallback(cl_program, void*);
*/
import "C"
import (
	"unsafe"
)

func init() {
	mochHolder = make(map[*memObjectCallbackHolder]struct{})
	pobchHolder = make(map[*programObjectBuildCompleteHolder]struct{})
}

const (
	SUCCESS                                   = C.CL_SUCCESS
	DEVICE_NOT_FOUND                          = C.CL_DEVICE_NOT_FOUND
	DEVICE_NOT_AVAILABLE                      = C.CL_DEVICE_NOT_AVAILABLE
	COMPILER_NOT_AVAILABLE                    = C.CL_COMPILER_NOT_AVAILABLE
	MEM_OBJECT_ALLOCATION_FAILURE             = C.CL_MEM_OBJECT_ALLOCATION_FAILURE
	OUT_OF_RESOURCES                          = C.CL_OUT_OF_RESOURCES
	OUT_OF_HOST_MEMORY                        = C.CL_OUT_OF_HOST_MEMORY
	PROFILING_INFO_NOT_AVAILABLE              = C.CL_PROFILING_INFO_NOT_AVAILABLE
	MEM_COPY_OVERLAP                          = C.CL_MEM_COPY_OVERLAP
	IMAGE_FORMAT_MISMATCH                     = C.CL_IMAGE_FORMAT_MISMATCH
	IMAGE_FORMAT_NOT_SUPPORTED                = C.CL_IMAGE_FORMAT_NOT_SUPPORTED
	BUILD_PROGRAM_FAILURE                     = C.CL_BUILD_PROGRAM_FAILURE
	MAP_FAILURE                               = C.CL_MAP_FAILURE
	MISALIGNED_SUB_BUFFER_OFFSET              = C.CL_MISALIGNED_SUB_BUFFER_OFFSET
	EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST = C.CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST
	COMPILE_PROGRAM_FAILURE                   = C.CL_COMPILE_PROGRAM_FAILURE
	LINKER_NOT_AVAILABLE                      = C.CL_LINKER_NOT_AVAILABLE
	LINK_PROGRAM_FAILURE                      = C.CL_LINK_PROGRAM_FAILURE
	DEVICE_PARTITION_FAILED                   = C.CL_DEVICE_PARTITION_FAILED
	KERNEL_ARG_INFO_NOT_AVAILABLE             = C.CL_KERNEL_ARG_INFO_NOT_AVAILABLE
	INVALID_VALUE                             = C.CL_INVALID_VALUE
	INVALID_DEVICE_TYPE                       = C.CL_INVALID_DEVICE_TYPE
	INVALID_PLATFORM                          = C.CL_INVALID_PLATFORM
	INVALID_DEVICE                            = C.CL_INVALID_DEVICE
	INVALID_CONTEXT                           = C.CL_INVALID_CONTEXT
	INVALID_QUEUE_PROPERTIES                  = C.CL_INVALID_QUEUE_PROPERTIES
	INVALID_COMMAND_QUEUE                     = C.CL_INVALID_COMMAND_QUEUE
	INVALID_HOST_PTR                          = C.CL_INVALID_HOST_PTR
	INVALID_MEM_OBJECT                        = C.CL_INVALID_MEM_OBJECT
	INVALID_IMAGE_FORMAT_DESCRIPTOR           = C.CL_INVALID_IMAGE_FORMAT_DESCRIPTOR
	INVALID_IMAGE_SIZE                        = C.CL_INVALID_IMAGE_SIZE
	INVALID_SAMPLER                           = C.CL_INVALID_SAMPLER
	INVALID_BINARY                            = C.CL_INVALID_BINARY
	INVALID_BUILD_OPTIONS                     = C.CL_INVALID_BUILD_OPTIONS
	INVALID_PROGRAM                           = C.CL_INVALID_PROGRAM
	INVALID_PROGRAM_EXECUTABLE                = C.CL_INVALID_PROGRAM_EXECUTABLE
	INVALID_KERNEL_NAME                       = C.CL_INVALID_KERNEL_NAME
	INVALID_KERNEL_DEFINITION                 = C.CL_INVALID_KERNEL_DEFINITION
	INVALID_KERNEL                            = C.CL_INVALID_KERNEL
	INVALID_ARG_INDEX                         = C.CL_INVALID_ARG_INDEX
	INVALID_ARG_VALUE                         = C.CL_INVALID_ARG_VALUE
	INVALID_ARG_SIZE                          = C.CL_INVALID_ARG_SIZE
	INVALID_KERNEL_ARGS                       = C.CL_INVALID_KERNEL_ARGS
	INVALID_WORK_DIMENSION                    = C.CL_INVALID_WORK_DIMENSION
	INVALID_WORK_GROUP_SIZE                   = C.CL_INVALID_WORK_GROUP_SIZE
	INVALID_WORK_ITEM_SIZE                    = C.CL_INVALID_WORK_ITEM_SIZE
	INVALID_GLOBAL_OFFSET                     = C.CL_INVALID_GLOBAL_OFFSET
	INVALID_EVENT_WAIT_LIST                   = C.CL_INVALID_EVENT_WAIT_LIST
	INVALID_EVENT                             = C.CL_INVALID_EVENT
	INVALID_OPERATION                         = C.CL_INVALID_OPERATION
	INVALID_GL_OBJECT                         = C.CL_INVALID_GL_OBJECT
	INVALID_BUFFER_SIZE                       = C.CL_INVALID_BUFFER_SIZE
	INVALID_MIP_LEVEL                         = C.CL_INVALID_MIP_LEVEL
	INVALID_GLOBAL_WORK_SIZE                  = C.CL_INVALID_GLOBAL_WORK_SIZE
	INVALID_PROPERTY                          = C.CL_INVALID_PROPERTY
	INVALID_IMAGE_DESCRIPTOR                  = C.CL_INVALID_IMAGE_DESCRIPTOR
	INVALID_COMPILER_OPTIONS                  = C.CL_INVALID_COMPILER_OPTIONS
	INVALID_LINKER_OPTIONS                    = C.CL_INVALID_LINKER_OPTIONS
	INVALID_DEVICE_PARTITION_COUNT            = C.CL_INVALID_DEVICE_PARTITION_COUNT
	VERSION_1_0                               = C.CL_VERSION_1_0
	VERSION_1_1                               = C.CL_VERSION_1_1
	VERSION_1_2                               = C.CL_VERSION_1_2
	FALSE                                     = C.CL_FALSE
	TRUE                                      = C.CL_TRUE
	BLOCKING                                  = C.CL_BLOCKING
	NON_BLOCKING                              = C.CL_NON_BLOCKING
	PLATFORM_PROFILE                          = C.CL_PLATFORM_PROFILE
	PLATFORM_VERSION                          = C.CL_PLATFORM_VERSION
	PLATFORM_NAME                             = C.CL_PLATFORM_NAME
	PLATFORM_VENDOR                           = C.CL_PLATFORM_VENDOR
	PLATFORM_EXTENSIONS                       = C.CL_PLATFORM_EXTENSIONS
	DEVICE_TYPE_DEFAULT                       = C.CL_DEVICE_TYPE_DEFAULT
	DEVICE_TYPE_CPU                           = C.CL_DEVICE_TYPE_CPU
	DEVICE_TYPE_GPU                           = C.CL_DEVICE_TYPE_GPU
	DEVICE_TYPE_ACCELERATOR                   = C.CL_DEVICE_TYPE_ACCELERATOR
	DEVICE_TYPE_CUSTOM                        = C.CL_DEVICE_TYPE_CUSTOM
	DEVICE_TYPE_ALL                           = C.CL_DEVICE_TYPE_ALL
	DEVICE_TYPE                               = C.CL_DEVICE_TYPE
	DEVICE_VENDOR_ID                          = C.CL_DEVICE_VENDOR_ID
	DEVICE_MAX_COMPUTE_UNITS                  = C.CL_DEVICE_MAX_COMPUTE_UNITS
	DEVICE_MAX_WORK_ITEM_DIMENSIONS           = C.CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS
	DEVICE_MAX_WORK_GROUP_SIZE                = C.CL_DEVICE_MAX_WORK_GROUP_SIZE
	DEVICE_MAX_WORK_ITEM_SIZES                = C.CL_DEVICE_MAX_WORK_ITEM_SIZES
	DEVICE_PREFERRED_VECTOR_WIDTH_CHAR        = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR
	DEVICE_PREFERRED_VECTOR_WIDTH_SHORT       = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT
	DEVICE_PREFERRED_VECTOR_WIDTH_INT         = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT
	DEVICE_PREFERRED_VECTOR_WIDTH_LONG        = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG
	DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT       = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT
	DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE      = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE
	DEVICE_MAX_CLOCK_FREQUENCY                = C.CL_DEVICE_MAX_CLOCK_FREQUENCY
	DEVICE_ADDRESS_BITS                       = C.CL_DEVICE_ADDRESS_BITS
	DEVICE_MAX_READ_IMAGE_ARGS                = C.CL_DEVICE_MAX_READ_IMAGE_ARGS
	DEVICE_MAX_WRITE_IMAGE_ARGS               = C.CL_DEVICE_MAX_WRITE_IMAGE_ARGS
	DEVICE_MAX_MEM_ALLOC_SIZE                 = C.CL_DEVICE_MAX_MEM_ALLOC_SIZE
	DEVICE_IMAGE2D_MAX_WIDTH                  = C.CL_DEVICE_IMAGE2D_MAX_WIDTH
	DEVICE_IMAGE2D_MAX_HEIGHT                 = C.CL_DEVICE_IMAGE2D_MAX_HEIGHT
	DEVICE_IMAGE3D_MAX_WIDTH                  = C.CL_DEVICE_IMAGE3D_MAX_WIDTH
	DEVICE_IMAGE3D_MAX_HEIGHT                 = C.CL_DEVICE_IMAGE3D_MAX_HEIGHT
	DEVICE_IMAGE3D_MAX_DEPTH                  = C.CL_DEVICE_IMAGE3D_MAX_DEPTH
	DEVICE_IMAGE_SUPPORT                      = C.CL_DEVICE_IMAGE_SUPPORT
	DEVICE_MAX_PARAMETER_SIZE                 = C.CL_DEVICE_MAX_PARAMETER_SIZE
	DEVICE_MAX_SAMPLERS                       = C.CL_DEVICE_MAX_SAMPLERS
	DEVICE_MEM_BASE_ADDR_ALIGN                = C.CL_DEVICE_MEM_BASE_ADDR_ALIGN
	DEVICE_MIN_DATA_TYPE_ALIGN_SIZE           = C.CL_DEVICE_MIN_DATA_TYPE_ALIGN_SIZE
	DEVICE_SINGLE_FP_CONFIG                   = C.CL_DEVICE_SINGLE_FP_CONFIG
	DEVICE_GLOBAL_MEM_CACHE_TYPE              = C.CL_DEVICE_GLOBAL_MEM_CACHE_TYPE
	DEVICE_GLOBAL_MEM_CACHELINE_SIZE          = C.CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE
	DEVICE_GLOBAL_MEM_CACHE_SIZE              = C.CL_DEVICE_GLOBAL_MEM_CACHE_SIZE
	DEVICE_GLOBAL_MEM_SIZE                    = C.CL_DEVICE_GLOBAL_MEM_SIZE
	DEVICE_MAX_CONSTANT_BUFFER_SIZE           = C.CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE
	DEVICE_MAX_CONSTANT_ARGS                  = C.CL_DEVICE_MAX_CONSTANT_ARGS
	DEVICE_LOCAL_MEM_TYPE                     = C.CL_DEVICE_LOCAL_MEM_TYPE
	DEVICE_LOCAL_MEM_SIZE                     = C.CL_DEVICE_LOCAL_MEM_SIZE
	DEVICE_ERROR_CORRECTION_SUPPORT           = C.CL_DEVICE_ERROR_CORRECTION_SUPPORT
	DEVICE_PROFILING_TIMER_RESOLUTION         = C.CL_DEVICE_PROFILING_TIMER_RESOLUTION
	DEVICE_ENDIAN_LITTLE                      = C.CL_DEVICE_ENDIAN_LITTLE
	DEVICE_AVAILABLE                          = C.CL_DEVICE_AVAILABLE
	DEVICE_COMPILER_AVAILABLE                 = C.CL_DEVICE_COMPILER_AVAILABLE
	DEVICE_EXECUTION_CAPABILITIES             = C.CL_DEVICE_EXECUTION_CAPABILITIES
	DEVICE_QUEUE_PROPERTIES                   = C.CL_DEVICE_QUEUE_PROPERTIES
	DEVICE_NAME                               = C.CL_DEVICE_NAME
	DEVICE_VENDOR                             = C.CL_DEVICE_VENDOR
	DRIVER_VERSION                            = C.CL_DRIVER_VERSION
	DEVICE_PROFILE                            = C.CL_DEVICE_PROFILE
	DEVICE_VERSION                            = C.CL_DEVICE_VERSION
	DEVICE_EXTENSIONS                         = C.CL_DEVICE_EXTENSIONS
	DEVICE_PLATFORM                           = C.CL_DEVICE_PLATFORM
	DEVICE_DOUBLE_FP_CONFIG                   = C.CL_DEVICE_DOUBLE_FP_CONFIG
	DEVICE_PREFERRED_VECTOR_WIDTH_HALF        = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF
	DEVICE_HOST_UNIFIED_MEMORY                = C.CL_DEVICE_HOST_UNIFIED_MEMORY
	DEVICE_NATIVE_VECTOR_WIDTH_CHAR           = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR
	DEVICE_NATIVE_VECTOR_WIDTH_SHORT          = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT
	DEVICE_NATIVE_VECTOR_WIDTH_INT            = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_INT
	DEVICE_NATIVE_VECTOR_WIDTH_LONG           = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG
	DEVICE_NATIVE_VECTOR_WIDTH_FLOAT          = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT
	DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE         = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE
	DEVICE_NATIVE_VECTOR_WIDTH_HALF           = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF
	DEVICE_OPENCL_C_VERSION                   = C.CL_DEVICE_OPENCL_C_VERSION
	DEVICE_LINKER_AVAILABLE                   = C.CL_DEVICE_LINKER_AVAILABLE
	DEVICE_BUILT_IN_KERNELS                   = C.CL_DEVICE_BUILT_IN_KERNELS
	DEVICE_IMAGE_MAX_BUFFER_SIZE              = C.CL_DEVICE_IMAGE_MAX_BUFFER_SIZE
	DEVICE_IMAGE_MAX_ARRAY_SIZE               = C.CL_DEVICE_IMAGE_MAX_ARRAY_SIZE
	DEVICE_PARENT_DEVICE                      = C.CL_DEVICE_PARENT_DEVICE
	DEVICE_PARTITION_MAX_SUB_DEVICES          = C.CL_DEVICE_PARTITION_MAX_SUB_DEVICES
	DEVICE_PARTITION_PROPERTIES               = C.CL_DEVICE_PARTITION_PROPERTIES
	DEVICE_PARTITION_AFFINITY_DOMAIN          = C.CL_DEVICE_PARTITION_AFFINITY_DOMAIN
	DEVICE_PARTITION_TYPE                     = C.CL_DEVICE_PARTITION_TYPE
	DEVICE_REFERENCE_COUNT                    = C.CL_DEVICE_REFERENCE_COUNT
	DEVICE_PREFERRED_INTEROP_USER_SYNC        = C.CL_DEVICE_PREFERRED_INTEROP_USER_SYNC
	DEVICE_PRINTF_BUFFER_SIZE                 = C.CL_DEVICE_PRINTF_BUFFER_SIZE
	DEVICE_IMAGE_PITCH_ALIGNMENT              = C.CL_DEVICE_IMAGE_PITCH_ALIGNMENT
	DEVICE_IMAGE_BASE_ADDRESS_ALIGNMENT       = C.CL_DEVICE_IMAGE_BASE_ADDRESS_ALIGNMENT
	FP_DENORM                                 = C.CL_FP_DENORM
	FP_INF_NAN                                = C.CL_FP_INF_NAN
	FP_ROUND_TO_NEAREST                       = C.CL_FP_ROUND_TO_NEAREST
	FP_ROUND_TO_ZERO                          = C.CL_FP_ROUND_TO_ZERO
	FP_ROUND_TO_INF                           = C.CL_FP_ROUND_TO_INF
	FP_FMA                                    = C.CL_FP_FMA
	FP_SOFT_FLOAT                             = C.CL_FP_SOFT_FLOAT
	FP_CORRECTLY_ROUNDED_DIVIDE_SQRT          = C.CL_FP_CORRECTLY_ROUNDED_DIVIDE_SQRT
	NONE                                      = C.CL_NONE
	READ_ONLY_CACHE                           = C.CL_READ_ONLY_CACHE
	READ_WRITE_CACHE                          = C.CL_READ_WRITE_CACHE
	LOCAL                                     = C.CL_LOCAL
	GLOBAL                                    = C.CL_GLOBAL
	EXEC_KERNEL                               = C.CL_EXEC_KERNEL
	EXEC_NATIVE_KERNEL                        = C.CL_EXEC_NATIVE_KERNEL
	QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE       = C.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE
	QUEUE_PROFILING_ENABLE                    = C.CL_QUEUE_PROFILING_ENABLE
	CONTEXT_REFERENCE_COUNT                   = C.CL_CONTEXT_REFERENCE_COUNT
	CONTEXT_DEVICES                           = C.CL_CONTEXT_DEVICES
	CONTEXT_PROPERTIES                        = C.CL_CONTEXT_PROPERTIES
	CONTEXT_NUM_DEVICES                       = C.CL_CONTEXT_NUM_DEVICES
	CONTEXT_PLATFORM                          = C.CL_CONTEXT_PLATFORM
	CONTEXT_INTEROP_USER_SYNC                 = C.CL_CONTEXT_INTEROP_USER_SYNC
	DEVICE_PARTITION_EQUALLY                  = C.CL_DEVICE_PARTITION_EQUALLY
	DEVICE_PARTITION_BY_COUNTS                = C.CL_DEVICE_PARTITION_BY_COUNTS
	DEVICE_PARTITION_BY_COUNTS_LIST_END       = C.CL_DEVICE_PARTITION_BY_COUNTS_LIST_END
	DEVICE_PARTITION_BY_AFFINITY_DOMAIN       = C.CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN
	DEVICE_AFFINITY_DOMAIN_NUMA               = C.CL_DEVICE_AFFINITY_DOMAIN_NUMA
	DEVICE_AFFINITY_DOMAIN_L4_CACHE           = C.CL_DEVICE_AFFINITY_DOMAIN_L4_CACHE
	DEVICE_AFFINITY_DOMAIN_L3_CACHE           = C.CL_DEVICE_AFFINITY_DOMAIN_L3_CACHE
	DEVICE_AFFINITY_DOMAIN_L2_CACHE           = C.CL_DEVICE_AFFINITY_DOMAIN_L2_CACHE
	DEVICE_AFFINITY_DOMAIN_L1_CACHE           = C.CL_DEVICE_AFFINITY_DOMAIN_L1_CACHE
	DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE = C.CL_DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE
	QUEUE_CONTEXT                             = C.CL_QUEUE_CONTEXT
	QUEUE_DEVICE                              = C.CL_QUEUE_DEVICE
	QUEUE_REFERENCE_COUNT                     = C.CL_QUEUE_REFERENCE_COUNT
	QUEUE_PROPERTIES                          = C.CL_QUEUE_PROPERTIES
	MEM_READ_WRITE                            = C.CL_MEM_READ_WRITE
	MEM_WRITE_ONLY                            = C.CL_MEM_WRITE_ONLY
	MEM_READ_ONLY                             = C.CL_MEM_READ_ONLY
	MEM_USE_HOST_PTR                          = C.CL_MEM_USE_HOST_PTR
	MEM_ALLOC_HOST_PTR                        = C.CL_MEM_ALLOC_HOST_PTR
	MEM_COPY_HOST_PTR                         = C.CL_MEM_COPY_HOST_PTR
	MEM_HOST_WRITE_ONLY                       = C.CL_MEM_HOST_WRITE_ONLY
	MEM_HOST_READ_ONLY                        = C.CL_MEM_HOST_READ_ONLY
	MEM_HOST_NO_ACCESS                        = C.CL_MEM_HOST_NO_ACCESS
	MIGRATE_MEM_OBJECT_HOST                   = C.CL_MIGRATE_MEM_OBJECT_HOST
	MIGRATE_MEM_OBJECT_CONTENT_UNDEFINED      = C.CL_MIGRATE_MEM_OBJECT_CONTENT_UNDEFINED
	R                                         = C.CL_R
	A                                         = C.CL_A
	RG                                        = C.CL_RG
	RA                                        = C.CL_RA
	RGB                                       = C.CL_RGB
	RGBA                                      = C.CL_RGBA
	BGRA                                      = C.CL_BGRA
	ARGB                                      = C.CL_ARGB
	INTENSITY                                 = C.CL_INTENSITY
	LUMINANCE                                 = C.CL_LUMINANCE
	Rx                                        = C.CL_Rx
	RGx                                       = C.CL_RGx
	RGBx                                      = C.CL_RGBx
	DEPTH                                     = C.CL_DEPTH
	DEPTH_STENCIL                             = C.CL_DEPTH_STENCIL
	SNORM_INT8                                = C.CL_SNORM_INT8
	SNORM_INT16                               = C.CL_SNORM_INT16
	UNORM_INT8                                = C.CL_UNORM_INT8
	UNORM_INT16                               = C.CL_UNORM_INT16
	UNORM_SHORT_565                           = C.CL_UNORM_SHORT_565
	UNORM_SHORT_555                           = C.CL_UNORM_SHORT_555
	UNORM_INT_101010                          = C.CL_UNORM_INT_101010
	SIGNED_INT8                               = C.CL_SIGNED_INT8
	SIGNED_INT16                              = C.CL_SIGNED_INT16
	SIGNED_INT32                              = C.CL_SIGNED_INT32
	UNSIGNED_INT8                             = C.CL_UNSIGNED_INT8
	UNSIGNED_INT16                            = C.CL_UNSIGNED_INT16
	UNSIGNED_INT32                            = C.CL_UNSIGNED_INT32
	HALF_FLOAT                                = C.CL_HALF_FLOAT
	FLOAT                                     = C.CL_FLOAT
	UNORM_INT24                               = C.CL_UNORM_INT24
	MEM_OBJECT_BUFFER                         = C.CL_MEM_OBJECT_BUFFER
	MEM_OBJECT_IMAGE2D                        = C.CL_MEM_OBJECT_IMAGE2D
	MEM_OBJECT_IMAGE3D                        = C.CL_MEM_OBJECT_IMAGE3D
	MEM_OBJECT_IMAGE2D_ARRAY                  = C.CL_MEM_OBJECT_IMAGE2D_ARRAY
	MEM_OBJECT_IMAGE1D                        = C.CL_MEM_OBJECT_IMAGE1D
	MEM_OBJECT_IMAGE1D_ARRAY                  = C.CL_MEM_OBJECT_IMAGE1D_ARRAY
	MEM_OBJECT_IMAGE1D_BUFFER                 = C.CL_MEM_OBJECT_IMAGE1D_BUFFER
	MEM_TYPE                                  = C.CL_MEM_TYPE
	MEM_FLAGS                                 = C.CL_MEM_FLAGS
	MEM_SIZE                                  = C.CL_MEM_SIZE
	MEM_HOST_PTR                              = C.CL_MEM_HOST_PTR
	MEM_MAP_COUNT                             = C.CL_MEM_MAP_COUNT
	MEM_REFERENCE_COUNT                       = C.CL_MEM_REFERENCE_COUNT
	MEM_CONTEXT                               = C.CL_MEM_CONTEXT
	MEM_ASSOCIATED_MEMOBJECT                  = C.CL_MEM_ASSOCIATED_MEMOBJECT
	MEM_OFFSET                                = C.CL_MEM_OFFSET
	IMAGE_FORMAT                              = C.CL_IMAGE_FORMAT
	IMAGE_ELEMENT_SIZE                        = C.CL_IMAGE_ELEMENT_SIZE
	IMAGE_ROW_PITCH                           = C.CL_IMAGE_ROW_PITCH
	IMAGE_SLICE_PITCH                         = C.CL_IMAGE_SLICE_PITCH
	IMAGE_WIDTH                               = C.CL_IMAGE_WIDTH
	IMAGE_HEIGHT                              = C.CL_IMAGE_HEIGHT
	IMAGE_DEPTH                               = C.CL_IMAGE_DEPTH
	IMAGE_ARRAY_SIZE                          = C.CL_IMAGE_ARRAY_SIZE
	IMAGE_BUFFER                              = C.CL_IMAGE_BUFFER
	IMAGE_NUM_MIP_LEVELS                      = C.CL_IMAGE_NUM_MIP_LEVELS
	IMAGE_NUM_SAMPLES                         = C.CL_IMAGE_NUM_SAMPLES
	ADDRESS_NONE                              = C.CL_ADDRESS_NONE
	ADDRESS_CLAMP_TO_EDGE                     = C.CL_ADDRESS_CLAMP_TO_EDGE
	ADDRESS_CLAMP                             = C.CL_ADDRESS_CLAMP
	ADDRESS_REPEAT                            = C.CL_ADDRESS_REPEAT
	ADDRESS_MIRRORED_REPEAT                   = C.CL_ADDRESS_MIRRORED_REPEAT
	FILTER_NEAREST                            = C.CL_FILTER_NEAREST
	FILTER_LINEAR                             = C.CL_FILTER_LINEAR
	SAMPLER_REFERENCE_COUNT                   = C.CL_SAMPLER_REFERENCE_COUNT
	SAMPLER_CONTEXT                           = C.CL_SAMPLER_CONTEXT
	SAMPLER_NORMALIZED_COORDS                 = C.CL_SAMPLER_NORMALIZED_COORDS
	SAMPLER_ADDRESSING_MODE                   = C.CL_SAMPLER_ADDRESSING_MODE
	SAMPLER_FILTER_MODE                       = C.CL_SAMPLER_FILTER_MODE
	MAP_READ                                  = C.CL_MAP_READ
	MAP_WRITE                                 = C.CL_MAP_WRITE
	MAP_WRITE_INVALIDATE_REGION               = C.CL_MAP_WRITE_INVALIDATE_REGION
	PROGRAM_REFERENCE_COUNT                   = C.CL_PROGRAM_REFERENCE_COUNT
	PROGRAM_CONTEXT                           = C.CL_PROGRAM_CONTEXT
	PROGRAM_NUM_DEVICES                       = C.CL_PROGRAM_NUM_DEVICES
	PROGRAM_DEVICES                           = C.CL_PROGRAM_DEVICES
	PROGRAM_SOURCE                            = C.CL_PROGRAM_SOURCE
	PROGRAM_BINARY_SIZES                      = C.CL_PROGRAM_BINARY_SIZES
	PROGRAM_BINARIES                          = C.CL_PROGRAM_BINARIES
	PROGRAM_NUM_KERNELS                       = C.CL_PROGRAM_NUM_KERNELS
	PROGRAM_KERNEL_NAMES                      = C.CL_PROGRAM_KERNEL_NAMES
	PROGRAM_BUILD_STATUS                      = C.CL_PROGRAM_BUILD_STATUS
	PROGRAM_BUILD_OPTIONS                     = C.CL_PROGRAM_BUILD_OPTIONS
	PROGRAM_BUILD_LOG                         = C.CL_PROGRAM_BUILD_LOG
	PROGRAM_BINARY_TYPE                       = C.CL_PROGRAM_BINARY_TYPE
	PROGRAM_BINARY_TYPE_NONE                  = C.CL_PROGRAM_BINARY_TYPE_NONE
	PROGRAM_BINARY_TYPE_COMPILED_OBJECT       = C.CL_PROGRAM_BINARY_TYPE_COMPILED_OBJECT
	PROGRAM_BINARY_TYPE_LIBRARY               = C.CL_PROGRAM_BINARY_TYPE_LIBRARY
	PROGRAM_BINARY_TYPE_EXECUTABLE            = C.CL_PROGRAM_BINARY_TYPE_EXECUTABLE
	BUILD_SUCCESS                             = C.CL_BUILD_SUCCESS
	BUILD_NONE                                = C.CL_BUILD_NONE
	BUILD_ERROR                               = C.CL_BUILD_ERROR
	BUILD_IN_PROGRESS                         = C.CL_BUILD_IN_PROGRESS
	KERNEL_FUNCTION_NAME                      = C.CL_KERNEL_FUNCTION_NAME
	KERNEL_NUM_ARGS                           = C.CL_KERNEL_NUM_ARGS
	KERNEL_REFERENCE_COUNT                    = C.CL_KERNEL_REFERENCE_COUNT
	KERNEL_CONTEXT                            = C.CL_KERNEL_CONTEXT
	KERNEL_PROGRAM                            = C.CL_KERNEL_PROGRAM
	KERNEL_ATTRIBUTES                         = C.CL_KERNEL_ATTRIBUTES
	KERNEL_ARG_ADDRESS_QUALIFIER              = C.CL_KERNEL_ARG_ADDRESS_QUALIFIER
	KERNEL_ARG_ACCESS_QUALIFIER               = C.CL_KERNEL_ARG_ACCESS_QUALIFIER
	KERNEL_ARG_TYPE_NAME                      = C.CL_KERNEL_ARG_TYPE_NAME
	KERNEL_ARG_TYPE_QUALIFIER                 = C.CL_KERNEL_ARG_TYPE_QUALIFIER
	KERNEL_ARG_NAME                           = C.CL_KERNEL_ARG_NAME
	KERNEL_ARG_ADDRESS_GLOBAL                 = C.CL_KERNEL_ARG_ADDRESS_GLOBAL
	KERNEL_ARG_ADDRESS_LOCAL                  = C.CL_KERNEL_ARG_ADDRESS_LOCAL
	KERNEL_ARG_ADDRESS_CONSTANT               = C.CL_KERNEL_ARG_ADDRESS_CONSTANT
	KERNEL_ARG_ADDRESS_PRIVATE                = C.CL_KERNEL_ARG_ADDRESS_PRIVATE
	KERNEL_ARG_ACCESS_READ_ONLY               = C.CL_KERNEL_ARG_ACCESS_READ_ONLY
	KERNEL_ARG_ACCESS_WRITE_ONLY              = C.CL_KERNEL_ARG_ACCESS_WRITE_ONLY
	KERNEL_ARG_ACCESS_READ_WRITE              = C.CL_KERNEL_ARG_ACCESS_READ_WRITE
	KERNEL_ARG_ACCESS_NONE                    = C.CL_KERNEL_ARG_ACCESS_NONE
	KERNEL_ARG_TYPE_NONE                      = C.CL_KERNEL_ARG_TYPE_NONE
	KERNEL_ARG_TYPE_CONST                     = C.CL_KERNEL_ARG_TYPE_CONST
	KERNEL_ARG_TYPE_RESTRICT                  = C.CL_KERNEL_ARG_TYPE_RESTRICT
	KERNEL_ARG_TYPE_VOLATILE                  = C.CL_KERNEL_ARG_TYPE_VOLATILE
	KERNEL_WORK_GROUP_SIZE                    = C.CL_KERNEL_WORK_GROUP_SIZE
	KERNEL_COMPILE_WORK_GROUP_SIZE            = C.CL_KERNEL_COMPILE_WORK_GROUP_SIZE
	KERNEL_LOCAL_MEM_SIZE                     = C.CL_KERNEL_LOCAL_MEM_SIZE
	KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE = C.CL_KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE
	KERNEL_PRIVATE_MEM_SIZE                   = C.CL_KERNEL_PRIVATE_MEM_SIZE
	KERNEL_GLOBAL_WORK_SIZE                   = C.CL_KERNEL_GLOBAL_WORK_SIZE
	EVENT_COMMAND_QUEUE                       = C.CL_EVENT_COMMAND_QUEUE
	EVENT_COMMAND_TYPE                        = C.CL_EVENT_COMMAND_TYPE
	EVENT_REFERENCE_COUNT                     = C.CL_EVENT_REFERENCE_COUNT
	EVENT_COMMAND_EXECUTION_STATUS            = C.CL_EVENT_COMMAND_EXECUTION_STATUS
	EVENT_CONTEXT                             = C.CL_EVENT_CONTEXT
	COMMAND_NDRANGE_KERNEL                    = C.CL_COMMAND_NDRANGE_KERNEL
	COMMAND_TASK                              = C.CL_COMMAND_TASK
	COMMAND_NATIVE_KERNEL                     = C.CL_COMMAND_NATIVE_KERNEL
	COMMAND_READ_BUFFER                       = C.CL_COMMAND_READ_BUFFER
	COMMAND_WRITE_BUFFER                      = C.CL_COMMAND_WRITE_BUFFER
	COMMAND_COPY_BUFFER                       = C.CL_COMMAND_COPY_BUFFER
	COMMAND_READ_IMAGE                        = C.CL_COMMAND_READ_IMAGE
	COMMAND_WRITE_IMAGE                       = C.CL_COMMAND_WRITE_IMAGE
	COMMAND_COPY_IMAGE                        = C.CL_COMMAND_COPY_IMAGE
	COMMAND_COPY_IMAGE_TO_BUFFER              = C.CL_COMMAND_COPY_IMAGE_TO_BUFFER
	COMMAND_COPY_BUFFER_TO_IMAGE              = C.CL_COMMAND_COPY_BUFFER_TO_IMAGE
	COMMAND_MAP_BUFFER                        = C.CL_COMMAND_MAP_BUFFER
	COMMAND_MAP_IMAGE                         = C.CL_COMMAND_MAP_IMAGE
	COMMAND_UNMAP_MEM_OBJECT                  = C.CL_COMMAND_UNMAP_MEM_OBJECT
	COMMAND_MARKER                            = C.CL_COMMAND_MARKER
	COMMAND_ACQUIRE_GL_OBJECTS                = C.CL_COMMAND_ACQUIRE_GL_OBJECTS
	COMMAND_RELEASE_GL_OBJECTS                = C.CL_COMMAND_RELEASE_GL_OBJECTS
	COMMAND_READ_BUFFER_RECT                  = C.CL_COMMAND_READ_BUFFER_RECT
	COMMAND_WRITE_BUFFER_RECT                 = C.CL_COMMAND_WRITE_BUFFER_RECT
	COMMAND_COPY_BUFFER_RECT                  = C.CL_COMMAND_COPY_BUFFER_RECT
	COMMAND_USER                              = C.CL_COMMAND_USER
	COMMAND_BARRIER                           = C.CL_COMMAND_BARRIER
	COMMAND_MIGRATE_MEM_OBJECTS               = C.CL_COMMAND_MIGRATE_MEM_OBJECTS
	COMMAND_FILL_BUFFER                       = C.CL_COMMAND_FILL_BUFFER
	COMMAND_FILL_IMAGE                        = C.CL_COMMAND_FILL_IMAGE
	COMPLETE                                  = C.CL_COMPLETE
	RUNNING                                   = C.CL_RUNNING
	SUBMITTED                                 = C.CL_SUBMITTED
	QUEUED                                    = C.CL_QUEUED
	BUFFER_CREATE_TYPE_REGION                 = C.CL_BUFFER_CREATE_TYPE_REGION
	PROFILING_COMMAND_QUEUED                  = C.CL_PROFILING_COMMAND_QUEUED
	PROFILING_COMMAND_SUBMIT                  = C.CL_PROFILING_COMMAND_SUBMIT
	PROFILING_COMMAND_START                   = C.CL_PROFILING_COMMAND_START
	PROFILING_COMMAND_END                     = C.CL_PROFILING_COMMAND_END
)

/*=================================================================================================================
=======================================================Platform Api================================================
=================================================================================================================*/

type PlatformID uintptr

func GetPlatformIDs(numentries uint32, ids *PlatformID, numplatform *uint32) int32 {
	return int32(C.clGetPlatformIDs(C.cl_uint(numentries), (*C.cl_platform_id)(unsafe.Pointer(ids)), (*C.cl_uint)(numplatform)))
}

//paramName is one of [CL_PLATFORM_PROFILE,CL_PLATFORM_VERSION,CL_PLATFORM_NAME,CL_PLATFORM_VENDOR,CL_PLATFORM_EXTENSIONS]
func GetPlatformInfo(pid PlatformID, paramName uint32, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) int32 {
	return int32(C.clGetPlatformInfo(C.cl_platform_id(unsafe.Pointer(pid)), C.cl_platform_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
========================================================Device Api=================================================
=================================================================================================================*/

type DeviceId uintptr

func GetDeviceIDs(pid PlatformID, deviceType uint64, numentries uint32, devices *DeviceId, numdevices *uint32) int32 {
	return int32(C.clGetDeviceIDs(C.cl_platform_id(unsafe.Pointer(pid)), C.cl_device_type(deviceType), C.cl_uint(numentries), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.cl_uint)(numdevices)))
}

func GetDeviceInfo(did DeviceId, paramName uint32, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) int32 {
	return int32(C.clGetDeviceInfo(C.cl_device_id(unsafe.Pointer(did)), C.cl_device_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

//idk if properties is the right type
func CreateSubDevices(did DeviceId, properties *uint64, numDevices uint32, devices *DeviceId, numDevicesRet *uint32) int32 {
	return int32(C.clCreateSubDevices(C.cl_device_id(unsafe.Pointer(did)), (*C.cl_device_partition_property)(unsafe.Pointer(properties)), C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.cl_uint)(numDevicesRet)))
}
func RetainDevice(did DeviceId) int32 {
	return int32(C.clRetainDevice(C.cl_device_id(unsafe.Pointer(did))))
}
func ReleaseDevice(did DeviceId) int32 {
	return int32(C.clReleaseDevice(C.cl_device_id(unsafe.Pointer(did))))
}

/*=================================================================================================================
========================================================Context Api================================================
=================================================================================================================*/

type Context struct {
	clContext   C.cl_context
	errCallback func()
	userdata    interface{}
}

//export contextErrorCallback
func contextErrorCallback(errinfo *C.char, privateinfo unsafe.Pointer, cb C.size_t, user_data unsafe.Pointer) {
}

//same issue as CreateSubDevices
func CreateContext(properties *uint64, numDevices uint32, devices *DeviceId, errcb func(), userdata interface{}, errcode *int32) Context {
	ctx := Context{nil, errcb, userdata}
	ctx.clContext = C.clCreateContext((*C.cl_context_properties)(unsafe.Pointer(properties)),
		C.cl_uint(numDevices),
		(*C.cl_device_id)(unsafe.Pointer(devices)),
		(*[0]byte)(C.contextErrorCallback),
		unsafe.Pointer(&ctx),
		(*C.cl_int)(unsafe.Pointer(errcode)))
	return ctx
}

func CreateContextFromType(properties *uint64, deviceType uint64, errcb func(), userdata interface{}, errcode *int32) Context {
	ctx := Context{nil, errcb, userdata}
	ctx.clContext = C.clCreateContextFromType((*C.cl_context_properties)(unsafe.Pointer(properties)),
		C.cl_device_type(deviceType),
		(*[0]byte)(C.contextErrorCallback),
		unsafe.Pointer(&ctx),
		(*C.cl_int)(unsafe.Pointer(errcode)))
	return ctx
}

func RetainContext(context Context) int32 {
	return int32(C.clRetainContext(context.clContext))
}

func ReleaseContext(context Context) int32 {
	return int32(C.clReleaseContext(context.clContext))
}

func GetContextInfo(context Context, paramName uint32, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) int32 {
	return int32(C.clGetContextInfo(context.clContext, C.cl_context_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
=====================================================Command queue Api=============================================
=================================================================================================================*/

type CommandQueue C.cl_command_queue

func CreateCommandQueue(context Context, did DeviceId, properties uint64, errcode *int32) CommandQueue {
	return CommandQueue(C.clCreateCommandQueue(context.clContext, C.cl_device_id(unsafe.Pointer(did)), C.cl_command_queue_properties(properties), (*C.cl_int)(errcode)))
}

func RetainCommandQueue(cq CommandQueue) {
	C.clRetainCommandQueue(cq)
}

func ReleaseCommandQueue(cq CommandQueue) {
	C.clReleaseCommandQueue(cq)
}

func GetCommandQueueInfo(cq CommandQueue, paramName uint32, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) int32 {
	return int32(C.clGetCommandQueueInfo(cq, C.cl_command_queue_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
=====================================================Memory Object Api=============================================
=================================================================================================================*/

type ImageFormat struct {
	imageChannelOrder, imageChannelDataType uint32
}

func (imf ImageFormat) toC() *C.cl_image_format {
	return &C.cl_image_format{
		image_channel_order:     C.cl_channel_order(imf.imageChannelOrder),
		image_channel_data_type: C.cl_channel_type(imf.imageChannelDataType)}
}

type ImageDesc struct {
	imageType, numMipLevels, numSamples                                                 uint32
	imageWidth, imageHeight, imageDepth, imageArraySize, imageRowPitch, imageSlicePitch uint64
	buffer                                                                              Mem
}

func (imde ImageDesc) toC() *C.cl_image_desc {
	return &C.cl_image_desc{
		image_type:        C.cl_mem_object_type(imde.imageType),
		image_width:       C.size_t(imde.imageWidth),
		image_height:      C.size_t(imde.imageHeight),
		image_depth:       C.size_t(imde.imageDepth),
		image_array_size:  C.size_t(imde.imageArraySize),
		image_row_pitch:   C.size_t(imde.imageRowPitch),
		image_slice_pitch: C.size_t(imde.imageSlicePitch),
		num_mip_levels:    C.cl_uint(imde.numMipLevels),
		num_samples:       C.cl_uint(imde.numSamples),
		buffer:            C.cl_mem(imde.buffer)}
}

type Mem C.cl_mem

func CreateBuffer(context Context, flags, paramValueSize uint64, hostPtr unsafe.Pointer, errcode *int32) Mem {
	return Mem(C.clCreateBuffer(context.clContext, C.cl_mem_flags(flags), C.size_t(paramValueSize), hostPtr, (*C.cl_int)(unsafe.Pointer(errcode))))
}

func CreateSubBuffer(mem Mem, flags uint64, bufferCreateType uint32, bufferCreateInfo unsafe.Pointer, errcode *int32) Mem {
	return Mem(C.clCreateSubBuffer(mem, C.cl_mem_flags(flags), C.cl_buffer_create_type(bufferCreateType), bufferCreateInfo, (*C.cl_int)(unsafe.Pointer(errcode))))
}

func CreateImage(context Context, flags uint64, imageFormat ImageFormat, imageDesc ImageDesc, hostPtr unsafe.Pointer, errcode *int32) Mem {
	return Mem(C.clCreateImage(context.clContext, C.cl_mem_flags(flags), imageFormat.toC(), imageDesc.toC(), hostPtr, (*C.cl_int)(unsafe.Pointer(errcode))))
}

func RetainMemObject(mem Mem) int32 {
	return int32(C.clRetainMemObject(mem))
}

func ReleaseMemObject(mem Mem) int32 {
	return int32(C.clReleaseMemObject(mem))
}

func GetSupportedImageFormats(context Context, flags uint64, memObjectType uint32, numEntries uint32, imageformats *ImageFormat, numImageFormat *uint32) int32 {
	return int32(C.clGetSupportedImageFormats(context.clContext,
		C.cl_mem_flags(flags),
		C.cl_mem_object_type(memObjectType),
		C.cl_uint(numEntries),
		(*C.cl_image_format)(unsafe.Pointer(imageformats)),
		(*C.cl_uint)(numImageFormat)))
}

func GetMemObjectInfo(mem Mem, paramName uint32, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) int32 {
	return int32(C.clGetMemObjectInfo(mem, C.cl_mem_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

func GetImageInfo(mem Mem, paramName uint32, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) int32 {
	return int32(C.clGetImageInfo(mem, C.cl_image_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

type memObjectCallbackHolder struct {
	cbfunc   func(Mem, interface{})
	userdata interface{}
}

var mochHolder map[*memObjectCallbackHolder]struct{}

//export memObjectDestroyCallback
func memObjectDestroyCallback(mem C.cl_mem, userData unsafe.Pointer) {
	moch := (*memObjectCallbackHolder)(userData)
	moch.cbfunc(Mem(mem), moch.userdata)
	delete(mochHolder, moch)
}

func SetMemObjectDestructorCallback(mem Mem, destroyCb func(Mem, interface{}), userData interface{}) int32 {
	cbh := memObjectCallbackHolder{destroyCb, userData}
	mochHolder[&cbh] = struct{}{}
	return int32(C.clSetMemObjectDestructorCallback(mem, (*[0]byte)(C.memObjectDestroyCallback), unsafe.Pointer(&cbh)))
}

/*=================================================================================================================
========================================================Sampler Api================================================
=================================================================================================================*/

type Sampler C.cl_sampler

func CreateSampler(context Context, normalizedCoords uint32, addressingMode, filterMode uint32, errcode *int32) Sampler {
	return Sampler(C.clCreateSampler(context.clContext, C.cl_bool(normalizedCoords), C.cl_addressing_mode(addressingMode), C.cl_filter_mode(filterMode), (*C.cl_int)(unsafe.Pointer(errcode))))
}

func RetainSampler(sampler Sampler) int32 {
	return int32(C.clRetainSampler(sampler))
}

func ReleaseSampler(sampler Sampler) int32 {
	return int32(C.clReleaseSampler(sampler))
}

func GetSamplerInfo(sampler Sampler, paramName uint32, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) int32 {
	return int32(C.clGetSamplerInfo(sampler, C.cl_sampler_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
=====================================================Program Object Api============================================
=================================================================================================================*/

type Program C.cl_program

func CreateProgramWithSource(context Context, count uint32, src **uint8, lengths *uint64, errcode *int32) Program {
	return Program(C.clCreateProgramWithSource(context.clContext, C.cl_uint(count), (**C.char)(unsafe.Pointer(src)), (*C.size_t)(unsafe.Pointer(lengths)), (*C.cl_int)(unsafe.Pointer(errcode))))
}

func CreateProgramWithBinary(context Context, numDevices uint32, devices *DeviceId, lengths *uint64, binaries **uint8, binaryStatus *int32, errcode *int32) Program {
	return Program(C.clCreateProgramWithBinary(context.clContext,
		C.cl_uint(numDevices),
		(*C.cl_device_id)(unsafe.Pointer(devices)),
		(*C.size_t)(unsafe.Pointer(lengths)),
		(**C.uchar)(unsafe.Pointer(binaries)),
		(*C.cl_int)(unsafe.Pointer(binaryStatus)),
		(*C.cl_int)(unsafe.Pointer(errcode))))
}

func CreateProgramWithBuiltInKernels(context Context, numDevices uint32, devices *DeviceId, kernelNames *uint8, errcode *int32) Program {
	return Program(C.clCreateProgramWithBuiltInKernels(context.clContext, C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.char)(unsafe.Pointer(kernelNames)), (*C.cl_int)(unsafe.Pointer(errcode))))
}

func RetainProgram(prog Program) int32 {
	return int32(C.clRetainProgram(prog))
}
func ReleaseProgram(prog Program) int32 {
	return int32(C.clReleaseProgram(prog))
}

type programObjectBuildCompleteHolder struct {
	cbfunc   func(Program, interface{})
	userdata interface{}
}

var pobchHolder map[*programObjectBuildCompleteHolder]struct{}

//export programObjectBuildCompleteCallback
func programObjectBuildCompleteCallback(prog C.cl_program, userdata unsafe.Pointer) {
	pobch := (*programObjectBuildCompleteHolder)(userdata)
	pobch.cbfunc(Program(prog), pobch.userdata)
	delete(pobchHolder, pobch)
}

func BuildProgram(prog Program, numDevices uint32, devices *DeviceId, options *uint8, buildcomplete func(Program, interface{}), userdata interface{}) int32 {
	pobch := programObjectBuildCompleteHolder{buildcomplete, userdata}
	pobchHolder[&pobch] = struct{}{}
	return int32(C.clBuildProgram(prog, C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.char)(unsafe.Pointer(options)), (*[0]byte)(C.programObjectBuildCompleteCallback), unsafe.Pointer(&pobch)))
}

var code = `
extern CL_API_ENTRY cl_int CL_API_CALL
clCompileProgram(cl_program           /* program */,
                 cl_uint              /* num_devices */,
                 const cl_device_id * /* device_list */,
                 const char *         /* options */, 
                 cl_uint              /* num_input_headers */,
                 const cl_program *   /* input_headers */,
                 const char **        /* header_include_names */,
                 void (CL_CALLBACK *  /* pfn_notify */)(cl_program /* program */, void * /* user_data */),
                 void *               /* user_data */) CL_API_SUFFIX__VERSION_1_2;

extern CL_API_ENTRY cl_program CL_API_CALL
clLinkProgram(cl_context           /* context */,
              cl_uint              /* num_devices */,
              const cl_device_id * /* device_list */,
              const char *         /* options */, 
              cl_uint              /* num_input_programs */,
              const cl_program *   /* input_programs */,
              void (CL_CALLBACK *  /* pfn_notify */)(cl_program /* program */, void * /* user_data */),
              void *               /* user_data */,
              cl_int *             /* errcode_ret */ ) CL_API_SUFFIX__VERSION_1_2;


extern CL_API_ENTRY cl_int CL_API_CALL
clUnloadPlatformCompiler(cl_platform_id /* platform */) CL_API_SUFFIX__VERSION_1_2;

extern CL_API_ENTRY cl_int CL_API_CALL
clGetProgramInfo(cl_program         /* program */,
                 cl_program_info    /* param_name */,
                 size_t             /* param_value_size */,
                 void *             /* param_value */,
                 size_t *           /* param_value_size_ret */) CL_API_SUFFIX__VERSION_1_0;

extern CL_API_ENTRY cl_int CL_API_CALL
clGetProgramBuildInfo(cl_program            /* program */,
                      cl_device_id          /* device */,
                      cl_program_build_info /* param_name */,
                      size_t                /* param_value_size */,
                      void *                /* param_value */,
                      size_t *              /* param_value_size_ret */) CL_API_SUFFIX__VERSION_1_0;                 
`
