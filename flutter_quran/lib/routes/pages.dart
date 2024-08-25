import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/main_navigation/view.dart';
import 'package:flutter_quran/routes/routes.dart';

class Pages {
  static const initialRoute = Routes.mainPageRoute;
  static Map<String, WidgetBuilder> routes = {
    // Routes.splashScreenRoute: (context) => const SplashScreen(),
    Routes.mainPageRoute: (context) => const MainNavigation(),
    // Routes.createAssetTrackingScreenRoute: (context) =>
    //     const CreateAssetTrackingScreen(),
    // Routes.voiceRecordRoute: (context) => const VoiceRecordScreen(),
    // Routes.voiceDetailsRoute: (context) => const VoiceDetailsScreen(),
  };
}
