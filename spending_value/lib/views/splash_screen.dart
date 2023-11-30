import 'package:flutter/material.dart';
import 'package:spending_value/routes/routes.dart';

class SplashScreen extends StatelessWidget {
  const SplashScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).splashColor,
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Image.asset(
              'assets/logo/company_logo.jpg',
              width: 200,
              height: 200,
              fit: BoxFit.contain,
            ),
            const Text("Company Name"),
            const CircularProgressIndicator(),
            ElevatedButton(
                onPressed: () =>
                    Navigator.popAndPushNamed(context, Routes.homeScreenRoute),
                child: const Text("Click"))
          ],
        ),
      ),
    );
  }
}
