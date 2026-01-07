import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:periodic/components/painter/painter_controller.dart';
import 'package:periodic/config/config.dart' as config;

class GroupBox extends StatelessWidget {
  GroupBox({super.key});

  final c = Get.find<PainterController>();

  @override
  Widget build(BuildContext context) {
    return Obx(() => body());
  }

  Widget body() {
    if (c.groupbox == false) {
      return Container();
    }

    List<TableRow> lists = [];

    final widget = TableRow(children: [
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(5),
          child: const Text('순번')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(5),
          child: const Text('그룹/번호')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(5),
          child: const Text('지정')),
    ]);
    lists.add(widget);

    for (var i = 0; i < c.points.length; i++) {
      final item = c.points[i];

      if (item.type != DrawType.number && item.type != DrawType.numberLine) {
        continue;
      }

      if (!item.selected) {
        continue;
      }

      final widget = TableRow(children: [
        Container(
            alignment: Alignment.center,
            padding: const EdgeInsets.all(5),
            child: Text((i + 1).toString())),
        Container(
            alignment: Alignment.center,
            padding: const EdgeInsets.all(5),
            child: Text(item.number.toString())),
        Container(
          alignment: Alignment.center,
          height: 20,
          padding: const EdgeInsets.only(top: 6),
          child: Checkbox(
            value: c.points[i].grouped,
            onChanged: (value) {
              c.points[i].grouped = value;
              c.points.refresh();
            },
          ),
        ),
      ]);
      lists.add(widget);
    }

    return Positioned(
        top: 0.0,
        right: 0.0,
        child: Container(
          margin: const EdgeInsets.only(left: 10.0),
          padding: const EdgeInsets.all(10.0),
          decoration: BoxDecoration(
              border: Border.all(width: 1.0),
              borderRadius: const BorderRadius.all(
                Radius.circular(5.0),
              ),
              color: Colors.white),
          child: Column(children: [
            Container(
                width: 240,
                padding: const EdgeInsets.only(
                    left: 10, top: 5, right: 5, bottom: 5),
                decoration: BoxDecoration(
                    color: config.titleBackgroundColor,
                    borderRadius: BorderRadius.circular(5)),
                child: const Text('그룹 설정',
                    style: TextStyle(
                        fontSize: 18,
                        color: Colors.black,
                        fontWeight: FontWeight.w700))),
            Container(
              width: 240,
              margin: const EdgeInsets.only(top: 5, bottom: 5),
              child: Table(
                border: TableBorder.all(color: Colors.grey[400]!),
                children: lists,
              ),
            ),
            Row(children: [
              OutlinedButton(
                  style: OutlinedButton.styleFrom(
                    backgroundColor: const Color.fromRGBO(64, 158, 255, 1.000),
                    foregroundColor: Colors.amberAccent, //<-- SEE HERE
                  ),
                  onPressed: clickGroupApply,
                  child: const Text('적용',
                      style: TextStyle(
                          fontSize: 15,
                          fontWeight: FontWeight.bold,
                          color: Colors.white))),
              const SizedBox(width: 10),
              OutlinedButton(
                  style: OutlinedButton.styleFrom(
                    backgroundColor: const Color.fromRGBO(245, 108, 108, 1.000),
                    foregroundColor: Colors.amberAccent, //<-- SEE HERE
                  ),
                  onPressed: clickGroupCancel,
                  child: const Text('해제',
                      style: TextStyle(
                          fontSize: 15,
                          fontWeight: FontWeight.bold,
                          color: Colors.white))),
              const SizedBox(width: 10),
              OutlinedButton(
                  onPressed: () => c.groupbox = false,
                  child: Text('닫기',
                      style: TextStyle(
                          fontSize: 15,
                          fontWeight: FontWeight.bold,
                          color: Colors.grey[700]))),
            ]),
          ]),
        ));
  }

  clickGroupApply() {
    c.groupApply();

    c.groupbox = false;
  }

  clickGroupCancel() {
    c.groupCancel();

    c.groupbox = false;
  }
}
