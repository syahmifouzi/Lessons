import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';

class CreateAssetTrackingStore extends ChangeNotifier {
  final List<XFile> _receiptImages = [];

  List<XFile> get receiptImages => _receiptImages;

  void addReceiptImage(List<XFile> x) {
    _receiptImages.addAll(x);
    notifyListeners();
  }

  void clearReceiptImages() {
    _receiptImages.clear();
    notifyListeners();
  }
}
