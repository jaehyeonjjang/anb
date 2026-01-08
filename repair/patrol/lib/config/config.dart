import 'dart:io';

import 'package:flutter/material.dart';

// Android 에뮬레이터에서는 10.0.2.2 사용
const serverUrl = 'http://10.0.2.2:9108';
//const serverUrl = 'https://service.anbweb.kr';
//const serverUrl = 'http://service.anbweb.kr';

const primaryColor = Color.fromARGB(255, 33, 56, 190);
const selectColor = Color.fromARGB(255, 232, 31, 99);
const backgroundColor = Color.fromARGB(255, 101, 192, 240);
const titleBackgroundColor = Color.fromARGB(255, 101, 192, 240);

String platform() {
  try {
    if (Platform.isAndroid) {
      return 'android';
    } else if (Platform.isIOS) {
      return 'ios';
    }
  } catch (e) {
    return 'web';
  }

  return '';
}
