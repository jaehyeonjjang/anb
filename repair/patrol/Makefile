all:

release:
	flutter build apk --release --target-platform=android-arm64 --dart-define=API_BASE_URL=https://www.anbweb.kr:4443

deploy: release
	cp build/app/outputs/flutter-apk/app-release.apk ~/googledrive/anb/patrol-V$(version).apk
	scp build/app/outputs/flutter-apk/app-release.apk root@service.anbweb.kr:/home/repair/back/webdata/apk/patrol-V$(version).apk
