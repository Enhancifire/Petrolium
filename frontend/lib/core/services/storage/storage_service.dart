import 'package:hive_flutter/adapters.dart';
import 'package:petrolium/core/constants/storage_constants.dart';

class LocalStorageService {
  LocalStorageService._privateConstructor();

  static final LocalStorageService _instance =
      LocalStorageService._privateConstructor();

  factory LocalStorageService() {
    return _instance;
  }

  static initialize() async {
    await Hive.initFlutter();
    await Hive.openBox(StorageConstants.localStorageDatabaseKey);
  }

  final box = Hive.box(StorageConstants.localStorageDatabaseKey);
}
