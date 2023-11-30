import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:spending_value/controllers/assetTracking/create_asset_tracking_store.dart';
import 'package:spending_value/routes/routes.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          context.read<CreateAssetTrackingStore>().clearReceiptImages();
          Navigator.pushNamed(context, Routes.createAssetTrackingScreenRoute);
        },
        tooltip: 'Create New',
        child: const Icon(Icons.add),
      ),
      body: Container(
        color: Colors.red,
        alignment: Alignment.center,
        child: ElevatedButton(
            onPressed: () {
              // Navigator.pushNamed(context, Routes.voiceRecordRoute);
            },
            child: const Text('Hello')),
      ),
    );
  }
}
