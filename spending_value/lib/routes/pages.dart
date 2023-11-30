import 'package:flutter/material.dart';
import 'package:spending_value/views/assetTracking/create_asset_tracking_screen.dart';
import 'package:spending_value/routes/routes.dart';
import 'package:spending_value/views/home/home_screen.dart';
import 'package:spending_value/views/record_audio_screen.dart';
import 'package:spending_value/views/splash_screen.dart';

class Pages {
  static const initialRoute = Routes.splashScreenRoute;
  static Map<String, WidgetBuilder> routes = {
    Routes.splashScreenRoute: (context) => const SplashScreen(),
    Routes.homeScreenRoute: (context) => const HomeScreen(),
    Routes.createAssetTrackingScreenRoute: (context) =>
        const CreateAssetTrackingScreen(),
    Routes.recordAudioRoute: (context) => const RecordAudioScreen(),
  };
}
