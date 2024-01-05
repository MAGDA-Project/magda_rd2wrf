# MAGDA Weather stations to WRF

This module can be used to convert radar data in netcdf 
MAGDA format into WRF ascii format.

## Installation

The module use go-netcdf to read CAPPI netcdf files.

In order to use it, you need the developer version of the
library provided by your distribution installed.

On ubuntu you can install it with:

```bash
sudo apt install libnetcdf-dev
```

On Typhoon, it can be loaded with the WRF-KIT2 module:

```bash
module load gcc-8.3.1/WRF-KIT2
```

## Usage on CIMA Typhoon

`magda_rd2wrf` is already present in /data/safe/home/wrfprod/bin/magda_rd2wrf

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

