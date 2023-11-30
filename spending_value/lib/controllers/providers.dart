import 'package:provider/provider.dart';
import 'package:provider/single_child_widget.dart';
import 'package:spending_value/controllers/assetTracking/create_asset_tracking_store.dart';
import 'package:spending_value/controllers/splash_store.dart';

class Providers {
  static List<SingleChildWidget> providers = [
    ChangeNotifierProvider<SplashStore>(create: (context) => SplashStore()),
    ChangeNotifierProvider<CreateAssetTrackingStore>(
        create: (context) => CreateAssetTrackingStore()),
  ];
}
