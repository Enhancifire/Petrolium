import 'package:flutter/material.dart';
import 'package:petrolium/core/constants/color_constants.dart';

ThemeData getLightModeData(ColorScheme? lightDynamicData) {
  final ThemeData lightModeThemeData = ThemeData(
    useMaterial3: true,
    colorScheme: lightDynamicData ??
        ColorScheme.fromSeed(
          seedColor: Colors.deepPurple,
          brightness: Brightness.light,
        ),
    elevatedButtonTheme: ElevatedButtonThemeData(
      style: ElevatedButton.styleFrom(
        elevation: 0,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(8),
          side: const BorderSide(
            color: AppColors.borderLight,
          ),
        ),
        padding: const EdgeInsets.all(16),
      ),
    ),
  );
  return lightModeThemeData;
}

ThemeData getDarkModeData(ColorScheme? darkDynamicData) {
  final ThemeData darkModeThemeData = ThemeData(
    useMaterial3: true,
    colorScheme: darkDynamicData ??
        ColorScheme.fromSeed(
          seedColor: Colors.deepPurple,
          brightness: Brightness.dark,
        ),
    textTheme: TextTheme(
      displayLarge: TextStyle(),
    ),
    elevatedButtonTheme: ElevatedButtonThemeData(
      style: ElevatedButton.styleFrom(
        elevation: 0,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(8),
          side: const BorderSide(
            color: AppColors.borderDark,
          ),
        ),
        padding: const EdgeInsets.all(16),
      ),
    ),
  );
  return darkModeThemeData;
}
