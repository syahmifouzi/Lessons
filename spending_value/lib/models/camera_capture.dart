import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'package:provider/provider.dart';
import 'package:spending_value/controllers/assetTracking/create_asset_tracking_store.dart';

class CameraCapture extends StatefulWidget {
  const CameraCapture({super.key});

  @override
  State<CameraCapture> createState() => _CameraCaptureState();
}

class _CameraCaptureState extends State<CameraCapture> {
  List<XFile>? _mediaFileList;
  final _picker = ImagePicker();

  void _setImageFileListFromFile(XFile? value) {
    _mediaFileList = value == null ? null : <XFile>[value];
  }

  Future _imgFromCamera() async {
    final XFile? pickedFile =
        await _picker.pickImage(source: ImageSource.camera, imageQuality: 100);

    setState(() {
      _setImageFileListFromFile(pickedFile);
      context
          .read<CreateAssetTrackingStore>()
          .addReceiptImage(_mediaFileList ?? []);
    });
  }

  Future _imgFromGallery() async {
    final List<XFile> pickedFileList =
        await _picker.pickMultiImage(imageQuality: 100);

    setState(() {
      _mediaFileList = pickedFileList;
      context
          .read<CreateAssetTrackingStore>()
          .addReceiptImage(_mediaFileList ?? []);
    });
  }

  void _showPicker(context) {
    showModalBottomSheet(
      context: context,
      builder: (BuildContext context) {
        return SafeArea(
          child: Wrap(
            children: <Widget>[
              ListTile(
                leading: const Icon(Icons.photo_library),
                title: const Text('Photo Library'),
                onTap: () {
                  _imgFromGallery();
                  Navigator.of(context).pop();
                },
              ),
              ListTile(
                leading: const Icon(Icons.photo_camera),
                title: const Text('Camera'),
                onTap: () {
                  _imgFromCamera();
                  Navigator.of(context).pop();
                },
              ),
            ],
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return IconButton(
        onPressed: () {
          _showPicker(context);
        },
        icon: const Icon(Icons.camera_alt));
  }
}
