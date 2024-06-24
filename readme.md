# MAGDA radars to WRF

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

Template files to regrid radar for Romania and Switzerland is already present in path `~/.magda_rd2wrf`

## Command line usage

This module implements a console command
that can be used to convert radar observations from
netcdf to WRF ascii format.

Usage of `magda_rd2wrf`:

```
GRD_TEMPL=~/.magda_rd2wrf/swiss.grid.template
magda_rd2wrf <inputdir> <outfilename> YYYYMMDDHHNN
Options:
  -inputdir string
        the directory containing all CAPPI netcdf files
  -outfilename string
        name of the output file
  -YYYYMMDDHHNN date/time
        date and hour of the radar data to convert [YYYYMMDDHH]
```

GRD_TEMPL is an environment variable that must be set to the CDO template file to use for regridding. 
On Typhoon, there are two grid template already defined in ~/.magda_rd2wrf:

* swiss.grid.template for Switzerland radars
* romania.grid.template for Romania radars


## Build

1) If not already installed, follow the instructions here to install the Go Workbench::
[Install-go-environment](https://github.com/meteocima/documentazione/wiki/Installare-ambiente-di-lavoro-go)

2) Clone this repository 

```bash
git clone https://github.com/MAGDA-Project/magda_rd2wrf.git
cd magda_rd2wrf
```

5) Compile the executable with this command:

```bash
go build ./cli/magda_rd2wrf
```