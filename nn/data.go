package nn

import (
	"fmt"
	"log"

	"github.com/sugarme/gotch"
	ts "github.com/sugarme/gotch/tensor"
)

// Iter2 is an iterator over a pair of tensors which have the same first dimension
// size.
// The typical use case is to iterate over batches. Each batch is a pair
// containing a (potentially random) slice of each of the two input
// tensors.
type Iter2 struct {
	xs                   ts.Tensor
	ys                   ts.Tensor
	batchIndex           int64
	batchSize            int64
	totalSize            int64
	device               gotch.Device
	returnSmallLastBatch bool
}

// NewIter2 returns a new iterator.
//
// This takes as input two tensors which first dimension must match. The
// returned iterator can be used to range over mini-batches of data of
// specified size.
// An error is returned if `xs` and `ys` have different first dimension
// sizes.
//
// # Arguments
//
// * `xs` - the features to be used by the model.
// * `ys` - the targets that the model attempts to predict.
// * `batch_size` - the size of batches to be returned.
func NewIter2(xs, ys ts.Tensor, batchSize int64) (retVal Iter2, err error) {

	totalSize := xs.MustSize()[0]
	if ys.MustSize()[0] != totalSize {
		err = fmt.Errorf("Different dimension for the two inputs: %v - %v", xs.MustSize(), ys.MustSize())
		return retVal, err
	}

	retVal = Iter2{
		xs:                   xs.MustShallowClone(),
		ys:                   ys.MustShallowClone(),
		batchIndex:           0,
		batchSize:            batchSize,
		totalSize:            totalSize,
		returnSmallLastBatch: false,
	}

	return retVal, nil
}

// MustNewIter2 returns a new iterator.
//
// This takes as input two tensors which first dimension must match. The
// returned iterator can be used to range over mini-batches of data of
// specified size.
// Panics if `xs` and `ys` have different first dimension sizes.
//
// # Arguments
//
// * `xs` - the features to be used by the model.
// * `ys` - the targets that the model attempts to predict.
// * `batch_size` - the size of batches to be returned.
func MustNewIter2(xs, ys ts.Tensor, batchSize int64) (retVal Iter2) {
	retVal, err := NewIter2(xs, ys, batchSize)

	if err != nil {
		log.Fatal(err)
	}

	return retVal
}

// Shuffle shuffles the dataset.
//
// The iterator would still run over the whole dataset but the order in
// which elements are grouped in mini-batches is randomized.
func (it Iter2) Shuffle() (retVal Iter2) {
	index := ts.MustRandperm(it.totalSize, gotch.Int64, gotch.CPU)

	it.xs = it.xs.MustIndexSelect(0, index)
	it.ys = it.ys.MustIndexSelect(0, index)
	return it
}

// ToDevice transfers the mini-batches to a specified device.
func (it Iter2) ToDevice(device gotch.Device) (retVal Iter2) {
	it.device = device
	return it
}

// ReturnSmallLastBatch when set, returns the last batch even if smaller than the batch size.
func (it Iter2) ReturnSmallLastBatch() (retVal Iter2) {
	it.returnSmallLastBatch = true
	return it
}

type Iter2Item struct {
	Images ts.Tensor // should change to Data
	Labels ts.Tensor
}

// Next implements iterator for Iter2
func (it *Iter2) Next() (item Iter2Item, ok bool) {
	start := it.batchIndex * it.batchSize
	size := it.batchSize
	if it.totalSize-start < it.batchSize {
		size = it.totalSize - start
	}

	if (size <= 0) || (!it.returnSmallLastBatch && size < it.batchSize) {
		// err = fmt.Errorf("Last small batch error")
		return item, false
	} else {
		it.batchIndex += 1

		// Indexing
		narrowIndex := ts.NewNarrow(start, start+size)

		return Iter2Item{
			Images: it.xs.Idx(narrowIndex),
			Labels: it.ys.Idx(narrowIndex),
		}, true
	}
}