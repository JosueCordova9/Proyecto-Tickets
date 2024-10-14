import 'package:flutter/material.dart';
import 'package:tickets/login.dart';
import 'package:tickets/registro.dart';

void main() {
  runApp(const MainApp());
}

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      initialRoute: '/',
      routes: {
        '/': (context) => const Login(),
        '/registro': (context) => const Registro(),
      },
    );
  }
}