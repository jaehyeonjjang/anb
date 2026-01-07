import 'dart:math';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:periodic/components/painter/painter_controller.dart';

class PainterDrawer extends CustomPainter {
  final PainterController c = Get.put(PainterController());

  Paint paintBlack = Paint()
    ..style = PaintingStyle.stroke
    ..color = Colors.black
    ..strokeWidth = 2;

  Paint paintRed = Paint()
    ..style = PaintingStyle.stroke
    ..color = Colors.red
    ..strokeWidth = 2;

  Paint paintGreen = Paint()
    ..style = PaintingStyle.stroke
    ..color = Colors.green
    ..strokeWidth = 2;

  Paint paintBlue = Paint()
    ..style = PaintingStyle.stroke
    ..color = const Color(0xff0000ff)
    ..strokeWidth = 2;

  Paint paintLightblue = Paint()
    ..style = PaintingStyle.stroke
    ..color = Color.fromRGBO(0, 0, 255, 1.0)
    ..strokeWidth = 2;

  Paint paintSelectbox = Paint()
    ..style = PaintingStyle.stroke
    ..color = const Color(0xff7ad7f0)
    ..strokeWidth = 2;

  Paint paintSelect = Paint()
    ..style = PaintingStyle.stroke
    ..color = const Color(0xff7ad7f0)
    ..strokeWidth = 2;

  Paint paintFillSelect = Paint()
    ..style = PaintingStyle.fill
    ..color = const Color(0xff7ad7f0);

  Paint paintViolet = Paint()
    ..style = PaintingStyle.stroke
    ..color = const Color(0xffa000a0)
    ..strokeWidth = 2;

  Paint paintFillViolet = Paint()
    ..style = PaintingStyle.fill
    ..color = const Color(0xffa000a0);

  Paint paintFillLightblue = Paint()
    ..style = PaintingStyle.fill
    ..color = Color.fromRGBO(0, 0, 255, 1.0);

  Paint paintFillRed = Paint()
    ..style = PaintingStyle.fill
    ..color = Colors.red;

  Paint paintFillBlue = Paint()
    ..style = PaintingStyle.fill
    ..color = const Color(0xff0000ff);

  Paint paintFillWhite = Paint()
    ..style = PaintingStyle.fill
    ..color = Colors.white;

  Paint paintFillGray = Paint()
    ..style = PaintingStyle.fill
    ..color = Color.fromRGBO(200, 200, 200, 1.0);

  @override
  void paint(Canvas canvas, Size size) {
    var zoom = c.zoom;
    var dx = -1 * (c.sx + c.currentSx) * zoom;
    var dy = -1 * (c.sy + c.currentSy) * zoom;

    Rect rect = const Offset(0.0, 0.0) & size;
    canvas.drawRect(rect, paintFillGray);

    if (c.isLoaded == true) {
      paintImage(
          canvas: canvas,
          rect:
              Rect.fromLTWH(dx, dy, c.imageWidth * zoom, c.imageHeight * zoom),
          image: c.image,
          fit: BoxFit.fill,
          repeat: ImageRepeat.noRepeat,
          scale: 1.0,
          alignment: Alignment.center,
          flipHorizontally: false,
          filterQuality: FilterQuality.high);
    }

    for (var i = 0; i < c.points.length; i++) {
      Path path = Path();

      var points = c.points[i];

      if (c.iconset == 1) {
        if (points.icon >= 200) {
          continue;
        }
      } else if (c.iconset == 2) {
        if (points.icon < 200 || points.icon >= 300) {
          continue;
        }
      } else if (c.iconset == 3) {
        if (points.icon < 300 || points.icon >= 400) {
          continue;
        }
      } else if (c.iconset == 4) {
        if (points.icon < 400 || points.icon >= 500) {
          continue;
        }
      }

      if (c.iconset != 2 && c.iconset != 4) {
        if (points.type == DrawType.icon ||
            points.type == DrawType.number ||
            points.type == DrawType.material) {
          drawIcon(canvas, points);

          //continue;
        }
      }

      Paint paint = paintFillRed;

      if (points.icon == basicHorizontalLine ||
          points.icon == basicHorizontalBreak) {
        paint = paintFillBlue;
        points.color = LineColor.blue;
      } else if (points.icon == basicVerticalLine ||
          points.icon == basicVerticalBreak) {
        paint = paintFillRed;
        points.color = LineColor.red;
      }

      if (points.icon == inclinationHorizontal) {
        if (points.items.length > 1) {
          var point = points.items[0];
          var x = point.dx * zoom + dx;
          var y = point.dy * zoom + dy;
          path.moveTo(x, y);

          point = points.items[points.items.length - 1];

          x = point.dx * zoom + dx;
          path.lineTo(x, y);

          x = point.dx * zoom + dx;
          y = point.dy * zoom + dy;
          path.lineTo(x, y);
        }
      } else if (points.icon == inclinationVertical) {
        if (points.items.length > 1) {
          var point = points.items[0];
          var x = point.dx * zoom + dx;
          var y = point.dy * zoom + dy;
          path.moveTo(x, y);

          point = points.items[points.items.length - 1];

          y = point.dy * zoom + dy;
          path.lineTo(x, y);

          x = point.dx * zoom + dx;
          y = point.dy * zoom + dy;
          path.lineTo(x, y);
        }
      } else {
        for (var j = 0; j < points.items.length; j++) {
          var point = points.items[j];
          var x = point.dx * zoom + dx;
          var y = point.dy * zoom + dy;
          if (j == 0) {
            path.moveTo(x, y);
          } else {
            path.lineTo(x, y);
          }
        }
      }

      var strokeWidth = 2.0;
      if (points.icon == basicVertical ||
          points.icon == basicHorizontal ||
          points.icon == basicVerticalLine ||
          points.icon == basicHorizontalLine ||
          points.icon == basicVerticalBreak ||
          points.icon == basicHorizontalBreak) {
        strokeWidth = 1.5;
      }

      if (points.icon == crackLineRed ||
          points.icon == crackLineBlue ||
          points.icon == crackLineViolet ||
          points.icon == crackCurveRed ||
          points.icon == crackCurveBlue ||
          points.icon == crackCurveViolet) {}

      paint.strokeWidth = strokeWidth;
      paintLightblue.strokeWidth = strokeWidth;
      paintRed.strokeWidth = strokeWidth;
      paintViolet.strokeWidth = strokeWidth;
      paintGreen.strokeWidth = strokeWidth;
      paintBlack.strokeWidth = strokeWidth;
      paintBlue.strokeWidth = strokeWidth;

      if (points.selected == true) {
        canvas.drawPath(path, paintSelect);
      } else if (points.icon == 31 ||
          points.icon == 41 ||
          points.icon == crackLineBlue ||
          points.icon == crackCurveBlue) {
        canvas.drawPath(path, paintLightblue);
      } else if (points.icon == 32 ||
          points.icon == 42 ||
          points.icon == crackLineRed ||
          points.icon == crackCurveRed) {
        canvas.drawPath(path, paintRed);
      } else if (points.icon == crackLineViolet ||
          points.icon == crackCurveViolet) {
        canvas.drawPath(path, paintViolet);
      } else if (points.icon == lineGreen || points.icon == curveGreen) {
        canvas.drawPath(path, paintGreen);
      } else if (points.icon == 34 || points.icon == 44) {
        canvas.drawPath(path, paintViolet);
      } else if (points.color == LineColor.black) {
        canvas.drawPath(path, paintBlack);
      } else if (points.color == LineColor.red) {
        canvas.drawPath(path, paintRed);
        paint = paintFillRed;
      } else if (points.color == LineColor.blue) {
        canvas.drawPath(path, paintBlue);
        paint = paintFillBlue;
      } else if (points.color == LineColor.green) {
        canvas.drawPath(path, paintGreen);
      } else if (points.color == LineColor.lightblue) {
        canvas.drawPath(path, paintLightblue);
      }

      if (points.isCrack()) {
        drawIcon(canvas, points);
      }

      if (points.isInclination() == true || points.isBasicLine() == true) {
        if (points.icon == basicVerticalBreak ||
            points.icon == basicHorizontalBreak) {
          drawBreakArrow(canvas, points, paint);
        } else {
          drawArrow(canvas, points, paint);
        }
        drawIcon(canvas, points);
      }

      if (points.isBasicLine() == true) {
        Color color = Colors.red;
        if (points.icon == basicHorizontalLine ||
            points.icon == basicHorizontalBreak) {
          color = Color.fromRGBO(0, 0, 255, 1.0);
        }
        var x = points.items[points.items.length - 1].dx * zoom + dx;
        var y = points.items[points.items.length - 1].dy * zoom + dy;
        var step = (c.numberZoom / 2.0) * zoom;
        var offset = Offset(x, y);
        var txt = points.member;
        if (points.shape != '') {
          txt = '$txt(${points.shape})';
        }
        var textSpan = TextSpan(
          text: txt,
          style: TextStyle(
              fontSize: step * 1.2 * 1.5,
              fontWeight: FontWeight.w600,
              color: color),
        );

        if (points.selected == true) {
          textSpan = TextSpan(
            text: txt,
            style: TextStyle(
                fontSize: step * 1.2 * 1.5,
                fontWeight: FontWeight.w600,
                color: const Color(0xff7ad7f0)),
          );
        }

        final textPainter = TextPainter()
          ..text = textSpan
          ..textDirection = TextDirection.ltr
          ..textAlign = TextAlign.center
          ..layout();

        if (points.items[0].dx > points.items[points.items.length - 1].dx) {
          offset = Offset(
              x - step * 1.5 - textPainter.width, y - textPainter.height / 2);
        } else {
          offset = Offset(x + step * 1.5, y - textPainter.height / 2);
        }

        textPainter.paint(canvas, offset);
      }
    }

    if (c.mode == Mode.select) {
      Path path = Path();

      path.moveTo(c.startSx, c.startSy);
      path.lineTo(c.endSx, c.startSy);
      path.lineTo(c.endSx, c.endSy);
      path.lineTo(c.startSx, c.endSy);
      path.lineTo(c.startSx, c.startSy);
      canvas.drawPath(path, paintSelectbox);
    }
  }

  drawArrow(canvas, points, color) {
    var zoom = c.zoom;
    var dx = -1 * (c.sx + c.currentSx) * zoom;
    var dy = -1 * (c.sy + c.currentSy) * zoom;
    //var x = point.dx * zoom + dx;
    //var y = point.dy * zoom + dy;

    var y1 = points.items[0].dy;
    var y2 = points.items[0].dy;
    var x1 = points.items[0].dx;
    var x2 = points.items[0].dx;

    if (points.items.length > 1) {
      y2 = points.items[1].dy;
      x2 = points.items[1].dx;
    }

    for (var i = 2; i < points.items.length; i++) {
      if (x1 == x2 && y1 == y2) {
        y2 = points.items[i].dy;
        x2 = points.items[i].dx;
      }
    }

    var sy = y2 - y1;
    var sx = x2 - x1;

    var angle = atan(sy / sx) * (180.0 / 3.14);

    angle += 180;

    if (sx < 0.0) {
      angle += 180.0;
    } else {
      if (sy < 0.0) angle += 360.0;
    }

    var r = (c.numberZoom / 2.0) * zoom * 0.5 / 2.0;

    if (points.icon == inclinationHorizontal) {
      if (points.items.length < 2) {
        return;
      }

      var point = points.items[0];
      var point2 = points.items[points.items.length - 1];

      if (point.dx > point2.dx) {
        angle = 0;
      } else {
        angle = 180;
      }
    } else if (points.icon == inclinationVertical) {
      if (points.items.length < 2) {
        return;
      }

      var point = points.items[0];
      var point2 = points.items[points.items.length - 1];

      if (point.dy > point2.dy) {
        angle = 90;
      } else {
        angle = 270;
      }
    }

    Path path = Path();

    double x;
    double y;

    x = cos(angle * 3.14 / 180) * r;
    y = sin(angle * 3.14 / 180) * r;

    x1 -= x;
    y1 -= y;

    path.moveTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    angle += 120;

    x = cos(angle * 3.14 / 180) * r;
    y = sin(angle * 3.14 / 180) * r;

    path.lineTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    path.lineTo((x1) * zoom + dx, (y1) * zoom + dy);

    angle += 120;

    x = cos(angle * 3.14 / 180) * r;
    y = sin(angle * 3.14 / 180) * r;

    path.lineTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    angle += 120;

    x = cos(angle * 3.14 / 180) * r;
    y = sin(angle * 3.14 / 180) * r;

    path.lineTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    if (points.selected == true) {
      canvas.drawPath(path, paintFillSelect);
    } else {
      canvas.drawPath(path, color);
    }
  }

  drawIcon(canvas, points) {
    var zoom = c.zoom;
    var dx = -1 * (c.sx + c.currentSx) * zoom;
    var dy = -1 * (c.sy + c.currentSy) * zoom;
    if (points.items.length == 0) {
      return;
    }
    var point = points.items[points.items.length - 1];
    var x = point.dx * zoom + dx;
    var y = point.dy * zoom + dy;

    Paint paint;

    Path path = Path();

    var step = (c.iconZoom / 2.0) * zoom;
    var stepMiddle = step * 0.66;
    var stepSmall = step * 0.33;

    // var strokeWidth = step * 0.15;
    var strokeWidth = 2.0;

    if (points.icon == crackLineRed ||
        points.icon == crackLineBlue ||
        points.icon == crackLineViolet ||
        points.icon == crackCurveRed ||
        points.icon == crackCurveBlue ||
        points.icon == crackCurveViolet) {
      step = (c.crackZoom / 2.0) * zoom;
      stepMiddle = step * 0.66;
      stepSmall = step * 0.33;
    }

    var stepNumber = (c.numberZoom / 2.0) * zoom;

    if (points.icon == basicVertical ||
        points.icon == basicHorizontal ||
        points.icon == materialVertical ||
        points.icon == materialHorizontal ||
        points.isInclination() == true ||
        points.isBasicLine() == true) {
      Color color = Colors.red;

      if (points.icon == basicVertical ||
          points.icon == materialVertical ||
          points.icon == basicVerticalLine ||
          points.icon == basicVerticalBreak) {
        paint = paintRed;
        color = Colors.red;
      } else if (points.icon == basicHorizontal ||
          points.icon == materialHorizontal ||
          points.icon == basicHorizontalLine ||
          points.icon == basicHorizontalBreak) {
        paint = paintBlue;
        color = const Color.fromRGBO(0, 0, 255, 1.0);
      } else {
        paint = paintRed;
        color = Colors.red;
      }

      if (points.selected == true) {
        paint = paintSelect;
      }

      var offset = Offset(x, y);
      var textSpan = TextSpan(
        text: points.number.toString(),
        style: TextStyle(
            fontSize: stepNumber * 1.2,
            fontWeight: FontWeight.w600,
            color: color),
      );

      if (points.selected == true) {
        canvas.drawCircle(offset, stepNumber, paintFillSelect);
        canvas.drawCircle(offset, stepNumber, paintSelect);

        textSpan = TextSpan(
          text: points.number.toString(),
          style: TextStyle(
              fontSize: stepNumber * 1.2,
              fontWeight: FontWeight.w600,
              color: Colors.white),
        );
      } else {
        // paint.strokeWidth = stepNumber * 0.15;
        paint.strokeWidth = 2.0;
        canvas.drawCircle(offset, stepNumber, paintFillWhite);
        canvas.drawCircle(offset, stepNumber, paint);
      }

      final textPainter = TextPainter()
        ..text = textSpan
        ..textDirection = TextDirection.ltr
        ..textAlign = TextAlign.center
        ..layout();

      offset = Offset(x - textPainter.width / 2, y - textPainter.height / 2);

      textPainter.paint(canvas, offset);

      paint.strokeWidth = strokeWidth;

      return;
    } else if (points.icon == 31 || points.icon == 41) {
      paint = paintLightblue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;
    } else if (points.icon == 32 || points.icon == 42) {
      paint = paintRed;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;
    } else if (points.icon == lineGreen || points.icon == curveGreen) {
      paint = paintGreen;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;
    } else if (points.icon == 34 || points.icon == 44) {
      paint = paintViolet;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;
    } else if (points.icon == 101) {
      paint = paintLightblue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      path.moveTo(x - step, y - step);
      path.lineTo(x + step, y + step);

      path.moveTo(x + step, y - step);
      path.lineTo(x - step, y + step);
    } else if (points.icon == 130) {
      paint = paintRed;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      path.moveTo(x - step, y - step);
      path.lineTo(x + step, y + step);

      path.moveTo(x + step, y - step);
      path.lineTo(x - step, y + step);
    } else if (points.icon == 131) {
      paint = paintGreen;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      path.moveTo(x - step, y - step);
      path.lineTo(x + step, y + step);

      path.moveTo(x + step, y - step);
      path.lineTo(x - step, y + step);
    } else if (points.icon == 102) {
      paint = paintLightblue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      path.moveTo(x - step, y);
      path.lineTo(x, y - step);
      path.lineTo(x + step, y);
      path.lineTo(x, y + step);
      path.lineTo(x - step, y);
    } else if (points.icon == 132) {
      paint = paintRed;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      path.moveTo(x - step, y);
      path.lineTo(x, y - step);
      path.lineTo(x + step, y);
      path.lineTo(x, y + step);
      path.lineTo(x - step, y);
    } else if (points.icon == 103) {
      paint = paintFillLightblue;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      paint.strokeWidth = strokeWidth;

      path.moveTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(150.0 * 3.14 / 180.0) * step,
          y + sin(150.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(270.0 * 3.14 / 180.0) * step,
          y + sin(270.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
    } else if (points.icon == 133) {
      paint = paintFillRed;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      paint.strokeWidth = strokeWidth;

      path.moveTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(150.0 * 3.14 / 180.0) * step,
          y + sin(150.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(270.0 * 3.14 / 180.0) * step,
          y + sin(270.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
    } else if (points.icon == 104) {
      paint = paintFillRed;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      paint.strokeWidth = strokeWidth;

      path.moveTo(x - step, y - step);
      path.lineTo(x + step, y - step);
      path.lineTo(x + step, y + step);
      path.lineTo(x - step, y + step);
      path.lineTo(x - step, y - step);
    } else if (points.icon == 134) {
      paint = paintFillLightblue;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      paint.strokeWidth = strokeWidth;

      path.moveTo(x - step, y - step);
      path.lineTo(x + step, y - step);
      path.lineTo(x + step, y + step);
      path.lineTo(x - step, y + step);
      path.lineTo(x - step, y - step);
    } else if (points.icon == 105) {
      paint = paintBlue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), stepMiddle, paint);
    } else if (points.icon == 106) {
      paint = paintBlue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), step, paint);

      path.moveTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(150.0 * 3.14 / 180.0) * step,
          y + sin(150.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(270.0 * 3.14 / 180.0) * step,
          y + sin(270.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
    } else if (points.icon == 107 ||
        points.icon == 108 ||
        points.icon == 109 ||
        points.icon == 110) {
      paint = paintBlue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), stepSmall, paint);

      var angles = [135.0, 90.0, 45.0, 0.0];
      var angle = angles[points.icon - 107];

      path.moveTo(x + cos(angle * 3.14 / 180.0) * step,
          y + sin(angle * 3.14 / 180.0) * step);
      angle += 180.0;
      path.lineTo(x + cos(angle * 3.14 / 180.0) * step,
          y + sin(angle * 3.14 / 180.0) * step);
    } else if (points.isCrack()) {
      paint = paintRed;

      if (points.icon == crackLineBlue || points.icon == crackCurveBlue) {
        paint = paintBlue;
      } else if (points.icon == crackLineViolet ||
          points.icon == crackCurveViolet) {
        paint = paintViolet;
      }

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      var x1 = points.items[0].dx;
      var y1 = points.items[0].dy;
      var x2 = points.items[points.items.length - 1].dx;
      var y2 = points.items[points.items.length - 1].dy;

      var cx = x1;
      var cy = y1;
      if (x1 > x2) {
        cx = x2 + (x1 - x2) / 2;
      } else {
        cx = x1 + (x2 - x1) / 2;
      }

      if (y1 > y2) {
        cy = y2 + (y1 - y2) / 2;
      } else {
        cy = y1 + (y2 - y1) / 2;
      }

      if (points.isCrackCurve() && points.items.length > 2) {
        var lengths = List<double>.empty(growable: true);
        double total = 0.0;
        var old = points.items[0];
        for (var i = 1; i < points.items.length; i++) {
          var cu = points.items[i];
          double length = sqrt((old.dx - cu.dx).abs() + (old.dy - cu.dy).abs());
          lengths.add(length);
          total += length;

          old = cu;
        }

        double half = total / 2;
        double current = 0.0;
        for (var i = 0; i < lengths.length; i++) {
          current += lengths[i];

          if (current >= half) {
            x1 = points.items[i].dx;
            y1 = points.items[i].dy;
            x2 = points.items[i + 1].dx;
            y2 = points.items[i + 1].dy;

            cx = x1;
            cy = y1;
            if (x1 > x2) {
              cx = x2 + (x1 - x2) / 2;
            } else {
              cx = x1 + (x2 - x1) / 2;
            }

            if (y1 > y2) {
              cy = y2 + (y1 - y2) / 2;
            } else {
              cy = y1 + (y2 - y1) / 2;
            }

            break;
          }
        }
      }

      canvas.drawCircle(
          Offset(cx * zoom + dx, cy * zoom + dy), stepSmall, paint);
    } else if (points.icon == 115) {
      paint = paintViolet;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), stepMiddle, paint);
    } else if (points.icon == 116) {
      paint = paintViolet;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), step, paint);

      path.moveTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(150.0 * 3.14 / 180.0) * step,
          y + sin(150.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(270.0 * 3.14 / 180.0) * step,
          y + sin(270.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
    } else if (points.icon == 117 ||
        points.icon == 118 ||
        points.icon == 119 ||
        points.icon == 120) {
      paint = paintViolet;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), stepSmall, paint);

      var angles = [135.0, 90.0, 45.0, 0.0];
      var angle = angles[points.icon - 117];

      path.moveTo(x + cos(angle * 3.14 / 180.0) * step,
          y + sin(angle * 3.14 / 180.0) * step);
      angle += 180.0;
      path.lineTo(x + cos(angle * 3.14 / 180.0) * step,
          y + sin(angle * 3.14 / 180.0) * step);
    } else if (points.icon == 111) {
      paint = paintFillRed;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), stepMiddle, paint);
    } else if (points.icon == 112) {
      paint = paintFillBlue;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), stepMiddle, paint);
    } else if (points.icon == fiberHorizontal) {
      paint = paintFillBlue;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), stepMiddle, paint);
    } else if (points.icon == fiberVertical) {
      paint = paintFillRed;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      paint.strokeWidth = strokeWidth;

      canvas.drawCircle(Offset(x, y), stepMiddle, paint);
    } else {
      paint = paintBlue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      paint.strokeWidth = strokeWidth;
    }

    canvas.drawPath(path, paint);
  }

  @override
  bool shouldRepaint(PainterDrawer oldDelegate) {
    return true;
  }

  drawBreakArrow(canvas, points, color) {
    var zoom = c.zoom;

    var r = (c.numberZoom / 2.0) * zoom * 0.5 / 2.0;

    var dx = -1 * (c.sx + c.currentSx) * zoom;
    var dy = -1 * (c.sy + c.currentSy) * zoom;
    //var x = point.dx * zoom + dx;
    //var y = point.dy * zoom + dy;

    var y1 = points.items[0].dy;
    var y2 = points.items[0].dy;
    var x1 = points.items[0].dx;
    var x2 = points.items[0].dx;

    if (points.items.length <= 1) {
      return;
    }

    y2 = points.items[1].dy;
    x2 = points.items[1].dx;

    for (var i = 2; i < points.items.length; i++) {
      if (x1 == x2 && y1 == y2) {
        y2 = points.items[i].dy;
        x2 = points.items[i].dx;
      }
    }

    var sy = y2 - y1;
    var sx = x2 - x1;

    var m = sy / sx;
    var d = 50;
    var bx = d / sqrt(1 + m * m);
    var by = (d * m) / sqrt(1 + m * m);

    x2 = bx;
    y2 = by;

    Paint color2;

    if (points.icon == basicVerticalBreak) {
      color2 = paintRed;
    } else {
      color2 = paintBlue;
    }

    var angle = atan(sy / sx) * (180.0 / 3.14);

    angle += 180;

    if (sx < 0.0) {
      angle += 180.0;
    } else {
      if (sy < 0.0) angle += 360.0;
    }

    if (points.icon == inclinationHorizontal) {
      if (points.items.length < 2) {
        return;
      }

      var point = points.items[0];
      var point2 = points.items[points.items.length - 1];

      if (point.dx > point2.dx) {
        angle = 0;
      } else {
        angle = 180;
      }
    } else if (points.icon == inclinationVertical) {
      if (points.items.length < 2) {
        return;
      }

      var point = points.items[0];
      var point2 = points.items[points.items.length - 1];

      if (point.dy > point2.dy) {
        angle = 90;
      } else {
        angle = 270;
      }
    }

    double x;
    double y;

    x = cos(angle * 3.14 / 180) * r;
    y = sin(angle * 3.14 / 180) * r;

    var originalX = x1;
    var originalY = y1;

    x1 -= x;
    y1 -= y;

    Path path = Path();

    // 반대선
    var cx = cos((angle + 90) * 3.14 / 180) * r;
    var cy = sin((angle + 90) * 3.14 / 180) * r;
    path.moveTo((x1 + x + cx) * zoom + dx, (y1 + y + cy) * zoom + dy);

    cx = cos((angle + 90 + 180) * 3.14 / 180) * r;
    cy = sin((angle + 90 + 180) * 3.14 / 180) * r;
    path.lineTo((x1 + x + cx) * zoom + dx, (y1 + y + cy) * zoom + dy);

    // 연장선
    cx = cos((angle) * 3.14 / 180) * r * 2;
    cy = sin((angle) * 3.14 / 180) * r * 2;
    path.moveTo((x1 + x + cx) * zoom + dx, (y1 + y + cy) * zoom + dy);

    cx = cos((angle + 180) * 3.14 / 180) * r * 2;
    cy = sin((angle + 180) * 3.14 / 180) * r * 2;
    path.lineTo((x1 + x + cx) * zoom + dx, (y1 + y + cy) * zoom + dy);

    if (points.selected == true) {
      canvas.drawPath(path, paintSelect);
    } else {
      canvas.drawPath(path, color2);
    }

    // path = Path();

    // path.moveTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    // angle += 120;

    // x = cos(angle * 3.14 / 180) * r;
    // y = sin(angle * 3.14 / 180) * r;

    // path.lineTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    // path.lineTo((x1) * zoom + dx, (y1) * zoom + dy);

    // angle += 120;

    // x = cos(angle * 3.14 / 180) * r;
    // y = sin(angle * 3.14 / 180) * r;

    // path.lineTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    // angle += 120;

    // x = cos(angle * 3.14 / 180) * r;
    // y = sin(angle * 3.14 / 180) * r;

    // path.lineTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    // if (points.selected == true) {
    //   canvas.drawPath(path, paintFillSelect);
    // } else {
    //   canvas.drawPath(path, color);
    // }

    // 반대 화살표
    path = Path();

    angle += 180;

    x = cos(angle * 3.14 / 180) * r;
    y = sin(angle * 3.14 / 180) * r;

    x1 = originalX - x;
    y1 = originalY - y;

    path.moveTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    angle += 120;

    x = cos(angle * 3.14 / 180) * r;
    y = sin(angle * 3.14 / 180) * r;

    path.lineTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    path.lineTo((x1) * zoom + dx, (y1) * zoom + dy);

    angle += 120;

    x = cos(angle * 3.14 / 180) * r;
    y = sin(angle * 3.14 / 180) * r;

    path.lineTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    angle += 120;

    x = cos(angle * 3.14 / 180) * r;
    y = sin(angle * 3.14 / 180) * r;

    path.lineTo((x1 + x) * zoom + dx, (y1 + y) * zoom + dy);

    if (points.selected == true) {
      canvas.drawPath(path, paintFillSelect);
    } else {
      canvas.drawPath(path, color);
    }
  }
}
