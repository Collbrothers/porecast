
# Porecast  (Polybar-Forecast)

OpenWeather integration for [polybar](https://github.com/polybar/polybar)  to get current forecast

## Installation
Clone the repository:
```
git clone https://github.com/Collbrothers/porecast
```
Move into the directory and build it
```
cd porecast && go build
```
This has currently only been tried using go version 1.22.3 on linux/amd64, but should work in most cases.

## Create & modify the configuration file
```
./porecast init
```
This will create a "config" file in your home directory, under .config/porecast and write a template that you will need to edit:
```
nano $HOME/.config/porecast/config
```
You will need to grab an [API key from OpenWeather](https://home.openweathermap.org/api_keys) **(This service requires you to create an account, but it is free)**.
You are also going to grab the coordinates of the location you want the forecast of, this can be acquired from sites like LatLong.net.

## Modify your polybar configuration file
You need to add a section to your polybar `config.ini:
```bash
nano $HOME/.config/polybar/config.ini
```
Create a module:
```ini
[module/porecast]
type = custom/script 
exec = PATH/TO/PORECAST/EXECUTABLE run
exec-if = ping openweathermap.org -c 1
label = "%output% °C"
interval = 100
```
This example works for Celsius, modify the symbol to correlate with the unit you set in the porecast configuration file.

You may also add color to the label to fit with the default polybar theme:
```ini
[module/porecast]
type = custom/script 
exec = PATH/TO/PORECAST/EXECUTABLE run
exec-if = ping openweathermap.org -c 1
label = "%output% %{F#F0C674}°C%{F-}"
interval = 100
```

**Notice the run command is being used, so when you change the path to the script, do not delete the command**

Make sure to add the module to the bar:
```ini
[bar/example]
modules-right = ... porecast ... 
```

## Upcoming features
*	Future forecast support
*	Icon support
