import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:petrolium/core/utils/dimensions/spacers.dart';

class CenteredElevatedButton extends StatelessWidget {
  const CenteredElevatedButton(
      {super.key, required this.text, required this.onPressed});

  final String text;
  final Function onPressed;

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: () {
        onPressed();
      },
      child: Center(
        child: Text(text),
      ),
    );
  }
}

class CenteredIconElevatedButton extends StatelessWidget {
  const CenteredIconElevatedButton({
    super.key,
    required this.text,
    required this.onPressed,
    required this.icon,
  });

  final String text;
  final Function onPressed;
  final Widget icon;

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: () {
        onPressed();
      },
      // icon: icon,
      child: Center(
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            icon,
            addHorizontalSpacer(20),
            Text(text),
          ],
        ),
      ),
    );
  }
}
