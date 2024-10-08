import "package:flutter/material.dart";

class Login extends StatefulWidget {
  const Login({super.key});

  @override
  State<Login> createState() => _LoginState();
}

class _LoginState extends State<Login> {
  final formKey = GlobalKey<FormState>();

  String usuario = '';
  String password = '';

  submitForm() {
    if (formKey.currentState!.validate()) {}
    // print('Usuario: $usuario, Password: $password');
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Center(child: Text("Inicio de Sesion")),
        ),
        body: Center(
            child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Form(
            key: formKey,
            child: Column(
              children: [
                TextFormField(
                  decoration: const InputDecoration(labelText: "Usuario"),
                  onChanged: (value) {
                    usuario = value;
                  },
                  validator: (value) {
                    if (value == null || value.isEmpty) {
                      return "Ingrese el nombre de usuario";
                    }
                    return null;
                  },
                ),
                TextFormField(
                  decoration: const InputDecoration(labelText: "Contraseña"),
                  onChanged: (value) {
                    password = value;
                  },
                  validator: (value) {
                    if (value == null || value.isEmpty) {
                      return "Ingrese su contraseña";
                    }
                    return null;
                  },
                ),
                TextButton(
                    onPressed: submitForm, child: const Text("Iniciar Sesión")),
                TextButton(
                    onPressed: () {
                      Navigator.pushNamed(context, '/registro');
                    },
                    child: const Text("Página de registro")),
              ],
            ),
          ),
        )),
      ),
    );
  }
}
