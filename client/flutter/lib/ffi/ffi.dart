import 'dart:ffi';
import 'package:ffi/ffi.dart';
import 'dart:io'; // For Platform.isX

DynamicLibrary findDynamicLibrary(String name, String dir) {
  if (!dir.endsWith('/')) dir = dir + '/';
  if (Platform.isAndroid) return DynamicLibrary.open('lib$name.so');
  if (Platform.isLinux) return DynamicLibrary.open('${dir}lib$name.so');
  if (Platform.isMacOS) return DynamicLibrary.open('${dir}lib$name.dylib');
  if (Platform.isWindows) return DynamicLibrary.open('$dir$name.dll');
  return DynamicLibrary.process();
}

final DynamicLibrary nativeAddLib = findDynamicLibrary("rust",'libraries');
final DynamicLibrary nativeHelloLib = findDynamicLibrary("hello",'libraries');


final Pointer<Utf8> Function(Pointer<Utf8> x) rustGreeting =
nativeAddLib
    .lookup<NativeFunction<Pointer<Utf8> Function(Pointer<Utf8>)>>("rust_greeting")
    .asFunction();

String callRustGreeting(){
  final String myString = "😎👿💬";
  final Pointer<Utf8> charPointer = rustGreeting(myString.toNativeUtf8());
  return charPointer.toDartString();
}

final Pointer<Utf8> Function(Pointer<Utf8> x) goPrint =
nativeHelloLib
    .lookup<NativeFunction<Pointer<Utf8> Function(Pointer<Utf8>)>>("goprint")
    .asFunction();

String callGoPrint(String s){
  final Pointer<Utf8> charPointer = goPrint(s.toNativeUtf8());
  return charPointer.toDartString();
}