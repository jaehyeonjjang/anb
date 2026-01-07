
import 'package:common_control/common_control.dart';

class UploadManager {
  static Future<String> image(String path) async {
    var url = '/api/upload/periodic';
    return await Http.upload(url, 'file', path);    
  }
}
