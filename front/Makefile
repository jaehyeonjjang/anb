all: release

release: dummy
	cordova build android --release --buildConfig=build.json

run: dummy
	cordova run --emulator android

device: dummy
	cordova run --device android

dummy:
