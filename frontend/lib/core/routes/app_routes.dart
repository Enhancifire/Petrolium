import 'package:petrolium/features/auth/views/login_screen.dart';
import 'package:petrolium/features/home/views/home_screen.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:go_router/go_router.dart';

part 'app_routes.g.dart';

abstract class AppRoutes {
  static String initialRoute = home;

  static String home = '/home';
  static String login = '/login';
  static String register = '/register';
}

@riverpod
class Router extends _$Router {
  @override
  GoRouter build() {
    return GoRouter(
      initialLocation: AppRoutes.initialRoute,
      routes: [
        GoRoute(
          path: AppRoutes.home,
          builder: (context, state) => const HomeScreen(),
        ),
        GoRoute(
          path: AppRoutes.login,
          builder: (context, state) => const LoginScreen(),
        ),
      ],
    );
  }
}
