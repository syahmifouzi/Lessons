# spending_value

A new Flutter project.

## Getting Started

This project is a starting point for a Flutter application.

A few resources to get you started if this is your first Flutter project:

- [Lab: Write your first Flutter app](https://docs.flutter.dev/get-started/codelab)
- [Cookbook: Useful Flutter samples](https://docs.flutter.dev/cookbook)

For help getting started with Flutter development, view the
[online documentation](https://docs.flutter.dev/), which offers tutorials,
samples, guidance on mobile development, and a full API reference.

To add a plugin:
1. flutter pub add <plugin-name>
2. flutterfire configure OR flutterfire configure --project=spending-value
3. overwrite the file during the configure process

To release apk:
(Make sure phone in dev mode and enable usb debugging)
flutter clean
flutter pub get
flutter build apk --split-per-abi
flutter install --use-application-binary=D:\GitHub\Lessons\spending_value\build\app\outputs\flutter-apk\app-arm64-v8a-release.apk