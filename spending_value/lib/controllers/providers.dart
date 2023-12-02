import 'package:provider/provider.dart';
import 'package:provider/single_child_widget.dart';
import 'package:spending_value/controllers/assetTracking/create_asset_tracking_store.dart';
import 'package:spending_value/controllers/audio_store.dart';
import 'package:spending_value/controllers/audiorecord_isolate_store.dart';
import 'package:spending_value/controllers/splash_store.dart';
import 'package:spending_value/controllers/surahname_store.dart';

class Providers {
  static List<SingleChildWidget> providers = [
    ChangeNotifierProvider<SplashStore>(create: (context) => SplashStore()),
    ChangeNotifierProvider<CreateAssetTrackingStore>(
        create: (context) => CreateAssetTrackingStore()),
    ChangeNotifierProvider<AudioStore>(create: (context) => AudioStore()),
    ChangeNotifierProvider<SurahnameStore>(
        create: (context) => SurahnameStore()),
    ChangeNotifierProvider<AudiorecordIsolateStore>(
        create: (context) => AudiorecordIsolateStore()),
  ];
}
