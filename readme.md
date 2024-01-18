# MAGDA Weather stations to WRF

This module can be used to convert radar data in netcdf 
MAGDA format into WRF ascii format.

## Installation

The module use go-netcdf to read CAPPI netcdf files.

In order to use it, you need the developer version of the
library provided by your distribution installed.

Furthermore, netcdf files are preprocessed using CDO and NCO tools.

On ubuntu you can install these dependencies with:

```bash
sudo apt install libnetcdf-dev cdo nco
```

On Typhoon, `wrfprod` user can load them with these modules:

```bash
module load CDO NCO
```

## Usage on CIMA Typhoon

`magda_rd2wrf` is already present in /data/safe/home/wrfprod/bin/magda_rd2wrf
A template file to regrid radar filse is already present in path `~/temp_Romania.nc`

## Command line usage

This module implements a console command
that can be used to convert radar observations from
netcdf to WRF ascii format.

Usage of `magda_rd2wrf`:

```
magda_rd2wrf <inputdir> <outfilename> YYYYMMDDHHNN
Options:
  -inputdir string
        the directory containing all CAPPI netcdf files
  -outfilename string
        name of the output file
  -YYYYMMDDHHNN date/time
        date and hour of the radar data to convert [YYYYMMDDHH]
```

