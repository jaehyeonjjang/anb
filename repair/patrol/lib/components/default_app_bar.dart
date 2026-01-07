import 'package:common_control/common_control.dart';
import 'package:patrol/controllers/auth_controller.dart';

class DefaultAppBar extends CWidget {
  DefaultAppBar({super.key, this.leading, this.actions});

  final c = Get.find<AuthController>();

  final Widget? leading;
  final List<Widget>? actions;

  @override
  build(BuildContext context) {
    const titleBackgroundColor = Color.fromARGB(255, 101, 192, 240);

    var title = c.title;

    if (title == '') {
      if (c.periodic.id > 0) {
        title = c.periodic.apt.name;

        if (title == '') {
          title = c.apt.name;
        }
      } else if (c.apt.id > 0) {
        title = c.apt.name;
      }
    }

    return AppBar(
      iconTheme: IconThemeData(
        color: Theme.of(context).primaryColor,
      ),
      title: Obx(() => Text(c.title == '' ? title : c.title,
          style: const TextStyle(fontSize: 20, fontWeight: FontWeight.w500))),
      backgroundColor: titleBackgroundColor,
      titleTextStyle: const TextStyle(color: Colors.black87),
      bottomOpacity: 0.0,
      elevation: 0.0,
      centerTitle: true,
      leading: leading,
      actions: actions,
    );
  }
}
