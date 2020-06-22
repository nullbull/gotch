// NOTE: this is a sample for OCaml generated code for `c-generated.go`
package libtch

//#include "stdbool.h"
//#include "torch_api.h"
import "C"

import (
	"unsafe"
)

// void atg_eq1(tensor *, tensor self, tensor other);
func AtgEq1(ptr *Ctensor, self Ctensor, other Ctensor) {
	C.atg_eq1(ptr, self, other)
}

// void atg_matmul(tensor *, tensor self, tensor other);
func AtgMatmul(ptr *Ctensor, self Ctensor, other Ctensor) {
	C.atg_matmul(ptr, self, other)
}

// void atg_to(tensor *, tensor self, int device);
func AtgTo(ptr *Ctensor, self Ctensor, device int) {
	cdevice := *(*C.int)(unsafe.Pointer(&device))
	C.atg_to(ptr, self, cdevice)
}

// int at_device(tensor);
func AtDevice(ts Ctensor) int {
	cint := C.at_device(ts)
	return *(*int)(unsafe.Pointer(&cint))
}

// void atg_grad(tensor *, tensor self);
func AtgGrad(ptr *Ctensor, self Ctensor) {
	C.atg_grad(ptr, self)
}

// void atg_detach_(tensor *, tensor self);
func AtgDetach_(ptr *Ctensor, self Ctensor) {
	C.atg_detach_(ptr, self)
}

// void atg_zero_(tensor *, tensor self);
func AtgZero_(ptr *Ctensor, self Ctensor) {
	C.atg_zero_(ptr, self)
}

// void atg_set_requires_grad(tensor *, tensor self, int r);
func AtgSetRequiresGrad(ptr *Ctensor, self Ctensor, r int) {
	cr := *(*C.int)(unsafe.Pointer(&r))
	C.atg_set_requires_grad(ptr, self, cr)
}

// void atg_mul(tensor *, tensor self, tensor other);
func AtgMul(ptr *Ctensor, self Ctensor, other Ctensor) {
	C.atg_mul(ptr, self, other)
}

// void atg_mul_(tensor *, tensor self, tensor other);
func AtgMul_(ptr *Ctensor, self Ctensor, other Ctensor) {
	C.atg_mul_(ptr, self, other)
}

// void atg_mul1(tensor *, tensor self, scalar other);
func AtgMul1(ptr *Ctensor, self Ctensor, other Cscalar) {
	C.atg_mul1(ptr, self, other)
}

// void atg_add(tensor *, tensor self, tensor other);
func AtgAdd(ptr *Ctensor, self Ctensor, other Ctensor) {
	C.atg_add(ptr, self, other)
}

// void atg_add_(tensor *, tensor self, tensor other);
func AtgAdd_(ptr *Ctensor, self Ctensor, other Ctensor) {
	C.atg_add_(ptr, self, other)
}

// id atg_add1(tensor *, tensor self, scalar other);
func AtgAdd1(ptr *Ctensor, self Ctensor, other Cscalar) {
	C.atg_add1(ptr, self, other)
}

// void atg_totype(tensor *, tensor self, int scalar_type);
func AtgTotype(ptr *Ctensor, self Ctensor, scalar_type int32) {
	cscalar_type := *(*C.int)(unsafe.Pointer(&scalar_type))
	C.atg_totype(ptr, self, cscalar_type)
}

// void atg_unsqueeze(tensor *, tensor self, int64_t dim);
func AtgUnsqueeze(ptr *Ctensor, self Ctensor, dim int64) {
	cdim := *(*C.int64_t)(unsafe.Pointer(&dim))
	C.atg_unsqueeze(ptr, self, cdim)
}

// void atg_select(tensor *, tensor self, int64_t dim, int64_t index);
func AtgSelect(ptr *Ctensor, self Ctensor, dim int64, index int64) {
	cdim := *(*C.int64_t)(unsafe.Pointer(&dim))
	cindex := *(*C.int64_t)(unsafe.Pointer(&index))
	C.atg_select(ptr, self, cdim, cindex)
}

// void atg_narrow(tensor *, tensor self, int64_t dim, int64_t start, int64_t length);
func AtgNarrow(ptr *Ctensor, self Ctensor, dim int64, start int64, length int64) {
	cdim := *(*C.int64_t)(unsafe.Pointer(&dim))
	cstart := *(*C.int64_t)(unsafe.Pointer(&start))
	clength := *(*C.int64_t)(unsafe.Pointer(&length))
	C.atg_narrow(ptr, self, cdim, cstart, clength)
}

// void atg_index_select(tensor *, tensor self, int64_t dim, tensor index);
func AtgIndexSelect(ptr *Ctensor, self Ctensor, dim int64, index Ctensor) {
	cdim := *(*C.int64_t)(unsafe.Pointer(&dim))
	C.atg_index_select(ptr, self, cdim, index)
}

// void atg_zeros(tensor *, int64_t *size_data, int size_len, int options_kind, int options_device);
func AtgZeros(ptr *Ctensor, sizeData []int64, sizeLen int, optionsKind, optionsDevice int32) {
	// just get pointer of the first element of the shape(sizeData)
	csizeDataPtr := (*C.int64_t)(unsafe.Pointer(&sizeData[0]))
	csizeLen := *(*C.int)(unsafe.Pointer(&sizeLen))
	coptionsKind := *(*C.int)(unsafe.Pointer(&optionsKind))
	coptionsDevice := *(*C.int)(unsafe.Pointer(&optionsDevice))

	C.atg_zeros(ptr, csizeDataPtr, csizeLen, coptionsKind, coptionsDevice)
}

// void atg_ones(tensor *, int64_t *size_data, int size_len, int options_kind, int options_device);
func AtgOnes(ptr *Ctensor, sizeData []int64, sizeLen int, optionsKind, optionsDevice int32) {
	// just get pointer of the first element of the shape(sizeData)
	csizeDataPtr := (*C.int64_t)(unsafe.Pointer(&sizeData[0]))
	csizeLen := *(*C.int)(unsafe.Pointer(&sizeLen))
	coptionsKind := *(*C.int)(unsafe.Pointer(&optionsKind))
	coptionsDevice := *(*C.int)(unsafe.Pointer(&optionsDevice))

	C.atg_ones(ptr, csizeDataPtr, csizeLen, coptionsKind, coptionsDevice)
}

// void atg_uniform_(tensor *, tensor self, double from, double to);
func AtgUniform_(ptr *Ctensor, self Ctensor, from float64, to float64) {
	cfrom := *(*C.double)(unsafe.Pointer(&from))
	cto := *(*C.double)(unsafe.Pointer(&to))

	C.atg_uniform_(ptr, self, cfrom, cto)
}

// void atg_zeros_like(tensor *, tensor self);
func AtgZerosLike(ptr *Ctensor, self Ctensor) {
	C.atg_zeros_like(ptr, self)
}

// void atg_fill_(tensor *, tensor self, scalar value);
func AtgFill_(ptr *Ctensor, self Ctensor, value Cscalar) {
	C.atg_fill_(ptr, self, value)
}

// void atg_randn_like(tensor *, tensor self);
func AtgRandnLike(ptr *Ctensor, self Ctensor) {
	C.atg_rand_like(ptr, self)
}

// void atg_log_softmax(tensor *, tensor self, int64_t dim, int dtype);
func AtgLogSoftmax(ptr *Ctensor, self Ctensor, dim int64, dtype int32) {
	cdim := *(*C.int64_t)(unsafe.Pointer(&dim))
	cdtype := *(*C.int)(unsafe.Pointer(&dtype))

	C.atg_log_softmax(ptr, self, cdim, cdtype)
}

// void atg_nll_loss(tensor *, tensor self, tensor target, tensor weight, int64_t reduction, int64_t ignore_index);
func AtgNllLoss(ptr *Ctensor, self Ctensor, target Ctensor, weight Ctensor, reduction int64, ignoreIndex int64) {
	creduction := *(*C.int64_t)(unsafe.Pointer(&reduction))
	cignoreIndex := *(*C.int64_t)(unsafe.Pointer(&ignoreIndex))

	C.atg_nll_loss(ptr, self, target, weight, creduction, cignoreIndex)
}

// void atg_argmax(tensor *, tensor self, int64_t dim, int keepdim);
func AtgArgmax(ptr *Ctensor, self Ctensor, dim int64, keepDim int) {
	cdim := *(*C.int64_t)(unsafe.Pointer(&dim))
	ckeepDim := *(*C.int)(unsafe.Pointer(&keepDim))

	C.atg_argmax(ptr, self, cdim, ckeepDim)
}

// void atg_mean(tensor *, tensor self, int dtype);
func AtgMean(ptr *Ctensor, self Ctensor, dtype int32) {
	cdtype := *(*C.int)(unsafe.Pointer(&dtype))

	C.atg_mean(ptr, self, cdtype)
}

// void atg_permute(tensor *, tensor self, int64_t *dims_data, int dims_len);
func AtgPermute(ptr *Ctensor, self Ctensor, dims []int64, dimLen int) {
	// just get pointer of the first element of the shape
	cdimsPtr := (*C.int64_t)(unsafe.Pointer(&dims[0]))
	cdimLen := *(*C.int)(unsafe.Pointer(&dimLen))

	C.atg_permute(ptr, self, cdimsPtr, cdimLen)
}

// void atg_squeeze1(tensor *, tensor self, int64_t dim);
func AtgSqueeze1(ptr *Ctensor, self Ctensor, dim int64) {
	cdim := *(*C.int64_t)(unsafe.Pointer(&dim))

	C.atg_squeeze1(ptr, self, cdim)
}

// void atg_squeeze_(tensor *, tensor self);
func AtgSqueeze_(ptr *Ctensor, self Ctensor) {
	C.atg_squeeze_(ptr, self)
}

// void atg_stack(tensor *, tensor *tensors_data, int tensors_len, int64_t dim);
func AtgStack(ptr *Ctensor, tensorsData []Ctensor, tensorsLen int, dim int64) {
	tensorsDataPtr := (*Ctensor)(unsafe.Pointer(&tensorsData[0]))
	ctensorsLen := *(*C.int)(unsafe.Pointer(&tensorsLen))
	cdim := *(*C.int64_t)(unsafe.Pointer(&dim))

	C.atg_stack(ptr, tensorsDataPtr, ctensorsLen, cdim)
}

// void atg_mm(tensor *, tensor self, tensor mat2);
func AtgMm(ptr *Ctensor, self Ctensor, mat2 Ctensor) {
	C.atg_mm(ptr, self, mat2)
}

// void atg_view(tensor *, tensor self, int64_t *size_data, int size_len);
func AtgView(ptr *Ctensor, self Ctensor, sizeData []int64, sizeLen int) {
	sizeDataPtr := (*C.int64_t)(unsafe.Pointer(&sizeData[0]))
	csizeLen := *(*C.int)(unsafe.Pointer(&sizeLen))

	C.atg_view(ptr, self, sizeDataPtr, csizeLen)
}

// void atg_div1(tensor *, tensor self, scalar other);
func AtgDiv1(ptr *Ctensor, self Ctensor, other Cscalar) {
	C.atg_div1(ptr, self, other)
}

// void atg_randperm(tensor *, int64_t n, int options_kind, int options_device);
func AtgRandperm(ptr *Ctensor, n int64, optionKind int32, optionDevice int32) {
	cn := *(*C.int64_t)(unsafe.Pointer(&n))
	coptionKind := *(*C.int)(unsafe.Pointer(&optionKind))
	coptionDevice := *(*C.int)(unsafe.Pointer(&optionDevice))

	C.atg_randperm(ptr, cn, coptionKind, coptionDevice)
}

// void atg_clamp_(tensor *, tensor self, scalar min, scalar max);
func AtgClamp_(ptr *Ctensor, self Ctensor, min Cscalar, max Cscalar) {
	C.atg_clamp_(ptr, self, min, max)
}

// void atg_relu(tensor *, tensor self);
func AtgRelu(ptr *Ctensor, self Ctensor) {
	C.atg_relu(ptr, self)
}

// void atg_relu_(tensor *, tensor self);
func AtgRelu_(ptr *Ctensor, self Ctensor) {
	C.atg_relu_(ptr, self)
}

// void atg_t(tensor *, tensor self);
func AtgT(ptr *Ctensor, self Ctensor) {
	C.atg_t(ptr, self)
}

// void atg_t_(tensor *, tensor self);
func AtgT_(ptr *Ctensor, self Ctensor) {
	C.atg_t_(ptr, self)
}

// void atg_mse_loss(tensor *, tensor self, tensor target, int64_t reduction);
func AtgMseLoss(ptr *Ctensor, self Ctensor, target Ctensor, reduction int) {
	creduction := *(*C.int64_t)(unsafe.Pointer(&reduction))

	C.atg_mse_loss(ptr, self, target, creduction)
}

// void atg_exp(tensor *, tensor self);
func AtgExp(ptr *Ctensor, self Ctensor) {
	C.atg_exp(ptr, self)
}

// void atg_exp_(tensor *, tensor self);
func AtgExp_(ptr *Ctensor, self Ctensor) {
	C.atg_exp_(ptr, self)
}

// void atg_pow(tensor *, tensor self, scalar exponent);
func AtgPow(ptr *Ctensor, self Ctensor, exponent Cscalar) {
	C.atg_pow(ptr, self, exponent)
}

// void atg_sum(tensor *, tensor self, int dtype);
func AtgSum(ptr *Ctensor, self Ctensor, dtype int32) {
	cdtype := *(*C.int)(unsafe.Pointer(&dtype))

	C.atg_sum(ptr, self, cdtype)
}

// void atg_sub(tensor *, tensor self, tensor other);
func AtgSub(ptr *Ctensor, self Ctensor, other Ctensor) {
	C.atg_sub(ptr, self, other)
}

// void atg_sub1(tensor *, tensor self, scalar other);
func AtgSub1(ptr *Ctensor, self Ctensor, other Cscalar) {
	C.atg_sub1(ptr, self, other)
}

// void atg_sub_(tensor *, tensor self, tensor other);
func AtgSub_(ptr *Ctensor, self Ctensor, other Ctensor) {
	C.atg_sub_(ptr, self, other)
}

// void atg_conv1d(tensor *, tensor input, tensor weight, tensor bias, int64_t *stride_data, int stride_len, int64_t *padding_data, int padding_len, int64_t *dilation_data, int dilation_len, int64_t groups);
func AtgConv1d(ptr *Ctensor, input Ctensor, weight Ctensor, bias Ctensor, strideData []int64, strideLen int, paddingData []int64, paddingLen int, dilationData []int64, dilationLen int, groups int64) {
	cstrideDataPtr := (*C.int64_t)(unsafe.Pointer(&strideData[0]))
	cstrideLen := *(*C.int)(unsafe.Pointer(&strideLen))
	cpaddingDataPtr := (*C.int64_t)(unsafe.Pointer(&paddingData[0]))
	cpaddingLen := *(*C.int)(unsafe.Pointer(&paddingLen))
	cdilationDataPtr := (*C.int64_t)(unsafe.Pointer(&dilationData[0]))
	cdilationLen := *(*C.int)(unsafe.Pointer(&dilationLen))
	cgroups := *(*C.int64_t)(unsafe.Pointer(&groups))

	C.atg_conv1d(ptr, input, weight, bias, cstrideDataPtr, cstrideLen, cpaddingDataPtr, cpaddingLen, cdilationDataPtr, cdilationLen, cgroups)
}

// void atg_conv2d(tensor *, tensor input, tensor weight, tensor bias, int64_t *stride_data, int stride_len, int64_t *padding_data, int padding_len, int64_t *dilation_data, int dilation_len, int64_t groups);
func AtgConv2d(ptr *Ctensor, input Ctensor, weight Ctensor, bias Ctensor, strideData []int64, strideLen int, paddingData []int64, paddingLen int, dilationData []int64, dilationLen int, groups int64) {
	cstrideDataPtr := (*C.int64_t)(unsafe.Pointer(&strideData[0]))
	cstrideLen := *(*C.int)(unsafe.Pointer(&strideLen))
	cpaddingDataPtr := (*C.int64_t)(unsafe.Pointer(&paddingData[0]))
	cpaddingLen := *(*C.int)(unsafe.Pointer(&paddingLen))
	cdilationDataPtr := (*C.int64_t)(unsafe.Pointer(&dilationData[0]))
	cdilationLen := *(*C.int)(unsafe.Pointer(&dilationLen))
	cgroups := *(*C.int64_t)(unsafe.Pointer(&groups))

	C.atg_conv2d(ptr, input, weight, bias, cstrideDataPtr, cstrideLen, cpaddingDataPtr, cpaddingLen, cdilationDataPtr, cdilationLen, cgroups)
}

// void atg_conv3d(tensor *, tensor input, tensor weight, tensor bias, int64_t *stride_data, int stride_len, int64_t *padding_data, int padding_len, int64_t *dilation_data, int dilation_len, int64_t groups);
func AtgConv3d(ptr *Ctensor, input Ctensor, weight Ctensor, bias Ctensor, strideData []int64, strideLen int, paddingData []int64, paddingLen int, dilationData []int64, dilationLen int, groups int64) {
	cstrideDataPtr := (*C.int64_t)(unsafe.Pointer(&strideData[0]))
	cstrideLen := *(*C.int)(unsafe.Pointer(&strideLen))
	cpaddingDataPtr := (*C.int64_t)(unsafe.Pointer(&paddingData[0]))
	cpaddingLen := *(*C.int)(unsafe.Pointer(&paddingLen))
	cdilationDataPtr := (*C.int64_t)(unsafe.Pointer(&dilationData[0]))
	cdilationLen := *(*C.int)(unsafe.Pointer(&dilationLen))
	cgroups := *(*C.int64_t)(unsafe.Pointer(&groups))

	C.atg_conv3d(ptr, input, weight, bias, cstrideDataPtr, cstrideLen, cpaddingDataPtr, cpaddingLen, cdilationDataPtr, cdilationLen, cgroups)
}

// void atg_max_pool2d(tensor *, tensor self, int64_t *kernel_size_data, int kernel_size_len, int64_t *stride_data, int stride_len, int64_t *padding_data, int padding_len, int64_t *dilation_data, int dilation_len, int ceil_mode);
func AtgMaxPool2d(ptr *Ctensor, self Ctensor, kernelSizeData []int64, kernelSizeLen int, strideData []int64, strideLen int, paddingData []int64, paddingLen int, dilationData []int64, dilationLen int, ceilMode int) {

	ckernelSizeDataPtr := (*C.int64_t)(unsafe.Pointer(&kernelSizeData[0]))
	ckernelSizeLen := *(*C.int)(unsafe.Pointer(&kernelSizeLen))
	cstrideDataPtr := (*C.int64_t)(unsafe.Pointer(&strideData[0]))
	cstrideLen := *(*C.int)(unsafe.Pointer(&strideLen))
	cpaddingDataPtr := (*C.int64_t)(unsafe.Pointer(&paddingData[0]))
	cpaddingLen := *(*C.int)(unsafe.Pointer(&paddingLen))
	cdilationDataPtr := (*C.int64_t)(unsafe.Pointer(&dilationData[0]))
	cdilationLen := *(*C.int)(unsafe.Pointer(&dilationLen))
	cceilMode := *(*C.int)(unsafe.Pointer(&ceilMode))

	C.atg_max_pool2d(ptr, self, ckernelSizeDataPtr, ckernelSizeLen, cstrideDataPtr, cstrideLen, cpaddingDataPtr, cpaddingLen, cdilationDataPtr, cdilationLen, cceilMode)
}

// void atg_dropout(tensor *, tensor input, double p, int train);
func AtgDropout(ptr *Ctensor, input Ctensor, p float64, train int) {
	cp := *(*C.double)(unsafe.Pointer(&p))
	ctrain := *(*C.int)(unsafe.Pointer(&train))

	C.atg_dropout(ptr, input, cp, ctrain)
}

// void atg_dropout_(tensor *, tensor self, double p, int train);
func AtgDropout_(ptr *Ctensor, self Ctensor, p float64, train int) {
	cp := *(*C.double)(unsafe.Pointer(&p))
	ctrain := *(*C.int)(unsafe.Pointer(&train))

	C.atg_dropout_(ptr, self, cp, ctrain)
}
