package radar

import (
	"time"

	"github.com/fhs/go-netcdf/netcdf"
)

// GetDimensionLen ...
func GetDimensionLen(ds *netcdf.Dataset, name string) (int64, error) {
	dim, err := ds.Dim(name)
	if err != nil {
		return 0, err
	}

	dimlen, err := dim.Len()
	if err != nil {
		return 0, err
	}

	return int64(dimlen), nil
}

func ReadDoubleVar(ds *netcdf.Dataset, name string) ([]float64, error) {
	var v netcdf.Var
	var vlen uint64
	var err error

	if v, err = ds.Var(name); err != nil {
		return nil, err
	}

	if vlen, err = v.Len(); err != nil {
		return nil, err
	}

	res := make([]float64, vlen)
	if err = v.ReadFloat64s(res); err != nil {
		return nil, err
	}

	return res, nil
}

func ReadFloatVar(ds *netcdf.Dataset, name string) ([]float32, error) {
	var v netcdf.Var
	var vlen uint64
	var err error

	if v, err = ds.Var(name); err != nil {
		return nil, err
	}

	if vlen, err = v.Len(); err != nil {
		return nil, err
	}

	res := make([]float32, vlen)
	if err = v.ReadFloat32s(res); err != nil {
		return nil, err
	}

	return res, nil
}

// ReadTimeVar ...
func ReadTimeVar(ds *netcdf.Dataset, name string) ([]time.Time, error) {

	var varDs netcdf.Var
	var err error
	var varlen uint64

	if varDs, err = ds.Var(name); err != nil {
		return nil, err
	}

	if varlen, err = varDs.Len(); err != nil {
		return nil, err
	}

	varval := make([]int32, varlen)

	if err = varDs.ReadInt32s(varval); err != nil {
		return nil, err
	}

	res := make([]time.Time, varlen)
	for i, inst := range varval {
		res[i] = time.Unix(int64(inst), 0).UTC()
	}
	return res, nil
}

// MosaicData ...
type MosaicData struct {
	Lat    []float32
	Lon    []float32
	Width  int64
	Height int64
	//Instants                       []time.Time
	Cappi2, Cappi3, Cappi4, Cappi5 []float32
}
