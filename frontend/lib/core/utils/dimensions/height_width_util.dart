import 'package:flutter/widgets.dart';

(double width, double height) getWidthHeight(BuildContext context) {
  final width = MediaQuery.of(context).size.width;
  final height = MediaQuery.of(context).size.height;

  return (width, height);
}
