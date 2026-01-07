import 'package:common_control/common_control.dart';
import 'package:flutter/cupertino.dart';
import 'package:patrol/components/painter/painter_controller.dart';
import 'package:patrol/controllers/auth_controller.dart';

class ToolBox extends StatelessWidget {
  ToolBox({super.key});

  final c = Get.find<PainterController>();
  final authController = Get.find<AuthController>();

  Widget drawButton(int index, String title) {
    if (index == c.index && c.mode == Mode.draw) {
      return InkWell(
          onTap: () {
            c.setMode(Mode.draw);
            c.setIndex(index);            
          },
          child: Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(5.0),
                color: Colors.yellow,
              ),
              padding: const EdgeInsets.all(5.0),
              child: Image.asset(title)));
    }

    return InkWell(
        onTap: () {
          c.setMode(Mode.draw);
          c.setIndex(index);
        },
        child: Container(
            padding: const EdgeInsets.all(5.0), child: Image.asset(title)));
  }

  Widget toolBoxNumber(int index, Color color) {
    Color background = Colors.white;

    if (index == c.index) {
      background = Colors.yellow;
    }

    return InkWell(
        onTap: () {
          c.setMode(Mode.draw);
          c.setIndex(index);
        },
        child: Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(5.0),
              color: background,
            ),
            padding: const EdgeInsets.symmetric(horizontal: 4.0, vertical: 1.0),
            child: Container(
                decoration: BoxDecoration(
                    color: Colors.white,
                    shape: BoxShape.circle,
                    border: Border.all(width: 1.5, color: color)),
                padding: const EdgeInsets.all(5),
                child:
                    Text('1', style: TextStyle(fontSize: 13, color: color)))));
  }

  @override
  Widget build(BuildContext context) {
    return Obx(() => body());
  }

  Widget body() {
    if (c.toolbox == false) {
      return Container();
    }

    return Container(
        margin: const EdgeInsets.only(left: 5.0, top: 5.0),
        padding: const EdgeInsets.only(left:10.0, top:10.0, bottom:10.0, right: 5.0),
        width: 83,
        decoration: BoxDecoration(
            border: Border.all(width: 1.0),
            borderRadius: const BorderRadius.all(
              Radius.circular(5.0),
            ),
            color: Colors.white),
        child: Column(mainAxisSize: MainAxisSize.min, children: [
          Row(children: [
            toolBoxNumber(1, Colors.red),
            toolBoxNumber(2, const Color.fromRGBO(0, 0, 255, 1.0)),
          ]),
          
          const SizedBox(height: 20.0),
          Row(children: [
            drawButton(31, 'assets/imgs/curve_blue.png'),  
            drawButton(41, 'assets/imgs/line_blue.png'),            
          ]),
          Row(children: [
            drawButton(32, 'assets/imgs/curve_red.png'),
            drawButton(42, 'assets/imgs/line_red.png'),            
          ]),
          Row(children: [
            drawButton(33, 'assets/imgs/curve_green.png'),
            drawButton(43, 'assets/imgs/line_green.png'),
          ]),
          Row(children: [
          drawButton(101, 'assets/imgs/i001.png'),
          drawButton(102, 'assets/imgs/i002.png'),
          ]),
          Row(children: [
          drawButton(103, 'assets/imgs/i003.png'),
          drawButton(104, 'assets/imgs/i004.png'),
          ]),
          Row(children: [
          drawButton(105, 'assets/imgs/i005.png'),
          drawButton(106, 'assets/imgs/i006.png'),
          ]),
          Row(children: [
          drawButton(107, 'assets/imgs/i007.png'),
          drawButton(108, 'assets/imgs/i008.png'),
          ]),
          Row(children: [
          drawButton(109, 'assets/imgs/i009.png'),
          drawButton(110, 'assets/imgs/i010.png'),
          ]),
          Row(children: [
            drawButton(111, 'assets/imgs/i011.png'),
            Container(),
          ]),
          const SizedBox(height: 20.0),
          Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
            const SizedBox(width: 1),            
          InkWell(
              onTap: () {
                c.setIconZoom(c.iconZoom + authController.zoomlevel.toDouble());
              },
              child: const Icon(CupertinoIcons.add, size: 24.0)),
          const SizedBox(height: 5.0),
          InkWell(
              onTap: () {
                c.setIconZoom(c.iconZoom - authController.zoomlevel.toDouble());
              },
              child: const Icon(CupertinoIcons.minus, size: 24.0)),
              const SizedBox(width: 2),            
          ]),
        ]));
  }
}
