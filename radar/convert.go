package radar

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/fhs/go-netcdf/netcdf"
)

func filenameForLev(dirname string, cappilev int, dt time.Time) string {
	pt := fmt.Sprintf("%s/%s_cappi%dkm.nc", dirname, dt.Format("20060102_1504"), cappilev)
	return pt
}

func writeRadarData(f io.Writer, val float32, height float64) {

	if val < 0 {
		// write(301,'(3x,f12.1,2(f12.3,i4,f12.3,2x))')
		// hgt(i,m), rv_data(i,m), rv_qc(i,m), rv_err(i,m), rf_data(i,m), rf_qc(i,m), rf_err(i,m)
		fmt.Fprintf(f, "   %12.1f -888888.000 -88 -888888.000   -888888.000 -88 -888888.000\n", height)
		return
	}

	fmt.Fprintf(
		f,
		// write(301,'(3x,f12.1,2(f12.3,i4,f12.3,2x))')
		// hgt(i,m), rv_data(i,m), rv_qc(i,m), rv_err(i,m), rf_data(i,m), rf_qc(i,m), rf_err(i,m)
		"   %12.1f -888888.000 -88 -888888.000  %12.3f   0       5.000\n",
		height,
		val,
	)

}

func writeConvertedDataTo(resultW io.WriteCloser, dims *MosaicData, dtRequested time.Time) {
	defer resultW.Close()
	result := bufio.NewWriterSize(resultW, 1000000)
	defer result.Flush()

	maxLon := float64(-1)
	maxLat := float64(-1)
	instant := dtRequested.Format("2006-01-02_15:04")
	totObs := 0

	totCappi := 0

	if dims.Cappi2 != nil {
		totCappi++
	}

	if dims.Cappi3 != nil {
		totCappi++
	}

	if dims.Cappi4 != nil {
		totCappi++
	}

	if dims.Cappi5 != nil {
		totCappi++
	}

	if dims.Cappi6 != nil {
		totCappi++
	}

	if dims.Cappi7 != nil {
		totCappi++
	}

	if dims.Cappi8 != nil {
		totCappi++
	}

	if dims.Cappi9 != nil {
		totCappi++
	}

	if dims.Cappi2 != nil ||
		dims.Cappi3 != nil ||
		dims.Cappi4 != nil ||
		dims.Cappi5 != nil ||
		dims.Cappi6 != nil ||
		dims.Cappi7 != nil ||
		dims.Cappi8 != nil ||
		dims.Cappi9 != nil {
		for _, l := range dims.Lon {
			if l > maxLon {
				maxLon = l
			}
		}
		for _, l := range dims.Lat {
			if l > maxLat {
				maxLat = l
			}
		}
	} else {
		maxLon = float64(1)
		maxLat = float64(1)
	}

	for i := int64(0); i < dims.Width*dims.Height; i++ {
		f2 := float32(-1)
		f3 := float32(-1)
		f4 := float32(-1)
		f5 := float32(-1)
		f6 := float32(-1)
		f7 := float32(-1)
		f8 := float32(-1)
		f9 := float32(-1)

		if dims.Cappi2 != nil {
			f2 = dims.Cappi2[i]
		}

		if dims.Cappi3 != nil {
			f3 = dims.Cappi3[i]
		}

		if dims.Cappi4 != nil {
			f4 = dims.Cappi4[i]
		}

		if dims.Cappi5 != nil {
			f5 = dims.Cappi5[i]
		}

		if dims.Cappi6 != nil {
			f6 = dims.Cappi6[i]
		}

		if dims.Cappi7 != nil {
			f7 = dims.Cappi7[i]
		}

		if dims.Cappi8 != nil {
			f8 = dims.Cappi8[i]
		}

		if dims.Cappi9 != nil {
			f9 = dims.Cappi9[i]
		}

		if f2 >= 0 || f3 >= 0 || f4 >= 0 || f5 >= 0 || f6 >= 0 || f7 >= 0 || f8 >= 0 || f9 >= 0 {
			totObs++
		}
	}

	fmt.Fprintf(result, "TOTAL NUMBER =  1\n")
	fmt.Fprintf(result, "#-----------------#\n")
	fmt.Fprintf(result, "\n")
	//  write(301,'(a5,2x,a12,2(f8.3,2x),f8.1,2x,a19,2i6)') 'RADAR', &
	//  radar_name, rlonr(irad), rlatr(irad), raltr(irad)*1000., &
	//  trim(radar_date), np, imdv_nz(irad)
	fmt.Fprintf(result, "RADAR              %8.3f  %8.3f     100.0  %s:00%6d     %d\n",
		maxLon,
		maxLat,
		instant,
		totObs,
		totCappi,
	)

	fmt.Fprintf(result, "#-------------------------------------------------------------------------------#\n")
	fmt.Fprintf(result, "\n")

	if dims.Cappi2 == nil &&
		dims.Cappi3 == nil &&
		dims.Cappi4 == nil &&
		dims.Cappi5 == nil &&
		dims.Cappi6 == nil &&
		dims.Cappi7 == nil &&
		dims.Cappi8 == nil &&
		dims.Cappi9 == nil {
		return
	}

	//instant = dims.Instants[0].Format("2006-01-02_15:04")

	for x := int64(0); x < dims.Width; x++ {
		for y := int64(dims.Height) - 1; y >= int64(0); y-- {
			//
			lat := dims.Lat[y]
			lon := dims.Lon[x]

			f2 := float32(-1)
			f3 := float32(-1)
			f4 := float32(-1)
			f5 := float32(-1)
			f6 := float32(-1)
			f7 := float32(-1)
			f8 := float32(-1)
			f9 := float32(-1)
			i := x + y*dims.Width

			totCappi := 0

			if dims.Cappi2 != nil {
				f2 = dims.Cappi2[i]
				totCappi++
			}

			if dims.Cappi3 != nil {
				f3 = dims.Cappi3[i]
				totCappi++
			}

			if dims.Cappi4 != nil {
				f4 = dims.Cappi4[i]
				totCappi++
			}

			if dims.Cappi5 != nil {
				f5 = dims.Cappi5[i]
				totCappi++
			}

			if dims.Cappi6 != nil {
				f6 = dims.Cappi6[i]
				totCappi++
			}

			if dims.Cappi7 != nil {
				f7 = dims.Cappi7[i]
				totCappi++
			}

			if dims.Cappi8 != nil {
				f8 = dims.Cappi8[i]
				totCappi++
			}

			if dims.Cappi9 != nil {
				f9 = dims.Cappi9[i]
				totCappi++
			}

			if f2 >= 0 || f3 >= 0 || f4 >= 0 || f5 >= 0 || f6 >= 0 || f7 >= 0 || f8 >= 0 || f9 >= 0 {
				fmt.Fprintf(
					result,
					//!----Write data
					//do i = 1,np ! np: # of total horizontal data points
					//write(301,'(a12,3x,a19,2x,2(f12.3,2x),f8.1,2x,i6)') 'FM-128 RADAR', &
					// trim(radar_date), plat(i), plon(i), raltr(irad)*1000, count_nz(i)

					//"FM-128 RADAR   %s:00       %7.3f      %8.3f     100.0       3\n",
					"FM-128 RADAR   %s:00  %12.3f  %12.3f     100.0       %d\n",
					instant,
					lat,
					lon,
					totCappi,
				)

				if dims.Cappi2 != nil {
					writeRadarData(result, f2, 2000.0)
				}
				if dims.Cappi3 != nil {
					writeRadarData(result, f3, 3000.0)
				}
				if dims.Cappi4 != nil {
					writeRadarData(result, f4, 4000.0)
				}
				if dims.Cappi5 != nil {
					writeRadarData(result, f5, 5000.0)
				}
				if dims.Cappi6 != nil {
					writeRadarData(result, f6, 6000.0)
				}
				if dims.Cappi7 != nil {
					writeRadarData(result, f7, 7000.0)
				}
				if dims.Cappi8 != nil {
					writeRadarData(result, f8, 8000.0)
				}
				if dims.Cappi9 != nil {
					writeRadarData(result, f9, 9000.0)
				}
			}
		}
	}
}

func readDataFromFile(mos *MosaicData, gridTmpl string, dirname string, dt time.Time, dest *[]float32, cappilev int) error {
	var err error
	var ds netcdf.Dataset

	fname := filenameForLev(dirname, cappilev, dt)

	if _, err := os.Stat(fname); os.IsNotExist(err) {
		*dest = nil
		return nil
	}

	fnameRegrid := fname + ".regrid.nc"

	// regrid to the same resolution as domain
	out, err := exec.Command("cdo", "remapbil,"+gridTmpl, fname, fnameRegrid).Output()
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("cannot regrid CAPPI file %s: %w\nCommand output:\n%s", fname, e, string(e.Stderr)+"\n"+string(out))
		} else {
			return fmt.Errorf("cannot regrid CAPPI file %s: %w", fname, err)
		}
	}
	defer os.Remove(fnameRegrid)

	fnameFiltered := fname + ".filtered.nc"
	operator := "where(DBZH < 10) DBZH=-9999"

	if out, err := exec.Command("ncap2", "-s", operator, fnameRegrid, fnameFiltered).Output(); err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("cannot filter CAPPI file %s: %w\nCommand output:\n%s", fname, e, string(e.Stderr)+"\n"+string(out))
		} else {
			return fmt.Errorf("cannot filter CAPPI file %s: %w", fname, err)
		}
	}
	defer os.Remove(fnameFiltered)

	if ds, err = netcdf.OpenFile(fnameFiltered, netcdf.FileMode(netcdf.NOWRITE)); err != nil {
		return fmt.Errorf("cannot open CAPPI file %s: %w", fname, err)
	}
	if mos.Width == -1 {
		//fmt.Println("MosaicData dimensions not initialized.")

		if mos.Width, err = GetDimensionLen(&ds, "lon"); err != nil {
			return fmt.Errorf("cannot get dimension lon from CAPPI file %s: %w", fname, err)
		}
		//fmt.Println("Width", mos.Width)

		if mos.Height, err = GetDimensionLen(&ds, "lat"); err != nil {
			return fmt.Errorf("cannot get dimension lat from CAPPI file %s: %w", fname, err)
		}
		//fmt.Println("Height", mos.Height)

		if mos.Lat, err = ReadDoubleVar(&ds, "lat"); err != nil {
			return fmt.Errorf("cannot read lat from CAPPI file %s: %w", fname, err)
		}
		//fmt.Println("Latitude", mos.Lat)

		if mos.Lon, err = ReadDoubleVar(&ds, "lon"); err != nil {
			return fmt.Errorf("cannot read lon from CAPPI file %s: %w", fname, err)
		}
		//fmt.Println("Longitude", mos.Lon)

		//if mos.Instants, err = ReadTimeVar(&ds, "time"); err != nil {
		//	return fmt.Errorf("cannot read time from CAPPI file %s: %w", fname, err)
		//}
		//fmt.Println("Time", mos.Instants)
	}
	if *dest, err = ReadFloatVar(&ds, "DBZH"); err != nil {
		return fmt.Errorf("cannot read DBZH from CAPPI file %s: %w", fname, err)
	}

	return ds.Close()
}

// Convert ...
func Convert(dirname string, gridTmpl string, dt time.Time) (io.Reader, error) {
	mos := MosaicData{
		Width: -1,
	}

	var err error
	if err = readDataFromFile(&mos, gridTmpl, dirname, dt, &mos.Cappi2, 2); err != nil {
		return nil, fmt.Errorf("cannot read CAPPI level 2 file: %w", err)
	}

	if err = readDataFromFile(&mos, gridTmpl, dirname, dt, &mos.Cappi3, 3); err != nil {
		return nil, fmt.Errorf("cannot read CAPPI level 3 file: %w", err)
	}

	if err = readDataFromFile(&mos, gridTmpl, dirname, dt, &mos.Cappi4, 4); err != nil {
		return nil, fmt.Errorf("cannot read CAPPI level 4 file: %w", err)
	}

	if err = readDataFromFile(&mos, gridTmpl, dirname, dt, &mos.Cappi5, 5); err != nil {
		return nil, fmt.Errorf("cannot read CAPPI level 5 file: %w", err)
	}

	if err = readDataFromFile(&mos, gridTmpl, dirname, dt, &mos.Cappi6, 6); err != nil {
		return nil, fmt.Errorf("cannot read CAPPI level 6 file: %w", err)
	}

	if err = readDataFromFile(&mos, gridTmpl, dirname, dt, &mos.Cappi7, 7); err != nil {
		return nil, fmt.Errorf("cannot read CAPPI level 7 file: %w", err)
	}

	if err = readDataFromFile(&mos, gridTmpl, dirname, dt, &mos.Cappi8, 8); err != nil {
		return nil, fmt.Errorf("cannot read CAPPI level 8 file: %w", err)
	}

	if err = readDataFromFile(&mos, gridTmpl, dirname, dt, &mos.Cappi9, 9); err != nil {
		return nil, fmt.Errorf("cannot read CAPPI level 9 file: %w", err)
	}

	reader, result := io.Pipe()

	go writeConvertedDataTo(result, &mos, dt)
	return reader, err
}
