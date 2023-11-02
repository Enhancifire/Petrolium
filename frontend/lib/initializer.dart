import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/services.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:loggy/loggy.dart';
import 'package:petrolium/core/constants/storage_constants.dart';
import 'package:petrolium/core/services/theme/theme_service.dart';
import 'package:petrolium/core/services/storage/storage_service.dart';
import 'package:petrolium/firebase_options.dart';

abstract class Initializer {
  static Future<void> initApp() async {
    WidgetsFlutterBinding.ensureInitialized();
    _initializeLogger();
    await _initFirebase();
    await _initStorage();
    await _initTheme();
    _initLooks();
  }

  static Future<void> _initFirebase() async {
    try {
      await Firebase.initializeApp(
        options: DefaultFirebaseOptions.currentPlatform,
      );
    } catch (e) {
      logError(e);
    }
  }

  static void _initializeLogger() {
    Loggy.initLoggy(
      logPrinter: const PrettyPrinter(showColors: true),
    );
  }

  static Future<void> _initStorage() async {
    await LocalStorageService.initialize();
  }

  static Future<void> _initTheme() async {
    final storage = LocalStorageService();
    final key = storage.box.get(StorageConstants.userThemeKey);

    if (key != null) {
      final themeService = ThemeService();
      themeService.fromThemeKey(key);
    }
  }

  static void _initLooks() {
    SystemChrome.setEnabledSystemUIMode(
      SystemUiMode.edgeToEdge,
    );
  }
}
