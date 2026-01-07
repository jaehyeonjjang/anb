all:

release:
	flutter build apk --release --target-platform=android-arm64 --dart-define=API_BASE_URL=https://service.anbweb.kr

localdeploy: release
	cp build/app/outputs/flutter-apk/app-release.apk ~/Desktop/periodic-V$(version).apk

deploy: release
	#cp build/app/outputs/flutter-apk/app-release.apk ~/googledrive/anb/periodic-V$(version).apk
	cp build/app/outputs/flutter-apk/app-release.apk ~/Desktop/periodic-V$(version).apk
	scp build/app/outputs/flutter-apk/app-release.apk root@service.anbweb.kr:/mnt/webdata/apk/periodic-V$(version).apk
	ssh root@db.service.anbweb.kr "/root/update_program_version.sh $(version)"
