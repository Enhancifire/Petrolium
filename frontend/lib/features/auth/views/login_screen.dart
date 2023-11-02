import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:petrolium/core/services/theme/theme_service.dart';
import 'package:petrolium/core/utils/dimensions/height_width_util.dart';
import 'package:petrolium/core/utils/dimensions/spacers.dart';
import 'package:petrolium/ui/widgets/buttons.dart';

class LoginScreen extends ConsumerStatefulWidget {
  const LoginScreen({super.key});

  @override
  ConsumerState<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends ConsumerState<LoginScreen> {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    final (_, height) = getWidthHeight(context);
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.all(24.0),
          child: Column(
            children: [
              Text(
                'Petrolium',
                style: Theme.of(context).textTheme.displayLarge,
              ),
              addVerticalSpacer(height * 0.1),
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  "Log in",
                  style: Theme.of(context).textTheme.headlineMedium,
                ),
              ),
              addVerticalSpacer(height * 0.02),
              TextField(
                controller: _emailController,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  hintText: 'Enter your email',
                ),
              ),
              addVerticalSpacer(height * 0.02),
              TextField(
                controller: _passwordController,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  hintText: 'Enter your password',
                ),
                obscureText: true,
              ),
              addVerticalSpacer(height * 0.01),
              Align(
                alignment: Alignment.centerRight,
                child: TextButton(
                  onPressed: () {},
                  child: Text('Forgot Password?'),
                ),
              ),
              addVerticalSpacer(height * 0.03),
              CenteredElevatedButton(
                onPressed: () {
                  final themeService = ref.watch(themeServiceProvider.notifier);
                  themeService.flipTheme();
                },
                text: "Log in",
              ),
              addVerticalSpacer(height * 0.02),
              Stack(
                alignment: Alignment.center,
                children: [
                  const Divider(),
                  Container(
                    color: Theme.of(context).scaffoldBackgroundColor,
                    padding: const EdgeInsets.all(8),
                    child: Text(
                      "Or",
                      style: Theme.of(context).textTheme.bodyLarge,
                    ),
                  ),
                ],
              ),
              addVerticalSpacer(height * 0.02),
              CenteredIconElevatedButton(
                icon: const FaIcon(FontAwesomeIcons.google),
                text: "Sign in with Google",
                onPressed: () {},
              ),
              addVerticalSpacer(height * 0.01),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    "No Account?",
                    style: Theme.of(context).textTheme.bodyMedium,
                  ),
                  TextButton(
                    style: TextButton.styleFrom(
                      padding: const EdgeInsets.all(5),
                    ),
                    onPressed: () {},
                    child: const Text("Create one"),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
