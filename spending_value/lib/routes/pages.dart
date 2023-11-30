import 'package:flutter/material.dart';
import 'package:spending_value/views/assetTracking/create_asset_tracking_screen.dart';
import 'package:spending_value/routes/routes.dart';
import 'package:spending_value/views/voice/voice_details_screen.dart';
import 'package:spending_value/views/home/home_screen.dart';
import 'package:spending_value/views/voice/voice_record_screen.dart';
import 'package:spending_value/views/splash_screen.dart';

class Pages {
  static const initialRoute = Routes.splashScreenRoute;
  static Map<String, WidgetBuilder> routes = {
    Routes.splashScreenRoute: (context) => const SplashScreen(),
    Routes.homeScreenRoute: (context) => const HomeScreen(),
    Routes.createAssetTrackingScreenRoute: (context) =>
        const CreateAssetTrackingScreen(),
    Routes.voiceRecordRoute: (context) => const VoiceRecordScreen(),
    Routes.voiceDetailsRoute: (context) => const VoiceDetailsScreen(),
  };
}
