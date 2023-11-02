import 'package:dynamic_color/dynamic_color.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:petrolium/core/constants/storage_constants.dart';
import 'package:petrolium/core/routes/app_routes.dart';
import 'package:petrolium/core/services/storage/storage_service.dart';
import 'package:petrolium/core/services/theme/theme_service.dart';
import 'package:petrolium/initializer.dart';
import 'package:petrolium/ui/theme.dart';

void main() async {
  await Initializer.initApp();
  runApp(
    const ProviderScope(
      child: Petrolium(),
    ),
  );
}

class Petrolium extends ConsumerStatefulWidget {
  const Petrolium({super.key});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() => _PetroliumState();
}

class _PetroliumState extends ConsumerState<Petrolium> {
  @override
  void initState() {
    super.initState();
    final themeService = ref.read(themeServiceProvider.notifier);
    final LocalStorageService storage = LocalStorageService();
    final key = storage.box.get(StorageConstants.userThemeKey);
    if (key != null) {
      themeService.fromThemeKey(key);
    }
  }

  @override
  Widget build(BuildContext context) {
    final themeMode = ref.watch(themeServiceProvider);
    final routes = ref.watch(routerProvider);

    return DynamicColorBuilder(
      builder: (lightColorScheme, darkColorScheme) {
        return MaterialApp.router(
          title: 'Petrolium',
          theme: getLightModeData(lightColorScheme),
          darkTheme: getDarkModeData(darkColorScheme),
          themeMode: themeMode,
          debugShowCheckedModeBanner: false,
          routerConfig: routes,
        );
      },
    );
  }
}
