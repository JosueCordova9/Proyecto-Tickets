// import 'dart:convert';
import "package:flutter/material.dart";
// import 'package:http/http.dart' as http;



class Registro extends StatefulWidget {
  const Registro({super.key});

  @override
  State<Registro> createState() => _RegistroState();
}

class _RegistroState extends State<Registro> {
  String nombre = '';
  String usuario = '';
  String password = '';

  final formKey = GlobalKey<FormState>();

  registrarUsuario() {
    if (formKey.currentState!.validate()) {
    //   final response = http.post(
    //     Uri.parse('localhost:8080/registrar'),
    //     headers: <String, String>{
    //       'Content-Type': 'application/json; charset=UTF-8',
    //     },
    //     body: jsonEncode(<String, String>{
    //       'nombre': nombre,
    //       'usuario': usuario,
    //       'password': password
    //     }),
    //   );
    //   if (response.statusCode == 201) {
    //   // Registro exitoso
    //   print('Usuario registrado exitosamente');
    // } else {
    //   // Manejar error
    //   print('Error al registrar usuario: ${response.body}');
    // }
      Navigator.pushNamed(context, '/');
    }
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Center(child: Text("Página de registro")),
        ),
        body: Center(
          child: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Form(
              key: formKey,
              child: Column(
                children: [
                  TextFormField(
                    decoration: const InputDecoration(labelText: "Nombre"),
                    onChanged: (value) {
                      nombre = value;
                    },
                    validator: (value) {
                      if (value!.isEmpty) {
                        return "Ingrese un nombre";
                      }
                      return null;
                    },
                  ),
                  TextFormField(
                    decoration: const InputDecoration(labelText: "Usuario"),
                    onChanged: (value) {
                      usuario = value;
                    },
                    validator: (value) {
                      if (value!.isEmpty) {
                        return "Ingrese un usuario";
                      }
                      return null;
                    },
                  ),
                  TextFormField(
                    decoration: const InputDecoration(labelText: "Contraseña"),
                    onChanged: (value) {
                      usuario = value;
                    },
                    validator: (value) {
                      if (value!.isEmpty) {
                        return "Ingrese una Contraseña";
                      }
                      return null;
                    },
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      TextButton(
                          onPressed: registrarUsuario(),
                          child: const Text("Registrarse")),
                      TextButton(
                        onPressed: () {
                          Navigator.pushNamed(context, '/');
                        },
                        child: const Text("Regresar"),
                      ),
                    ],
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}