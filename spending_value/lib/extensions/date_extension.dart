import 'package:intl/intl.dart';

extension MyDateFormat on DateTime {
  String toYMD() {
    return DateFormat('yyyy/MM/dd').format(this);
  }

  String toDMY() {
    return DateFormat('dd/MM/yyyy').format(this);
  }
}
